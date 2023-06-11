package scs

import (
	"github.com/go-faster/city"
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

func NormalizePaths(strList []string) []string {
	var list []string
	for _, str := range strList {
		if strings.HasPrefix(str, "/") {
			list = append(list, str[1:])
		} else {
			list = append(list, str)
		}
	}

	return list
}

func defaults(a string, b string) string {
	if a == "" {
		return b
	}

	return a
}

func hashed(path string) uint64 {
	return city.CH64([]byte(path))
}
