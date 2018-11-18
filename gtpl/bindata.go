// Code generated by go-bindata. DO NOT EDIT.
// sources:
// response/comment.md (21B)
// response/greeting.md (6B)
// response/improvement.md (33B)
// response/newcomer_greeting.md (66B)
// response/praise.md (15B)
// response/questions.md (79B)
// response/todo.md (73B)
// tools/compile.md (229B)
// tools/not_formatted.md (354B)
// tools/not_linted.md (523B)
// tools/pass_tests.md (497B)
// topic/benchmarking.md (1.635kB)
// topic/pprof-allocations.md (1.9kB)
// topic/regex.md (1.209kB)

package gtpl

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _responseCommentMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe2\x52\x2d\x56\xc8\xcf\x53\x28\xc9\xc8\x2c\x56\x48\xad\x48\x2d\x4a\xce\x2c\x4e\xb5\x02\x04\x00\x00\xff\xff\x8f\x03\x06\xcf\x15\x00\x00\x00")

func responseCommentMdBytes() ([]byte, error) {
	return bindataRead(
		_responseCommentMd,
		"response/comment.md",
	)
}

func responseCommentMd() (*asset, error) {
	bytes, err := responseCommentMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "response/comment.md", size: 21, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x95, 0xff, 0xaf, 0xf3, 0x2f, 0x3f, 0xad, 0x81, 0x55, 0xf4, 0xd, 0x9a, 0x98, 0x2d, 0x83, 0xef, 0x3e, 0xf5, 0x1c, 0x7a, 0xe8, 0x9, 0x36, 0x2e, 0xab, 0xeb, 0x7e, 0x6a, 0x9, 0xbe, 0x23, 0x4f}}
	return a, nil
}

var _responseGreetingMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xf2\xc8\x54\x50\x2d\x56\x04\x04\x00\x00\xff\xff\xc0\x2d\x24\x10\x06\x00\x00\x00")

func responseGreetingMdBytes() ([]byte, error) {
	return bindataRead(
		_responseGreetingMd,
		"response/greeting.md",
	)
}

func responseGreetingMd() (*asset, error) {
	bytes, err := responseGreetingMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "response/greeting.md", size: 6, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb6, 0x4d, 0xb8, 0x95, 0xfc, 0x5c, 0x3e, 0x1, 0x38, 0xee, 0x81, 0xa5, 0xe7, 0x8, 0xfb, 0xde, 0xf5, 0x17, 0x32, 0x9d, 0xa, 0xbd, 0x7f, 0xc2, 0xa0, 0x1d, 0x87, 0xea, 0xd1, 0xb, 0x32, 0xdb}}
	return a, nil
}

var _responseImprovementMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe2\xf2\x48\x2d\x4a\x55\x50\x2d\x56\x48\xcb\x2f\x52\x48\x2b\x2d\x2a\xc9\x48\x2d\x52\xc8\xcc\x2d\x28\xca\x2f\x4b\xcd\x4d\xcd\x2b\xb1\x02\x04\x00\x00\xff\xff\x8b\x2d\xa2\x17\x21\x00\x00\x00")

func responseImprovementMdBytes() ([]byte, error) {
	return bindataRead(
		_responseImprovementMd,
		"response/improvement.md",
	)
}

func responseImprovementMd() (*asset, error) {
	bytes, err := responseImprovementMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "response/improvement.md", size: 33, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x24, 0x47, 0x5, 0x78, 0x1d, 0x7e, 0x8f, 0xa1, 0x3e, 0x6, 0x5b, 0x34, 0xee, 0x69, 0x7b, 0x7, 0xb3, 0xb8, 0xa1, 0xc5, 0xfc, 0xf7, 0x39, 0xa2, 0x81, 0xe6, 0x2b, 0x55, 0x94, 0xe8, 0x66, 0xcb}}
	return a, nil
}

var _responseNewcomer_greetingMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe2\x0a\x4f\xcd\x49\xce\xcf\x4d\x55\x28\xc9\x57\x70\xad\x48\x2d\x4a\xce\x2c\xce\x55\x48\xcc\x4b\x51\x28\x47\x88\xe7\xa4\x26\x16\xe5\x65\xe6\xa5\x2b\xb8\xe7\x2b\x2a\xf8\x65\x26\x83\x05\x73\x53\x53\x4b\x14\x2a\xf3\x4b\x15\x01\x01\x00\x00\xff\xff\xd0\x93\xfd\x53\x42\x00\x00\x00")

func responseNewcomer_greetingMdBytes() ([]byte, error) {
	return bindataRead(
		_responseNewcomer_greetingMd,
		"response/newcomer_greeting.md",
	)
}

func responseNewcomer_greetingMd() (*asset, error) {
	bytes, err := responseNewcomer_greetingMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "response/newcomer_greeting.md", size: 66, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb3, 0xa7, 0x40, 0x9e, 0xd3, 0xbd, 0x95, 0x71, 0x56, 0x20, 0xff, 0xff, 0xf4, 0xe5, 0xa9, 0x84, 0x38, 0x78, 0xa2, 0xf, 0xd6, 0x63, 0xe2, 0xec, 0xee, 0xdb, 0x80, 0x3c, 0xf1, 0x3b, 0x59, 0xa1}}
	return a, nil
}

var _responsePraiseMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe2\x0a\xc9\xc8\x2c\x56\xc8\xc9\xcf\xcf\x2e\x56\x50\x2d\x56\x04\x04\x00\x00\xff\xff\x8c\xda\x01\x83\x0f\x00\x00\x00")

func responsePraiseMdBytes() ([]byte, error) {
	return bindataRead(
		_responsePraiseMd,
		"response/praise.md",
	)
}

func responsePraiseMd() (*asset, error) {
	bytes, err := responsePraiseMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "response/praise.md", size: 15, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x59, 0xcb, 0x83, 0x60, 0x69, 0x93, 0xba, 0xce, 0x32, 0xd9, 0xbc, 0xd8, 0x1b, 0x64, 0x86, 0xeb, 0x1c, 0xba, 0x31, 0x65, 0x69, 0x34, 0x6d, 0xfc, 0xbb, 0xf3, 0x4c, 0x12, 0x6d, 0x4e, 0xf, 0x23}}
	return a, nil
}

