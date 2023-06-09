package scs

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

type Buffer struct {
	bytes    []uint8
	readable []uint8
}

func NewBuffer(bytes []uint8) *Buffer {
	return &Buffer{
		bytes:    bytes,
		readable: bytes,
	}
}

func (this *Buffer) next8() uint64 {
	value := binary.LittleEndian.Uint64(this.readable[:8])
	this.readable = this.readable[8:]
	return value
}

func (this *Buffer) next4() uint32 {
	value := binary.LittleEndian.Uint32(this.readable[:4])
	this.readable = this.readable[4:]
	return value
}

func (this *Buffer) next2() uint16 {
	value := binary.LittleEndian.Uint16(this.readable[:2])
	this.readable = this.readable[2:]
	return value
}

func (this *Buffer) next() uint8 {
	return this.nextN(1)[0]
}

func (this *Buffer) nextN(length uint32) []uint8 {
	bytes := this.readable[:length]
	this.readable = this.readable[length:]
	return bytes
}

func (this *Buffer) nextString(length uint32) string {
	bytes := this.nextN(length)
	return string(bytes)
}

func (this *Buffer) nextBuffer(length uint32) *Buffer {
	bytes := this.readable[:length]
	this.readable = this.readable[length:]
	return NewBuffer(bytes)
}

func (this *Buffer) buffer(offset uint64, size uint32) *Buffer {
	bytes := this.bytes[offset : offset+uint64(size)]
	return NewBuffer(bytes)
}

func (this *Buffer) nextLine() (string, bool) {
	for i := uint32(0); i < this.length(); i++ {
		if this.readable[i] == '\n' {
			str := this.nextString(i)
			this.next()
			return str, this.length() > 0
		}
	}
	return string(this.readable), false
}

func (this *Buffer) deflate() *Buffer {
	z, err := zlib.NewReader(bytes.NewReader(this.bytes))
	if err != nil {
		return &Buffer{}
	}
	b, err := ioutil.ReadAll(z)
	if err != nil {
		return &Buffer{}
	}
	return NewBuffer(b)
}

func (this *Buffer) textual() bool {
	for i := uint32(0); i < this.length(); i++ {
		if this.bytes[i] > 127 {
			return false
		}
	}
	return true
}

func (this *Buffer) length() uint32 {
	return uint32(len(this.readable))
}

func (this *Buffer) size() uint64 {
	return uint64(len(this.bytes))
}

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
