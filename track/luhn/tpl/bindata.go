// Code generated by go-bindata. DO NOT EDIT.
// sources:
// mustcompile.md (168B)
// one-loop.md (1.731kB)
// regex-in-func.md (307B)
// regex-to-fast.md (1.753kB)

package tpl

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

var _mustcompileMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcc\x41\xb2\x82\x30\x10\x84\xe1\xfd\x3b\x45\xef\x78\x56\x21\x97\x70\xed\x8e\x03\x30\x86\x18\xba\x0c\x0c\x95\x99\x88\xdc\xde\x02\xb7\x5d\xfd\xfd\x57\xf4\x8a\xa0\xf3\xca\x1c\x21\x30\x17\x67\x40\x89\x29\x7e\x5a\xd0\x1b\x83\x20\xa9\x8e\xe0\x18\x05\xae\xa8\x16\x31\xdc\xab\xf9\xed\x87\xfe\x2f\x43\x87\x7e\xa2\x61\x63\xce\x08\x72\x1c\x7c\x8a\x58\x8b\xa6\x22\xf3\x61\x56\x59\x18\xc0\xe7\xb9\x9f\x6d\xd0\xc0\xe5\x2d\x99\x63\x0b\x53\xec\x5a\x9b\x9c\xf1\x5a\x74\x83\x3c\xb4\x3a\xe8\x28\x4c\x93\x43\x36\xd9\xbb\xbf\x6f\x00\x00\x00\xff\xff\xfa\x98\x2c\xcd\xa8\x00\x00\x00")

func mustcompileMdBytes() ([]byte, error) {
	return bindataRead(
		_mustcompileMd,
		"mustcompile.md",
	)
}

func mustcompileMd() (*asset, error) {
	bytes, err := mustcompileMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mustcompile.md", size: 168, mode: os.FileMode(420), modTime: time.Unix(1542541121, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x75, 0xb2, 0x40, 0x36, 0x90, 0xb3, 0x55, 0x4e, 0x7b, 0xa1, 0xe1, 0x86, 0xd4, 0x8f, 0x22, 0xe9, 0x1a, 0x8f, 0x1e, 0x8c, 0x8a, 0xf2, 0x6e, 0x83, 0xf0, 0xfe, 0x6f, 0x3a, 0xd0, 0x94, 0x91, 0x52}}
	return a, nil
}