var _responseQuestionsMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xcb\x4b\x0a\x80\x20\x00\x04\xd0\x7d\xa7\x98\x56\xed\x3a\x46\xd0\x31\x84\x46\x14\xcd\x29\xbf\x74\xfb\xa8\xfd\x7b\xd3\x46\x46\xd8\x4c\xa2\x0a\xa6\x04\x78\x8b\x47\x0d\xce\x74\xe2\x6e\x2c\xd5\x2b\x15\x28\x63\x98\x54\x3f\x14\x92\x06\x4e\x65\xae\xd8\x97\x03\x51\xfd\xbf\x8e\xf1\x9a\xdf\x00\x00\x00\xff\xff\x93\x42\xc4\x70\x4f\x00\x00\x00")

func responseQuestionsMdBytes() ([]byte, error) {
	return bindataRead(
		_responseQuestionsMd,
		"response/questions.md",
	)
}

func responseQuestionsMd() (*asset, error) {
	bytes, err := responseQuestionsMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "response/questions.md", size: 79, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa2, 0x9d, 0x81, 0x1f, 0x6c, 0x35, 0xc9, 0x79, 0xe6, 0xd6, 0xd0, 0xb7, 0x8a, 0x8d, 0x9e, 0xeb, 0x21, 0x6, 0x3e, 0x9b, 0x33, 0x60, 0xbe, 0xda, 0x27, 0xd2, 0xc, 0x5b, 0x59, 0x38, 0x86, 0x8b}}
	return a, nil
}

var _responseTodoMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xca\xbb\x0d\x84\x30\x10\x45\xd1\x7c\xab\xb8\xc9\xd6\x41\x40\x44\x19\x23\x30\xd8\x62\xe4\x67\xe1\x0f\xa2\x7b\x24\xf2\xf3\x9b\xd5\x7d\xe3\x51\x27\xda\x08\x18\x2e\x9d\x58\xa3\xc5\xc0\x2e\x77\xdd\x29\x1f\xfc\x2b\x55\x2c\xac\x96\xb1\x52\x2e\x8d\xf0\x89\x2a\xef\x2d\x29\x4f\x6f\x00\x00\x00\xff\xff\xf1\x50\x3a\x1d\x49\x00\x00\x00")

func responseTodoMdBytes() ([]byte, error) {
	return bindataRead(
		_responseTodoMd,
		"response/todo.md",
	)
}

func responseTodoMd() (*asset, error) {
	bytes, err := responseTodoMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "response/todo.md", size: 73, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa9, 0x52, 0x39, 0xc9, 0xcb, 0x75, 0x93, 0x60, 0xe8, 0x21, 0x22, 0x7b, 0xc0, 0x66, 0xed, 0x94, 0x3f, 0xd1, 0xef, 0x2, 0x3e, 0xf8, 0xf1, 0xfd, 0x19, 0x67, 0x7c, 0xdc, 0x7b, 0x4e, 0x6b, 0x7a}}
	return a, nil
}

var _toolsCompileMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8e\x31\x4e\xc5\x40\x0c\x44\x7b\x4e\x31\x54\x34\x1f\xae\xf0\x1b\x1a\x0a\x4a\x0e\xe0\x6c\x9c\x64\x85\x63\x47\xde\x59\xd0\xde\x1e\x25\x12\xbf\x1c\x8d\x66\xde\x7b\xc5\x97\x2f\x91\xec\x2e\x54\x1b\x37\x70\xab\x0d\x2d\xac\xb3\x86\x63\x0e\x6d\xf0\x20\x4a\xec\x47\x35\x45\x89\x4c\x2d\xb4\xf1\x86\xf7\x3a\x63\x44\x47\x76\x07\x37\x05\xb5\xb1\xdd\xf1\x29\x63\xd2\xab\x90\x52\xea\xac\x4e\x31\x1b\x68\x7d\xda\x2b\xa9\x33\xc4\x51\xfd\xfc\x33\xa5\x3e\x50\x77\x7c\x2c\xe7\xea\x25\x15\x9b\xfc\x54\x5f\xc1\x8c\x3e\x99\x62\x55\xf2\xcc\x23\x7a\xe2\xc8\x58\x53\x76\x30\xfe\x9d\x6e\x17\xac\x88\x43\xec\x57\x46\x83\xb4\x6f\xec\x8a\x25\x12\x9b\xda\xf1\xfc\xf4\x17\x00\x00\xff\xff\x15\x59\x6b\x94\xe5\x00\x00\x00")

func toolsCompileMdBytes() ([]byte, error) {
	return bindataRead(
		_toolsCompileMd,
		"tools/compile.md",
	)
}

func toolsCompileMd() (*asset, error) {
	bytes, err := toolsCompileMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tools/compile.md", size: 229, mode: os.FileMode(420), modTime: time.Unix(1542628451, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x21, 0x51, 0x8f, 0x7e, 0xbe, 0x46, 0x22, 0xaf, 0x24, 0x5c, 0xe6, 0xb5, 0x90, 0x7e, 0x57, 0x7d, 0xf5, 0x14, 0xc6, 0x77, 0x88, 0x6e, 0x23, 0x80, 0xf, 0x3c, 0x87, 0x9c, 0x9, 0x10, 0xe9, 0x6f}}
	return a, nil
}

