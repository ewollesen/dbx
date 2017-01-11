// Code generated by go-bindata.
// sources:
// golang.create.tmpl
// golang.delete-all.tmpl
// golang.delete-world.tmpl
// golang.delete.tmpl
// golang.footer.tmpl
// golang.get-all.tmpl
// golang.get-count.tmpl
// golang.get-has.tmpl
// golang.get-last.tmpl
// golang.get-limit-offset.tmpl
// golang.get-paged.tmpl
// golang.get.tmpl
// golang.header.tmpl
// golang.misc.tmpl
// golang.update.tmpl
// DO NOT EDIT!

package templates

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

var _golangCreateTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x92\x41\x6b\xdc\x30\x10\x85\xcf\xd6\xaf\x98\x2e\x14\x6c\xd8\xe8\x07\x04\xf6\x50\x42\x0a\x81\xb0\x90\x6c\x4a\x8f\x46\x59\x8d\x8c\x5a\xaf\xe4\x8e\xc6\x75\x16\xa1\xff\x5e\xc6\x66\x83\x73\x68\x0e\x7b\x30\xd8\xb2\x78\xdf\x7b\x6f\x26\xe7\x1b\xb0\xe8\x7c\x40\xd8\x24\xdf\x05\xc3\x23\xe1\x06\x6e\x4a\x51\x77\x84\x86\x31\x67\xd0\x87\xd1\x39\xff\x06\xa5\xd4\x39\xc3\x60\xc8\x9c\x40\x7f\xa3\x2e\x41\x29\x0d\xac\xce\x9e\x91\x47\x0a\x50\xca\x16\x90\x48\x9e\x48\x8d\x12\x04\x06\x3b\x6b\xaa\x35\xef\x35\xda\xf3\x8c\xca\xaa\x92\x73\xef\x40\xef\x11\x6d\xda\xc7\x09\x4a\x51\x55\xdb\x86\x38\xc1\xed\x0e\xf6\x71\xaa\x1b\xfd\xe3\xe5\xae\x6e\xe4\xea\xbb\x9c\xbc\xfb\xe0\x39\xe0\x04\xfa\xbb\xc7\xde\x8a\x27\xa5\xaa\x63\x0c\x89\xa1\x6d\x13\x9f\x78\x27\x06\xc9\x07\x76\xb0\xf9\xfa\x67\x03\xfa\xf0\xf4\x38\xcb\xc7\xd7\x5f\xba\x8f\xdd\x81\x4f\x5c\x2f\x57\xb7\x90\x33\x18\xea\x56\x62\x8d\x5a\x28\x4e\x6a\x18\x86\x48\x9c\x96\x98\x3e\x74\xb0\xb2\xb0\x4a\xaf\x2a\x49\xbf\x03\xd1\xb7\xe4\xff\x22\xe9\xa7\x11\xe9\xfc\x1c\xa7\x4f\x38\xfa\x70\x34\x41\xca\x34\xd6\x52\x74\x50\xbb\xde\x30\x63\xb8\x08\x37\xb3\x99\xca\xbb\xb9\xdb\x2f\x3b\x08\xbe\x87\xac\xaa\x8a\x16\x6e\xf0\xfd\x16\x7e\x92\x19\xee\x89\x6a\x24\x6a\x54\x55\xd4\xe5\xe7\x05\xb7\x9a\x50\xf0\xfd\xd2\x65\x9f\x70\x29\xb3\x6d\x09\xd3\x32\xba\xdb\x0f\xee\xef\xdf\xf0\xf8\x1f\xe7\xd7\x78\x6a\xdb\xe1\xf7\x3b\x66\x86\xea\x47\x93\xf8\x21\x24\x24\x7e\xb0\xf5\xd5\x29\xc5\x72\x87\x2c\x62\xb2\xb6\x4b\x58\xfd\x72\x1e\x50\x76\x57\xb0\x1f\xd7\xa7\xa8\xd5\xc7\xbf\x00\x00\x00\xff\xff\x50\x3f\x0b\xb4\x0a\x03\x00\x00")

func golangCreateTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangCreateTmpl,
		"golang.create.tmpl",
	)
}

func golangCreateTmpl() (*asset, error) {
	bytes, err := golangCreateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.create.tmpl", size: 778, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangDeleteAllTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x90\x31\x6b\xf3\x40\x0c\x86\xe7\xd3\xaf\xd0\x67\xf8\xc0\x81\xc4\x74\x28\x1d\x0a\x1e\x02\xcd\xd6\xa5\xcd\xd0\xd1\x38\xb6\x2e\x5c\xb9\xdc\xa5\x3a\x5d\x9a\x22\xee\xbf\x17\x3b\xa1\xed\x56\x3a\x68\x92\xde\xe7\x79\x91\xea\x0a\x47\xb2\x2e\x10\x56\xc9\xed\x43\x2f\x99\xa9\xc2\x55\x29\xf0\x40\x9e\x84\x54\xb1\xd9\x66\x6b\xdd\x19\x4b\xa9\x55\xf1\xd8\x73\x7f\xc0\x66\xcd\xfb\x84\xa5\x2c\xb0\x1e\x62\x0e\x82\x2e\xc8\xdd\xed\x12\x89\x79\x9a\xc8\x0b\x98\xd0\x14\xc6\x99\x05\x3f\x3d\xbb\x38\x7e\xcc\x0a\x05\x33\xc4\x90\x04\xbb\x2e\xc9\x41\xda\x89\xce\x2e\x88\xc5\xea\xff\x5b\x85\xcd\xf6\xe9\x11\x4b\x01\x13\x77\xaf\x8d\x8f\xfb\xad\x1c\xa4\xbe\x9c\x2e\x51\x15\xad\x23\x3f\x9e\x7a\x9f\xe9\xbb\x0e\x80\xe9\x3a\xa6\x74\x69\x72\xdf\xe2\x94\x1d\xd9\x9d\x88\x9b\xcd\x99\x86\x5f\xf3\xc6\xd9\x39\xfa\xaf\xc5\xe0\x3c\x2a\x18\xc3\x24\x99\x03\xde\x2c\xf1\x85\xfb\xe3\x86\xb9\x26\xe6\x05\x98\x02\x53\xff\x1c\xe4\x22\x6b\x71\x36\x37\xcf\xf1\x3d\xad\xad\xa5\x41\x68\xac\xff\x0c\xbc\xee\xae\xdc\xe0\x3c\x14\x50\xfd\x7a\xe4\x67\x00\x00\x00\xff\xff\x15\x63\x97\xa8\xb0\x01\x00\x00")

func golangDeleteAllTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangDeleteAllTmpl,
		"golang.delete-all.tmpl",
	)
}

