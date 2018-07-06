// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/model.tpl
// templates/routes.json.tpl
// templates/routes.tpl
// templates/view_create.tpl
// templates/view_list.tpl
// templates/view_update.tpl

package cmd_gen

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

var _templatesModelTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x94\x51\x8b\xd4\x30\x10\xc7\x9f\x9b\x4f\x31\x2e\x28\xad\xec\xc5\x77\xa5\x0f\xe2\x21\x2c\x2b\xf7\xe0\x9d\xbe\x88\x0f\xb9\x76\xba\x04\xd3\xb4\x9b\xa4\x42\x19\xe6\xbb\x4b\xd3\x5a\xbb\x72\x3d\xce\x65\xf1\xde\x96\xd9\xc9\xfc\x7f\xf9\x4d\xdb\x56\x15\x3f\xd4\x01\xa1\x6e\x4a\x34\x5e\x08\x5d\xb7\x8d\x0b\x90\x8a\x64\x53\xaa\xa0\xee\x95\xc7\x37\xfe\x68\x36\x22\x13\x22\xf4\x2d\x02\x91\xbc\x51\x35\x32\x83\x0f\xae\x2b\x02\x10\x10\x39\x65\x0f\x08\xf2\xa3\x46\x53\x7a\x66\x22\x79\x1b\xff\x8c\x05\xe6\xe1\xd0\x5d\xdf\x22\xf3\x3b\x22\xb4\x43\x85\x85\xa8\x3a\x5b\x40\x5a\xc3\xeb\x79\x64\x06\x77\xea\xde\x60\x9a\x0d\xb3\xb5\x3d\x00\x81\xc3\xd0\x39\x0b\x44\xad\xd3\x36\x54\xb0\x79\x79\xdc\x80\x8c\x6d\xeb\x53\x76\xd6\xa3\x0b\x69\x06\x69\xd1\x98\xae\xb6\x1e\xbe\x7d\x1f\x27\x6e\xe1\xa7\x32\x1d\x0e\x05\x6d\x03\xba\x4a\x15\x48\x9c\x01\x89\xe4\x77\x6b\x3e\x37\x93\x48\x12\xa2\x2b\x38\xb9\x1d\x5c\x31\xc7\x3a\xe8\x0a\x6c\x13\x40\xde\x34\x63\x1e\x0c\x17\x87\x13\xce\x0f\x71\x26\x30\x6f\xa7\x7b\x4f\x13\xd1\x96\xe3\x1c\x16\xc9\x04\x94\x9f\x22\xfd\x7b\x76\x2d\xff\xd6\xfe\x48\xe8\xa8\x55\xb0\x10\x44\xba\x02\x79\x8d\x95\xea\x4c\xd8\x63\x1f\x5b\x1e\x92\xfa\xbe\x0a\xe8\x26\xb3\x0e\x7d\x67\x02\xf8\xa3\x91\x9f\xe3\xcf\x68\x50\x97\x5b\x40\xe7\xe0\x6d\x0e\x63\x83\xfc\xa4\x7c\x18\x8f\xec\xca\x34\x13\x89\xae\x62\xc3\x8b\x1c\xac\x36\xc3\x91\xa4\x55\x56\x17\x29\x3a\x97\x45\xae\x5a\xee\xae\x21\x07\x5d\x0a\x16\x13\x33\xaf\x2c\xf9\x4b\x5b\xaa\x80\xff\x6f\xc9\x63\xde\xf3\x2c\x79\xce\x3e\x73\xc9\x0f\xf9\xbb\x45\x83\xc5\xca\x4b\xa2\x6d\x68\x2e\x6b\x6f\x4c\x3b\xd3\x5e\xc4\x39\xd7\xdd\x9c\xfc\xea\x82\xf2\xf6\xd8\x8f\xe0\x3e\xcd\x66\x19\x7f\x3e\x57\x0b\x3d\xeb\x8c\xc3\x7b\xb7\xc7\xfe\x49\x4a\x96\x74\xc3\x67\x6f\x85\xe9\x6b\x7c\xcc\x22\xd2\x42\xd5\x92\xeb\xd4\xe0\x13\xe0\x06\x69\xb0\xb4\xf6\x18\xd7\xaf\x00\x00\x00\xff\xff\x1d\x64\x7e\x1e\x4e\x06\x00\x00")

func templatesModelTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesModelTpl,
		"templates/model.tpl",
	)
}

func templatesModelTpl() (*asset, error) {
	bytes, err := templatesModelTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/model.tpl", size: 1614, mode: os.FileMode(420), modTime: time.Unix(1530470276, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf6, 0x4a, 0xef, 0x83, 0x59, 0x80, 0xaa, 0xa4, 0x6f, 0xe7, 0x8d, 0x78, 0x68, 0x93, 0x5e, 0xd, 0x98, 0xbb, 0x25, 0x66, 0x60, 0xb9, 0x23, 0xe4, 0x89, 0x4d, 0x8f, 0xf6, 0xec, 0x69, 0xbc, 0xd0}}
	return a, nil
}

var _templatesRoutesJsonTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x93\x4f\x6f\xda\x3e\x18\xc7\xcf\xf6\xab\xf0\x2f\x87\xca\xae\x90\x53\xfd\xd4\xd3\xa6\x1e\xd6\x35\xec\x8f\x3a\x3a\x01\xdd\x8e\x93\x89\x1f\xc0\x9d\x63\x87\xc7\x8e\x32\x14\xe5\xbd\x4f\x01\xca\xba\x51\xb5\x05\x75\x3b\x6c\x37\x8c\x65\x3f\xdf\x7c\xbe\x1f\x97\x2a\xff\xaa\x66\xc0\x0a\x65\x1c\xa5\xa6\x28\x3d\x46\xc6\x29\x49\xa6\x45\x4c\x28\x49\x1c\xc4\x74\x1e\x63\xd9\xfd\x36\x3e\x35\xbe\x8a\xc6\x76\x0b\xad\xa2\x9a\xa8\x00\x69\x58\xac\xd6\xe0\x72\xaf\x8d\x9b\xa5\x37\xc1\xbb\x84\x52\x92\xcc\x4c\x9c\x57\x13\x99\xfb\x22\x2d\xf2\xdc\x5b\x7b\x83\xa9\x8a\x51\xe5\x73\x48\xa8\xa0\x74\x5a\xb9\x9c\xf1\x9c\x1d\x37\x8d\x7c\xed\x5d\x84\x6f\x71\xbc\x2c\xa1\x6d\x05\x7b\x93\x8d\xbf\x34\x8d\xfc\xe0\x35\x58\x39\x50\x05\xb4\xed\xa5\x09\x91\xd7\xac\xcb\x22\x87\x10\x4a\xef\x02\x7c\x46\x13\x01\x7b\x0c\xd9\xf1\xe6\xff\x45\x05\x21\x0a\xd6\x50\xa2\xac\xed\x31\x40\x64\x2f\xce\x58\x2e\x2f\xce\xb9\x90\xaf\xac\xe5\xdd\x50\x2e\xd8\x26\x87\x1c\x45\x8f\x6a\x62\xa1\x61\x08\xb1\x42\xc7\x1c\xd4\xbc\xe8\xc6\x06\xf9\x4b\x00\xc1\x5a\x41\x89\x99\xae\x2e\xfd\xef\x8c\x39\x63\xd9\xd1\xd1\xed\x2a\x2c\xac\xcc\x10\x07\x7e\xe8\xeb\xd0\xcd\x27\xb7\x23\x32\x44\x8f\x7d\x15\x95\xe5\x80\x28\x28\x69\x29\xdd\x6e\x0e\xc1\x69\xc0\xf7\xa3\xab\x01\xaf\x7b\x4c\x59\x2b\x68\xbb\x2f\x99\x7d\xa8\x18\xdd\xf1\x40\xd9\xf7\x58\x7c\x52\xb6\x02\x9e\x18\x9d\x08\x4a\xa2\xc2\x19\xc4\x6e\xf3\x01\x02\xdb\xcf\xff\xc1\xb4\x6f\x9c\xe6\xeb\xc3\x3d\x66\xb4\x78\x79\x17\x4f\x87\x61\x73\xe2\xec\x3e\x44\x3f\x33\xe2\xa7\x27\xa7\x82\x12\xd2\x52\x42\x0e\xe3\xb7\xce\xf1\x18\xc2\x8f\x57\xa3\x1d\x86\x03\xa8\xf7\xc1\x38\xf1\x7a\xb9\xb5\x6b\xfd\x26\xe4\x10\x94\xee\x0c\x43\x79\xee\xf5\x72\x47\x95\x47\x95\x78\x62\x03\x77\x3b\xe8\x9e\x9a\xbc\x76\x85\xc2\x30\x57\x96\xaf\x53\x6d\x20\xec\x14\xf1\xe0\xf0\x9d\x5e\xdf\xb9\x00\x18\xf9\x41\x97\xd5\x72\xc5\xee\x2d\x28\x0d\xc8\xff\x3f\x39\x39\xa4\x91\x7f\xd1\xea\xe7\xd7\x8a\xfc\x0e\x5f\xc8\x7d\xc2\x5c\x97\x5a\x45\xf8\x43\xc2\x5c\x64\x97\xd9\x38\xfb\x0b\x95\x19\x55\x79\x0e\x21\xf0\xa7\x1a\xb3\x13\xe3\x02\x2c\x3c\x67\x0d\xdf\x03\x00\x00\xff\xff\x9b\x63\x83\x81\x20\x08\x00\x00")

func templatesRoutesJsonTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesRoutesJsonTpl,
		"templates/routes.json.tpl",
	)
}

func templatesRoutesJsonTpl() (*asset, error) {
	bytes, err := templatesRoutesJsonTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/routes.json.tpl", size: 2080, mode: os.FileMode(420), modTime: time.Unix(1530712516, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x51, 0x64, 0xe3, 0x93, 0x9b, 0x23, 0x85, 0xa3, 0x6e, 0x7, 0x11, 0x8e, 0x2e, 0x34, 0xa5, 0xbc, 0x6e, 0x2e, 0x25, 0x2d, 0x3e, 0xc8, 0x11, 0xc1, 0xfc, 0x58, 0xb3, 0x15, 0xde, 0x7e, 0xf9, 0x63}}
	return a, nil
}

var _templatesRoutesTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x93\x5d\x6f\x13\x3d\x10\x85\xaf\xd7\xbf\xc2\xef\x4a\x6f\x65\x57\x2b\xa7\x42\xbd\x02\x45\x08\x48\x42\x11\x6d\xa9\x9a\x00\x97\xc8\xb5\x27\xa9\xc1\xfb\xd1\xf1\x2c\xa1\x8a\xfc\xdf\x91\x9b\x34\x29\x2d\x6a\xbb\x41\x42\x02\xee\xf6\x43\x9e\x73\xfc\x9c\x33\x8d\x36\x5f\xf4\x0c\x78\xa9\x5d\xc5\x98\x2b\x9b\x1a\x89\x0b\x96\xe5\xd3\x92\x72\x96\xe5\x15\x50\xef\x9c\xa8\x49\xcf\x56\x93\x3e\xd3\x01\x7a\xe1\xc2\xe7\x8c\x65\xf9\xcc\xd1\x79\x7b\xa6\x4c\x5d\xf6\x4a\x63\x6a\xef\x3f\x63\x4f\x13\x69\x73\x0e\x39\x93\x8c\x4d\xdb\xca\x70\x61\xf8\xee\x62\xa1\x5e\xd5\x15\xc1\x37\x9a\x5c\x36\x10\xa3\xe4\xaf\x87\x93\x4f\x8b\x85\x3a\xaa\x2d\x78\x75\xac\x4b\x88\xf1\x18\xe6\x62\xce\x93\x98\x3a\x85\xd0\xd4\x55\x80\x8f\xe8\x08\xb0\xe0\xc8\x77\x57\xdf\x2f\x5a\x08\x24\xf9\x82\x65\x2b\x21\x75\x0a\x95\x05\x3c\x98\x1c\x1d\x0a\x53\xf0\x7c\x3d\x74\xa2\xcf\x3c\xc4\xa8\x0c\x82\x26\xc8\x0b\x3e\x2f\x78\xe5\xbc\x64\xb1\xab\xb1\x43\x17\xa8\x93\x33\xef\x0b\x0e\x88\xfc\x69\x9f\x1b\x35\x78\x29\xa4\x7a\xe1\xbd\x48\xa2\x42\xf2\x6b\xdf\x63\xaa\x31\x59\x5c\x70\x04\x6a\xb1\xe2\x15\xcc\x45\x99\x64\x83\xba\x65\x40\xf2\x28\x59\xe6\xa6\x57\x43\xff\xeb\xa7\x6b\xf0\x9d\x9d\xeb\xb7\x70\xe1\xd5\x10\xf1\xb8\x3e\xad\xe7\x21\xe9\xaf\xd1\x0c\x11\x6b\x1c\x69\xd2\x5e\x00\xa2\x64\x59\x64\x8f\xe7\xe6\x5d\xa0\x25\x35\xed\xb7\xa1\xd6\x85\x98\xb3\x89\x15\xaa\x51\x8d\xe5\x07\xed\x5b\x10\xb9\xb3\xb9\x64\x19\x69\x9c\x01\xa5\x9f\xf7\xd0\x59\xa3\xd9\xf0\x1e\xb9\xca\x8a\xe5\xe1\x82\x3b\x2b\x9f\xdd\x44\x97\x10\xad\x4e\xf4\x7f\x86\xef\x47\x7e\x62\x7f\x6f\x5f\xb2\x2c\x8b\x2c\xcb\x1e\x60\xfb\x78\xb8\x6d\x63\xd7\xa5\xdc\x59\xfa\x7c\x08\xf1\xc9\xbb\xf1\xaf\xae\xcc\x06\x13\xaa\x13\x8d\x01\x12\x6f\x71\x17\xce\xbd\xfd\x79\x64\x24\x37\xd5\xae\xe7\x25\xb9\x01\x98\xda\xc2\x3a\x9b\x65\xe6\x1d\x2d\xdc\x89\xfb\x4d\x15\x00\x69\x35\xb4\xe3\xb0\x4d\x64\xd6\x21\x18\x3a\xd1\x33\x10\xd3\x92\xd4\xb8\x41\x57\xd1\x54\xe4\xbd\xdb\xf1\x3d\x77\xb6\xff\xff\xd7\xbc\xe0\x4b\xc1\xcd\xed\xdf\xc2\xe5\x98\xb0\x35\x34\x72\xe0\x6d\x8c\x72\x9b\x50\xff\xc1\xc5\xf9\x9d\x5d\x79\x7f\xb5\x7a\x7f\x60\x57\x06\xc3\xc3\xe1\x64\xf8\x17\xb6\x65\xdc\x1a\x03\x21\x88\xee\x65\x59\xd9\x18\x80\x87\x6d\x23\x9d\xab\x2b\x62\x07\xa0\x2d\xa0\x78\xb2\xb7\x27\x59\xfc\x1e\x00\x00\xff\xff\xa2\xdf\x80\x90\x15\x09\x00\x00")

func templatesRoutesTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesRoutesTpl,
		"templates/routes.tpl",
	)
}

func templatesRoutesTpl() (*asset, error) {
	bytes, err := templatesRoutesTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/routes.tpl", size: 2325, mode: os.FileMode(420), modTime: time.Unix(1530472668, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf9, 0xaf, 0x62, 0xf2, 0x23, 0xec, 0x3d, 0xa0, 0x1c, 0x31, 0x60, 0xaa, 0x16, 0xfd, 0xe2, 0x46, 0xaf, 0x3c, 0xeb, 0xd4, 0x3a, 0xde, 0xd7, 0xd5, 0x54, 0x96, 0xd7, 0xd4, 0x68, 0xc7, 0x7f, 0xde}}
	return a, nil
}

var _templatesView_createTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x91\x41\x6f\xa3\x30\x10\x85\xcf\xf0\x2b\x46\x73\x07\xfe\x80\xe1\xb2\xd2\x4a\x7b\xc9\x65\x7b\xb3\xac\xca\xe0\x21\xb5\x6a\xc6\xc8\x0c\x4d\x23\xc4\x7f\xaf\x42\x48\x93\xb4\xb7\xf1\xf3\x9b\xe7\x4f\xcf\xcb\xe2\xa8\xf7\x4c\x80\xe2\x25\x10\xae\xeb\x81\x4e\xa0\x75\x79\xb0\x03\x19\xb3\x2c\xc4\x6e\x5d\xf3\xbb\xad\x8d\xee\x8c\xeb\x9a\x67\xaa\x8f\x69\x00\xb6\x03\xd5\xc8\x74\x7a\xd5\xba\x7c\xb1\x6d\x20\x63\x10\x06\x92\xb7\xe8\x6a\x1c\xe3\x24\x08\xb6\x13\x1f\xb9\xc6\xea\x6e\xa9\x98\x4e\x08\x5d\xb0\xd3\x54\xe3\x38\x27\x2a\xb6\xb4\xef\xa9\x98\xc4\x76\xef\xe4\xb0\xc9\xb3\x4c\xf5\x9e\x82\x9b\x48\x2e\x87\x4c\x05\x3a\x12\xbb\xe6\x89\x53\x55\xbb\x7a\x71\x68\x9d\x2c\x1f\x09\xca\xbf\xdb\x9e\x31\x57\xb1\x00\xdf\x03\x47\x81\xf2\x10\xff\xf1\x44\x49\xae\x37\x2a\xd8\x96\x02\xf4\x31\xd5\xa8\x75\xf9\x5f\xd2\xdc\xc9\xb6\x6a\x0c\x36\x3f\x15\x55\x6d\xf6\x2b\x8a\xe7\x71\x16\x90\xf3\x48\x35\x0a\x7d\x0a\xee\x7d\xfc\x8a\x81\xaa\xb9\x41\x10\x3b\x28\x6e\x4c\xc4\x6e\x87\x78\x8c\x9a\xe6\x76\xf0\x82\xf0\x61\xc3\x4c\x35\xfe\x49\x64\x85\x9e\xeb\x6a\x67\x91\xc8\xf0\x30\x17\x63\xf2\x83\x4d\x67\xdc\x9e\x52\xd5\x43\x69\xaa\xba\x74\xda\xe4\xfb\x6f\x7e\x05\x00\x00\xff\xff\x03\x76\x6d\xf3\xf4\x01\x00\x00")