var _toolsNot_formattedMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8f\xb1\x4e\x2b\x31\x10\x45\xfb\x7c\xc5\x2d\xdf\x2b\x92\xed\xd3\x02\x42\xd4\x20\x24\x14\x21\x65\xb2\x9e\xb5\x47\xb1\x3d\x91\x67\x1c\xb1\x7f\x8f\xbc\x14\xe9\xa6\xb9\xe7\x9c\xd9\xe3\x49\x03\x63\xd1\x56\xc8\x5d\x6a\x84\x54\xbc\x2a\xc4\x70\xe7\xb6\x42\xca\x4d\x9b\x53\x75\xb8\xc2\x13\x63\xd6\x52\x7a\x15\x5f\x0f\xf8\xa0\x2b\x83\x90\x55\xaf\x20\xc7\x39\xea\x52\xfc\x0c\xaa\x01\xad\x57\x88\x43\x2b\x56\xed\x0d\xb3\x06\x3e\xe0\x6b\x9c\x1c\xc4\xb5\x41\x1b\xde\x9e\x5f\x60\x49\x7b\x0e\xb8\x30\xe8\x92\x79\x38\xc2\xd0\x88\x8d\xa2\xb1\x05\x75\xd7\x42\x2e\x33\xe5\xbc\x1e\x77\x3b\x00\xd8\xe3\x53\xac\x53\xc6\xbb\xf7\x20\xfa\xf7\x42\x22\xc3\x29\x36\x26\x1f\xfd\xd6\x6f\x23\xfc\xfb\x5f\x72\xbf\xd9\x71\x9a\xb6\x84\xfb\x36\xb3\x6d\x75\x98\xb5\x4c\x41\x67\x9b\x32\xd5\xd8\x29\xb2\x4d\x51\xff\x3f\x0c\x05\x94\x4d\x37\x2e\xe1\xb4\x50\x75\x32\x97\x79\xd0\xf9\xc7\xb9\x9a\x68\x7d\xf0\xa3\x78\xea\x97\x0d\xba\x90\x4b\x9a\xee\x52\xf6\x83\xf7\x1b\x00\x00\xff\xff\x25\x14\x54\x18\x62\x01\x00\x00")

func toolsNot_formattedMdBytes() ([]byte, error) {
	return bindataRead(
		_toolsNot_formattedMd,
		"tools/not_formatted.md",
	)
}

func toolsNot_formattedMd() (*asset, error) {
	bytes, err := toolsNot_formattedMdBytes()
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	info := bindataFileInfo{name: "tools/not_formatted.md", size: 354, mode: os.FileMode(420), modTime: time.Unix(1542628451, 0)}
=======
	info := bindataFileInfo{name: "tools/not_formatted.md", size: 354, mode: os.FileMode(420), modTime: time.Unix(1542540514, 0)}
>>>>>>> More tweaks
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xea, 0x5c, 0xf5, 0x28, 0x4d, 0x86, 0x30, 0x8, 0x28, 0x99, 0xc0, 0x75, 0xb9, 0x45, 0x20, 0xe3, 0x34, 0x28, 0x5a, 0x3b, 0xc2, 0x2e, 0x5c, 0xcf, 0x3d, 0xfc, 0x4a, 0x35, 0xb1, 0xf3, 0x4a, 0x6a}}
	return a, nil
}

var _toolsNot_lintedMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x91\x31\xaf\x13\x31\x10\x84\xfb\xfc\x8a\xa1\xca\x8b\x04\x49\x0f\x15\x15\x7f\xe0\x49\x14\x08\x29\xbe\xf3\xda\x5e\x3d\xdb\x7b\xf2\xae\x39\xfc\xef\x91\x2f\x40\x5e\xe7\x62\x3c\xf3\xcd\xec\x27\xbc\x26\x6a\x74\x56\x38\xc4\x46\xce\x60\x22\x19\xab\xcb\x99\x3c\x7e\xdc\xa3\x64\xae\x76\xff\xf9\x92\xcc\x36\xfd\x7c\xbb\x45\xb6\xd4\x97\xeb\x2a\xe5\x16\x25\xbb\x1a\x6f\x53\x70\xc1\x9e\x78\x4d\xd8\x39\x67\xd0\x6f\x57\xb8\x12\x86\xf4\x86\x55\x3c\x21\xc8\x7c\x94\x22\x15\x5b\x93\x25\x53\x51\xb8\xea\xa1\x36\x32\x81\x55\x3b\xe9\x15\xaf\x6d\xa0\xf5\x5a\xb9\x46\xfc\x0b\x86\xd4\xa7\xcf\x17\xb0\x3d\x22\x8a\x7b\x23\xa8\x14\x42\x57\x0a\x3d\x43\x7b\x8c\xa4\xc6\x52\xf5\x48\x1b\xd2\xaf\xa7\x13\x80\xc3\xd5\x04\x91\x0c\x5c\x4d\x60\x89\x90\xdc\xc2\x06\x09\x70\x79\x77\x43\xb1\x26\x5a\xdf\x66\xec\x13\x79\x67\x4b\x4f\x8a\x97\xa3\x40\x0d\x1c\x7b\xfb\xaf\x23\xcf\x26\x6d\x9a\x7b\x99\x64\x7f\x73\x2f\x57\x7c\x4f\x74\x60\x9f\x1b\x61\x6f\x6c\xf3\xcb\x37\x81\x4a\xb0\xdd\xb5\xc7\x1e\x5b\x13\xdf\xd7\x49\xfc\x11\xc5\xd5\x81\xa5\x73\xf6\xd8\x78\xa3\xcc\x95\xf4\x51\x34\x38\xce\xc7\x74\x6c\x0a\xd7\x4d\x8a\x33\x9e\xd7\x19\xe0\x30\xbb\x0c\x78\xa9\x67\xc3\xe6\x54\xdf\xf1\xce\x75\xef\x51\x42\xb1\xfb\xe5\x8a\xaf\xbf\x84\x3d\x2c\xb1\x62\x19\x98\x92\x49\x74\xf4\x9c\x4d\x94\x72\xf8\x70\xfa\x13\x00\x00\xff\xff\x0a\xea\x2c\x5e\x0b\x02\x00\x00")

func toolsNot_lintedMdBytes() ([]byte, error) {
	return bindataRead(
		_toolsNot_lintedMd,
		"tools/not_linted.md",
	)
}