func golangDeleteAllTmpl() (*asset, error) {
	bytes, err := golangDeleteAllTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.delete-all.tmpl", size: 432, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangDeleteWorldTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x90\x31\x4f\xc3\x30\x10\x85\x67\xdf\xaf\x38\x22\x21\x25\xa2\xb5\x18\x10\x5b\x86\x4a\x74\x63\xa1\x0c\x8c\x55\x9a\x9c\x2b\x23\xeb\xdc\x9e\x9d\x02\xb2\xfc\xdf\x51\x1c\x52\x75\x65\xf0\x62\xf9\x7d\x9f\xdf\x4b\x69\x8d\x03\x19\xcb\x84\x55\xb0\x47\xee\xe2\x28\x54\xe1\x3a\x67\x78\x21\x47\x91\x36\xce\xd5\x0d\xd6\xbd\x1f\x39\xa2\xe5\xf8\xfc\xb4\x42\x12\x99\x8e\x97\x06\xa6\x3c\xf1\x50\x02\x70\x0b\x3b\xf8\xe1\xa7\x70\x12\xa8\x4b\x27\xb8\xdf\x0b\x05\x0c\x67\xa7\x77\x14\x46\x17\x97\xdb\x1b\xf0\x0c\x90\x8e\x8f\x84\xfa\xfd\xed\x35\x60\xce\xa0\x4a\x70\x76\xb6\xe8\x0f\x9f\x7a\x10\x7b\x21\xd1\xdb\x6f\xea\xeb\x94\xf0\x24\x96\xa3\xc1\xea\xfe\x5c\xa1\xc6\x9c\x1b\x50\xd6\x94\xe7\x77\x2d\xb2\x75\x98\x40\x29\xa1\x38\x0a\xe3\xe3\x0a\x3f\xa4\x3b\x6d\x45\x6a\x12\x69\x40\x65\x98\x04\xe5\x0f\x8b\xa2\xf8\xf4\xce\x7f\x85\x8d\x31\xd4\x47\x1a\xea\x7f\x22\xd5\xdc\xe9\xa1\x5d\xea\x5d\x57\x9a\x46\x5a\x82\x7f\x52\xb6\x0e\x32\xa4\x74\x5d\xf1\x37\x00\x00\xff\xff\xb2\x3f\x06\xc6\x92\x01\x00\x00")

func golangDeleteWorldTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangDeleteWorldTmpl,
		"golang.delete-world.tmpl",
	)
}

func golangDeleteWorldTmpl() (*asset, error) {
	bytes, err := golangDeleteWorldTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.delete-world.tmpl", size: 402, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangDeleteTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x90\xb1\x6a\xf3\x30\x14\x85\x67\xe9\x29\xce\x6f\xf8\xc1\x01\xc7\x74\x2e\xa4\x10\x68\xb6\x2e\x6d\x86\x8e\xc6\xb1\xae\x82\x8a\x22\xa5\x57\x72\x9a\x72\xd1\xbb\x17\x3b\x21\xed\x56\xe8\xa0\x49\xe7\x7c\xdf\xe1\x8a\x2c\x61\xc8\xba\x40\xa8\x92\xdb\x87\x3e\x8f\x4c\x15\x96\xa5\xe8\x47\xf2\x94\x49\x04\xed\x76\xb4\xd6\x9d\x51\x4a\x2d\x82\x63\xcf\xfd\x01\xed\x9a\xf7\x09\xa5\x2c\x50\x9b\x39\x67\xb0\x8b\xd1\x37\x20\xe6\xe9\x45\x5e\xe8\x89\x4d\xc1\xcc\x30\xfd\x53\xb4\x8b\xe6\x73\x76\x88\x56\x43\x0c\x29\xa3\xeb\x52\x3e\xe4\xd5\x84\x67\x17\xb2\x45\xf5\xff\xbd\x42\xbb\x7d\x7e\x42\x29\x5a\xc5\xdd\x5b\xeb\xe3\x7e\x9b\x0f\xb9\xbe\x44\x1b\x88\xc0\x3a\xf2\xe6\xd4\xfb\x91\xbe\xf7\x68\xad\xba\x8e\x29\x5d\x96\xdc\xaf\x30\x75\x0d\xbb\x13\x71\xbb\x39\xd3\xf0\x6b\x5f\x39\x3b\x57\xff\xad\x10\x9c\x87\x68\xa5\x98\xf2\xc8\x01\xb6\xf7\x89\x1a\xbc\x72\x7f\xdc\x30\xd7\xc4\xbc\xd0\xaa\xcc\xc2\x21\x8e\x21\xdf\x94\xf3\x80\xf6\x25\x7e\xa4\xb5\xb5\x34\x64\x32\xf5\x9f\xb8\xd7\xff\x2b\x1e\x0f\xb8\x6b\xa6\xae\x2e\x5a\xe4\x76\xd9\xaf\x00\x00\x00\xff\xff\x2c\x28\x22\xad\xc2\x01\x00\x00")

func golangDeleteTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangDeleteTmpl,
		"golang.delete.tmpl",
	)
}

func golangDeleteTmpl() (*asset, error) {
	bytes, err := golangDeleteTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.delete.tmpl", size: 450, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangFooterTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xce\xb1\x0a\xc2\x30\x10\x06\xe0\x39\xf7\x14\x37\xb6\x42\xf3\x10\x22\x6e\x4e\x66\x70\x8d\xc9\xd9\x16\x34\xd6\xe4\xc4\x48\xb8\x77\x97\x22\xa4\x4b\xb6\x83\xfb\xff\x8f\x1f\xf8\xbb\x10\x9a\xcb\x89\x78\x7a\xfa\x84\x73\x60\x8a\x37\xeb\x08\x0b\x94\x32\x60\xb4\x61\x24\xd4\xc7\x77\x70\x09\x45\x40\x95\x82\x7a\x3d\xd6\x27\x05\x8f\x83\x08\x08\xfc\x19\xce\x2d\x46\x55\xbd\x06\x0f\xfb\x66\xf0\xec\x26\x7a\xd8\xae\xc7\xc4\x71\x0e\x63\xab\xe9\xaf\xcd\x66\x05\x01\xd4\x27\xda\xc5\xe4\x8e\x33\xee\xd2\xeb\xae\x4d\xee\xb7\x61\x20\xf0\x0b\x00\x00\xff\xff\x6d\x27\x2d\x7a\xf2\x00\x00\x00")

func golangFooterTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangFooterTmpl,
		"golang.footer.tmpl",
	)
}