var _oneLoopMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x55\x4d\x6f\xdb\x46\x10\xbd\xeb\x57\xbc\xe4\x22\x1b\xb0\x58\xbb\xb7\x2a\x4e\x81\x14\x41\x01\x5f\x92\x4b\x01\x23\x28\x0a\x70\x44\x0e\xc5\xad\x57\x33\xc4\xee\x50\xb4\xd0\xf4\xbf\x07\xb3\xa4\x64\x27\xb9\x51\xe4\xce\xcc\x9b\xf7\xb1\x5a\x6d\x36\x9b\xd5\x17\x1d\x13\xb2\xc6\xd1\x82\x0a\xa6\x60\x3d\x0e\x63\xb4\x30\x44\x46\x54\x1d\x32\x42\x46\x62\x6a\x69\x17\x19\x24\x2d\x4c\x8d\x62\x3c\xa1\x0b\xc2\x15\x1e\x3a\x9c\x74\xc4\x44\x62\x30\x05\x1d\x35\xb4\xa5\x30\xc8\x1e\x7a\xe4\x04\xeb\x19\xd9\x92\xff\xbe\x74\xb6\x70\xe0\x7c\x83\x9e\x13\xaf\x33\x08\x2d\x1b\x85\xc8\x2d\xf8\x79\x88\x24\x54\xc0\x68\x07\x15\xc6\xa0\x39\x07\x1f\x3e\xd1\xc9\x47\xb4\x8a\x60\xdb\xd5\xea\xae\xc2\x9f\x21\x65\xbb\x41\x90\x46\xd3\xa0\x89\x8c\xcb\x34\x51\xd9\xb4\x61\x1f\x0c\x3d\x49\x1b\x7d\x72\x10\xd3\xf2\xed\x40\x41\x0a\xbe\xed\x6a\x05\x00\x1b\x3c\xf6\xc1\x38\x0f\xd4\xf0\x16\xd3\xe5\xd9\xd7\xa6\x18\x75\xe2\x16\x41\x70\xa4\x18\xfc\x61\x18\x2d\xdf\x20\x2b\xa6\x9e\x05\x13\x83\xa5\xd1\x51\x8c\xd3\xab\xda\x1b\xff\xd0\x90\x60\xcc\x8c\xba\x51\xb1\x20\x23\xd7\x0e\xfe\xdf\xf1\x30\x60\x81\x22\xfc\x6c\x08\xc6\xe9\xb2\xae\xbf\x75\x6c\xd5\x19\xdb\x67\xeb\x39\xbd\xec\x93\xb7\xde\xb9\x4e\x6c\x63\x12\x74\x14\x33\xd7\x08\xdd\xf7\x38\xac\xe7\x83\x43\xbe\x34\xc3\x17\x1d\x0b\x9c\xd6\xe7\x92\xcd\x2a\xd7\xa3\x84\x46\x5b\xae\x1e\xf2\x47\xef\x5d\x1a\xb9\x94\x94\x18\x93\xa6\x27\xa7\xad\x9c\x4c\xa3\x70\xae\x0a\xe1\x8f\xfe\xdb\xfa\x90\x5f\x33\xd5\xf4\x24\xfb\x79\xe9\x20\x96\xb4\x1d\x1b\x6e\x41\x18\x92\xee\x22\x1f\xb6\x05\x49\x90\x96\x9f\x51\x87\xfa\xf5\x9e\x4e\xb2\xa8\xa1\x40\xf7\x79\x6f\x1b\x4d\x89\x1b\x8b\xa7\xb7\x20\x39\xe1\xa0\x89\x2b\x3c\xce\x6c\x36\x89\x5d\x61\x82\xf0\x84\x23\xa5\x50\x2c\x59\x2f\x7b\x17\x7a\x9f\x98\x07\x58\xa2\xe6\xc9\xc7\xf4\x3a\xe1\xe0\x6d\x66\xf2\x30\xf1\xfa\xc8\xc8\xcc\x72\x53\x9c\x1c\xa4\x49\x7c\x60\x71\x15\xfc\xe0\x58\x7c\xfd\x8a\x3a\x5c\x51\xb7\x50\xfa\x5a\x47\xaf\x5d\x44\xa8\xd1\x50\xe6\x7c\x5d\xe1\xa3\xca\xda\xd0\x69\xda\x73\x49\x42\xd3\x73\xf3\x84\x97\xfa\x79\xdf\xee\x82\x17\xbf\xe3\xae\x60\xce\xcc\x8b\x86\x7b\xb5\xb2\xb1\xab\x24\xc5\xfb\x05\xf8\x4c\xfd\x07\xd1\x62\x86\x85\xd5\x1f\xd8\x9e\x66\x61\xf8\x67\x5d\x9c\xe3\x59\xf5\xc2\x62\xa1\x9b\x62\x33\xc6\x73\x5a\xea\xc8\xb2\xb7\xbe\xc6\x8e\xbb\x79\xf8\x02\x96\xe4\x34\xf3\xff\xc7\x78\xae\x86\xe8\xb4\x58\x96\xb1\xa3\xe6\x69\xa2\xd4\xe6\x9f\x72\x9e\x83\x34\xec\x25\x3d\x1d\xfd\xd2\x28\xea\x0f\x2c\xad\x73\x3d\x27\x73\x61\x61\x8b\xba\xd3\x84\x80\xed\x7b\x44\x96\xab\x6c\xe9\x1a\x1b\xdc\xbd\xc3\x2d\xee\xdf\x23\xbc\x43\xd8\x6c\xf0\x5f\x55\x55\xff\xd7\x4b\x2a\xfe\xea\x17\x5e\x8a\x31\xe7\xf5\x58\x50\x27\xef\xe1\xaf\xbc\xc9\xdf\xe1\x9f\xeb\xfa\xc5\xf9\xc6\xd9\x0a\xc0\x23\xc5\x91\xdd\x1b\xdf\xd9\xc6\x25\x98\x7a\x2e\xfc\x6a\x2a\x96\xf4\xc3\xcd\x98\xd2\x0b\xe2\xdc\xeb\x18\x5b\xec\x18\xad\x8e\xbb\xc8\x6d\xb5\x5a\x7d\xb0\x39\x0d\x83\x06\x29\x8f\xb2\xcf\xe7\x83\x9e\x21\xd0\xde\xef\x1b\x37\xcc\x8e\x67\x2f\x96\xab\x0f\x1d\x65\xe3\xf4\x66\xb5\xda\xe0\xb3\x30\x22\xe5\xa5\x7c\x7b\x76\x7b\xe2\x21\xba\x8a\xf5\x9b\x1f\x82\x7a\x95\xae\xeb\x25\xc2\x09\xf7\x58\xdf\xae\xf1\xf5\x2b\xd6\xbf\xad\x71\x8f\x34\x9b\xd3\xbd\xc4\x06\x5a\x2c\x53\xf0\x96\xfb\xb8\x43\x1e\xb8\x20\x07\x3e\xa9\xf9\x8d\x77\x49\x74\x36\x1e\xce\x32\xb7\x49\x87\x99\x02\x95\x23\xa7\xec\xf7\x93\x29\x6a\xa7\xb7\xae\xf0\x60\x65\xbb\x3c\x6b\x4e\x07\x5e\xf0\xec\x4e\xc6\x2e\xd3\x27\x9d\xbc\xd3\x0b\x63\x4b\xa6\x28\x31\x15\xf2\x7f\xbd\xdd\xdc\xdd\xde\x42\xf2\x2f\x3a\xd4\x98\xbd\x51\x00\xce\xe7\x7a\x4a\xed\x44\x8b\x17\x77\x2c\x4d\x7f\x20\x1f\xe7\xaf\xd2\xe8\xd1\xa8\x56\xe5\xdf\xeb\x5b\x00\x00\x00\xff\xff\xe8\x69\x65\x6a\xc3\x06\x00\x00")

