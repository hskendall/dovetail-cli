// Code generated by go-bindata.
// sources:
// resources/import.template
// resources/shim.template
// resources/shim_support.template
// DO NOT EDIT!

package contract

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

var _resourcesImportTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\x4d\xcc\xcc\xe3\xe2\xca\xcc\x2d\xc8\x2f\x2a\x51\xd0\xe0\xe2\xac\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xcb\xd5\x51\x50\x29\x4a\x4d\x53\xb0\xb2\x55\xd0\x73\x2b\xcd\x4b\x2e\xc9\xcc\xcf\x2b\xae\xad\xe5\xe2\x8c\x57\x50\xaa\xae\x06\x49\xd5\xd6\x2a\x81\x34\xa4\xe6\xa5\xd4\xd6\x72\x69\x02\x02\x00\x00\xff\xff\x23\xe1\x09\x64\x52\x00\x00\x00")

func resourcesImportTemplateBytes() ([]byte, error) {
	return bindataRead(
		_resourcesImportTemplate,
		"resources/import.template",
	)
}

func resourcesImportTemplate() (*asset, error) {
	bytes, err := resourcesImportTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/import.template", size: 82, mode: os.FileMode(420), modTime: time.Unix(1540263332, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesShimTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x4d\x6f\xe3\x36\x10\x3d\x8b\xbf\x62\x2a\xc0\x85\x58\x38\x52\xcf\x2e\x74\x68\xdd\xdd\x36\xc0\x22\x2d\xea\xdc\xb6\x8b\x82\xa2\x86\x32\x11\x89\x14\x86\x23\x27\x41\xe0\xff\x5e\x90\x92\x1b\x6d\xeb\xa6\x1f\x37\x9b\x9e\xf7\xde\xbc\x37\x33\x1e\x95\x7e\x50\x1d\xc2\xa0\xac\x13\xc2\x0e\xa3\x27\x86\x42\x64\xb9\x19\x38\x17\xd9\xb1\x37\x90\x77\x96\x8f\x53\x53\x6a\x3f\x54\xf7\xb7\xdf\xed\x7f\x3a\x78\xc3\x8f\x8a\xb0\x6a\xfd\x09\x59\xd9\xfe\x46\x7b\xc7\x64\x9b\xaa\xe9\xbd\x7e\xd0\x47\x65\x5d\x75\x7c\x1e\x91\x7a\x6c\x3b\xa4\x1b\xa3\x1a\xb2\x3a\x17\x59\xcb\x4c\xb6\xeb\x90\xfe\x0b\x69\x18\x14\x71\xfa\xa2\x34\xdf\x74\xbe\xa2\xc9\xb1\x1d\xb0\x5a\xb8\x72\x91\xad\xd9\x56\xc2\xd5\x2c\x5c\x69\x4f\x58\xa5\xb6\xb4\x6f\xb1\x0a\x47\x3b\xfc\x33\x68\x24\xcf\x3e\x54\x23\x46\x05\x29\x04\x3f\x8f\x08\x2f\x2f\xe5\x7e\x7f\xa7\x06\x3c\x9f\xf7\x91\x6f\xef\x5b\x84\xc0\x34\x69\x86\x17\x91\xed\x97\x2e\x63\x45\x7c\xb6\xae\x13\xd9\xfd\x93\xbb\x5f\x5c\x03\xfc\x91\x40\x79\x88\xb6\x2e\x80\xa5\x40\x9c\x85\x30\x93\xd3\x50\x18\xf8\xea\x9a\x96\x84\x5b\x67\xb9\x08\x3c\x35\x10\x6d\x94\xfb\x8b\xab\x03\x4f\xcd\xad\x63\x24\xa3\x34\x4a\x88\x6d\x97\xbf\x60\x18\xbd\x0b\x18\x5b\x23\xe4\x89\xdc\x0c\x3a\x4c\x5a\x63\x08\x85\xb3\xbd\xfc\x57\x92\x27\xff\x80\xff\x47\xf4\xd8\x1b\xcd\x8e\x60\x57\xc3\xb1\x37\xe5\x1d\x3e\xfe\xf8\x1a\xf4\xfb\x94\x73\x4c\x40\x59\x87\x74\x40\x3a\x59\x3d\xcb\x6c\x21\x5f\xb5\x92\x4b\x91\xf5\x3e\x05\x38\x13\x45\xce\xf2\x07\xe4\x0f\xbe\xbb\xa0\xa4\xc8\x8c\xdb\x82\xa2\x2e\xc4\xa2\x48\x12\x2b\xde\x4f\x4e\xb3\xf5\xee\x5b\xd7\xfe\xac\x48\x0d\xc8\x48\xa1\x90\xe2\x42\x58\x7e\x8f\xcd\xd4\x15\x66\xe0\xf2\x30\x92\x75\x6c\x8a\xdc\xb8\x7a\x13\x66\xaa\x7a\x13\x7e\x75\xf9\x16\x2e\xd4\x1f\xbf\xfe\x24\x23\x98\x9f\xdc\x5b\x9e\xee\x49\xb9\xa0\x92\xf0\xe7\xae\x22\xc7\xcc\x86\x4e\x35\x3d\x1e\x50\x4f\x64\xf9\x59\x8a\x2c\xb0\xe2\x29\x6c\xa1\x55\xac\xb6\x80\x94\xbc\x9a\xf2\x75\x79\xca\x65\x0c\x8b\xff\x2d\xf0\x93\x93\x6f\xf9\x98\x19\xeb\xcd\x69\x26\x4d\x1f\x90\xa8\xde\x9c\x92\xa7\xbf\x08\x4a\x29\x32\x6b\x92\x74\x5d\x83\xb3\x7d\x9c\x60\x7c\x89\x15\xeb\xa7\xbf\xdf\xa5\x2c\x3b\x03\xf6\xf3\xec\xaf\x97\x7d\xfc\xd4\x3c\x33\x16\x91\xb2\x2c\xe6\x0b\x91\x32\x01\xc5\x1a\xbb\x86\xbe\x23\xf2\xf4\xa7\x11\x29\xdb\x63\x0b\xec\x01\x9f\x50\x4f\x8c\xc0\xaf\x91\xef\x20\x8e\x0f\x23\x6a\xb7\x78\x4d\x89\xcf\x0e\xcf\x71\xe1\x4f\x8a\xc0\xf4\xfe\x51\x6b\xa8\xaf\xde\xf4\xcb\xfa\x92\x77\x9f\x6f\xe3\x39\xc1\x97\x85\xac\xe7\x1e\xef\xf0\xf1\x43\x7a\x28\xd6\xa5\xbf\xc5\x63\x7d\xe7\x3a\xeb\x30\x97\xcb\x99\xc5\x3f\xda\x42\x46\x93\x4b\xd6\xbb\x85\xe2\xc0\x8a\xb8\xf8\x72\x6e\x4b\x7e\x93\x7e\xfb\xe2\x35\xf4\x65\xce\x57\xc2\x48\x4f\x71\x9c\xc4\xd6\x75\x90\x3c\x40\xbc\x4e\xd8\x84\x98\x45\xf4\x9f\x48\xcb\xb5\xa9\x75\x20\xbf\x07\x00\x00\xff\xff\xf4\xb8\xb2\x46\x05\x06\x00\x00")