func golangFooterTmpl() (*asset, error) {
	bytes, err := golangFooterTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.footer.tmpl", size: 242, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetAllTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x92\xbf\x6a\xc3\x30\x10\xc6\x67\xe9\x29\xae\x86\x82\x0c\x89\x1f\x20\xc5\x43\x29\xa5\x4b\x29\xa4\x1e\x3a\x1a\x25\x3a\x19\x15\x45\x72\xcf\x72\xfe\x20\xf4\xee\x45\x76\x92\x76\x28\x81\x0e\x1e\x6c\x9f\xbe\xdf\xef\x3e\x3b\xc6\x25\x28\xd4\xc6\x21\x14\x83\xe9\x9c\x0c\x23\x61\x01\xcb\x94\xf8\x0b\x86\x18\xa1\x6a\x46\xad\xcd\x11\x52\x12\x31\x42\x2f\x49\xee\xa0\x7a\xa4\x6e\x80\x94\x4a\x10\xe4\x0f\x03\xc4\x08\x83\x35\x5b\xf4\x1a\xaa\x77\x7f\x80\x94\x16\x80\x44\xf9\xf2\x54\xf2\xcc\x40\xa7\xa6\x50\xfe\x1b\xb8\xf1\xea\x54\x40\x4a\x91\xb3\xad\x77\x43\x80\xb6\x1d\xc2\x2e\xd4\x19\x44\xc6\x05\x0d\xc5\xfd\x57\x01\x55\xb3\x7e\x85\x94\x38\xf3\x9b\xcf\xca\xfa\xae\x09\xbb\x20\xe6\xd1\x45\x66\x6b\x83\x56\xed\xa5\x1d\xf1\xc7\x8c\x73\xd6\xb6\x59\x6e\x36\x59\xd5\x90\x0f\x2b\x32\x7b\xa4\x6a\x3d\x22\x9d\x6e\x25\xe4\x00\x66\xf4\x74\xf4\xae\x06\x67\x2c\x44\xce\x18\x61\x18\xc9\xe5\xdb\x05\x7c\x90\xec\x9f\x89\x04\x12\x95\x9c\x25\xce\x14\x6a\x24\x98\xa1\xd5\x93\xf5\x03\x8a\x6c\xa1\xfd\xf5\xe1\x1b\x1e\x83\x28\xa7\xa4\x18\xc1\x38\x13\x1c\x1e\x2e\x8d\x71\xc6\x32\xad\xbe\x0c\x37\x5b\xe9\x72\xe3\x52\x29\xf2\x1a\x84\xb6\x32\x04\x74\xd3\x78\x39\x6d\xc8\xfe\x30\xbc\xa1\x98\x1d\xd9\xf4\xb9\x6a\x90\x7d\x8f\x4e\x89\xb9\x9f\x0c\xa1\xee\xe2\x31\x2f\x73\x4e\x5e\x5d\x75\x72\x4e\xf9\xf0\xaf\x42\xce\x2f\x67\x88\x33\x96\x27\x1e\xe3\xf5\x47\xf8\x0e\x00\x00\xff\xff\xa3\x31\x84\x54\x79\x02\x00\x00")

func golangGetAllTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetAllTmpl,
		"golang.get-all.tmpl",
	)
}

func golangGetAllTmpl() (*asset, error) {
	bytes, err := golangGetAllTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-all.tmpl", size: 633, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetCountTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x8f\xbd\x4e\xec\x30\x10\x85\x6b\xcf\x53\xcc\x8d\x74\x51\x22\x65\x2d\x0a\x44\x97\x02\x21\x3a\x9a\x25\x05\xe5\xca\x9b\xd8\x91\x51\x62\x2f\x13\x7b\x61\x35\x9a\x77\x47\x0e\xe2\xa7\xa3\x70\xe5\xf3\x9d\x39\x1f\xf3\x0e\x47\xeb\x7c\xb0\x58\xad\x7e\x0a\x26\x65\xb2\x15\xee\x44\xe0\x3e\xe6\x90\x98\x51\xf7\xd9\x39\xff\x8e\x22\x35\x33\x9e\x0c\x99\x05\xf5\x1d\x4d\x2b\x8a\x34\x58\x0f\x25\x86\x3e\xa4\xdb\x9b\x16\x2d\x51\x79\x91\x1a\x28\xcd\x36\x8c\x5b\x15\xfc\x3e\x73\x8c\xe3\xa5\x42\x11\x06\x35\xc4\xb0\x26\x3c\x1c\xd6\xb4\xa4\xae\x94\x93\x0f\xc9\x61\xf5\xff\xb5\x42\xdd\xef\x1f\x51\x04\x54\x3c\xbe\xe8\x39\x4e\x7d\x5a\x52\xfd\x19\x6d\x91\x19\x9d\xb7\xf3\x78\x36\x73\xb6\x3f\x6b\x00\x54\x59\xd0\x61\x61\x46\xf2\x67\x4b\x7a\x9f\x2d\x5d\x9e\xe2\xdb\x5f\xac\xee\x07\x13\xea\xab\x4d\xa7\x01\xe5\xdd\x26\xf3\xaf\xc3\xe0\x67\x64\x50\x8a\x6c\xca\x14\xf0\xba\xc5\x67\x32\xa7\x07\xa2\xda\x12\x35\xa0\x04\xe0\xeb\x6f\x83\xdb\x42\x80\x00\xf3\xb7\xff\x47\x00\x00\x00\xff\xff\x65\x8f\x08\xe6\x66\x01\x00\x00")

func golangGetCountTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetCountTmpl,
		"golang.get-count.tmpl",
	)
}

func golangGetCountTmpl() (*asset, error) {
	bytes, err := golangGetCountTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-count.tmpl", size: 358, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetHasTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x8f\xbf\x4e\x03\x31\x0c\xc6\xe7\xf8\x29\xcc\x49\xa0\x3b\xe9\x7a\x6f\xd0\x81\x01\x89\x81\xa5\xdc\xc0\x58\xa5\x3d\xa7\x0d\xca\x25\xc5\xc9\x15\x2a\xcb\xef\x8e\x52\xc4\x9f\x8d\xc1\x93\xfd\xfb\xf9\xfb\x44\x56\x38\x91\xf3\x91\xb0\xc9\xfe\x10\x6d\x59\x98\x1a\x5c\xa9\xc2\xa3\xcd\x22\x38\x8c\x8b\x73\xfe\x03\x55\x5b\x11\x3c\x59\xb6\x33\x0e\xf7\x7c\xc8\xa8\xda\x61\x7b\xb4\x19\x77\x29\x85\x1e\x89\xb9\x4e\xe2\x0e\xaa\x94\xe2\x74\xb5\xc0\xdf\x0f\xbb\x34\x5d\x1a\x54\x15\x30\xfb\x14\x73\xc1\xed\x36\x97\xb9\xac\xab\x99\x7d\x2c\x0e\x9b\xdb\xb7\x06\x87\x71\xf3\x84\xaa\x60\xd2\xee\x75\x08\xe9\x30\x96\xb9\xb4\x5f\xa7\x3d\x8a\xa0\xf3\x14\xa6\xb3\x0d\x0b\xfd\x46\x01\x30\x35\xc1\x1a\x2b\x33\xb1\x3f\x13\x0f\x9b\x85\xf8\xf2\x9c\xde\xff\x63\x87\x71\x6f\x63\x7b\x77\xb4\xb9\x03\xe3\xdd\xb5\xca\xcd\x1a\xa3\x0f\x28\x60\x0c\x53\x59\x38\xa2\xb3\x21\x53\x8f\x2f\x6c\x4f\x0f\xcc\x2d\x31\x77\x60\x14\xbe\xd7\x47\x9b\xfb\x8a\x80\x82\xc8\x4f\xfd\xcf\x00\x00\x00\xff\xff\xcc\xbc\x00\xa4\x60\x01\x00\x00")

func golangGetHasTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetHasTmpl,
		"golang.get-has.tmpl",
	)
}

