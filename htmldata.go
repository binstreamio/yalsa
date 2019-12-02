// Code generated by go-bindata. DO NOT EDIT.
// sources:
// html/index.html (8.176kB)

package main

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

var _htmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x19\x7f\x6f\xda\x48\xf6\x7f\x3e\xc5\x1c\xab\xca\xa0\x33\x83\x0d\x84\x06\x07\x72\xd7\xb4\xe5\x1a\x65\xab\x56\xa1\x5a\xdd\x5e\x14\x45\x83\x3d\xc6\xd3\xda\x33\xd6\xcc\x40\xa0\x2b\xbe\xfb\x69\x6c\x03\xfe\x31\x26\x65\xab\xae\xaa\x68\xad\x28\xb6\xdf\x7b\xf3\x7e\xbf\xe7\x37\xc3\xf8\x1f\x6f\x3e\xbc\xfe\xf4\xfb\xc7\xb7\x20\x90\x51\x78\xd9\x18\xa7\xb7\xc6\x38\xc0\xc8\xbb\x6c\x00\x00\xc0\x38\xc2\x12\x01\x37\x40\x5c\x60\x39\x69\x2e\xa5\xdf\x39\x6f\x66\x28\xe1\x72\x12\x4b\x20\xb8\x3b\x69\x06\x52\xc6\xc2\xe9\x76\x5d\x8f\xc2\xcf\xc2\xc3\x21\x59\x71\x48\xb1\xec\xd2\x38\xea\x62\xb5\x5e\x8a\x7f\x0f\xe0\x00\x5a\x5d\x8f\x08\xb9\x03\xc1\x88\x28\xfa\xe6\xe5\xb8\x9b\x72\xbb\x6c\x8c\xbb\xa9\xf4\x46\x63\x3c\x67\xde\x26\x93\x35\xe7\xd9\x83\x47\x56\x80\x78\x93\xa6\x24\x32\xc4\x99\x26\x09\x22\xb0\x81\x90\x9b\x10\x4f\x9a\x12\xaf\x65\x07\x85\x64\x41\x1d\xe0\x62\x2a\x31\x6f\x5e\xfe\x4a\x56\x18\xcc\x24\xc7\x28\x02\xef\x19\x25\x92\xf1\x71\x37\xb0\x33\xa6\x5d\x8f\xac\xea\x04\x71\x8c\xc2\x07\x49\x22\xdc\xdc\xf1\x0f\x30\x59\x04\xd2\xe9\x9d\x59\xf1\xfa\x02\x78\x44\xc4\x21\xda\x38\x60\xc1\x89\x77\x91\xfc\xef\x48\x1c\xc5\x21\x92\xb8\xe3\xb2\x70\x19\x51\xe1\x80\xe1\x0b\xd0\x3b\x7f\x01\xec\xfc\xed\x22\xaf\xbe\xd2\x20\xa7\x87\x5e\x85\x07\x3f\x16\x7b\x35\xe6\xc8\xfd\xb2\xe0\x6c\x49\x3d\x25\x86\x71\x07\xfc\x82\xe7\xbe\x8d\xe7\x17\x4d\x1d\xa7\x6f\xe0\x3e\x27\x92\x23\x89\x7f\xa0\x84\x2f\x78\xe3\x73\x14\xe1\x93\xad\x28\x85\xa8\x1a\x26\x3f\x16\x0f\x01\x11\x92\xf1\x4d\x39\x50\x7d\x4b\x05\x6a\xcf\x6e\xdc\x4d\xf3\xaa\xb1\x4b\x60\xb9\x89\xb3\xac\xe9\x7e\x46\x2b\x94\x42\xb3\xd8\xac\x10\x07\xeb\x57\x6b\x22\xde\x20\x89\xc0\x04\xdc\xdd\x5f\xec\xe1\x01\x11\x53\x5b\x07\xb3\xac\x2a\xf4\x16\x4b\x4e\xb0\xa8\x22\xde\x20\x89\x33\xe8\x1e\xcc\x65\x85\xb1\x02\x95\xd9\x72\x79\x75\x5b\x01\xdd\x4c\x67\x15\x58\x59\x06\x8b\x25\x61\x14\x4c\xc0\x1f\xfb\x60\x49\xc6\x42\x49\x62\x27\x07\x4a\xc0\x9c\x2c\x16\x98\x3b\xc0\x40\x6b\x22\x0c\xb3\x80\x8c\x99\x20\x8a\x91\x03\xfc\x25\x75\x13\x96\xad\x58\xb6\x4b\x2c\xd4\xc5\xb1\x5c\x72\x0a\xee\x62\x79\x67\xdd\x9b\xc0\xb0\xad\x17\x46\xa6\xe1\xee\xda\xee\xdf\xb6\x07\x29\x49\x95\x57\x74\xc2\x6b\xe9\x00\xa3\x28\x61\xfa\x71\x06\xde\xa5\xf1\xcf\x69\x99\x67\xc5\x58\x38\x67\xeb\x32\x33\x1f\x23\xb9\xe4\x15\x19\xea\xf2\x90\x44\xff\x63\x2c\xd2\xe1\xd4\xb5\x51\x79\x71\x4d\x3d\xbc\x76\x80\x41\x19\xc5\x46\x85\x6c\x6b\x6a\x5c\xa1\x94\x54\x02\x35\x48\x81\x56\xf8\x95\xb8\x8e\xd0\x22\x21\x78\xda\x41\x49\x6e\x56\x1c\xb4\x89\xb1\x03\x0c\x17\x49\xbc\x28\xfa\x43\x5d\x73\x55\x6c\x88\x6f\xfe\x83\x62\x07\xf8\x28\x14\xb8\x88\x57\x66\x3b\xbb\xcc\xd4\xc9\xdc\xa4\x32\xef\x0a\xab\xaa\x2e\xca\xb4\x58\xa1\x70\x89\x0d\x8d\xa9\x2e\x52\xa1\x95\x7c\x89\xab\x48\x8a\x22\xb5\x76\xfa\x71\xa6\x59\x19\xa1\xb5\x03\xce\x2c\x0d\x82\x50\x07\x68\xe0\xdd\x6e\xd1\xe8\x3b\x0b\xf6\x4c\x60\xc1\xd1\x7d\xd1\xc1\xe6\x8f\xb7\x28\x6b\x03\x3a\xab\x6a\x94\x2f\x6a\x6e\x2b\xbd\xcf\xee\x6b\x12\xe3\xfe\xb0\xfc\x90\xbc\x77\xda\xe4\x20\x54\x10\xaf\x6c\x85\x90\x88\x4b\x07\x8c\x4a\x6a\x60\xea\x39\xc0\xb6\xac\x5c\x32\x94\xbc\x93\x2d\xd4\xaf\x2b\x02\x03\x44\xbd\x10\x5f\xbb\xaa\x6d\x18\xef\x6d\x0b\xbe\x34\x6d\x1b\x8e\x56\x1d\x1b\xf6\xdf\x8d\x60\x7f\x65\xc3\xbe\xdb\x19\xc0\x91\x69\xc1\x7e\xe7\x1c\x9e\x9b\x03\x38\x48\xee\x23\x38\x70\x2d\xf3\xcc\xec\xc3\x91\x39\x82\xb6\x99\xc1\xd4\x82\xc0\x86\xfd\x84\x83\x3b\x80\xa3\x8e\x05\xfb\x0a\xd9\x19\xc0\x41\x72\x1f\xc1\xc1\x6b\x7b\x04\xcf\x4c\x7b\x08\xfb\xa6\x7d\x06\x87\xa6\xdd\x83\x3d\x73\x2f\xfc\x2b\x78\x6f\xf7\x61\xdf\xec\x0d\xe0\xe0\xdd\x10\xbe\xfc\xad\xd7\x0f\x86\x70\xf8\x9b\x7a\xdf\xe1\xec\x11\x1c\x2a\x9c\x12\x33\x48\xb0\x0a\xf2\xd5\xd0\x59\x37\x23\x5f\x95\x8f\xcf\xad\x17\x7a\xb4\xfa\x3c\xe9\xba\x4a\xf6\x05\x34\x7e\xf1\x7d\x5f\x97\x60\x01\xf2\xd8\xe3\x55\xb8\xe4\x0e\xe8\xd7\xa1\x5f\x67\x3c\xf8\x62\x8e\x5a\x96\x09\xb2\x3f\x38\x6c\xd7\x72\xfc\xe0\xfb\x02\xcb\xff\x3a\xa0\x77\x9c\xe2\x77\x07\xf4\xea\x5a\x52\x2e\xf5\x42\xbc\x48\x22\xff\x87\xa6\xad\xdc\xa9\x9a\x7e\xb0\x85\x61\x82\xf4\xc9\xb2\x92\xe7\x19\x5b\x72\x17\x3f\xf0\xac\x3c\xee\x75\x7d\x47\x60\x85\x7b\xba\xf1\x1c\x9a\x47\x22\xa8\xae\x8c\xe7\x88\xeb\x3c\x12\x31\x26\x83\xba\x2a\x16\x9b\x68\xce\xc2\x5d\xbb\xd7\xf5\xef\x28\x0e\x09\x5d\xa8\xef\xe5\x0a\x73\xb4\xd0\x11\x11\x89\xa3\xda\x14\x00\xb9\x34\xe0\x8b\x79\xab\x77\x76\x66\x82\x97\x96\x09\xec\xbe\xdd\xfe\xa6\x0f\xcc\xbe\x7f\x4f\xed\xd3\xda\x5b\xde\x6f\x49\x58\xea\x3c\x17\x12\xbd\xf1\x3f\x9b\xeb\x7a\x96\x09\x86\x96\x09\x7a\x7f\xc2\x75\xb9\x6e\xa7\x21\xae\x75\x5e\x29\x8f\x7f\x42\x17\xe6\xa7\x16\xfb\x3b\x3d\x3c\x18\xa4\xb9\xd9\xeb\xf5\x4e\x74\x70\xf6\x1d\xac\xf3\x71\x5a\xff\xdb\xc2\xb8\x7a\xfb\x69\x5a\x9c\x58\x8f\xcc\x86\xb7\x18\x85\xe0\x13\x89\xb0\x9a\x0b\x8d\xba\x71\xf0\xe7\x9b\x77\x7f\xd8\x38\x97\xee\x01\x8e\x4c\x73\x5a\x91\xba\x49\xa7\x34\x49\x25\xc6\x29\xeb\x4a\x93\x44\x1c\x12\xf9\x2b\xa1\xda\x34\x12\x01\x7b\x4c\x73\xfc\x69\x8f\xec\x9b\x7e\x91\xcd\xb1\x1e\x5f\xdb\xdf\x95\xdc\x59\x56\x43\x1a\x4f\x05\x6c\x85\xf9\x2b\x4a\x22\x94\x05\xb9\x4a\x72\xb4\x3a\xbe\xb1\x6d\x6f\xf5\xf1\xc9\x75\xeb\x1c\x45\xbd\xd5\x95\x06\x5d\xdf\x59\xfe\x42\xbb\xfb\xfd\x93\xcd\xce\x8d\x95\xba\xb2\x4f\xb6\xb8\xa7\x56\xfd\x55\x7a\x92\xf1\x77\xe5\x3f\xdb\xca\xdf\x45\xf8\x39\x94\xfe\xd5\xed\xf1\x12\x48\x8f\x74\x4e\xad\x81\x9b\xdd\x59\xdb\xdf\x55\xf0\x6c\xab\xe0\x10\xe3\xe7\x50\x07\x37\xd3\x59\x5d\x21\xac\x10\x07\x7e\x2c\xb2\x13\x3e\x30\x01\xbb\x1f\x10\x08\x25\xb2\xe5\x31\x77\x19\x61\x2a\xe1\x02\xcb\xb7\x21\x56\x8f\x57\x9b\x6b\xaf\x65\xe4\x4e\x85\x8d\x76\xfb\x22\xcf\x4a\x15\x4a\x52\x27\xdf\xca\xab\x70\x0a\x5f\xe0\x36\xe7\xdf\xc1\x2c\x3b\x74\x2f\x30\xfc\xe2\x7f\x8f\x7a\xfb\x43\xf6\x84\xe7\x9e\x69\x88\x84\xfc\xc8\x08\x95\x60\x02\x9a\x56\x33\x77\xda\x8c\xa9\x87\x79\x76\xc0\x7d\xa8\x6c\x15\x95\x7c\x6d\xab\x77\xe8\x33\xfe\x16\xb9\x41\xeb\x40\x86\xca\xf5\xaf\x58\x52\xf6\x08\x26\x80\xe2\x47\xa0\xca\xae\x85\xa0\x52\x4c\x48\x14\xc5\x0f\x91\x68\x57\xc8\x65\x24\x24\x07\x13\x70\x47\xd9\xa3\x32\xec\x1d\x5b\x72\xd1\x6a\x9b\x20\x7b\x7f\x4f\xe8\x52\xe2\x3c\x64\x86\x5d\x46\x3d\xd1\x6a\xdf\xc3\xcf\x8c\xd0\x96\xe1\x18\xed\x46\x31\x95\xd5\x2e\x0e\xc6\x4b\x11\xb4\x10\x44\xab\x85\x8a\xda\x83\x5d\x12\x9e\x6d\xf5\x2a\x64\x96\x55\x25\xcc\xb6\x2c\x3b\x5a\x51\xd8\xe9\x55\xa8\x95\xdd\x29\x69\x62\x5c\x49\x39\x35\xf7\x3c\xa5\x5b\x32\x1b\x55\x35\xab\x50\x5d\xdd\xee\x88\xb2\x54\x2a\x13\xdc\x4c\x67\x3b\x8a\x7d\x6a\x94\x69\x8e\xab\x4b\x7c\xd0\x4a\x54\x0e\x31\x5d\xc8\x00\x5c\x82\xa1\xa5\x6d\xfb\x8a\x46\x04\xc4\x97\xad\xb6\x16\x69\x1d\xc1\x26\x2a\xd4\xa3\xaf\x6e\x8f\x20\x95\x89\x3a\xec\xb6\x68\x47\xbe\x04\x8a\x39\x79\xe8\x3b\xed\x46\xba\x26\x5b\xea\x63\xe9\x06\x2d\xa3\xbb\xb2\xbb\xb1\x5a\x29\xfe\xa5\x98\x3c\x24\xcf\x13\x03\xfc\xf3\xc0\xf3\x20\x19\xca\x00\xd3\x16\xc7\x22\x66\x54\x60\x30\xb9\x04\xbb\x67\xf8\x59\x30\xda\x6a\x97\x49\xbd\xa4\xf6\x2e\x4b\x1e\x3d\xd4\x65\x5a\x8b\x05\x64\xae\x83\x41\x81\xe5\x87\x64\x50\x68\xed\x77\xca\xed\xe2\xc7\xf6\xd0\xa1\x2a\xc4\x57\xb7\x25\xda\x5c\xf3\xa9\x10\xdf\x4c\x67\x25\xea\x43\x53\xae\x10\xe7\x28\xb7\xbb\x46\x24\xb0\xbc\xa6\x12\xf3\x15\x0a\x73\x3d\x24\x9f\x4b\x7f\xc6\xe1\x27\x3a\xfd\x29\xc7\x3f\xe9\xfc\xd3\x03\x70\x6a\x10\x4e\x0f\xc4\x36\x7b\xdf\x9a\xc0\xb6\x2c\xab\xdd\x38\xfc\xa4\xde\x18\x77\x93\x5f\xf6\xff\x1f\x00\x00\xff\xff\xfd\x7b\x7e\x5d\xf0\x1f\x00\x00")

func htmlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlIndexHtml,
		"html/index.html",
	)
}

func htmlIndexHtml() (*asset, error) {
	bytes, err := htmlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/index.html", size: 8176, mode: os.FileMode(0644), modTime: time.Unix(1575257209, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xeb, 0xf5, 0x18, 0x61, 0xff, 0x0, 0x1e, 0xf0, 0x4c, 0xb5, 0x16, 0xb7, 0xeb, 0x92, 0x11, 0xd2, 0x67, 0x5c, 0x59, 0x34, 0xe3, 0xa0, 0x24, 0xad, 0x5e, 0xb3, 0x36, 0xa0, 0x11, 0x4a, 0x49, 0x8f}}
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
	"html/index.html": htmlIndexHtml,
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
	"html": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{htmlIndexHtml, map[string]*bintree{}},
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
