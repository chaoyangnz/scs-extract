package scs

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-faster/city"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type HashfsHeader struct {
	magic        string //uint32 // 4
	version      uint16 // 2
	salt         uint16 // 2
	hashMethod   string //uint32 // 4
	entriesCount uint32 // 4
	startOffset  uint32 // 4
	// padding
	pad0 uint64 // 8
	pad1 uint32 // 4
} // 32 bytes

func NewHeader(buffer *Buffer) HashfsHeader {
	header := HashfsHeader{}
	header.magic = buffer.nextString(4)
	header.version = buffer.next2()
	header.salt = buffer.next2()
	header.hashMethod = buffer.nextString(4)
	header.entriesCount = buffer.next4()
	header.startOffset = buffer.next4()
	header.pad0 = buffer.next8()
	header.pad1 = buffer.next4()

	return header
}

const (
	FLAG_DIR        = (1 << 0)
	FLAG_COMPRESSED = (1 << 1)
	FLAG_VERIFY     = (1 << 2)
	FLAG_ENCRYPTED  = (1 << 3)
)

type Flags struct {
	dir        bool
	compressed bool
	verify     bool
	encrypted  bool
}

func NewFlags(value uint32) Flags {
	return Flags{
		dir:        (value & (1 << 0)) != 0,
		compressed: (value & (1 << 1)) != 0,
		verify:     (value & (1 << 2)) != 0,
		encrypted:  (value & (1 << 3)) != 0,
	}
}

type HashfsEntry struct {
	hash           uint64 // 8
	offset         uint64 // 8
	flags          Flags  //uint32 // 4
	crc            uint32 // 4
	size           uint32 // 4
	compressedSize uint32 // 4
} // 32 bytes

func NewEntry(buffer *Buffer) HashfsEntry {
	entry := HashfsEntry{}
	entry.hash = buffer.next8()
	entry.offset = buffer.next8()
	entry.flags = NewFlags(buffer.next4())
	entry.crc = buffer.next4()
	entry.size = buffer.next4()
	entry.compressedSize = buffer.next4()

	return entry
}

func NewSCS(path string) *SCS {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic("Failed to read file")
	}

	scs := &SCS{
		buffer: NewBuffer(bytes),
	}
	scs.init()
	return scs
}

func (this *SCS) readTree(path string) HashfsNode {
	if entry, ok := this.GetEntryByPath(path); ok {
		if entry.flags.dir {
			dir := this.readDir(entry)
			dir.name = filepath.Base(path)
			dir.path = path
			for i, name := range dir.data {
				if strings.HasPrefix(name, "*") {
					p := path + "/" + name[1:]
					child := this.readTree(p).(HashfsDir)
					child.name = name[1:]
					child.path = p
					dir.children[i] = child
				} else {
					if e, o := this.GetEntryByPath(path + "/" + name); o {
						child := this.readFile(e)
						child.name = name
						child.path = path + "/" + name
						dir.children[i] = child
					}
				}
			}

			return dir
		} else {
			f := this.readFile(entry)
			f.name = filepath.Base(path)
			f.path = path

			return f
		}
	}

	return nil
}

func (this *SCS) GetEntryByPath(path string) (HashfsEntry, bool) {
	entry, ok := this.entryMap[hashed(path)]

	return entry, ok
}

func hashed(path string) uint64 {
	return city.CH64([]byte(path))
}

func (this *SCS) readDir(entry HashfsEntry) HashfsDir {

	dir := HashfsDir{
		hash: entry.hash,
		data: make([]string, 0),
	}

	content := this.buffer.buffer(entry.offset, entry.size)

	for {
		line, hasNext := content.nextLine()
		dir.data = append(dir.data, line)
		if !hasNext {
			break
		}
	}

	dir.children = make([]HashfsNode, len(dir.data))

	return dir
}

func (this *SCS) readFile(entry HashfsEntry) HashfsFile {
	var content *Buffer
	if entry.flags.compressed {
		content = this.buffer.buffer(entry.offset, entry.compressedSize).inflate()
	} else {
		content = this.buffer.buffer(entry.offset, entry.size)
	}

	file := HashfsFile{hash: entry.hash, textual: content.textual()}
	file.data = content.bytes
	if file.textual {
		file.text = string(content.bytes)
	}

	return file
}

func (this *SCS) Dump(base string) {

	os.MkdirAll(base, os.ModeDir)

	for _, entry := range this.entries {
		if !entry.flags.dir {
			f := this.readFile(entry)
			suffix := ""
			if f.textual && strings.Contains(f.text, "SiiNunit") {
				suffix = ".sii"
			}
			os.WriteFile(base+"/"+strconv.FormatUint(entry.hash, 10)+suffix, f.data, 0644)
		}
	}

}

