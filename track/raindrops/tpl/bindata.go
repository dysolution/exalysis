// Code generated by go-bindata. DO NOT EDIT.
// sources:
// concat-not-needed.md (77B)
// many-loops.md (146B)
// strings-builder.md (251B)

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

var _concatNotNeededMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xc1\x0d\x02\x21\x10\x85\xe1\x56\xfe\xbb\xd1\x0e\x2c\xc4\xdb\x12\xe6\xb9\x4b\x24\x4c\x32\x30\x31\x74\xbf\xdf\x93\x8f\x27\xd5\xb3\x1b\xf5\x52\xfd\xb1\x3d\x83\x9c\xe5\x14\xfe\xe5\x78\xbc\x0f\xda\x20\x74\x96\xb0\xc9\x72\xfe\x97\x42\xb4\x45\x9b\x84\x4a\xef\x9b\x21\x99\xec\x75\x07\x00\x00\xff\xff\x68\xe9\x4d\x1a\x4d\x00\x00\x00")

func concatNotNeededMdBytes() ([]byte, error) {
	return bindataRead(
		_concatNotNeededMd,
		"concat-not-needed.md",
	)
}

func concatNotNeededMd() (*asset, error) {
	bytes, err := concatNotNeededMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "concat-not-needed.md", size: 77, mode: os.FileMode(420), modTime: time.Unix(1540277586, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb9, 0xa8, 0x32, 0xed, 0x64, 0x1c, 0x58, 0x85, 0x29, 0xe4, 0x1d, 0x23, 0xf0, 0xed, 0xd5, 0xc1, 0x2e, 0xce, 0x54, 0xd3, 0x58, 0x88, 0xd1, 0xc6, 0x30, 0xbc, 0xa1, 0xa6, 0xd4, 0x42, 0xca, 0x42}}
	return a, nil
}

var _manyLoopsMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xcc\xc1\xad\x83\x40\x0c\x84\xe1\xfb\xab\x62\x1a\x78\xb4\x90\x03\x2d\xa4\x01\x47\x2c\x60\xc5\x78\xd0\xda\x16\xa2\xfb\x28\x7b\xc8\x71\x46\x9f\xfe\x7f\xcc\x5c\xd4\x37\x68\x40\x60\x4c\xc8\x8b\x95\x58\xd5\xc7\x2d\x08\x3d\x4e\x6b\x08\x5a\xa5\xd2\x27\x3c\x77\x8d\xdf\x44\x45\x0b\x1c\x65\xa9\x5f\x65\xe4\x19\x13\x66\x71\xdc\x2c\x64\xbf\x21\xbe\xe0\x62\x7f\xe3\xef\xd2\xdc\x47\xd0\x37\x6b\x58\xd9\x07\x07\x3b\x9c\xde\x20\x09\x31\x7b\x7c\x02\x00\x00\xff\xff\x31\x8f\x37\xf5\x92\x00\x00\x00")

func manyLoopsMdBytes() ([]byte, error) {
	return bindataRead(
		_manyLoopsMd,
		"many-loops.md",
	)
}

func manyLoopsMd() (*asset, error) {
	bytes, err := manyLoopsMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "many-loops.md", size: 146, mode: os.FileMode(420), modTime: time.Unix(1539184320, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5, 0x1f, 0x5b, 0xd7, 0xf4, 0xb2, 0xe, 0x23, 0x8, 0x59, 0x56, 0xba, 0x80, 0xba, 0xc8, 0xfc, 0x4a, 0xda, 0x52, 0xad, 0xee, 0x92, 0x2d, 0xea, 0x3a, 0x6, 0xe0, 0x4b, 0xa8, 0x3c, 0x4d, 0x6e}}
	return a, nil
}

var _stringsBuilderMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8f\xbd\x4e\xc4\x40\x0c\x84\xfb\x7b\x8a\xe9\x11\x5b\x20\x51\xd2\xd0\xdd\x23\xd0\xad\xb3\xf1\x5d\xac\xdb\x1f\xc9\xe3\x45\xa4\xe1\xd9\x51\x82\xa0\x73\xe1\xef\x9b\x99\x67\x5c\x41\x55\xec\x63\x42\x5c\x31\x69\xfd\x0e\x41\x66\xb8\xf5\x3b\xd3\xfb\xb4\xba\xaa\xe7\x84\x6b\x47\x6c\x46\xe8\x97\x7a\x31\x2a\x62\x93\x38\xe0\x46\x08\x16\x0b\x8c\x4f\xf5\x87\xd5\x0a\xe9\x2b\xf2\xd3\x5b\x86\x94\x98\x52\xeb\x8e\x4b\x19\x9d\xb3\x29\xf1\xfd\x92\x5e\x11\x76\x9c\x55\x49\x34\x6d\xc3\xf7\x13\xb1\x3f\xd1\x4d\x18\xea\x09\x1f\x63\xa2\x48\x87\xcf\x23\x5b\x91\x17\xed\x65\x6b\xe2\x0f\x66\x08\xb1\x2a\x8b\xdb\xa2\x2b\xec\xf7\xe1\xf2\x5f\xce\x3a\xc3\x67\x09\x1b\x9d\xa7\x3c\x7c\x87\xc5\x31\xd4\xa9\xf5\x96\xf0\x13\x00\x00\xff\xff\xe3\x47\xf0\xa4\xfb\x00\x00\x00")

func stringsBuilderMdBytes() ([]byte, error) {
	return bindataRead(
		_stringsBuilderMd,
		"strings-builder.md",
	)
}

func stringsBuilderMd() (*asset, error) {
	bytes, err := stringsBuilderMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "strings-builder.md", size: 251, mode: os.FileMode(420), modTime: time.Unix(1539184320, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x32, 0xb4, 0xb4, 0x9a, 0xe9, 0x90, 0x6d, 0xe3, 0x6d, 0x4c, 0x2f, 0x5c, 0xf9, 0x6e, 0x68, 0x92, 0x55, 0x96, 0x95, 0x9b, 0xe4, 0xf0, 0x5b, 0x50, 0x76, 0x28, 0xcb, 0x83, 0x35, 0xca, 0x19, 0xa8}}
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
	"concat-not-needed.md": concatNotNeededMd,

	"many-loops.md": manyLoopsMd,

	"strings-builder.md": stringsBuilderMd,
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
	"concat-not-needed.md": &bintree{concatNotNeededMd, map[string]*bintree{}},
	"many-loops.md":        &bintree{manyLoopsMd, map[string]*bintree{}},
	"strings-builder.md":   &bintree{stringsBuilderMd, map[string]*bintree{}},
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