func templatesView_createTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesView_createTpl,
		"templates/view_create.tpl",
	)
}

func templatesView_createTpl() (*asset, error) {
	bytes, err := templatesView_createTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/view_create.tpl", size: 500, mode: os.FileMode(420), modTime: time.Unix(1530470276, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbc, 0x29, 0x8b, 0x52, 0x2c, 0xb9, 0x5e, 0xf, 0x16, 0x94, 0xd5, 0x34, 0x89, 0xd8, 0x40, 0x7a, 0xf7, 0x27, 0x9d, 0x66, 0xde, 0x89, 0x48, 0x12, 0x7, 0x2b, 0x4a, 0x1e, 0xc5, 0xb1, 0x3f, 0x54}}
	return a, nil
}

var _templatesView_listTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x31\x6f\xfb\x20\x10\xc5\x67\xf8\x14\x08\xfd\x57\x1b\x65\xfd\xeb\x4c\xb7\x2e\xad\xb2\xa4\x9b\xc5\x40\xc2\xa5\x46\x72\xed\x0a\x5f\x86\x08\xf1\xdd\x2b\xc0\x49\x94\x28\x52\xd5\xc9\x07\xf7\xce\xef\xf7\x4e\xc4\xe8\xf0\xe8\x27\x14\x92\x3c\x8d\x28\x53\xea\xfb\x76\x6b\xbf\xd0\x18\xf1\xee\x17\x8a\x11\x27\x97\x12\xbf\xe9\xf6\xb3\x3b\xcb\x94\x38\x0c\x1b\x7d\xaf\x05\x35\x6c\x34\x07\xb2\xfb\x11\xc5\x61\xb4\xcb\xd2\xc9\xef\x53\xc0\xa6\xde\xdc\xca\x66\x3f\x07\x87\x01\x9d\xd4\x9c\x01\x0d\x68\x9d\xe6\x8c\x01\x85\xfc\xe9\xfb\x46\x04\x3b\x7d\xa2\x68\x5f\x3d\x8e\x6e\x31\x86\xb3\x7a\xed\x8f\x62\x9a\x49\xb4\xdb\x79\x87\x23\x1e\xa8\x76\x80\x86\x4c\xb2\xa3\x70\x3a\x50\x19\x31\x06\x14\x0d\xfa\x32\x86\x93\x13\x4d\xd1\xae\xa7\x52\x83\x2a\x7e\x59\x59\x01\x80\x72\x36\xcd\x59\x8c\xab\x7f\x4a\xf7\x58\xff\x6a\x92\xff\x9d\x68\x3f\x72\x75\xfd\xe7\x1f\x79\x4b\xdc\xab\xa4\x7d\xc3\x73\xe6\x03\x2b\x86\x80\xc7\x4e\xaa\xbe\xaf\x4e\xc6\xbc\x78\xd7\xc5\xd8\x3e\xc6\x4b\x49\xea\xfb\x64\x8c\x3d\x95\x3d\xf5\x51\xf6\x71\x18\xd4\xca\xf4\xeb\xba\x2e\x0f\x22\x9f\xeb\xba\x40\x15\x56\xcd\xd7\xd6\x4f\x00\x00\x00\xff\xff\x00\x8b\xcc\x44\x53\x02\x00\x00")