func golangGetHasTmpl() (*asset, error) {
	bytes, err := golangGetHasTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-has.tmpl", size: 352, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetLastTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x90\x31\x6f\xc2\x30\x10\x85\x67\xdf\xaf\xb8\x46\xaa\x94\x48\x90\xa9\xea\xc6\xd8\x8d\x05\xa8\xd4\x11\x19\x72\x89\x5c\x82\x9d\x5e\x8e\x56\xf4\x74\xff\xbd\x72\x10\x55\x18\xbc\x7d\xf7\xbd\xf7\xac\xba\xc4\x86\xda\x10\x09\x8b\x31\x74\xd1\xcb\x85\xa9\xc0\xa5\x19\x74\x24\x6b\x3f\x8a\x2a\xd6\x5b\x92\x0b\xc7\xfa\xfd\x3a\x10\x9a\x95\xc3\x09\x43\x94\xd7\x97\x0a\x4b\x55\x1c\x3c\xfb\xf3\x9d\x41\xb3\x05\x12\x73\x7e\x89\x2b\xc8\x7e\x8a\xcd\x24\x84\x79\xd8\x21\x35\xd7\x29\x47\xc1\x1d\x53\x1c\x05\xf7\xfb\x51\xce\xb2\xca\x46\x0e\x51\x5a\x2c\x9e\xbf\x0a\xac\x77\x9b\x35\x9a\x81\x4b\x87\xcf\xba\x4f\xdd\x4e\xce\x52\xde\xd0\x05\x0e\xa7\x0a\xc0\xa9\x62\x88\x41\x66\x15\xc0\xe5\x0a\x2b\xcc\x37\x0d\x87\x6f\xe2\x7a\x73\x21\xbe\x6e\xd3\xcf\xfc\xb6\xde\x1d\x7d\xcc\x13\x7c\xd3\x70\x6a\xb1\x6c\x7b\x2f\x42\xf1\x6e\xaa\xd0\xac\x02\x17\xda\x69\xd1\xd3\x0a\x63\xe8\x51\xc1\x39\xbe\x05\xa9\xe2\x2f\x71\x7a\xd8\xfe\xc1\x7e\x78\x63\x2e\x89\xb9\x02\x67\x30\x63\x3d\x77\x0f\x68\x0c\x3d\x18\xa8\xfe\x7f\xd0\x5f\x00\x00\x00\xff\xff\x05\x0f\x5d\xce\x8d\x01\x00\x00")

func golangGetLastTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetLastTmpl,
		"golang.get-last.tmpl",
	)
}

func golangGetLastTmpl() (*asset, error) {
	bytes, err := golangGetLastTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-last.tmpl", size: 397, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetLimitOffsetTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x92\xc1\x8a\xdb\x30\x10\x86\xcf\xd2\x53\xfc\x35\x14\x6c\xc8\xfa\x54\x7a\xd8\xe2\x43\x29\xa5\x97\x52\xd8\xe6\xd0\x63\xd0\x46\xa3\xa0\x22\x4b\xe9\x58\xde\xec\xa2\xce\xbb\x17\xd9\x49\x68\x69\x28\xf4\x60\xb0\xa4\xe1\xfb\x3f\xcd\xa8\x94\x3b\x58\x72\x3e\x12\x9a\xc9\x1f\xa2\xc9\x33\x53\x83\x3b\x11\xfd\x89\x72\x29\xe8\xb7\xb3\x73\xfe\x19\x22\x6d\x29\x38\x1a\x36\x23\xfa\xf7\x7c\x98\xf0\x13\xfb\x34\x8e\x06\x22\xc1\x8f\x3e\xc3\xc7\xbc\x41\x72\x6e\xa2\xe5\xff\xed\x9b\x0e\x2d\xa7\xd3\x84\x52\x30\x05\xbf\xa7\xe4\xd0\x7f\x4d\x27\x88\x6c\x40\xcc\xf5\x4b\xac\xab\x01\x45\xbb\x44\xea\xdf\x75\x1e\x93\x7d\x69\x20\x52\xb4\xda\xa7\x38\x65\xec\x76\x53\x1e\xf3\x50\x35\xd8\xc7\xec\xd0\xbc\xfe\xd1\xa0\xdf\x3e\x7c\x86\x88\x56\xe9\xf1\x7b\x1f\xd2\x61\x9b\xc7\xdc\xae\xa5\x9b\x1a\xed\x3c\x05\xfb\x64\xc2\x4c\xb7\xbd\x2f\xce\x9d\xd6\x6a\xb7\xab\xc2\xab\xdd\xfd\x80\x4a\xb4\xec\x9f\x88\xfb\x87\x99\xf8\xe5\x36\xf6\x0f\xea\x5f\x50\xe5\xdd\x82\x7b\x35\x20\xfa\x80\xa2\x95\x62\xca\x33\xc7\xba\xdc\xe0\x1b\x9b\xe3\x47\xe6\x96\x98\x3b\xad\x44\x2b\x4b\x8e\x18\xab\x48\xff\x21\xa4\x89\xda\x6a\xe6\xd2\x75\xf3\x0b\x3d\xe7\xb6\x5b\x48\xa5\xc0\x47\x9f\x23\x9d\x2e\x9d\xd5\x4a\xd5\xb4\xe1\x52\xbc\xdd\x9b\x58\x07\x67\xac\xe5\xe4\xd0\xba\x60\x72\xa6\xb8\x94\x77\x10\xe9\xb4\xba\x61\xf8\x0f\xc5\xea\xa8\x96\xb1\x0e\x30\xc7\x23\x45\xdb\xae\x3d\xab\x21\x7c\xb8\x78\xac\x97\x39\x93\xef\xaf\x3a\x95\xd3\xbd\xfb\xaf\x86\x9c\x0f\xd7\x90\xe8\x83\x16\x5d\xca\xf5\xc5\xfc\x0a\x00\x00\xff\xff\x2c\x2e\xe1\xc1\xc0\x02\x00\x00")

func golangGetLimitOffsetTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetLimitOffsetTmpl,
		"golang.get-limit-offset.tmpl",
	)
}

