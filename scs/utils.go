package scs

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-faster/city"
	"github.com/samber/lo"
	"path"
	"path/filepath"
	"strings"
)

func Unjoin(path string) (string, string, string) {
	//dir := filepath.Dir(path)
	fileName := filepath.Base(path)
	dir := strings.TrimSuffix(path, fileName)
	extension := filepath.Ext(fileName)
	name := strings.TrimSuffix(fileName, extension)

	return dir, name, extension
}

func Join(dir string, name string, extension string) string {
	return path.Join(dir, name+extension)
}

func toStrings(arr [][][]byte) [][]string {
	var lists [][]string
	for _, stringList := range arr {
		var list []string
		for _, str := range stringList {
			list = append(list, string(str))
		}
		lists = append(lists, list)
	}

	return lists
}

func NormalizePath(path string) string {
	path = lo.Ternary(strings.HasPrefix(path, "/"), path[1:], path)
	path = lo.Ternary(strings.HasSuffix(path, "/"), path[:len(path)-1], path)
	return path
}

func NormalizePaths(paths ...string) []string {
	return lo.Map[string, string](paths, func(path string, index int) string {
		return NormalizePath(path)
	})
}

func defaults(a string, b string) string {
	if a == "" {
		return b
	}

	return a
}

// path must be without leading /
func cityhash(path string) uint64 {
	return city.CH64([]byte(path))
}

func backtrackPaths(path string, paths mapset.Set[string]) {
	if path == "" {
		return
	}
	base := filepath.Base(path)
	dir := strings.TrimSuffix(path, base)
	if strings.HasSuffix(dir, "/") {
		dir = strings.TrimSuffix(dir, "/")
	}
	paths.Add(dir)
	backtrackPaths(dir, paths)
}