func oneLoopMdBytes() ([]byte, error) {
	return bindataRead(
		_oneLoopMd,
		"one-loop.md",
	)
}

func oneLoopMd() (*asset, error) {
	bytes, err := oneLoopMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "one-loop.md", size: 1731, mode: os.FileMode(420), modTime: time.Unix(1543061558, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf9, 0xa3, 0xd4, 0x9a, 0xe9, 0xd4, 0x4d, 0x75, 0x80, 0x85, 0x7a, 0xd1, 0x1a, 0xbe, 0xaa, 0x2d, 0x11, 0x68, 0xca, 0x7, 0x10, 0x47, 0xb3, 0xa3, 0x92, 0x3f, 0x76, 0x90, 0xfc, 0x6e, 0xd9, 0x75}}
	return a, nil
}

var _regexInFuncMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcd\x3b\x72\xc3\x30\x0c\x84\xe1\x3e\xa7\xd8\xce\x8d\xa3\x63\x24\x93\xd6\x37\xa0\xa8\x25\x8d\x31\x05\x68\x40\xd0\xb6\x6e\x9f\x51\x1e\x93\xf4\xff\x7e\xfb\x8a\x0b\xeb\x68\xc9\xc1\xe7\xe6\xec\x5d\x4c\x3b\x44\xf1\x6e\x78\x98\xdf\x90\xd0\x24\xa2\x11\x8b\x94\x42\xa7\x46\xdb\x51\xdc\x56\x58\x5c\xe9\x68\x49\xeb\x48\x95\x7d\xc2\x9b\x39\x66\x46\xd0\xb1\xd1\x8b\xf9\x9a\x34\xf3\x8c\x04\x67\xe5\x13\xd2\x91\x6d\xdd\xa4\x71\xc1\xcc\x62\x4e\x48\x9c\x3a\x46\xe7\x32\xe1\x22\xf5\x1a\x50\x7b\x9c\xb1\xdb\x38\x39\x7f\x62\xd1\x8a\xb8\xf2\xd7\xd0\x2e\x0b\x8f\xc2\x51\x86\xe6\x10\xd3\x33\xe6\x11\x90\x80\x69\xdb\xa1\xe4\xd2\x11\x86\x99\x7f\x77\xa6\x99\x13\x3e\xca\x7f\xa9\xa3\x47\x0a\xc9\x07\x86\x9c\x14\xab\xdd\xf9\x15\x7c\xcf\xd2\x61\x1f\xd0\x96\xf2\x2d\x55\xa2\xf1\xce\x36\xbd\x7c\x06\x00\x00\xff\xff\x5d\x1e\x2b\xac\x33\x01\x00\x00")

func regexInFuncMdBytes() ([]byte, error) {
	return bindataRead(
		_regexInFuncMd,
		"regex-in-func.md",
	)
}

func regexInFuncMd() (*asset, error) {
	bytes, err := regexInFuncMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "regex-in-func.md", size: 307, mode: os.FileMode(420), modTime: time.Unix(1543061558, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb9, 0xbc, 0x15, 0x8, 0x47, 0xce, 0xaa, 0xd3, 0x30, 0xee, 0xc5, 0x40, 0x2b, 0x17, 0xcb, 0x79, 0x80, 0x6f, 0xc0, 0x5e, 0xd4, 0x9c, 0x8a, 0x3e, 0xb8, 0x5c, 0xcb, 0x6a, 0xd1, 0xca, 0x1a, 0x6c}}
	return a, nil
}