func resourcesShimTemplateBytes() ([]byte, error) {
	return bindataRead(
		_resourcesShimTemplate,
		"resources/shim.template",
	)
}

func resourcesShimTemplate() (*asset, error) {
	bytes, err := resourcesShimTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/shim.template", size: 1541, mode: os.FileMode(420), modTime: time.Unix(1541470515, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesShim_supportTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\xc1\x6e\xe3\x36\x10\x3d\x8b\x5f\x31\x15\x16\x81\x94\xf5\x52\xe8\xb5\x80\x0f\x5b\x27\x5d\xa4\xc0\x66\x83\x24\x3d\x05\x41\x41\x4b\x23\x85\x5d\x89\x14\x48\x3a\x89\x21\xe8\xdf\x0b\x8a\xa4\x24\xdb\x49\xd6\x6d\x83\x9e\x0c\xd2\x9c\xc7\x37\x6f\xde\x8c\xd8\xb2\xfc\x3b\xab\x10\x1a\xc6\x05\x21\xbc\x69\xa5\x32\x90\x90\x28\x5e\x6f\x0d\xea\x98\x44\x31\x8a\x5c\x16\x5c\x54\xd9\x5f\x5a\x0a\xbb\x51\x36\xc6\xfe\x48\x1d\x13\x12\xc5\x15\x37\x0f\x9b\x35\xcd\x65\x93\xdd\x5e\xfc\xba\xfa\x76\x23\x4b\xf3\xc4\x14\x66\x65\x2d\x2b\xf9\xa9\xe6\xeb\x2c\x97\x0a\x33\xa3\x78\x55\xa1\x8a\x8f\x0b\x61\x6d\x7b\xfc\xc9\x4c\xa1\x96\x1b\x95\xe3\x91\x21\x03\x1f\x96\x1b\xfe\xc8\xcd\x36\x26\x51\x61\xfe\x54\x1b\x61\x78\x83\xf0\x7a\x78\x21\x1f\xd1\x30\x5e\x7f\xca\xa5\x30\x8a\xaf\x33\xdd\x30\x65\x86\x05\xcb\xcd\xa7\x4a\x66\x1e\xc3\x01\xfa\x74\xdf\x03\x70\x26\x5d\xd7\x29\x26\x2a\x84\x0f\x86\x2f\xe0\x83\x51\x58\xc2\x2f\x4b\xa0\xb7\xee\xc0\x35\x96\xba\xef\x87\xeb\xbb\x6e\xf8\x97\x5e\xb2\x06\xfb\x1e\xe2\xb0\xbe\xc6\xb2\xef\x07\x20\x14\x85\x3d\x3b\x22\x32\x8b\xc8\x02\xe2\x67\xaf\xce\x2e\x24\xdb\x83\x64\x2f\x41\xa6\x84\x3c\x32\x05\x28\xd8\xba\xc6\x1b\xcc\x37\x8a\x9b\x2d\x2c\xa1\xeb\xe8\xf9\xb0\x77\xfb\x2c\xc2\x76\xdf\x93\x72\x23\x72\xe0\x82\x9b\x24\x85\x8e\x44\x65\x63\xe8\x95\xe2\xc2\xd4\x22\x89\xed\x76\x9c\x92\xc8\xe2\xe5\x2d\x2c\xe1\xbc\x59\x63\x51\x60\x71\xa5\xe4\x23\x2f\x50\x25\x29\x21\x11\x6b\xdb\x95\x14\x25\xaf\x16\x80\x4a\x59\xfa\x79\x4b\xbf\xa0\xf9\xdc\xb6\x49\x4a\x22\x5e\x0e\xdb\x3f\x2d\x41\xf0\xda\x5e\xb1\x73\x07\x2a\x45\xcf\x95\x92\x2a\x49\x53\x12\x45\x52\xd3\xf3\x67\x6e\x92\x9f\x53\x12\xf5\x84\x44\x36\x72\x09\x0a\x2b\xae\x0d\x2a\xaf\x0a\x47\x9d\x8c\x97\xd2\x6b\x6f\x3e\xfd\xbe\x77\xf9\x9a\xea\xff\x9c\x42\x2d\x2d\x0e\xbd\x10\xa5\x4c\xe2\x95\x42\x66\x10\xce\xbc\xf7\xce\x45\xc5\x05\x52\x4a\xad\xca\x38\xea\x37\x35\x04\xbd\xc4\x27\x77\x68\xca\xf8\x5d\x09\x5d\x08\x6e\xf6\xe8\x40\xe0\x33\xe8\x81\xf4\x62\x30\xc7\xff\x70\x29\x9c\x49\x81\xb1\xb5\x54\x59\xcb\xa7\x3c\xa7\xb7\xcf\xc2\x97\x61\x20\xf2\x05\x8d\x5f\x26\x29\x4d\xa6\x26\xa7\x37\xb6\x73\x57\xbe\x73\xfd\x91\x94\xf4\xc4\x99\xfb\xb0\xa2\x36\x0f\xa9\x6c\x0a\x3b\xb4\xf6\x0e\x7a\x19\x66\x3d\x6a\x5b\xf4\x95\x9e\x57\xa8\x5b\x66\x1e\xec\x5f\xf1\x23\x8a\x42\xaa\xac\xeb\x3e\x8c\xfd\x19\x66\x08\xf5\x43\xdc\xfe\x0c\x13\x7e\xac\xf9\x67\xad\xd1\x24\x1e\xe6\x45\xb5\x15\x9a\x8d\x12\x76\xd7\x49\xda\xa0\x61\x05\x33\xcc\x46\x07\xf8\xaf\x7e\xaf\xeb\x43\xfd\xc2\x29\xfa\x87\x68\x98\xd2\x0f\xac\xfe\xfd\xe6\xdb\x65\x32\x12\x38\xea\xaa\x00\x7f\xed\x15\xfa\x8d\xe5\x46\xaa\x6d\x12\xcf\x73\x8c\x17\xe0\xc6\xd4\x34\xa5\xac\x7d\xc3\xd9\x93\xc0\xc4\xfa\x63\x9c\x7f\x3f\x2c\x41\x30\x45\xe0\x24\x78\x7d\x50\xd9\xd9\x5c\x50\xae\x47\x34\xdc\xdd\x9f\x86\xcf\x12\xf5\x7d\xf3\x83\xb2\x4f\x28\xae\xf2\x24\x62\xe3\x8e\xd5\xb8\x61\xdf\x31\x69\x58\x7b\xa7\x8d\xe2\xa2\xba\x77\x3f\x29\x89\xb2\xcc\x60\xd3\xd6\xcc\xe0\x2b\x5e\xd9\x9f\xe6\x13\xee\x5d\xbc\xa3\x57\x7c\x0f\x4b\xd8\x15\x75\xd2\xca\xc6\x05\x0d\x47\xdb\xd4\x92\x15\xa1\xe8\xc9\x84\x7b\x54\x55\xff\x2d\xef\xed\xe8\x83\xe4\xa5\x82\x87\xa8\x64\xa2\x7b\x90\xe6\xdc\x03\xc7\x94\xe3\x2d\x17\xec\x48\xa0\xb0\xd4\x70\x58\x24\x98\x17\xee\x74\xcc\xe3\xeb\x5c\x4e\xa9\xd2\x03\x6f\xcc\xb1\xfd\x3c\xc8\xbd\xc1\x5e\x70\xc4\x21\x70\x4a\xa2\x52\x2a\x10\xcd\x02\xbc\xa8\x4e\xe6\x81\xa6\xab\xc9\xc1\xdc\x88\xe1\xe3\x70\xf8\x23\xc4\xe3\x2b\x29\xcc\x8d\x23\x06\xc7\x0b\x85\x9f\xa9\xb6\x70\xe5\x1f\xea\x3f\x4c\x10\x0b\x71\x72\xc0\xdb\xce\x8f\xd9\x00\x79\x63\x78\xfc\x83\xeb\xbc\x72\x77\xa2\xb9\xf7\xb8\xd6\x87\x6f\x0a\xbe\x5f\x77\x0f\xb1\x08\x06\xc8\x32\x40\xff\x26\x71\x4d\x1e\x5e\x26\xc0\x9b\xb6\xc6\x06\x85\x61\x86\x4b\x01\xb2\x84\xdd\x03\xc4\x6c\x5b\x1c\x83\xc7\x30\x6d\xd4\x26\x37\xd0\x39\xf0\xfd\xf7\x0e\x38\x1a\x1a\x98\x00\xd6\xb6\x9e\x0e\x94\x4a\x36\xc0\x20\x97\x4d\xcb\x6b\x2c\xc0\xea\x03\x25\xaf\xd1\x19\xf4\xf0\xd5\x64\x83\xe9\x1e\xe1\x6e\x4c\xf2\x64\x9f\x55\xd7\xfb\x5c\xdd\xab\x6a\x64\x61\x1e\x70\x46\x63\xa3\x86\x54\xdd\x9d\x49\x01\xa7\xfb\x30\x29\x84\x57\x19\x24\xa7\x13\x83\x99\xf9\xdd\x63\xae\xac\x9c\x29\xc6\x13\x5d\xef\xbf\xc9\x36\xb1\x3d\xe3\xc5\x5d\x47\x57\x2b\xdf\xfd\x83\x4b\xdf\x9a\x3d\x93\x25\x7a\xf7\x11\xbc\x62\x4a\xe3\x00\x67\x57\x76\x7c\x9c\x61\x2e\xad\x48\x83\xc1\xec\xc6\x35\x32\xbb\x0e\xf7\xa7\xe3\xd3\x64\x8a\xa7\x2e\x28\x39\x71\xf4\x8f\x65\x30\x2a\xee\xc2\x82\xab\xfe\x0e\x00\x00\xff\xff\x77\x03\x39\x52\x94\x0d\x00\x00")

func resourcesShim_supportTemplateBytes() ([]byte, error) {
	return bindataRead(
		_resourcesShim_supportTemplate,
		"resources/shim_support.template",
	)
}

func resourcesShim_supportTemplate() (*asset, error) {
	bytes, err := resourcesShim_supportTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/shim_support.template", size: 3476, mode: os.FileMode(420), modTime: time.Unix(1541470505, 0)}
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
	"resources/import.template": resourcesImportTemplate,
	"resources/shim.template": resourcesShimTemplate,
	"resources/shim_support.template": resourcesShim_supportTemplate,
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
	"resources": &bintree{nil, map[string]*bintree{
		"import.template": &bintree{resourcesImportTemplate, map[string]*bintree{}},
		"shim.template": &bintree{resourcesShimTemplate, map[string]*bintree{}},
		"shim_support.template": &bintree{resourcesShim_supportTemplate, map[string]*bintree{}},
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