func toolsNot_lintedMd() (*asset, error) {
	bytes, err := toolsNot_lintedMdBytes()
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	info := bindataFileInfo{name: "tools/not_linted.md", size: 523, mode: os.FileMode(420), modTime: time.Unix(1542628451, 0)}
=======
	info := bindataFileInfo{name: "tools/not_linted.md", size: 523, mode: os.FileMode(420), modTime: time.Unix(1542540514, 0)}
>>>>>>> More tweaks
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5a, 0xc8, 0xf2, 0x8a, 0x23, 0x4f, 0xa3, 0x6d, 0xed, 0x67, 0x8a, 0xdb, 0x78, 0x68, 0x80, 0xc1, 0x58, 0xf3, 0xc6, 0xdb, 0x2a, 0x71, 0x35, 0xae, 0x86, 0x86, 0x7, 0x8, 0x12, 0x5f, 0x5, 0x9}}
	return a, nil
}

var _toolsPass_testsMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\x31\x92\xdb\x30\x0c\x45\x7b\x9d\xe2\x77\x6e\x14\x1d\x20\xcd\x36\x69\xb6\x4a\x93\x14\x29\x61\x11\x0c\x39\xa6\x00\x0f\x00\x59\xd1\xed\x33\xa2\x3c\x76\xd2\x02\xe0\x7f\x0f\xe0\x17\xfc\x94\xac\x16\xab\x50\x70\xdb\x47\x44\xa9\x0e\xd7\xb6\x46\x55\x41\x52\x76\x88\x06\xee\xe4\x8e\x28\x8c\x60\x0f\x9f\xf0\xad\x26\xec\xba\xc2\x56\x79\x97\x3f\xf0\xdd\x90\x9e\x1d\x9a\xe7\x9a\x58\x82\x5a\xdb\xe1\xeb\x75\xa9\x01\x12\x54\x99\x75\xb9\x37\x0e\x7e\x51\x3e\x86\xcf\x7c\x3e\x31\xee\x30\x5f\x8d\x51\x74\x43\xe8\x31\xf4\xe0\xd3\x8a\xff\xb0\xcd\xd5\x79\x44\x66\x6e\xc8\xc6\x7c\x4c\x90\xdf\xb0\x30\xb2\x1a\x4a\x95\xf0\x69\x18\x00\xe0\x33\xbf\xc5\xba\x3e\x27\x6c\x85\xa5\x93\x4e\x9f\xe0\xf4\xff\xbe\x63\x6f\x2e\xb4\xa3\xd0\x83\xb1\x15\x8a\x8b\xe3\x26\xba\x09\xc8\x41\xb8\xe4\x56\xe7\x1b\x5b\x95\xdf\x3d\xf9\xf2\x15\x2a\x87\x1e\x9d\x27\xe2\x23\x6b\xe1\xa8\x0b\xfb\x08\x92\x84\x4c\xb5\x39\x34\x0a\x1b\x7a\x79\xc2\x0f\xdb\x8f\xbb\x49\x0f\x79\x29\xc6\xa6\x50\x43\x94\xbe\xd6\x31\x89\x2a\x20\x98\x6e\x3d\xc8\x99\x51\x73\xc7\xf5\xcc\x09\xbf\xfe\x75\x25\x18\xcd\x8c\x59\x25\xd5\xfe\x75\x6a\x5d\xe5\x89\xbe\x9b\x5e\x1b\x2f\xd3\xf0\x37\x00\x00\xff\xff\x80\xe3\x5d\x8b\xf1\x01\x00\x00")

func toolsPass_testsMdBytes() ([]byte, error) {
	return bindataRead(
		_toolsPass_testsMd,
		"tools/pass_tests.md",
	)
}

func toolsPass_testsMd() (*asset, error) {
	bytes, err := toolsPass_testsMdBytes()
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	info := bindataFileInfo{name: "tools/pass_tests.md", size: 497, mode: os.FileMode(420), modTime: time.Unix(1542628451, 0)}
=======
	info := bindataFileInfo{name: "tools/pass_tests.md", size: 497, mode: os.FileMode(420), modTime: time.Unix(1542540514, 0)}
>>>>>>> More tweaks
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe2, 0xe1, 0xfb, 0x58, 0x97, 0x9d, 0xf6, 0x8f, 0x28, 0x7c, 0x49, 0xbc, 0xaa, 0x8e, 0xdb, 0xd5, 0x51, 0xb3, 0x40, 0xb, 0xcc, 0x46, 0x87, 0xb5, 0xda, 0xf7, 0x34, 0xa3, 0x26, 0xc8, 0xff, 0x9b}}
	return a, nil
}

