// Code generated by go-bindata.
// sources:
// config/config.go
// config/local.local
// config/local.memory
// DO NOT EDIT!

package config

import (
	"bytes"
	"compress/gzip"
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
	bytes []byte
	info  os.FileInfo
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

var _configDs_store = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\xd8\x31\x0a\x02\x31\x10\x85\xe1\x37\x31\x45\xc0\x26\xa5\x65\x1a\x0f\xe0\x0d\xc2\xb2\x9e\xc0\x0b\x58\x78\x05\xfb\x1c\x5d\x96\x79\x60\x60\xd5\x4e\x8c\xcb\xfb\x40\xfe\x05\x37\x2a\x16\x31\x23\x00\x9b\xee\xb7\x13\x90\x01\x24\x78\x71\xc4\x4b\x89\x8f\x95\xd0\x5d\x1b\x5f\x43\x44\x44\x44\xc6\x66\x9e\xb4\xff\xf5\x07\x11\x91\xe1\x2c\xfb\x43\x61\x2b\xdb\xbc\xc6\xe7\x03\x1b\xbb\x35\x99\x2d\x6c\x65\x9b\xd7\x78\x5f\x60\x23\x9b\xd8\xcc\x16\xb6\xb2\xcd\xcb\x4d\xcb\x38\x7c\x18\xdf\xd9\x38\xa1\x18\xa7\x10\x2b\x6c\xfd\xce\x77\x23\xf2\xef\x76\x9e\xbc\xfc\xfe\x9f\xdf\xcf\xff\x22\xb2\x61\x16\xe7\xcb\x3c\x3d\x07\x82\xf5\x0d\x00\xae\xdd\xf5\xa7\x43\x40\xf0\x3f\x0b\x0f\xdd\x5a\x1d\x04\x44\x06\xf3\x08\x00\x00\xff\xff\x6a\x00\x88\x6d\x04\x18\x00\x00")

func configDs_storeBytes() ([]byte, error) {
	return bindataRead(
		_configDs_store,
		"config/.DS_Store",
	)
}

func configDs_store() (*asset, error) {
	bytes, err := configDs_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/.DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1493128719, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configConfigGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func configConfigGoBytes() ([]byte, error) {
	return bindataRead(
		_configConfigGo,
		"config/config.go",
	)
}

func configConfigGo() (*asset, error) {
	bytes, err := configConfigGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/config.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1495635133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configLocalLocal = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x90\x41\x6b\x32\x31\x10\x86\xef\xfb\x2b\x42\xce\xf2\x7d\xd6\xd2\x52\xf6\xaa\xbd\x94\x2e\x8a\x2e\xf4\x3c\x66\x27\x31\x98\xcd\x2c\x93\x89\x54\x8a\xff\xbd\xc4\xba\x68\xb1\xb9\x65\x9e\x87\x99\x77\xe6\xab\x52\x4a\xcf\x43\x4e\x82\xac\x6b\x55\xbe\x4a\xe9\xf6\x38\xa0\xae\x95\x0e\x64\x20\xe8\x4a\xa9\xd3\xa4\x78\x1f\xc4\x7b\xe4\x74\xef\xf1\x60\xf4\xe4\xa7\xb4\xa2\x10\x7c\x74\x2b\x64\x4f\x5d\x61\xb3\xa7\x69\x9f\x46\xfa\x1a\x2d\xb1\xc1\x16\xd2\xbe\xf5\x3d\x52\x16\x5d\x2b\xe1\x8c\x17\xfe\x1b\xe8\xc7\x69\x7f\x9d\xbe\x31\x3b\xec\x72\x40\x9e\x53\xb4\xde\xdd\xa7\x48\x02\x82\x36\x87\x71\x58\x03\x9f\x6b\x14\xf6\x98\x56\xc8\xa5\xb3\x56\xb5\x9a\x5d\xe0\x02\xb7\xd9\x35\xd4\x61\x29\x5a\x08\x69\x8c\xb0\x46\x43\x07\xe4\x37\xda\xa6\x65\xdc\x08\xb0\xe4\xa1\x38\x37\x29\x17\x68\x21\x07\xb9\x09\xdb\xa4\xa2\x3c\xbc\x4c\xcb\x1b\x1b\xe5\x18\x91\x97\x07\xe4\x1d\x42\xd7\x94\xb3\x3d\x9f\xf9\x75\x23\x70\xf0\x4e\x7f\x6c\x62\x7d\xc0\x71\x8b\x85\x67\x34\x42\x7c\x2c\xe0\x5f\x32\x44\xd2\x81\xc0\xff\xe2\x24\x70\x10\xc8\x9d\x6f\x54\x9d\xaa\xef\x00\x00\x00\xff\xff\x43\xc9\xed\x70\xcd\x01\x00\x00")

func configLocalLocalBytes() ([]byte, error) {
	return bindataRead(
		_configLocalLocal,
		"config/local.local",
	)
}

func configLocalLocal() (*asset, error) {
	bytes, err := configLocalLocalBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/local.local", size: 461, mode: os.FileMode(420), modTime: time.Unix(1493393396, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configLocalMemory = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x90\xc1\x6a\xc3\x30\x10\x44\xef\xfe\x8a\x65\xcf\x3d\x24\x3d\x94\xa2\xab\x73\x2a\x35\x29\x89\x7f\x40\x89\xc7\x8e\xa9\x2c\x85\xd5\x6e\x68\x28\xf9\xf7\xa2\xd6\xa6\xd0\x46\x37\xcd\x3c\x1e\xcc\x7e\x56\x44\x5c\x07\xcb\x0a\x61\x47\xe5\x4b\xc4\xed\xf5\x0c\x76\xc4\x13\xa6\x24\x57\x7e\xf8\x49\xeb\x64\x51\xd9\xd1\x7a\x55\x11\xdd\x4a\xc8\xfb\xe3\x09\x9d\x05\x48\x9d\x62\x3f\x0e\xff\x0d\x59\xbd\xa2\xb7\xb0\x38\x1a\xff\xb1\x83\xca\x88\xfc\x06\x69\x7d\x7e\x67\x72\xf4\x38\x97\x1b\x1c\x6c\x68\x52\x87\x12\xf6\x3e\x64\xcc\xc5\x0e\xc7\x74\x81\xbc\xa4\x43\xde\xc6\xbd\x7a\x51\x3b\xff\x65\x36\xe8\xbd\x05\x2d\xce\x76\x9c\x90\x4c\x9b\x5c\x98\xf5\xf3\xaa\xbc\xc5\x64\x31\x42\xb6\x17\xc8\x09\xbe\x6b\x32\x3b\x7a\xfa\xee\x7f\x27\xf9\xc1\xbf\xa6\x3b\x53\xe6\x63\x14\xb0\xba\x7d\x05\x00\x00\xff\xff\x0a\xdb\x3f\x04\x38\x01\x00\x00")

func configLocalMemoryBytes() ([]byte, error) {
	return bindataRead(
		_configLocalMemory,
		"config/local.memory",
	)
}

func configLocalMemory() (*asset, error) {
	bytes, err := configLocalMemoryBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/local.memory", size: 312, mode: os.FileMode(420), modTime: time.Unix(1489433454, 0)}
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
	"config/config.go":    configConfigGo,
	"config/local.local":  configLocalLocal,
	"config/local.memory": configLocalMemory,
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
	"config": &bintree{nil, map[string]*bintree{
		"config.go":    &bintree{configConfigGo, map[string]*bintree{}},
		"local.local":  &bintree{configLocalLocal, map[string]*bintree{}},
		"local.memory": &bintree{configLocalMemory, map[string]*bintree{}},
	}},
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