func golangGetLimitOffsetTmpl() (*asset, error) {
	bytes, err := golangGetLimitOffsetTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-limit-offset.tmpl", size: 704, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetPagedTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x93\x51\x8b\xd4\x3e\x14\xc5\x9f\x93\x4f\x71\xff\x81\x3f\xb4\xd0\x29\xeb\xeb\x4a\x05\x11\xf1\x45\xd4\x75\x1e\x7c\x10\x19\xb2\x93\x9b\x12\x6d\x93\xf1\x36\xdd\xd9\x25\xe6\xbb\xcb\x4d\x66\xc6\x05\x15\xf5\x61\x60\x9a\x9e\x9c\xfc\xce\xcd\x69\x4a\x1b\x30\x68\x9d\x47\x50\x8b\x1b\xbd\x8e\x2b\xa1\x82\x4d\xce\xf2\x15\xc6\x94\xa0\xdf\xae\xd6\xba\x7b\xc8\xb9\x49\x09\x0e\x9a\xf4\x0c\xfd\x73\x1a\x17\xf8\x06\xfb\x30\xcf\x1a\x72\xde\xc7\xf0\x05\x3d\x2c\x91\x9c\x1f\x3b\x98\xdc\xec\x22\x38\x1f\x5b\x68\x28\x1c\x17\x48\x09\x96\xc9\xed\x31\x58\xe8\xdf\x87\x23\xe4\xdc\x41\xdd\x13\xd6\x78\xd9\x86\x44\xfc\x0b\xd4\x4a\xa6\x42\x6f\x0a\x86\x7c\x8c\x78\x1b\xcc\x83\x82\x9c\x93\x14\xce\x9e\x3c\x60\x18\x40\x29\x48\x52\x88\xf3\x02\xa8\x2b\x25\x45\x96\x52\xec\x83\x5f\x22\xec\x76\x4b\x9c\xe3\xc0\x01\xc8\xf9\x68\x41\xfd\xff\x55\x41\xbf\xbd\x79\x0d\x39\x4b\x11\x6e\x3f\xf7\x53\x18\xb7\x71\x8e\x4d\x95\x76\xcc\x6c\x1d\x4e\xe6\x4e\x4f\x2b\xfe\x26\xf1\x29\x6a\x2b\xa5\xd8\xed\x38\x69\x0d\x71\x3d\x00\x3b\x1a\x72\x77\x48\xfd\xcd\x8a\xf4\xf0\x37\xb6\x3f\xb9\x72\x44\xf6\xfb\x6f\x00\xef\xa6\x12\x90\x30\xae\xe4\xf9\xb1\x03\xa5\x3a\xf8\x40\xfa\xf0\x92\xa8\x41\xa2\x96\x03\x0b\x83\x16\x09\x2a\x4d\xff\x62\x0a\x0b\x36\x8c\x67\xc3\x65\xf1\x0d\xde\xc7\xa6\x2d\x6e\x29\x81\xf3\x2e\x7a\x3c\x9e\xef\x45\x0a\xc1\x27\x0e\x67\xf1\x76\xaf\x3d\xdf\xbb\x36\x86\x82\x85\xc6\x4e\x3a\x46\xf4\x45\xde\x42\xce\xad\x14\xbf\xa0\xfc\x03\x26\x73\x8a\x52\x8c\x01\xf4\xe1\x80\xde\x34\x75\x78\x7c\x10\x8d\x67\x96\x1a\xe8\xe4\x7e\x7d\x41\x62\x9f\xf6\xe9\xbf\x0e\xa6\x18\xd5\x62\x3e\x83\xab\xb2\x85\x17\xd0\x97\xa3\x5b\xee\x50\x7d\x5b\xf8\x7f\x94\x73\x00\x3b\xc7\x7e\x5b\x6a\x53\xa4\x1f\x8b\x6c\xf3\xe4\x53\xcf\x1f\xc7\x3b\x3d\xa2\x79\xeb\x4f\x93\xc8\x52\x64\xc0\x69\xc1\x47\x65\xac\x26\xf5\x7f\x05\x39\xa1\xd6\xc8\x17\x51\xc7\xec\x32\xcb\x94\x2e\xcd\xff\x1e\x00\x00\xff\xff\x95\xe8\xdd\x70\x9c\x03\x00\x00")

func golangGetPagedTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetPagedTmpl,
		"golang.get-paged.tmpl",
	)
}

func golangGetPagedTmpl() (*asset, error) {
	bytes, err := golangGetPagedTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get-paged.tmpl", size: 924, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangGetTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x90\x31\x6f\xe3\x30\x0c\x85\x67\xe9\x57\xf0\x0c\x1c\x60\x03\x89\xfe\x81\x87\x1b\x82\x5b\x8a\x02\x89\x87\x8e\x81\x12\x51\x86\x0a\x47\x4a\x68\x39\x69\x4a\xf0\xbf\x17\x72\xd1\x24\x5d\xda\x81\x0b\xc1\xf7\xf1\xbd\xc7\xbc\x04\x87\x3e\x44\x84\x6a\x0c\x7d\xb4\x79\x22\xac\x60\x29\xa2\xff\x63\x66\x06\xd3\x4d\xde\x87\x37\x10\xa9\x99\xe1\x68\xc9\x1e\xc0\xfc\xa3\x7e\x04\x91\x06\x1e\x76\x9b\x74\x01\x91\x05\x20\x51\x99\x44\x8d\x2e\x70\x8c\x6e\xa6\xe9\xc7\x4f\xbb\xe4\xae\x15\x88\xb0\x56\xfb\x14\xc7\x0c\xdb\xed\x98\x0f\xb9\x2d\x34\x0a\x31\x7b\xa8\xfe\x9e\x2a\x30\xdd\xfa\x09\x44\xb4\x4a\xbb\x57\x33\xa4\xbe\xcb\x87\x5c\x7f\x9e\x2e\x80\x19\x7c\xc0\xc1\x9d\xed\x30\xe1\xdd\x92\xd6\x8a\x19\x42\x0c\xf9\xcb\x92\x56\xc5\x52\x0b\x05\xe2\x28\x9c\x91\xcc\x7a\x42\xba\x6e\xd2\xe5\x37\x98\xe9\xf6\x36\x96\x8c\xd6\x39\x4a\x1e\x6a\x3f\xd8\x9c\x31\xce\xe8\x66\x7e\xa7\x82\x9f\x23\xb7\x2d\x8c\xa7\xc1\xac\x88\x9e\xd3\x26\x5d\x46\x60\xad\x14\x61\x9e\x28\x16\xfa\x3b\x52\xba\x77\x14\xc3\xa0\x95\xdc\xb4\x7f\xda\xb2\xf9\x51\xf1\x42\xf6\xb8\x22\xaa\x91\xa8\x99\xa5\xf7\x43\x4b\xfd\x77\xb2\x68\xe6\x5b\xef\x1f\x01\x00\x00\xff\xff\x05\x56\xfe\x25\xe1\x01\x00\x00")

func golangGetTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangGetTmpl,
		"golang.get.tmpl",
	)
}