var _topicBenchmarkingMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x55\x4d\x8f\xe3\x36\x0c\xbd\xfb\x57\xbc\x9e\x3a\x13\x8c\x9d\x99\xfd\x9c\xe6\xd6\x69\xd1\xed\x5c\x8a\x02\xbb\x97\xa2\x28\x60\xc6\xa6\x6d\x21\x92\x18\x48\x74\x3c\xfe\xf7\x85\xe4\x7c\x62\xdb\x9c\x14\x45\x7c\x7c\x8f\x7c\x64\x8a\xb2\x2c\x8b\xd5\x6a\xf5\xc2\xbe\x19\x1c\x85\x9d\xf1\xfd\x6a\xb5\x2a\x8a\x57\x0c\x74\x60\x38\x6a\x19\x51\x1c\x23\x8e\x7d\xcf\x51\x8d\xf8\x08\x1d\x48\x61\xfc\x41\xec\x81\xb1\xbd\x0a\xc5\x9d\xa6\x37\xbe\x87\x0e\x8c\x3d\x87\x4e\x82\x23\xdf\x30\xc8\xb7\xe0\xae\x33\x8d\x61\xdf\xcc\x90\x0e\xb3\x8c\x01\x8d\xb4\x7c\x5f\xe1\x35\x7f\xfd\x31\x30\xc8\x06\xa6\x76\x46\x47\xce\x58\x43\x01\x93\xd1\xe1\x36\xc7\x2c\x23\x1a\xf2\x88\x3b\xb3\x87\x0e\x26\x56\x45\x71\xe6\x1f\xcf\x08\xfc\x66\xa2\xa2\x93\x90\xdf\x80\xdf\x38\x34\x26\x72\x85\xbf\x8e\xf1\xfc\xc6\xcd\xa8\x9c\xa8\xba\x9c\x66\x53\xd4\x75\x5d\xf4\x82\x24\x02\xe5\x01\x65\x99\x13\xa3\x3a\x9d\x1c\xbb\xfc\xa6\xf8\x96\x20\x27\x63\x2d\x3a\x13\xa2\x22\x8c\x3e\x4b\x4e\x91\xf1\x21\xab\xd5\x81\x97\xbb\x33\xf9\xf8\x80\x7d\x90\x76\x6c\x92\x8a\x54\x54\x1d\xd2\xc9\x9a\x1d\x67\x8e\x9b\xe2\x48\x40\xe2\x06\x2d\x85\xc9\xf8\xa2\x17\x0a\xcd\xb0\x01\xb9\xf6\xd3\x87\x62\xbf\xeb\x37\xd8\x93\x0e\x6b\x95\xf5\x49\xd1\x45\xfc\x1f\xe4\xb8\x7c\x06\x9e\x1e\xd3\x07\x78\x7a\xfe\xf8\x13\x7c\x5c\xcb\x1e\xf8\xf8\x19\x2f\xf9\xf0\x19\x64\xad\x34\xe9\xb6\xf8\xf3\xe7\xaf\x5f\x0b\xd9\xe1\x3b\x4c\x00\xef\xaa\xc7\x0f\xef\xe2\x22\xf7\x37\x09\x60\x6a\xae\x1a\x91\x94\x05\x5e\x4a\xb0\x65\x10\xac\xf1\x8c\xa8\x14\x72\xfb\x73\xd7\x92\xf8\xfa\x86\x5c\x8d\x4e\xac\x95\x89\x5b\x6c\xe7\xfc\xbb\x1f\xdd\x96\x43\xf2\x43\x23\x81\x23\xee\xea\xe7\xfa\x1e\x74\x20\x63\x69\x6b\xb9\xc2\xb7\xf4\x88\xdf\xf4\xf4\xf2\xae\x5e\xd4\xd5\xf7\x30\xbe\x35\x0d\x29\x47\x0c\x32\xc1\x91\x9f\xa1\xc6\x71\xbc\xad\x3a\x26\x8a\xa9\x3f\x15\xbe\xc8\xc2\x97\x46\x15\x47\x6a\x1a\xb2\x76\xc6\x24\x61\x07\x19\xf5\x3b\x10\x39\x77\xf5\x4a\xb5\xa0\x67\x05\x25\xa5\x6a\xe2\x09\x63\x8c\xdc\x8d\x16\x81\xe3\x68\xb5\x4a\xf6\x38\x92\xd6\x21\xf0\x49\x64\x3c\x13\xde\x14\xc5\x53\xb5\xf8\xc5\x38\x86\x51\xa8\x48\x06\xbf\xf2\xe4\x25\xeb\x03\x8c\x47\x9d\xfb\x58\xe3\xce\x93\x97\xc8\x8d\xf8\x36\xa6\x09\x83\xec\x39\x50\x9a\xcb\xfb\x33\xaa\x63\x27\x21\xb1\xa2\x9e\x97\xe0\x97\x25\x76\x3b\xa7\x6a\x49\x77\x7a\x91\x9d\x40\xca\xed\xff\x21\x5d\xda\x73\x1b\x91\xd6\xc0\x82\x9c\x2f\x12\xfa\xe2\x12\xb2\x36\x05\xc6\xb3\xe8\x07\xa4\x7e\x07\x98\x88\x2d\xab\x72\x38\x96\xc7\x52\x3c\xf7\xb4\x5e\xbc\x56\x5f\x75\x34\x7b\x47\x45\xc9\x1e\x6b\x62\xc4\xe7\x6a\xd5\x79\xa4\x73\x9a\x34\x6b\x79\xd4\x2e\x13\x56\xe1\xb5\xf7\x12\x96\x79\x4a\xe6\x31\x31\x65\x5e\x79\xd1\x15\xa2\xe9\xbd\xe9\x4c\x43\x7e\xd9\x0b\x8e\x29\x8e\x21\x0f\xe3\x9e\xb9\xfd\x21\x7b\xed\x66\xd5\xa8\x88\x4d\x2a\xbf\x9c\x3b\x13\xd1\x51\x54\x0e\x57\x39\xe1\x52\x42\xe9\x34\xcd\xbb\x1c\x47\x9c\x97\xdb\xc0\xd6\x24\x1f\x1f\xad\x91\x36\x80\xc4\x68\xb6\x76\x86\xf1\x4d\x60\x8a\xa7\x65\xf9\x5f\x5a\xab\xa2\xf8\x35\x6d\xe1\x5f\x06\xf6\x3c\x63\xa0\x88\x29\x18\x4d\x89\x08\xbd\x48\x8b\xad\x95\x3e\x41\x2a\x24\xb3\xbc\x61\x3f\x0d\xa6\x19\xf2\xba\x74\x34\xa3\x33\xbe\x85\xf1\xca\x61\x59\xd1\x1b\xfc\xfd\xbb\x4c\x89\x71\xc2\xbc\x5e\x53\x8b\xe2\x7f\xee\x06\xd5\x7d\xdc\xac\xd7\x2d\x1d\xb8\x6a\x32\x87\xca\xb3\xae\xdf\x3d\x3e\xbd\x5f\x3f\x7e\x5a\xbf\x7f\x5c\x0f\x32\x95\x2a\x65\x46\x28\x2f\x08\xa5\xf1\x65\x2f\xf7\x45\xfa\x7b\xf9\x37\x00\x00\xff\xff\x5f\x4d\x95\x5f\x63\x06\x00\x00")

func topicBenchmarkingMdBytes() ([]byte, error) {
	return bindataRead(
		_topicBenchmarkingMd,
		"topic/benchmarking.md",
	)
}

