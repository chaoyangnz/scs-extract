package scs

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/chaoyangnz/scs-extract/structs"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"github.com/samber/lo"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
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
			names := strings.Split(content, "\n")
			for i, name := range names {
				if strings.HasPrefix(name, "*") {
					names[i] = name[1:]
				}
			}
			node.dir = names
		} else {
			if isCompressed {
				content, _ := entry.File()
				reader, _ := zlib.NewReader(bytes.NewReader(content))
				b, _ := io.ReadAll(reader)
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
			lo.ForEach(node.dir, func(name string, index int) {
				this.traverse(node.path + "/" + name)
			})
		}
	}
}

func (this *Hashfs) match(r *regexp.Regexp, handler func(m []string) []string, paths mapset.Set[string]) {
	for _, node := range this.nodes {
		if node.isDir || !r.Match(node.file) {
			continue
		}

		matches := r.FindAllSubmatch(node.file, -1)
		for _, match := range toStrings(matches) {
			paths.Append(handler(match)...)
		}

	}
}

func (this *Hashfs) NodeByPath(path string) (Node, bool) {
	hash := cityhash(path)
	if node, ok := this.nodesMap[hash]; ok {
		return node, true
	}
	return Node{}, false
}

func (this *Hashfs) seed() mapset.Set[string] {
	// match
	var paths = mapset.NewSet[string]()

	r := regexp.MustCompile(`/([\x20-\x7E]+(\.(pmd|pmg|pmc|pma|ppd|tobj|mat|soundref|sii|sui|dds|tga|png)))`)
	this.match(r, func(match []string) []string {
		return []string{match[1]}
	}, paths)

	r = regexp.MustCompile(`icon\s?:\s?"([\x20-\x21\x23-\x7E]+(\.(tobj|mat|dds|tga|png))?)"`)
	this.match(r, func(match []string) []string {
		return []string{"material/ui/accessory/" + match[1] + defaults(match[2], ".mat")}
	}, paths)

	r = regexp.MustCompile(`texture\s?:\s?"([\x20-\x21\x23-\x7E]+\.tobj)"`)
	this.match(r, func(match []string) []string {
		return []string{"material/ui/accessory/" + match[1]}
	}, paths)

	lo.ForEach(paths.ToSlice(), func(path string, _ int) {
		backtrackPaths(path, paths)
	})

	paths.Append(wellKnownPaths...)

	return paths
}

func (this *Hashfs) Dump(dest string) {
	os.MkdirAll(dest, os.ModeDir)

	for _, node := range this.nodes {
		if node.isDir {
			contents := []byte(strings.Join(node.dir, "\n"))
			os.WriteFile(filepath.Join(dest, "folder_"+strconv.FormatUint(node.hash, 10)), contents, 0644)
		} else {
			contents := node.file
			r := regexp.MustCompile("SiiNunit")
			suffix := lo.Ternary(r.Match(contents), ".sii", "")
			os.WriteFile(filepath.Join(dest, strconv.FormatUint(node.hash, 10)+suffix), contents, 0644)
		}
	}
}

func (this *Hashfs) Extract(dest string, additionalPaths ...string) {

	seeds := this.seed()
	seeds.Append(additionalPaths...)

	for _, path := range seeds.ToSlice() {
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

		icon := lo.Ternary(node.isDir, "📁", "📜")
		fmt.Printf("%s %s \n", icon, path)
	}

	fmt.Printf("\n%d mounted, %d extracted \n", len(this.nodes), len(this.extracted.ToSlice()))
}