func golangGetTmpl() (*asset, error) {
	bytes, err := golangGetTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.get.tmpl", size: 481, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x57\x5d\x8f\x9b\x38\x14\x7d\xc6\xbf\xe2\x2e\x6a\x2b\x18\xa5\x44\xfb\x9a\x55\x1e\x9a\x49\x76\x35\x52\x77\xd2\xce\x50\xad\x56\xaa\x34\x75\xc0\x30\x74\xc0\x26\xb6\x49\xa8\x58\xfe\xfb\xca\xc6\x80\xc9\x24\xb3\x53\xed\x13\xc4\xbe\xf7\xdc\x73\x3f\x8e\x71\xe6\x73\xf8\xf0\x25\xdc\xfe\xb1\xb9\xdd\xdc\x7d\x08\x37\x6b\x58\xfd\x0d\x29\x2b\x9f\xd2\x20\xa3\x73\x51\xe2\x88\x14\x8c\x3e\x91\x1f\x29\x9b\xc7\xbb\x3a\x38\xfc\x8a\xe6\x73\x58\x6f\xe1\x76\x1b\xc2\x66\x7d\x13\x06\x08\x95\x38\x7a\xc2\x29\x81\xa6\x81\xe0\x93\x79\x6f\x5b\x84\xb2\xa2\x64\x5c\x82\x87\x1c\x37\xc6\x12\xef\xb0\x20\x73\xb1\xcf\x5d\xe4\xb8\x49\x21\xd5\x43\x66\x05\x51\x4f\x21\x79\xc4\xe8\xc1\x45\xa8\x69\xde\x03\xc7\x34\x25\x10\x6c\x6a\xc9\xf1\x8d\xc6\x10\x0a\xcf\x51\xf8\x1f\xd4\x2b\xb8\xd3\x50\xae\x76\x23\x34\x56\x66\x3e\x52\x0c\x3f\x71\x72\x20\x54\x42\xc4\x68\x9c\xc9\x8c\x51\x9c\x43\x66\xb0\x12\xce\x0a\x88\x70\x25\x32\x9a\xc2\xae\xca\xf2\x18\x12\x9c\xe5\x15\x27\x02\x1d\x30\x87\x07\x58\x82\x61\x14\xdc\x48\x86\x91\x5e\xfd\x8b\xe3\x72\xc3\x39\x2c\x21\xa9\x68\xe4\x11\xce\x81\x70\xce\xb8\xdf\x3d\xa0\xe1\x44\x56\x9c\xaa\x5f\xad\x76\xf8\xc8\xd2\x94\xf0\xce\x3a\x61\xbc\xc0\x52\xa1\x66\x34\x9d\x01\xe6\xa9\x80\x20\x08\x32\x2a\x09\x4f\x70\x44\x9a\xd6\xd7\x3e\xb7\xec\x08\x4b\x50\x65\x09\x6e\xd9\x11\x21\xf9\xa3\x24\x10\xf3\xec\x40\x38\x0c\xc6\xd0\x20\x67\x53\x93\xc8\xdb\x57\x84\xff\x78\x11\x14\x3c\xb1\xcf\x83\x3b\x22\xaa\x5c\xce\x0c\x5f\xe4\x7c\x56\x7e\xaf\xf2\xbe\xd2\xee\xec\x28\x4e\x9c\xef\xd8\xf1\x35\xfe\xbd\x3b\x6a\x4d\x2a\xeb\x95\x72\xa8\x22\xa9\x72\xd0\xbb\xeb\x15\x72\xe2\xdd\x9f\x44\x3e\xb2\x58\x28\x3b\x55\x2f\xd8\x96\x84\x7a\x5d\xde\x33\x10\xac\xe2\x11\x31\x91\x7c\xf0\xe2\x1d\x5c\xad\x57\x9a\x51\xdf\x82\x06\x39\x62\x9f\x3f\xc4\xbb\x6e\x75\xb1\x04\x85\x7d\x06\xc5\x47\x4e\x96\x68\x9b\x5f\x96\x40\xb3\x5c\x79\x3a\xa6\x73\x34\xcb\x67\x7d\x9b\x55\x7f\x7d\xe4\xb4\xc8\x89\x49\xd2\x77\xb1\x0b\x01\x86\xb7\x8e\x7a\x06\xcd\x30\x09\xae\x73\x26\x88\xe7\x3b\x0e\x72\x14\x4e\x6b\xdc\x7d\x34\x50\xe8\x68\x2a\xdb\x4f\x19\x4d\x3d\xff\xb7\x9f\x22\x86\x1c\x35\x31\x59\x51\xe6\x30\x16\xd0\x11\xc7\x4c\x46\x8f\xfd\xcc\x34\xb6\xa0\xd6\x19\xce\x49\x64\xc4\x14\x61\xd1\x29\xf6\x16\x17\x04\xfe\x81\x92\x67\x54\x26\xe0\xbe\xdd\xbb\xd0\xb6\x0b\x95\x99\x42\x5e\x02\x25\xc7\xc1\xac\x1d\x93\x98\x08\x95\x44\x5b\xaa\x8a\xad\x91\x55\x76\x0f\x33\x38\xc9\x50\x0f\x6c\xd3\x4c\xe2\x04\xd0\xb6\xcf\x93\x7e\x21\x6b\xa7\xb5\x95\x6e\xbd\xaa\x2e\xe1\x2a\x97\x8b\x0b\x45\x4b\x0a\x19\x6c\xd4\xa8\x24\x9e\x5b\x51\x51\x95\xea\x24\x20\x71\x5f\xa6\xb7\xc2\x9d\x99\x77\x5f\x17\x17\xf5\x28\xef\xd6\x2b\xc5\x69\xbd\x5a\x98\x44\x66\xc8\x19\xe7\x75\xa1\xcb\x3f\x43\x4e\x3b\x53\xf1\x86\xf1\xf5\xd8\xee\xbb\x9a\x51\x1f\xcc\x10\x80\x37\x1d\x56\x83\xde\xd3\x63\xbb\xef\xc1\x7a\xd5\x4f\x8c\x7f\x06\x47\x8f\xb2\x12\x64\x58\xcf\x2c\x1c\x59\x0f\x03\x6f\x30\x56\x24\xcd\xa8\xf7\xd3\x63\x3e\x26\x1c\xd6\xca\x56\xd6\x0b\x90\xf5\x4c\xbf\x0d\xc9\xaa\x10\x47\x8e\xcb\xb0\xf6\x64\xed\x4f\xd2\xd6\xea\x0e\x6b\x4b\xdd\xb2\xee\x84\x12\xd6\x68\xc4\x18\x33\x53\xdb\x61\xed\xc3\x35\x2b\x8a\x4c\xfe\x67\x85\x64\x1d\xc8\x3a\xe8\x8d\xfd\xe7\x38\x77\x2c\xcf\x77\x38\x7a\x7a\x25\xd2\x68\xae\xb1\x2e\xa8\x04\x35\x0d\xbc\x89\x77\x3a\xb9\xc5\xf2\xb9\x56\xc4\x7a\xe5\x76\x93\x08\x6f\x64\x7d\xd9\x2c\xac\x07\x33\x35\x30\x97\x0d\x6f\x8a\x32\xd7\xa6\x5d\x41\x27\x0e\x6d\x6b\x55\xd7\x0c\x6e\xf7\x38\x99\x97\x13\x2f\x1f\x72\x96\xde\xcb\x42\x7a\x42\x16\xd3\x6f\x50\x10\x04\x30\x39\xb0\x9b\xee\x2b\x6b\xc4\xfe\xd1\xf2\x1b\x1c\x7c\x64\xb3\x33\xc5\x99\x72\x1b\x8e\x48\xe4\x9c\x92\x19\xa8\x9e\x1c\x2b\xf6\xb1\x7a\x35\x05\x1e\x5b\xf8\x6e\xb2\xd1\x68\x25\x2e\xa0\xd3\xe4\x49\x9c\x45\x67\x6c\xad\xe8\xa3\xa5\xab\x57\xef\xd3\xaa\x11\x3e\x53\xbc\x21\x84\x0f\xf7\xd1\x23\x29\xb0\xe7\x9b\xb2\x59\x64\xbe\x29\xfa\xdd\xf6\xfd\xe7\x8f\xd0\xb6\xdf\x5e\x46\x1a\x84\xd3\xcb\xc2\x87\x41\x16\xa7\x39\x9a\x59\xea\x48\x8f\x52\xfc\xb9\x1c\x3b\x9f\x21\xc7\xa1\x65\x03\xf8\x05\xb1\x5e\x6c\xd9\x85\xc9\x78\xf9\x0a\xd0\x20\x67\x3e\x87\x70\xbb\xde\x2e\x80\x13\x1a\x13\x0e\x65\x8e\x23\xf2\xc8\xf2\x98\x70\xa1\xcf\x29\x73\x53\xb2\x8e\xaa\x6e\xc5\x73\xc5\x3e\x5f\x7c\xa5\x6f\xc5\x57\xaa\xc0\xd5\xeb\xde\x9d\xc1\x38\x8f\xbe\x49\xce\xfa\x14\x28\x95\x19\x29\xdf\xeb\xfc\x44\x2f\x3d\x93\xee\x20\xbc\x89\xcc\xcc\xa6\x5d\x15\xeb\x50\xf8\x3d\x23\x79\x3c\xde\x42\x8d\xbb\xae\x48\x68\xaa\x64\x51\x30\x84\xb2\x04\x82\x6d\xd9\x5d\x3e\x6f\xa8\x20\x5c\x8e\x30\x43\xe0\xe0\x9a\x13\x2c\x49\x47\xb5\xc7\x3d\x47\xe1\x12\xd2\x84\xd0\x95\x1e\x4a\x1b\x6b\xca\x6b\x5a\xa7\x8e\xe2\x97\x52\xdf\xce\x73\x72\x8e\x9d\xde\x7c\x05\xbb\x33\x20\xff\x8b\xd8\xb3\xba\xab\xc5\x37\x89\xd5\xc2\x29\xd8\xd8\xc8\xe4\x79\x27\xe1\x60\x77\x0a\xda\x61\x9c\xc7\xae\x3f\xd8\x87\xd1\xc4\xdc\x3f\x45\x6d\xc0\x48\x75\xba\xde\x1c\x16\x70\x18\xc0\xbd\xea\x64\xdb\x87\x88\xe5\x55\x41\x55\x0c\xeb\x38\xe9\xb1\xf4\x3f\x9a\x6b\x6d\xa1\xfe\xd0\x58\x30\x57\xa7\x38\x07\x9c\x57\x0a\xc2\x92\x19\x34\xa0\x85\x54\xc1\xf2\xdc\xe7\x5e\x5f\x17\xcd\xcf\x2a\x38\xa0\xcb\x1c\xef\x48\xe2\xf9\xa7\x21\x47\x96\xef\x2a\xb0\xc4\xf6\x7e\xda\xc4\x7f\x03\x00\x00\xff\xff\xe6\x9b\x08\xb0\x4a\x0e\x00\x00")

func golangHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangHeaderTmpl,
		"golang.header.tmpl",
	)
}

func golangHeaderTmpl() (*asset, error) {
	bytes, err := golangHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.header.tmpl", size: 3658, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangMiscTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\x41\x4b\xc3\x40\x10\x85\xcf\xcd\xaf\x78\x84\x08\x2a\xa6\x3f\xa0\xe0\xa5\x07\x41\x04\x0f\x5a\x3c\x77\xcd\x4e\xc2\x48\xba\x8d\x9b\x8d\x50\xc6\xfd\xef\x32\xbb\xb2\x56\x7a\x5b\xde\x7c\xef\xbd\x99\x15\x69\x61\xa9\x67\x47\xa8\x2d\x75\x63\x8d\x18\xab\x7e\x71\x1d\xae\x8f\xef\x1f\xb8\x15\xc1\xfa\x85\x3a\xe2\x2f\xf2\x5b\x33\x13\x62\x7c\x3c\x4c\xe3\x0d\x74\xf0\xca\x83\x33\x61\xf1\xaa\x26\x61\x7b\xb4\x27\x0d\x10\x01\x39\x8b\x36\xc6\xaa\x3a\x6f\x98\x83\x5f\xba\x90\x3a\xc2\x69\xa2\xe4\x79\x36\x87\xe4\xcf\x33\x88\x9a\xbd\x71\x03\x61\xfd\xc0\x34\xda\x59\xe9\x95\x08\xb8\x2f\xf0\xb9\xef\xb7\x2b\x8b\x3b\x4d\x4d\x4f\xa5\x77\x66\x98\xd3\x12\xab\x7d\x09\x6d\xf8\x0e\x4d\xc0\xe6\xfe\x6f\x9c\xe9\x86\x2f\xd2\x9e\x48\xaf\xd9\xe8\xf3\xcd\x8c\x0b\xe1\x1b\x93\x67\x17\x7a\xd4\x57\x9f\x75\x86\x32\xbd\xd7\x0d\xdb\x72\x74\xd1\xab\x7f\x7f\xf1\x13\x00\x00\xff\xff\x1b\x36\x80\x57\x6d\x01\x00\x00")

func golangMiscTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangMiscTmpl,
		"golang.misc.tmpl",
	)
}