func topicBenchmarkingMd() (*asset, error) {
	bytes, err := topicBenchmarkingMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "topic/benchmarking.md", size: 1635, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4e, 0xee, 0x59, 0xb9, 0x81, 0x2a, 0x96, 0x24, 0xf0, 0xf8, 0xc1, 0xe2, 0x76, 0x7b, 0x64, 0x68, 0x9, 0x81, 0xd0, 0x9f, 0x57, 0x9, 0x62, 0x90, 0x10, 0x4c, 0x15, 0x14, 0x97, 0x34, 0xa3, 0x7a}}
	return a, nil
}

var _topicPprofAllocationsMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x55\xc1\x8e\xdc\x36\x0c\xbd\xfb\x2b\xde\x2d\xcd\x62\xed\x41\x0f\xbd\x6c\xd1\x43\x8a\x22\xe9\x29\x28\xda\x05\x72\x28\x0a\x98\x96\x69\x5b\x5d\x49\x34\x24\xd9\xce\xf4\xeb\x0b\xca\xf6\xec\x24\x48\xd1\x9b\x77\x45\xbe\x47\x3e\x3e\x72\xaa\xba\xae\xab\x87\x87\x87\x79\x8e\x32\x3c\xe1\x9d\x73\x62\x28\x5b\x09\xe9\xe1\xe1\xa1\xaa\x9e\x27\x46\x47\xc9\x1a\x64\x11\x87\x41\x22\x34\xd0\x3a\x1b\x46\x7c\x10\xfd\x63\x8c\xe4\x13\x6c\x42\x5b\x30\xda\x06\xcf\x13\x47\x06\x45\x46\x9e\x22\x33\x7a\x3b\x0c\x1c\x39\x64\xac\x1c\x93\x95\x00\x19\x90\x27\x3e\x33\x0a\xf6\x53\x55\xd5\xf8\x33\x2e\x21\x5b\xcf\x67\xe0\x5f\xdf\x4d\x39\xcf\xe9\xe9\x72\x19\xc5\x51\x18\x1b\x89\xe3\x65\x7e\x19\x2f\x47\xdc\xa5\x00\x5c\xde\x6a\x6a\xe0\x7c\xd1\xe8\xff\xcd\x3d\x03\xef\x93\x8d\x78\x4f\xa1\xaf\x9d\x0d\xdf\x20\xef\x9c\x8c\xcd\x3d\xca\x29\x41\x3d\x4a\x7d\x4a\xf0\x76\x57\xeb\x5b\x48\xda\xef\xd9\xab\x4d\x98\x29\xe6\x53\x82\x0f\x52\xba\x37\x13\xd9\xf0\x08\x0a\x3d\xae\xb2\xc0\x50\x40\x5c\x02\x6c\xc6\x66\xf3\x84\x76\xdc\xc3\x70\x80\x48\xc4\xdf\x4b\xca\xaf\xa0\xa1\x80\x65\x8e\xde\x06\x72\x0d\x3e\x31\x36\xeb\x1c\x96\xa4\x43\xb0\x09\x41\x36\x64\x81\x13\x79\x41\xcf\x3c\x73\x84\x0d\x59\x4a\x1a\xbd\x0e\x5d\xcb\x92\x25\xc2\x48\xcf\x4d\x55\xbd\xb7\x31\x65\x6c\x8c\xc0\xdc\x6b\xbe\x89\x4c\x99\x41\xf0\xec\x25\x5e\x0f\x33\x70\x21\xd4\xa2\x7b\xd9\xe9\xb6\xc9\x3a\xd6\x16\x82\x1a\xa5\xe3\x60\x26\x4f\xf1\x25\x3d\x55\x55\xdb\xb6\x95\xb6\xc3\x29\xa3\x5e\x51\xd7\xe5\x15\xcd\xf9\xe5\xd9\xa3\xae\x3d\xfb\x03\x5b\xa9\x1a\x59\xf2\x17\xff\x8c\x94\xf9\xa7\xef\x0b\x56\xf5\x89\x4b\x77\x13\xad\x5a\x58\x7b\x84\xb7\x28\xc9\x46\x42\x26\x5b\xaa\xf8\xcf\xa2\x6d\x58\x39\x65\x3b\x6a\x67\xa7\xe2\xf7\xf6\x3c\x46\x8a\x21\x8a\xff\x52\xe7\xbd\x9b\x12\x76\x96\xb9\xd7\xf4\x51\x36\x95\x4d\x77\xe0\x98\xcd\x2b\x58\x48\xa2\xe4\xbf\xea\x92\x6c\x7b\x05\xf9\x3a\x33\xd2\xcc\xc6\x0e\xd6\x7c\xcd\x9b\x1a\xfc\x91\xd5\x32\xbb\x15\x26\x76\xb3\xae\x0c\xd2\x24\x1b\x08\xce\xa6\xe2\x26\x72\x0e\xb4\x92\x75\xd4\x39\xbe\xcb\x3d\x9a\xcc\x13\x07\x24\x66\x6c\x13\xe5\x32\xe3\x2c\xf3\xa9\x88\x16\xb5\x78\x8e\xa9\x54\xbc\xf3\x64\x99\x7f\x68\x1f\xd5\x6a\xe5\x13\xb5\x59\xfc\xce\x2b\x31\x2b\x9c\x47\x77\x2d\xad\x99\xc5\x2f\x8e\xb2\x5d\x95\xd6\x2d\x3e\x34\x65\x28\xca\x4a\x2e\x09\xcc\xc4\xe6\x45\x2d\x61\xa6\x62\xac\xd3\x70\x9c\x0e\xfe\xdd\xf8\xda\x8e\x5f\xcc\x54\x0e\xc8\x4d\xf4\x56\xfb\x6b\x91\xe9\x85\x13\x08\x91\xc7\xc5\x51\x04\x7f\x9e\x23\xa7\xb2\x5c\x5a\x12\x53\x34\x93\x4a\x7d\xbd\xb9\x17\xcf\xf1\x8a\xd1\xae\x3a\x7a\x9b\x41\x18\x96\x60\xd4\xe4\x08\xe4\x59\xb3\xee\xe6\x7e\xf8\xb2\x68\x79\xc6\x7d\x24\xcf\xfb\x34\x9f\x8b\xa5\x75\x9f\x76\xcd\x9d\xdb\xfb\xd8\xca\xa1\x6b\xef\x13\x5a\x6c\x94\x30\xc8\x12\x7a\x64\x55\xda\x86\x55\xdc\x7a\x6b\xf5\x6e\xd9\x1a\xfc\xcc\x83\x94\x2b\xc9\xd0\x73\x91\xf6\x8c\x53\x9d\x33\x43\x0f\xc2\x4e\xce\x7c\x13\xe9\x7c\x54\xb6\x33\xa1\x6f\xaa\xea\x77\xf6\xec\x3b\x8e\x3b\x94\xda\xab\x6c\x6c\x5f\x48\xce\x8d\xea\xae\xb7\xd5\xa4\xd7\xe5\x7c\x44\x12\xf4\x12\xde\x64\x74\x3a\x23\x8a\x7e\x5f\x7a\xe5\x4d\xbc\x72\x24\x07\xcf\x23\x75\x57\x1d\x5d\xc7\x25\xff\xe4\x86\xba\x2a\x18\xfe\x71\xb7\x44\x91\x87\x12\xf8\x33\x9b\xa5\x3c\x63\xe0\x0d\x79\x92\x25\xe9\x58\xf5\x7a\xa7\x72\x61\x02\x39\x77\x7d\x84\xcd\x6f\x12\x98\xd2\x55\x19\x47\x0e\x65\xc5\x41\x08\xd6\x30\x56\x9b\x16\x72\xf6\x9f\xa2\xdb\x79\x3a\xbf\xe5\xdf\xdd\x3b\xed\xdc\xdf\xad\x6d\xd1\xee\x76\xb8\x7a\x4b\x7a\xae\x8b\xe3\x92\x9e\x0c\xf5\x86\x3a\xeb\xb7\x5f\xde\x63\x3f\x0b\xef\x5c\xe6\x18\x8a\xa1\xb5\xb2\xb2\xbd\x1b\x77\x5f\x21\xca\xcc\x87\xdd\xba\x28\x5b\xe2\xb8\x43\xea\x80\x34\xe3\xe4\xc9\xea\x91\xa6\xd2\xdf\xd9\x7f\x03\x00\x00\xff\xff\xaf\x47\x4f\x76\x6c\x07\x00\x00")

