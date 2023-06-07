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

func (this *SCS) ReadTree(path string) HashfsNode {
	if entry, ok := this.GetEntryByPath(path); ok {
		if entry.flags.dir {
			dir := this.readDir(entry)
			dir.name = filepath.Base(path)
			dir.path = path
			for i, name := range dir.data {
				if strings.HasPrefix(name, "*") {
					p := path + "/" + name[1:]
					child := this.ReadTree(p).(HashfsDir)
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
		content = this.buffer.buffer(entry.offset, entry.compressedSize).deflate()
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

// match text in SII files
func (this *SCS) match(r *regexp.Regexp, group int, icon bool) []string {
	set := mapset.NewSet[string]()

	for _, entry := range this.entries {
		if !entry.flags.dir {
			f := this.readFile(entry)
			text := strings.ReplaceAll(f.text, "\n", "")
			if f.textual { //&& strings.Contains(text, "SiiNunit") {
				if r.MatchString(text) {
					arr0 := r.FindAllStringSubmatch(text, -1)
					for _, arr1 := range arr0 {
						if icon {
							set.Add("material/ui/accessory/" + arr1[group] + ".tga")
							set.Add("material/ui/accessory/" + arr1[group] + ".mat")
							set.Add("material/ui/accessory/" + arr1[group] + ".tobj")
						} else {
							path := arr1[group]
							dir, name, extension := Unjoin(path)
							// model files
							if len(extension) == 4 && strings.HasPrefix(extension, ".pm") {
								set.Add(Join(dir, name, extension))
								set.Add(Join(dir, name, ".pmd"))
								set.Add(Join(dir, name, ".pmc"))
								set.Add(Join(dir, name, ".pmg"))
								set.Add(Join(dir, name, ".pma"))
							}
							// intermediate files
							if len(extension) == 4 && strings.HasPrefix(extension, ".pi") {
								set.Add(Join(dir, name, extension))
								set.Add(Join(dir, name, ".pit"))
								set.Add(Join(dir, name, ".pim"))
								set.Add(Join(dir, name, ".pic"))
								set.Add(Join(dir, name, ".pia"))
								set.Add(Join(dir, name, ".pip"))
							}
							// texture/material files
							if strings.HasPrefix(extension, ".tobj") ||
								strings.HasPrefix(extension, ".mat") {
								set.Add(Join(dir, name, extension))
								set.Add(Join(dir, name, ".mat"))
								set.Add(Join(dir, name, ".tobj"))
								set.Add(Join(dir, name, ".dds"))
								set.Add(Join(dir, name, ".tga"))
								set.Add(Join(dir, name, ".png"))
							}

						}

					}
				}
			}
		}
	}

	return set.ToSlice()
}

var extracted = mapset.NewSet[uint64]()

func (this *SCS) TryExtract(base string) {
	// well-known roots
	paths := wellKnownPaths

	// pmg, pmd...
	r := regexp.MustCompile("\"/([^\"]+)\"")
	paths = append(paths, this.match(r, 1, false)...)

	// icons in material/
	r = regexp.MustCompile("icon: \"([^\"]+)\"")
	paths = append(paths, this.match(r, 1, true)...)

	extracted.Clear()

	for _, path := range paths {
		tree := this.ReadTree(path)

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