func golangMiscTmpl() (*asset, error) {
	bytes, err := golangMiscTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.misc.tmpl", size: 365, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangUpdateTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x55\x41\x8b\xeb\x36\x10\x3e\xdb\xbf\x62\x6a\x28\xd8\x24\x11\xa5\xc7\x07\x3e\x2c\xaf\x69\x09\x6c\xc3\xee\xa6\xaf\x3d\x3c\x1e\x41\x89\x47\x46\xc1\x91\x1c\x49\x4e\x76\x31\xfe\xef\x65\x24\x27\x6b\x6f\xb2\xdb\x96\x5e\xde\x21\x07\x45\x33\xdf\xcc\x7c\xdf\xe8\x73\xdb\xce\xa0\x40\x21\x15\x42\x62\x65\xa9\xb8\x6b\x0c\x26\x30\xeb\xba\xf8\x4b\x5d\x70\x87\x6d\x0b\x6c\xd5\x08\x21\x9f\xa1\xeb\xd2\xb6\x85\x9a\x1b\xbe\x07\x76\x67\x4a\x0b\x5d\x37\x85\xc6\x87\x81\x8f\x73\xa6\xd9\x3a\x16\x12\xc3\x61\xc9\xf7\x08\x5d\x97\xc1\x20\xf5\x09\x5d\x63\x94\x4f\x46\x63\xe8\xa7\x4d\x16\x53\x27\xa8\x0a\x5f\x3a\x1e\xb6\xb5\xd1\xc5\x8b\xef\xa8\x8d\xa3\xad\x56\xd6\xc1\x7a\x6d\xdd\xde\xad\x6b\x83\x42\x3e\xe7\x04\x6c\xa4\x72\x02\x92\x1f\x0f\x09\xb0\xd5\xe3\xfd\x83\xbf\x81\xae\x7b\x93\x61\xfd\x20\xb7\x32\x2e\x23\xc6\x71\x74\xe4\x06\x0e\x16\xbe\x7e\xdb\xbc\x38\x0c\xc7\x23\xaf\x1a\xa4\xbf\xa4\x72\x68\x04\xdf\x62\x4b\x91\x6d\x0b\x52\x00\x7b\xd0\x56\x3a\xa9\x15\xaf\xee\x4c\xd9\xec\x51\x39\xeb\x6b\xef\xe0\x53\x4e\xc4\x54\xa8\x2e\x84\x11\x9e\xc7\xca\x81\xd7\x35\xaa\x22\x0d\xe7\x29\x05\x0a\x89\x55\xe1\xcf\x97\xf8\xcc\x57\xb9\x10\x13\xf1\xa2\x38\x17\x21\x74\xd1\xa8\x6d\xaa\x88\x64\xeb\x8c\x54\x65\x06\x6d\x1c\x45\x87\x01\xfc\xc1\x4e\x81\x02\x18\x63\x59\x1c\x45\xc4\xec\x47\x3d\x47\xbb\xc9\xe4\x06\x42\x02\x7e\x12\x76\x0e\xbe\x30\x9c\xf4\xb8\x6f\x13\xac\x33\x5b\xad\x8e\x6c\xe1\x34\x4f\x77\xd9\x3b\x51\xc9\x14\x92\x41\x63\x58\x59\x0c\x5d\xfc\xeb\xfa\x6f\x00\x54\xe1\xf3\x7b\x71\x0c\x57\x25\x8e\xf7\x92\x6f\x2a\xfc\x95\x68\x0e\xe3\x4a\xd1\x2f\x30\x23\xf4\x7e\x5b\xe1\x87\x1c\x94\xac\x3c\x95\x03\xbe\xd3\x84\x62\x3e\xeb\xaa\xd9\xd3\xf6\x26\x54\xf4\x3d\x31\xaf\x41\x99\xbf\x4a\xb3\x8c\xba\x3b\x6b\xda\x85\x46\x83\x24\x4b\xc4\xc2\x2e\xf5\xc9\x37\xb6\x5e\x2b\x7d\x22\x85\x97\xfa\x94\x66\xec\xcb\x1f\x9f\xd3\xf1\x2a\x0c\xe6\xbb\x6b\x9c\x1e\xcc\x74\xd5\x72\xdf\x01\x35\xfc\xc1\xf2\xb1\x85\x92\xee\x4f\x5e\x8d\x96\xae\x7b\x5d\x73\xa5\xdd\x55\x29\x29\x68\xb9\xd3\x83\xcd\x20\xcf\xe1\x27\xcf\x98\x09\xcf\x5b\xc9\x6a\x0a\x7f\x19\x5e\xcf\x8d\x49\xc5\xde\xb1\x39\xbd\x73\x91\x26\x0a\xb1\x00\xa7\xcf\xc6\xc1\x1d\x54\xc8\xad\x03\xad\x30\x3c\x80\xe4\x26\x47\x97\x16\x6e\x6d\xee\xec\xff\x3c\x2c\x5f\x20\x38\x04\x11\x3e\x72\x17\x98\xf4\x0f\x2b\x3d\xd8\xaf\x9f\xfa\x59\x67\x3f\x7f\xcb\x60\x32\x36\x95\x38\xd2\x9b\x1d\xab\x74\xb9\x72\x7b\x97\x86\xab\x69\xef\x1b\x7e\x41\xc3\x0c\x4a\xba\x81\x01\x5e\x1c\x64\xd5\xd4\xb5\x36\xce\x86\x1b\xa9\x4a\x7f\x49\xee\x98\x03\xe1\x16\x46\x1e\xd1\xb0\xc7\x06\xcd\xcb\x93\x3e\xdd\xc0\x67\xab\x2d\x57\x64\xb2\xbc\x28\x8c\x16\x90\x8a\x8a\x3b\x47\xc6\x13\x30\xb3\x30\xb2\x14\xde\x73\xf3\x1c\xec\xa1\x22\x49\x96\xfa\x49\x9f\xec\x95\x70\x4a\x56\x5e\x83\x3e\x7e\xf0\x20\x6e\xc9\x8b\xc6\xbc\x4a\x46\x6f\xd8\xeb\xb1\x0e\xfe\x3e\x9a\x60\xfe\x8c\xdb\x5b\xec\xfc\xd7\x42\x67\x5b\x2f\xd1\x79\xb0\x2b\x4f\xff\x0d\xdd\xea\xf1\xde\xd3\x38\x14\x26\x49\x17\xbf\x3f\xdc\x2f\xe6\xbf\x64\x90\x78\x09\xcf\x00\xef\xaf\xc9\xc7\x3a\xfc\x63\xfe\xf7\x23\xcc\xd9\x38\xfa\x40\x6a\xc9\x94\xa3\xcf\x31\x81\x77\xa3\x2f\xf1\xdf\x01\x00\x00\xff\xff\xa8\xad\x89\x4d\x20\x08\x00\x00")

func golangUpdateTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangUpdateTmpl,
		"golang.update.tmpl",
	)
}

func golangUpdateTmpl() (*asset, error) {
	bytes, err := golangUpdateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.update.tmpl", size: 2080, mode: os.FileMode(420), modTime: time.Unix(946710000, 0)}
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
	"golang.create.tmpl": golangCreateTmpl,
	"golang.delete-all.tmpl": golangDeleteAllTmpl,
	"golang.delete-world.tmpl": golangDeleteWorldTmpl,
	"golang.delete.tmpl": golangDeleteTmpl,
	"golang.footer.tmpl": golangFooterTmpl,
	"golang.get-all.tmpl": golangGetAllTmpl,
	"golang.get-count.tmpl": golangGetCountTmpl,
	"golang.get-has.tmpl": golangGetHasTmpl,
	"golang.get-last.tmpl": golangGetLastTmpl,
	"golang.get-limit-offset.tmpl": golangGetLimitOffsetTmpl,
	"golang.get-paged.tmpl": golangGetPagedTmpl,
	"golang.get.tmpl": golangGetTmpl,
	"golang.header.tmpl": golangHeaderTmpl,
	"golang.misc.tmpl": golangMiscTmpl,
	"golang.update.tmpl": golangUpdateTmpl,
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
	"golang.create.tmpl": &bintree{golangCreateTmpl, map[string]*bintree{}},
	"golang.delete-all.tmpl": &bintree{golangDeleteAllTmpl, map[string]*bintree{}},
	"golang.delete-world.tmpl": &bintree{golangDeleteWorldTmpl, map[string]*bintree{}},
	"golang.delete.tmpl": &bintree{golangDeleteTmpl, map[string]*bintree{}},
	"golang.footer.tmpl": &bintree{golangFooterTmpl, map[string]*bintree{}},
	"golang.get-all.tmpl": &bintree{golangGetAllTmpl, map[string]*bintree{}},
	"golang.get-count.tmpl": &bintree{golangGetCountTmpl, map[string]*bintree{}},
	"golang.get-has.tmpl": &bintree{golangGetHasTmpl, map[string]*bintree{}},
	"golang.get-last.tmpl": &bintree{golangGetLastTmpl, map[string]*bintree{}},
	"golang.get-limit-offset.tmpl": &bintree{golangGetLimitOffsetTmpl, map[string]*bintree{}},
	"golang.get-paged.tmpl": &bintree{golangGetPagedTmpl, map[string]*bintree{}},
	"golang.get.tmpl": &bintree{golangGetTmpl, map[string]*bintree{}},
	"golang.header.tmpl": &bintree{golangHeaderTmpl, map[string]*bintree{}},
	"golang.misc.tmpl": &bintree{golangMiscTmpl, map[string]*bintree{}},
	"golang.update.tmpl": &bintree{golangUpdateTmpl, map[string]*bintree{}},
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