func topicPprofAllocationsMdBytes() ([]byte, error) {
	return bindataRead(
		_topicPprofAllocationsMd,
		"topic/pprof-allocations.md",
	)
}

func topicPprofAllocationsMd() (*asset, error) {
	bytes, err := topicPprofAllocationsMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "topic/pprof-allocations.md", size: 1900, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6, 0x38, 0x5, 0x7a, 0xca, 0xbc, 0x82, 0x60, 0xe2, 0xc, 0xc7, 0x5f, 0x7d, 0x3f, 0x83, 0xa, 0x6d, 0x26, 0x60, 0xab, 0x1c, 0x6b, 0x52, 0x1d, 0x15, 0x9b, 0xc8, 0xc, 0xf, 0x22, 0x38, 0x25}}
	return a, nil
}

var _topicRegexMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x41\x6f\xe3\x46\x0c\x85\xcf\x99\x5f\xc1\x2e\x50\xc4\x16\x64\x19\xbd\xa6\xc8\x61\x5b\x74\x17\x7b\x28\x50\xac\x0f\x7b\x28\x8a\xd5\x58\xa2\x47\x03\x8f\x66\x04\x92\xb2\x93\x06\xf9\xef\x05\x67\x64\x3b\x08\x7a\xd8\xa3\x34\xe4\xe3\xe3\xf7\x68\x36\x9b\x8d\xa9\xaa\xea\x2b\xba\x39\x58\x82\x3f\x9e\x26\x42\x66\x9f\x22\x57\x55\x65\xcc\xc7\xae\x4b\xd4\xfb\xe8\x40\x12\xc8\x80\xf0\x77\x9f\xba\x79\xc4\x28\x56\x7c\x8a\x90\x0e\xf9\x2f\xa1\xc3\xa7\x09\x26\xdb\x1d\xad\xc3\x7f\x56\x83\xc8\xc4\x0f\xdb\xad\x4b\xc1\x46\xd7\x24\x72\xdb\xe9\xe8\xb6\xa5\x6c\xbb\xae\xe1\x73\xba\x67\xf0\xe3\x14\xf0\xa6\xe5\x19\xdc\x6c\xc9\x46\x41\xec\x75\x1e\xcd\x11\x7c\x84\xef\xc1\x47\xb4\x04\xe2\x47\xfc\xfe\x2b\xc8\x60\x05\x3c\xd7\xe0\x05\xfa\x84\x1c\xef\x05\x1c\x0a\x70\x48\x67\x24\xb0\x9c\x1d\xb1\xff\x17\x2f\xee\x7c\x9c\xe6\x5c\xc3\xb0\xf7\xce\x21\x35\xf0\x1b\x76\x76\xe6\xa5\x42\xc5\x68\x01\x80\x37\x00\x3a\xfb\x73\x82\xc1\x9e\x10\xe4\x9c\x80\x71\xb2\x64\x05\x81\xc5\x3a\xe4\x07\x68\xbb\x34\x4e\x3e\x64\xf7\x2d\xd8\xd8\x43\x3b\xb3\x75\xd8\x36\xc6\x54\xd5\xef\xb7\x47\x25\xf9\x6d\xc0\x08\xa5\x41\x71\x5a\x68\x59\xf7\xee\xda\x02\x4f\xb7\xf1\x0c\xb6\x3f\x79\xb6\xfb\x80\xba\xbf\x1a\x6c\x0b\xb3\xe6\xcf\x99\xa5\x28\xe2\x6a\xdd\xd6\x79\xda\x98\xd4\xd9\x80\xf0\xc6\x87\xf6\xb5\x4b\x0e\x2d\x04\x3c\x61\x68\xe0\xcb\x9b\x94\x74\x8a\x8f\x27\x1b\x7c\x5f\xe7\xbf\x13\x25\x47\x76\x84\xb3\x0f\x01\x26\x1b\x7d\x07\x29\xea\x8e\x24\xf3\x54\xc3\x68\x8f\xea\xd7\x0b\x9c\x90\x9e\x21\xed\x4f\x3e\xcd\x7c\x39\x87\x5e\x07\xa4\x09\xa9\xc4\x22\x03\x12\xde\x33\x58\x55\xdd\x07\x1c\x1b\x63\xda\xb6\x75\xc9\x9c\x2c\xc1\xca\xdc\x71\x1a\xf1\x6b\xb6\xf1\x08\xff\xb3\x59\xab\xef\xf9\x7f\xbb\x36\x6b\x63\x0e\x73\xec\x60\x97\x46\xfc\x34\xc7\x6e\xc5\xc0\x42\x3e\xba\x35\xbc\x98\x3b\x42\x9e\x83\xc0\xc3\x23\x5c\x35\x9b\x4f\x3e\xf6\x1f\x43\xd8\xe5\xaa\x15\xd7\xb0\xf9\x65\x6d\xee\xb6\xdb\xa6\x69\xcc\xab\x1a\x31\xe6\x3d\x8a\x98\x04\x4a\x10\xf9\xa0\x58\x94\xc2\x68\x8f\xc8\xc0\x18\x39\xe7\x50\xf0\xa2\xbe\xa7\x18\x9e\xe1\xac\x51\x46\xec\x90\xd9\xd2\x73\x03\x5f\x62\x3e\x22\xe8\x6c\x09\xec\x49\xd7\xba\x85\xa5\x63\xf6\x28\x82\x54\xeb\x75\x7a\x29\xa8\x09\x65\xa6\x08\x36\x02\x12\x25\x82\x43\x22\xfd\x58\xb2\x29\x0e\x6b\x20\xab\x4c\x95\x6e\x2c\xe1\x68\x1a\x57\xaa\xef\xf8\xd4\xa0\x17\x3a\x5e\x31\x15\xe1\x17\x03\x00\x45\x6f\x27\xa4\xc4\x0e\xa3\x34\xbb\x89\x7c\x94\xc3\xea\x43\x7e\xf8\xe6\x65\xf8\x4b\x7b\x7f\xe6\x0f\x8b\xca\x3a\xb7\x5d\xe1\xd6\xaa\xa6\xcd\x4b\x6c\x97\xfd\x2e\xba\xa5\xdc\x1f\x72\xd9\x4f\x8f\x10\x7d\x58\x26\x97\xe9\x79\x59\x24\xca\xbf\x5e\xcd\xe2\xe9\x07\x33\xd4\xe2\x37\x21\x6e\x36\x9b\xff\x02\x00\x00\xff\xff\x58\xeb\x13\x24\xb9\x04\x00\x00")

