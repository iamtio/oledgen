// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// sprites.png
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _spritesPng = []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\x10\x00\x00\x00`\b\x06\x00\x00\x00\x14\xd1\xf2\xa0\x00\x00\x01}IDATXG\xedXk\x17\xc2 \bu\xff\xffG\xaf\xa3\xc5\x06\xc8KfuZ\xeeK5\xe5\x86\xc0\xe5\xe1V\u8cd7R\xb6R\x8a\xf7yH\xd5\xcd\xf8\xf1\x04a]\x05`x\xfeOP\xb7\xee\xc4\xdf-I\"\xa3\x1dA\x03\x10\x8fP_b\r8h]\x93l\xd3d>\xa6\x01\x1c\tkrg\r\xbc\xe81\xe3\xc0\x13\xee\u05a7s\xe1{d\x8a\xe4\x00Q\xbb\xb4 \xf0C\"\u0390'\xa29@\x03ml\xec8\x1eT\xa1\xc9e\x00\x8813\x00\x84\xda\v\xe0\xf4B\xd0s=\x9b\xa7Eb\x9a\x13i\xc1\xa9d\xca\x1a\U000281d7\x01\xa0\xb8f\x80Rl|\x1f\x99x\x9f\xc0\xbb\x16\xa9\x9b!G\xb8\x01\xc0\xa8\x1b\xd3IUt\xe3\xe8\xbf\xc3\xfe\xaeK\x1b\x06\x9aB\xe7hO\xa0V\xe7a\xb5\xb1\xc0\x94\x9cxY\x83\xdf\xcb\a\xab\xbc\x9fN\x9fK\xe7LD\x12\r\xb2\u0478\xf2\xc1k\xe0\x96\f\xa81\xb4\xf3\x9643AM\xb4\x80\x0f9k\xe8\x92\x00\xba\xd6xt\xf8\xae\xa0d<\x88\x02`!\x17`\xd9\xe0\xe9<\x8d\u07b8\r\x127\xf2+\x0e\x1e\v\xff\x10\a\xb8\xff\u0478@\xdeg\x92\xe90\x80\x99\x1b,\r:\x9f\xb33\xb5u\f\xc0\xa7\xf8PLXt\x8e\\\v\xec\xda\x05\x845\x94\x93\xdc\xe0i\xe0\xf2b\xba\r\xb0\xa1\xc36 \xdd7B\xd0\x00D\x1bH\x9b\xc3q E\x1ao\xc2\xd5\x1a\x01F\xd44\xf0\xb8\xd2\xe2`\x8a\rBu\xd0\xe3\x82\xd6\xe6\xa4\xd9\x18\xea\x9b\x1e\x16\xac\x03cs\f&\xb7\x00\x00\x00\x00IEND\xaeB`\x82")

func spritesPngBytes() ([]byte, error) {
	return _spritesPng, nil
}

func spritesPng() (*asset, error) {
	bytes, err := spritesPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sprites.png", size: 438, mode: os.FileMode(436), modTime: time.Unix(1570700063, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"sprites.png": spritesPng,
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
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"sprites.png": &bintree{spritesPng, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