var _regexToFastMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x55\x4d\x8b\xdc\x46\x13\xbe\xcf\xaf\x78\xec\xcb\x78\x61\xa5\x77\xf7\xbd\x65\x6c\x07\x1c\x4c\x88\x2f\x36\x84\x80\x31\x21\xa0\x96\x54\x92\x9a\xed\xa9\x12\xdd\xd5\xa3\x1d\xe2\xfc\xf7\x50\x2d\xcd\xce\xda\xce\x69\x50\xd3\x55\xfd\x7c\x55\xcd\xae\xaa\xaa\xdd\x17\xc9\x11\x49\x42\x56\x2f\x8c\xc5\xeb\x84\x26\xd2\x48\x8f\x73\x03\x9f\x70\xa2\x78\x46\x24\xd7\xbb\x36\x10\x1c\xf7\x50\x51\x17\xc2\x19\x83\x67\xaa\xf1\x3b\x8d\x39\xb8\x08\x7a\x9c\x23\xa5\xe4\x85\x13\x5c\xa4\xb5\x6e\x96\x85\xe2\x90\xc3\x2d\xda\xac\xc8\x29\x97\xc2\x14\x64\xa9\xf1\x61\x80\x4e\x9e\xc7\x84\xce\x31\x5a\x42\x2f\x4c\x38\x3a\x5e\x2f\x15\x20\x91\x5c\x12\x2e\x2f\xd3\x30\x48\x54\xe8\x44\xa0\x47\x8d\x0e\x8b\xc4\x07\x03\x28\x83\x12\xdb\x97\x4e\xf0\x5a\xe3\x37\x8a\x64\xe7\x0e\x49\x69\xae\xda\x73\x65\xbf\x18\xb3\xef\x09\x2a\x70\x18\x5c\x52\xba\x72\x3e\xec\x76\xf7\x35\x7e\xf5\x31\xe9\x2d\x28\xf8\xa3\x67\xa7\x54\x5e\x8a\x3f\x90\x43\x7b\xc6\x51\x4e\x9e\xc7\x72\x21\xc8\xe8\x3b\x78\x56\xd9\x3e\x65\x3e\xec\x76\x00\x50\xe1\xf3\xe4\x95\xd2\xec\xba\x15\x4e\x08\xb2\x50\x7f\x8b\x24\x58\xa8\x70\x6e\x3a\x61\xf5\x9c\xa9\x81\x1f\xec\x90\xb8\x93\xcc\x86\x6d\x79\x56\xcb\x4f\xad\xeb\x4b\xeb\x4f\x3a\x51\x04\x0b\x57\xbd\x1f\xbd\xa6\x83\x55\x37\x91\x34\x47\xc6\xe0\x42\xfa\x8f\x96\x3a\xd1\xf1\x9b\x66\xf8\x22\xb9\xe0\xe8\x0d\xbc\xd3\xcd\xfb\xcc\xbe\x93\x9e\xea\x0f\xe9\xbd\xf5\x2e\x8d\xce\x92\x8b\xa9\xa6\xb9\x51\x5f\xcd\xc9\x4c\xa9\x2e\xe2\x7d\xb6\x6f\x9d\x7c\x7a\x0e\xbc\x9b\x1c\x8f\x74\x6b\x28\x3c\x6b\x94\x3e\x77\xd4\xc3\x61\x8e\xd2\x06\x3a\x1e\x0a\x12\xcf\x3d\x3d\xa2\xf1\x0d\x64\x78\x82\x66\x7a\xb1\x28\x0a\x74\x7b\xef\x65\x27\x31\x52\xa7\xe1\xfc\x12\x8e\xcd\x81\x48\x35\x3e\xaf\x32\x76\x91\xcc\x2f\x07\xa6\x05\x27\x17\x7d\xc9\x4b\xb3\xf1\x6e\xcc\xf3\x07\xa2\x19\x1a\x5d\xf7\x60\xcf\x4c\xb2\x58\xce\xce\x58\xc5\xc3\x42\xfb\x13\x21\x11\xf1\x6d\x09\xb8\xe7\x2e\xd2\x91\x58\xe1\xf5\x1a\xc8\x67\xd2\xe1\x95\x1b\x36\x49\x9f\x9b\x68\xb5\x9b\x09\x0d\x3a\x97\x28\xdd\xd4\x78\x2f\xbc\x57\x0c\x12\x47\x52\x83\xd2\x4d\xd4\x3d\xe0\x5a\xbf\xf2\x1d\x9e\xf0\xe2\x67\xdc\x17\xcc\x89\x68\xf3\x70\x14\x2d\x8c\xcd\x25\x86\x4d\x49\x01\xbe\x4a\xff\x8e\xa5\x84\x61\x53\xf5\x3b\xb5\x97\xd5\x18\xfa\xd1\x17\xd3\x78\x75\xbd\xa8\x58\xe4\x76\xa1\xcb\xe1\x92\xfd\x26\x10\x8f\x3a\x35\x68\x69\x58\x1f\xdf\xc0\x3a\x3e\xaf\xfa\xff\x92\x2f\xd5\x60\x59\xe0\x95\xa2\x15\xb7\xae\x7b\x58\x5c\xec\x13\xe4\xb4\x91\x4c\x1a\xcd\xc6\xe4\xb9\x23\x2b\x99\xdc\xc9\x76\x49\x71\x7f\x26\xee\x4d\xeb\xc2\x09\x9b\x0a\x07\x34\x83\x44\x78\x1c\xde\x22\x10\xbf\x4a\x1a\x6f\x50\xe1\xfe\x35\xee\xf0\xe6\x2d\xfc\x6b\xf8\xaa\xc2\xdf\x75\x5d\xff\xd3\x6c\x53\xf1\xc7\xb4\xe9\x52\x82\xb9\xd2\x23\x46\x13\xad\x87\x1d\x59\x93\x3f\xfd\x5f\x37\xcd\x35\xf9\x4a\x69\x5d\x29\x27\x17\x32\x59\x36\xbe\x89\x8d\x59\xb0\x4c\x54\xf4\x95\x58\x22\x69\x97\xbb\x1c\xe3\x15\x71\x9a\x24\x87\x7e\x5d\x60\xb9\x0d\xd4\xd7\xbb\xdd\x3b\x5d\xa7\x61\x16\xcf\x7a\x59\x73\xdb\xc5\xb2\xb7\xdc\xe8\x3c\x97\xc0\xb4\xb4\x66\x51\xfd\x91\xd2\xb6\x9a\x5e\xec\x76\x15\x3e\x31\x21\xb8\xb4\x95\x1f\x2e\x69\x8f\x34\x07\x73\xb1\x79\xf1\xdd\xa0\xbe\x8a\x37\xcd\x65\x7d\xe3\x0d\xf6\x77\x7b\x7c\xfd\x8a\xfd\x4f\x7b\xbc\x41\x5c\xc3\x69\x59\x22\x85\xdb\x22\x53\xf0\x9a\x2f\x32\x20\xcd\x54\x90\x03\x1f\x45\xe9\x70\x09\x8e\x4f\x65\x8b\x5e\x6c\xee\xa3\xcc\xab\x04\xc2\x27\x8a\x65\x21\xaa\xa0\x31\x79\x9b\x1a\x1f\xb4\xb0\x4b\xab\xe7\xee\x48\x1b\x9e\xf6\xac\x64\x36\x7d\x94\xc5\x3a\x5d\x15\xdb\x66\xca\x45\x72\x45\xfc\xff\xdf\x55\xf7\x77\x77\xe0\xf4\x3f\x99\x1b\xac\xd9\x28\x00\xd7\x7b\x93\x8b\xfd\xe2\xb6\x2c\xb6\xc4\xdd\x74\x74\xf6\x9c\x1d\xc5\x6c\xa3\x51\xef\xca\x7f\xda\xbf\x01\x00\x00\xff\xff\x80\xf1\x13\xb0\xd9\x06\x00\x00")