func topicRegexMdBytes() ([]byte, error) {
	return bindataRead(
		_topicRegexMd,
		"topic/regex.md",
	)
}

func topicRegexMd() (*asset, error) {
	bytes, err := topicRegexMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "topic/regex.md", size: 1209, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb2, 0xa0, 0xf6, 0x4c, 0x4f, 0x5e, 0x68, 0xa6, 0xdf, 0x62, 0xa7, 0xe7, 0x1b, 0x50, 0xaf, 0xb6, 0x68, 0x7c, 0x8e, 0xc8, 0x9a, 0x40, 0xfd, 0x98, 0x73, 0x93, 0xce, 0xf5, 0x24, 0x34, 0x8c, 0xd7}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"response/comment.md": responseCommentMd,

	"response/greeting.md": responseGreetingMd,

	"response/improvement.md": responseImprovementMd,

	"response/newcomer_greeting.md": responseNewcomer_greetingMd,

	"response/praise.md": responsePraiseMd,

	"response/questions.md": responseQuestionsMd,

	"response/todo.md": responseTodoMd,

	"tools/compile.md": toolsCompileMd,

	"tools/not_formatted.md": toolsNot_formattedMd,

	"tools/not_linted.md": toolsNot_lintedMd,

	"tools/pass_tests.md": toolsPass_testsMd,

	"topic/benchmarking.md": topicBenchmarkingMd,

	"topic/pprof-allocations.md": topicPprofAllocationsMd,

	"topic/regex.md": topicRegexMd,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"response": &bintree{nil, map[string]*bintree{
		"comment.md":           &bintree{responseCommentMd, map[string]*bintree{}},
		"greeting.md":          &bintree{responseGreetingMd, map[string]*bintree{}},
		"improvement.md":       &bintree{responseImprovementMd, map[string]*bintree{}},
		"newcomer_greeting.md": &bintree{responseNewcomer_greetingMd, map[string]*bintree{}},
		"praise.md":            &bintree{responsePraiseMd, map[string]*bintree{}},
		"questions.md":         &bintree{responseQuestionsMd, map[string]*bintree{}},
		"todo.md":              &bintree{responseTodoMd, map[string]*bintree{}},
	}},
	"tools": &bintree{nil, map[string]*bintree{
		"compile.md":       &bintree{toolsCompileMd, map[string]*bintree{}},
		"not_formatted.md": &bintree{toolsNot_formattedMd, map[string]*bintree{}},
		"not_linted.md":    &bintree{toolsNot_lintedMd, map[string]*bintree{}},
		"pass_tests.md":    &bintree{toolsPass_testsMd, map[string]*bintree{}},
	}},
	"topic": &bintree{nil, map[string]*bintree{
		"benchmarking.md":      &bintree{topicBenchmarkingMd, map[string]*bintree{}},
		"pprof-allocations.md": &bintree{topicPprofAllocationsMd, map[string]*bintree{}},
		"regex.md":             &bintree{topicRegexMd, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