func templatesView_listTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesView_listTpl,
		"templates/view_list.tpl",
	)
}

func templatesView_listTpl() (*asset, error) {
	bytes, err := templatesView_listTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/view_list.tpl", size: 595, mode: os.FileMode(420), modTime: time.Unix(1530470276, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x20, 0x7d, 0x80, 0xba, 0xcc, 0xd9, 0x52, 0x13, 0x48, 0x2f, 0x14, 0x46, 0x8, 0x53, 0x79, 0xe5, 0x29, 0x1e, 0x6e, 0xeb, 0x32, 0x1b, 0x1, 0xf7, 0xad, 0x95, 0x87, 0x14, 0x3c, 0xa5, 0xd9, 0x5}}
	return a, nil
}

var _templatesView_updateTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\xc1\xae\x9c\x30\x0c\x3c\xef\x7e\x85\xe5\x3b\xf0\x03\x09\x3d\xb5\x97\x4a\xef\xd2\xf6\x14\x45\x55\x20\xe6\x35\x6a\x48\x50\x62\xaa\x22\x94\x7f\xaf\x80\x7d\xdd\xe5\xad\x54\xa9\x37\x33\x4c\x3c\xe3\xb1\xd7\xd5\xd2\xe0\x02\x01\xb2\x63\x4f\x58\xca\x47\xeb\x18\x94\xaa\x5f\xcc\x48\x5a\xaf\x2b\x05\x5b\xca\xf5\xce\xeb\xa2\x5d\xb0\x94\xeb\x45\x0c\x31\x8d\x10\xcc\x48\x12\xc9\x3a\xfe\xae\x54\xfd\xd5\x74\x9e\xb4\x46\x18\x89\x7f\x44\x2b\x71\x8a\x99\x11\x4c\xcf\x2e\x06\x89\xcd\x9d\xf2\xc1\x59\xb9\xae\xb5\x52\xf5\x67\x5a\xbe\x70\x9a\x7b\xfe\xe4\xc8\x5b\xad\x4b\x41\xe8\xbd\xc9\x59\xe2\x34\x27\xaa\x76\x99\xbf\x55\x95\xd9\xf4\x3f\xc9\x62\x7b\xbd\x5c\xc4\xb0\x3d\xc9\xc4\xdb\xc7\x45\x78\x7a\xa5\x60\xdb\xf3\x04\xa2\xb9\xc1\x1b\x45\xa9\x64\xc2\x2b\x41\xbd\x6b\x65\xad\x0f\xb0\x02\x37\x40\x88\x0c\xf5\x4b\xfc\x36\x59\xc3\x74\xfc\x11\xde\x74\xe4\x61\x88\x49\xa2\x52\xf5\xc9\x26\xb6\xef\x11\xd1\xec\xf4\xc3\x8b\x0b\xd3\xcc\xc0\xcb\x44\x12\x99\x7e\x33\xde\x92\x7a\x6a\x03\xbf\x8c\x9f\x49\xe2\x11\xc6\x53\x12\x4a\xb9\x01\xb6\x8c\xb4\x4e\x64\x6c\x0c\x7e\x91\xc8\x69\x26\x54\x8a\x82\xd5\xba\x69\xdf\x66\x20\x9f\x4f\xbe\xff\xd7\xe0\xbf\x8d\xbc\x53\x7f\x90\x0d\x16\xaa\xb7\x24\x77\x4b\x4f\xfd\xf3\xdc\x8d\xee\xae\x70\x44\x7c\xde\x72\x37\x33\xc7\x00\x0f\x75\x35\x25\x37\x9a\xb4\x1c\x52\xa2\x79\xd8\xb5\x68\xb6\x53\x68\xaf\xb7\xeb\xfc\x13\x00\x00\xff\xff\x51\xb9\xea\xc1\xc5\x02\x00\x00")

