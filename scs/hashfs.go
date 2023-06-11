package scs

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/chaoyangnz/scs-extract/structs"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Node struct {
	hash  uint64
	isDir bool
	dir   []string
	file  []byte
	path  string
}

type Hashfs struct {
	nodes     []Node
	nodesMap  map[uint64]Node
	extracted mapset.Set[string]
}

func NewHashfs(scsFile string) *Hashfs {
	scs := structs.NewScs()

	b, ok := os.ReadFile(scsFile)
	if ok != nil {
		panic("Unable to read " + scsFile)
	}
	scs.Read(kaitai.NewStream(bytes.NewReader(b)), nil, scs)

	nodes := make([]Node, scs.EntryCount)
	nodesMap := make(map[uint64]Node)

	for i, entry := range scs.Entries {
		isDir, _ := entry.IsDir()
		isCompressed, _ := entry.IsCompressed()
		node := Node{
			hash:  entry.Hash,
			isDir: isDir,
		}
		if isDir {
			content, _ := entry.Dir()
			node.dir = strings.Split(content, "\n")
		} else {
			if isCompressed {
				content, _ := entry.File()
				reader, _ := zlib.NewReader(bytes.NewReader(content))
				b, _ := ioutil.ReadAll(reader)
				node.file = b
			}
		}
		nodes[i] = node
		nodesMap[entry.Hash] = node
	}

	extracted := mapset.NewSet[string]()

	return &Hashfs{nodes: nodes, nodesMap: nodesMap, extracted: extracted}
}

func (this *Hashfs) traverse(path string) {
	if node, ok := this.NodeByPath(path); ok {
		node.path = path
		this.extracted.Add(path)

		if node.isDir {
			for _, name := range node.dir {
				if strings.HasPrefix(name, "*") {
					name = name[1:]
				}
				this.traverse(node.path + "/" + name)
			}
		}
	}
}

func (this *Hashfs) Match(r *regexp.Regexp, handler func(match []string) []string, paths mapset.Set[string]) {
	for _, node := range this.nodes {
		if node.isDir {
			continue
		}

		if r.Match(node.file) {
			matches := r.FindAllSubmatch(node.file, -1)
			for _, match := range toStrings(matches) {
				paths.Append(handler(match)...)
			}
		}
	}
}

func (this *Hashfs) NodeByPath(path string) (Node, bool) {
	hash := hashed(path)
	if node, ok := this.nodesMap[hash]; ok {
		return node, true
	}
	return Node{}, false
}

func (this *Hashfs) seed() []string {
	// match
	var paths = mapset.NewSet[string]()

	r := regexp.MustCompile(`/([\x20-\x7E]+(\.(pmd|pmg|pmc|pma|ppd|tobj|mat|soundref|sii|sui|dds|tga|png)))`)
	this.Match(r, func(match []string) []string {
		return []string{match[1]}
	}, paths)

	r = regexp.MustCompile(`icon\s?:\s?"([\x20-\x21\x23-\x7E]+(\.(tobj|mat|dds|tga|png))?)"`)
	this.Match(r, func(match []string) []string {
		return []string{"material/ui/accessory/" + match[1] + defaults(match[2], ".mat")}
	}, paths)

	r = regexp.MustCompile(`texture\s?:\s?"([\x20-\x21\x23-\x7E]+\.tobj)"`)
	this.Match(r, func(match []string) []string {
		return []string{"material/ui/accessory/" + match[1]}
	}, paths)

	paths.Append(wellKnownPaths...)

	//fmt.Printf("%+v \n", paths.ToSlice())

	// seeds - all existing
	//var seeds []string
	//for _, path := range paths.ToSlice() {
	//	if node, ok := this.NodeByPath(path); ok {
	//		node.path = path
	//		seeds = append(seeds, path)
	//	}
	//}

	return paths.ToSlice()
}

func (this *Hashfs) Extract(dest string) {

	seeds := this.seed()

	for _, path := range seeds {
		this.traverse(path)
	}

	for _, path := range this.extracted.ToSlice() {
		node, _ := this.NodeByPath(path)
		if node.isDir {
			os.MkdirAll(filepath.Join(dest, node.path), os.ModeDir)
		} else {
			os.MkdirAll(filepath.Join(dest, filepath.Dir(path)), os.ModeDir)
			os.WriteFile(filepath.Join(dest, path), node.file, 0644)
		}

		icon := "📜"
		if node.isDir {
			icon = "📁"
		}
		fmt.Printf("%s %s \n", icon, path)
	}

	fmt.Printf("%d mounted, %d extracted \n", len(this.nodes), len(this.extracted.ToSlice()))
}