func regexToFastMdBytes() ([]byte, error) {
	return bindataRead(
		_regexToFastMd,
		"regex-to-fast.md",
	)
}

func regexToFastMd() (*asset, error) {
	bytes, err := regexToFastMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "regex-to-fast.md", size: 1753, mode: os.FileMode(420), modTime: time.Unix(1543061558, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe5, 0x1e, 0xe8, 0x7a, 0x91, 0x31, 0x2c, 0x69, 0xac, 0x72, 0x52, 0x8c, 0x41, 0x9c, 0x75, 0x2b, 0x17, 0x30, 0xbe, 0x13, 0x57, 0xb7, 0x6a, 0xb9, 0x14, 0x26, 0x8d, 0x7b, 0x52, 0x79, 0x85, 0xdf}}
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
	"mustcompile.md": mustcompileMd,

	"one-loop.md": oneLoopMd,

	"regex-in-func.md": regexInFuncMd,

	"regex-to-fast.md": regexToFastMd,
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
	"mustcompile.md":   &bintree{mustcompileMd, map[string]*bintree{}},
	"one-loop.md":      &bintree{oneLoopMd, map[string]*bintree{}},
	"regex-in-func.md": &bintree{regexInFuncMd, map[string]*bintree{}},
	"regex-to-fast.md": &bintree{regexToFastMd, map[string]*bintree{}},
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