func (this *SCS) searchPaths(r *regexp.Regexp, handler func(groups []string, set mapset.Set[string])) []string {
	set := mapset.NewSet[string]()
	for _, entry := range this.entries {
		if entry.flags.dir {
			continue
		}

		f := this.readFile(entry)

		if r.Match(f.data) {
			matches := r.FindAllSubmatch(f.data, -1)
			for _, groups := range toStrings(matches) {
				handler(groups, set)
			}
		}
	}

	return set.ToSlice()
}

var extracted = mapset.NewSet[uint64]()

func (this *SCS) TryExtract(base string, additionalPaths ...string) {
	// well-known roots
	paths := wellKnownPaths

	// any path in binary or textual files
	r := regexp.MustCompile(`/([\x20-\x21\x23-\x7E]+(\.(pmd|pmg|pmc|pma|ppd|tobj|mat|soundref|sii|sui|dds|tga|png)))`)
	paths = append(paths, this.searchPaths(r, func(groups []string, set mapset.Set[string]) {
		path := groups[1]
		set.Add(path)
	})...)

	// icons in some sii files which is a relative path without extension pointing to material/
	// example: icon: "hudgps_icon"
	r = regexp.MustCompile(`icon\s?:\s?"([\x20-\x21\x23-\x7E]+(\.(tobj|mat|dds|tga|png))?)"`)
	paths = append(paths, this.searchPaths(r, func(groups []string, set mapset.Set[string]) {
		path := groups[1]
		if strings.HasPrefix(path, "/") {
			set.Add(path)
			return
		}

		extension := ""
		if len(groups) >= 3 {
			extension = groups[2]
		}
		set.Add("material/ui/accessory/" + path + defaults(extension, ".mat"))
		set.Add("material/ui/accessory/" + path + defaults(extension, ".tobj"))
		set.Add("material/ui/accessory/" + path + defaults(extension, ".dds"))
		set.Add("material/ui/accessory/" + path + defaults(extension, ".png"))
	})...)

	// texture in .mat file which points to material/ which are not a full qualified path
	// example: texture : "hudgps_icon.tobj"
	r = regexp.MustCompile(`texture\s?:\s?"([\x20-\x21\x23-\x7E]+\.tobj)"`)
	paths = append(paths, this.searchPaths(r, func(groups []string, set mapset.Set[string]) {
		path := groups[1]
		if strings.HasPrefix(path, "/") {
			set.Add(path)
			return
		}
		set.Add("material/ui/accessory/" + path)
	})...)

	paths = append(paths, additionalPaths...)

	extracted.Clear()

	for _, path := range paths {
		tree := this.readTree(path)

		if tree != nil {
			tree.Save(base)
		}
	}

	fmt.Printf("\n[hashfs] %d mounted, %d extracted\n", this.header.entriesCount, len(extracted.ToSlice()))
}

func (this *SCS) init() {

	this.header = NewHeader(this.buffer.nextBuffer(32))
	if this.header.magic != "SCS#" {
		panic("Not a SCS hashfs format, may be a ZIP and use unzip instead")
	}

	this.entries = make([]HashfsEntry, this.header.entriesCount)
	this.entryMap = make(map[uint64]HashfsEntry)

	for i := range this.entries {
		entry := NewEntry(this.buffer.nextBuffer(32))
		this.entries[i] = entry
		this.entryMap[entry.hash] = entry
	}
}

type HashfsDir struct {
	name string
	path string
	hash uint64
	data []string

	children []HashfsNode
}

type HashfsFile struct {
	hash    uint64
	name    string
	path    string
	textual bool
	data    []byte
	text    string
}

func (this HashfsFile) Hash() uint64 {
	return this.hash
}

func (this HashfsFile) Save(base string) {

	if !extracted.Contains(this.hash) {
		fmt.Printf("📜 %s \n", this.path)
		extracted.Add(this.hash)
	}

	path := base + "/" + this.path
	os.MkdirAll(filepath.Dir(path), os.ModeDir)
	os.WriteFile(path, this.data, 0644)
}

func (this HashfsDir) Hash() uint64 {
	return this.hash
}

func (this HashfsDir) Save(base string) {
	if !extracted.Contains(this.hash) {
		fmt.Printf("📁 %s \n", this.path)
		extracted.Add(this.hash)
	}

	os.MkdirAll(base+"/"+this.path, os.ModeDir)
	for _, child := range this.children {
		if dir, ok := child.(HashfsDir); ok {
			dir.Save(base)
		}
		if f, ok := child.(HashfsFile); ok {
			f.Save(base)
		}
	}
}

type HashfsNode interface {
	Hash() uint64
	Save(base string)
}

type SCS struct {
	header   HashfsHeader
	entries  []HashfsEntry
	entryMap map[uint64]HashfsEntry
	buffer   *Buffer
}