func templatesView_updateTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesView_updateTpl,
		"templates/view_update.tpl",
	)
}

func templatesView_updateTpl() (*asset, error) {
	bytes, err := templatesView_updateTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/view_update.tpl", size: 709, mode: os.FileMode(420), modTime: time.Unix(1530470276, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdc, 0x40, 0x15, 0x3b, 0xc1, 0x1b, 0x57, 0xa9, 0xc2, 0x94, 0xc3, 0x96, 0xba, 0x53, 0x4b, 0xac, 0xb5, 0xd5, 0x1f, 0xee, 0xf0, 0xca, 0x60, 0x7, 0x2b, 0x25, 0xca, 0x9f, 0x81, 0x5d, 0xbf, 0x82}}
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
	"templates/model.tpl": templatesModelTpl,

	"templates/routes.json.tpl": templatesRoutesJsonTpl,

	"templates/routes.tpl": templatesRoutesTpl,

	"templates/view_create.tpl": templatesView_createTpl,

	"templates/view_list.tpl": templatesView_listTpl,

	"templates/view_update.tpl": templatesView_updateTpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"model.tpl":       &bintree{templatesModelTpl, map[string]*bintree{}},
		"routes.json.tpl": &bintree{templatesRoutesJsonTpl, map[string]*bintree{}},
		"routes.tpl":      &bintree{templatesRoutesTpl, map[string]*bintree{}},
		"view_create.tpl": &bintree{templatesView_createTpl, map[string]*bintree{}},
		"view_list.tpl":   &bintree{templatesView_listTpl, map[string]*bintree{}},
		"view_update.tpl": &bintree{templatesView_updateTpl, map[string]*bintree{}},
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
