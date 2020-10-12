// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
package internal

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

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x51\x5f\x6b\xdb\x3e\x14\x7d\xb6\x3e\xc5\xf9\x99\xfe\xa8\xdd\xa5\x4a\xdb\xb7\x0d\xf2\x50\xda\x0c\x32\xb6\x76\x90\xc2\x1e\xba\x52\x14\xfb\x3a\x11\x75\x24\xef\x4a\x29\x0b\x42\xdf\x7d\x48\x4e\xc2\xf6\x64\x4b\xe7\xdc\xf3\x47\x37\x84\xe9\x85\xb8\xb3\xc3\x9e\xf5\x7a\xe3\x71\x73\x75\xfd\xf1\x72\x60\x72\x64\x3c\x3e\xab\x86\x56\xd6\xbe\x61\x61\x1a\x89\xdb\xbe\x47\x26\x39\x24\x9c\xdf\xa9\x95\xe2\x69\xa3\x1d\x9c\xdd\x71\x43\x68\x6c\x4b\xd0\x0e\xbd\x6e\xc8\x38\x6a\xb1\x33\x2d\x31\xfc\x86\x70\x3b\xa8\x66\x43\xb8\x91\x57\x47\x14\x9d\xdd\x99\x56\x68\x93\xf1\xaf\x8b\xbb\xf9\xc3\x72\x8e\x4e\xf7\x84\xc3\x1d\x5b\xeb\xd1\x6a\xa6\xc6\x5b\xde\xc3\x76\xf0\x7f\x99\x79\x26\x92\xe2\x62\x1a\xa3\x10\x21\xa0\xa5\x4e\x1b\x42\xb9\x55\xda\x94\x88\x51\x4c\xa7\xb8\x4b\x79\xd6\x64\x88\x95\xa7\x16\xab\x3d\xce\xc9\xf8\xe6\x74\x75\x2e\x71\xff\x88\x87\xc7\x27\xcc\xef\x17\x4f\x52\x0c\xaa\x79\x53\x6b\x42\xd2\x10\x42\x6f\x07\xcb\x1e\x95\x28\x4a\xeb\x4a\x51\x94\xab\xbd\xa7\xf4\x13\x02\x3c\x6d\x87\x5e\x79\x42\x39\xb2\x5c\xb6\xcc\xd0\xc0\xda\xf8\x0e\xe5\xff\xbf\x4a\xc8\xef\x07\xc5\x18\x45\x9d\x63\x9e\xad\x94\x23\x7c\x9a\x21\x7f\x8f\x78\x9a\x7d\x57\x0c\xd7\x6c\x68\xab\x1c\x66\x78\x7e\x21\xe3\xe5\xc2\x78\xe2\x4e\x35\x14\xb2\x34\x2b\xb3\x26\x9c\xbd\x4e\x70\x66\xd4\x36\xcb\xc8\x07\xb5\x25\x97\xf4\x8b\x22\x84\xcb\x83\x7e\x8c\x32\x1d\x4e\x51\x5c\x88\xe5\x61\x26\xc6\x49\xd6\x22\xd3\xe2\x32\x46\x11\x85\xe8\x76\xa6\xc9\x9d\xab\x1a\x41\x14\x29\x48\xaf\x0d\x39\x3c\xbf\x3c\xbf\xa4\xd2\xa2\xe8\x2c\xe3\x75\x72\xc8\x97\x7c\xc7\x28\xc7\xbc\x41\x14\xc5\x6a\x02\x62\x4e\xd8\x37\xc5\x6e\xa3\xfa\x65\x06\xab\x91\x53\x8b\xa2\xd0\x5d\x66\xfc\x37\x83\xd1\x7d\x9e\x29\x3a\xa5\xfb\x8a\x98\x13\x9c\x2a\x8c\xbe\x33\xa8\x61\x20\xd3\x56\xf9\x38\xc1\xaa\x16\x09\xb5\x4e\x2e\x7d\x6b\x77\x5e\xfe\x60\xed\xa9\xca\xfb\x90\x5f\xac\x36\x47\xe2\x18\xb7\x2a\x7f\x9a\xb2\xae\xeb\x53\xb7\xa3\x4b\xb2\xb7\x9c\x4b\x8e\x5a\xc4\x3c\x6a\x2d\x3d\x6b\xb3\x4e\x1c\x39\x4f\x9c\xaa\xfe\x90\x45\x32\x71\xfe\x5b\xfb\xea\x3a\xcb\xfd\xb3\xfa\xb1\xd9\xb8\xf9\xc3\x8b\xc6\x28\xfe\x04\x00\x00\xff\xff\x95\x06\x0f\xa4\x50\x03\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 848, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x5a\x5f\x6f\xdc\x36\x12\x7f\xde\xfd\x14\x13\x03\x35\xa4\x60\x2b\xf7\x8a\xa2\xb8\xdb\xdc\x1e\x50\xb4\x29\xea\xeb\xd5\x0d\x9a\xa4\x2f\x41\xe0\xca\x12\xb9\xcb\x58\x22\xb7\x24\xd7\xb1\xeb\xfa\xbb\x1f\x38\x43\x52\x94\x56\xbb\xde\xc4\xf6\xbe\x58\x1a\x0e\xe7\xcf\x8f\x33\xc3\x21\xe5\x93\x13\xf8\x5e\xad\x6f\xb4\x58\xae\x2c\x7c\xfd\xd5\x3f\xfe\xf5\xe5\x5a\x33\xc3\xa4\x85\x1f\xcb\x8a\x5d\x28\x75\x09\xa7\xb2\x2a\xe0\xbb\xa6\x01\x64\x32\xe0\xc6\xf5\x15\xab\x8b\xe9\xc9\x09\xbc\x59\x09\x03\x46\x6d\x74\xc5\xa0\x52\x35\x03\x61\xa0\x11\x15\x93\x86\xd5\xb0\x91\x35\xd3\x60\x57\x0c\xbe\x5b\x97\xd5\x8a\xc1\xd7\xc5\x57\x61\x14\xb8\xda\xc8\xda\x89\x10\x12\x59\xfe\x77\xfa\xfd\xcb\xb3\xd7\x2f\x81\x8b\x86\x05\x9a\x56\xca\x42\x2d\x34\xab\xac\xd2\x37\xa0\x38\xd8\x44\x9f\xd5\x8c\x15\xd3\xe9\xba\xac\x2e\xcb\x25\x83\x46\x95\xf5\x74\x2a\xda\xb5\xd2\x16\xb2\xe9\xe4\x88\xc9\x4a\xd5\x42\x2e\x4f\x3e\x18\x25\x8f\xa6\x93\x23\xde\x5a\xf7\x47\x33\xde\xb0\xca\x1e\x4d\xa7\x93\xa3\xa5\xb0\xab\xcd\x45\x51\xa9\xf6\x84\x7b\x87\x4f\x98\x44\xb6\x1d\x43\x27\xa6\x5a\xb1\xb6\x3c\x61\xf5\x92\x1d\xc0\xc6\x05\x6b\xea\x03\xf8\x84\xac\xd9\xf5\xd1\x34\x9f\x3a\x48\x5e\x23\x0d\x34\xf3\x8b\x61\xa0\x94\xc0\xa4\x2d\xfc\x80\x5d\x95\x16\x3e\x96\x06\x7d\x66\x35\x70\xad\x5a\x28\xa1\x52\xed\xba\x11\x0e\x78\xc3\x34\x78\x5c\x8a\xa9\xbd\x59\xb3\x20\xd2\x58\xbd\xa9\x2c\xdc\x4e\x27\x67\x65\xcb\x00\xc0\x51\x84\x5c\x02\xfe\xfe\x70\x48\xcd\x8f\x64\xd9\xb2\x99\x6a\x85\x65\xed\xda\xde\x1c\xfd\x31\x9d\x7c\xaf\x24\x17\x4b\x40\x1b\xc2\xb3\x67\xae\xf0\xb5\xcf\xfe\xb2\x5e\x32\x03\x00\xef\xde\x3f\x77\x8f\xa9\x6c\x07\x9b\xe9\x73\xff\xe8\x20\x32\xc8\x8d\x8f\x09\x37\xa2\x37\x60\x3f\x75\x48\x31\xe3\xd8\xf1\x31\x61\x17\x34\xd4\xe7\xff\x49\xa9\x4b\x6f\xcc\x2b\x65\x84\x15\x4a\x06\xfe\x95\x1b\xea\x73\xbf\x52\x8d\xa8\x6e\x46\xb9\xd7\x38\xd4\x63\xbf\xc3\xe5\x8a\x8c\x35\x33\x95\x16\x17\xcc\x40\x09\xeb\x40\xf4\xb1\x4c\xeb\xec\x57\x23\xce\xe8\xd6\x23\xfa\x02\x20\xa4\x05\x38\x39\x01\x22\xf9\xf9\x08\xc5\x89\xb3\x18\x1a\x61\x6c\x31\x9d\xfc\x22\xae\x59\x7d\x2a\xdd\x8c\x0b\xa5\x1a\x3f\x43\x54\xa5\x65\x06\x04\x4f\xb4\x82\xba\xf8\xc0\x2a\x0a\x99\xd6\xcd\xfa\x52\x48\x12\x20\x64\x50\x42\x2a\x91\x04\x22\x55\xdc\x22\x89\x74\x92\xbf\xb4\x4a\xdb\xd1\x49\xf4\xcf\x08\x4e\x9a\xb8\x1d\x9b\xf4\x4b\x23\x34\xfd\xed\x8c\xd6\x53\xc9\x55\xc7\xf6\x1c\x91\x2b\xde\xdc\xac\x59\x6f\xc0\x4f\x77\x06\xf4\xa7\xbf\x29\x53\x65\xf7\x68\xb7\xe5\x20\xf6\x5f\x8b\xbf\x12\xdb\x9f\x0b\x69\xbf\xfd\x66\xe7\x6c\x23\xfe\x1a\x28\x7f\x29\x37\xad\x89\x6c\xef\xde\x13\x28\xb7\x70\x36\x83\xdf\x83\x2d\x77\x31\x99\x1c\x73\x7f\xfe\x5b\x29\xfe\xdc\x44\x03\x30\x2e\x46\x7e\x7e\xfe\x06\x99\xfb\x02\xce\x44\xd3\x94\x17\x0d\x3b\x48\x80\xf4\xcc\x7d\x11\xbf\xae\x5d\x6c\x97\xcd\x41\x22\x94\x67\xee\x8b\xf8\x81\xf1\x72\xd3\xd8\xc3\xdc\xa8\x89\x79\x54\xc2\xef\x65\xe3\xe0\x10\xd2\x32\xed\xea\xee\xed\xdd\x1e\x09\xe7\x57\x8e\x7b\x00\xe8\xba\x2e\x2d\x0b\xf6\xdc\x07\x28\x32\x9f\x8f\x1a\x74\xda\xb6\x1b\x1b\x91\xbd\x47\x90\x08\xcc\x7d\x19\xbf\x97\x8d\xa8\x4b\xab\x34\x86\x08\x26\xed\x6e\x19\x57\x91\x79\x10\xa1\x56\xe9\x72\xc9\x7e\x66\x37\x70\x7f\x7c\x1b\x62\x3e\xbf\x64\x37\xc3\x42\xe9\x4b\x18\xfe\x9e\xf7\x5f\x87\x52\x42\x31\x1c\x18\xc2\xa4\x23\x5f\x1d\x84\x88\x09\xcc\x03\x19\x58\xe0\x5c\x72\x3b\xde\xb6\x5c\xbf\x23\x87\xde\xf7\xfc\x0a\x32\x90\xf9\x7c\x3b\xe5\xbf\x93\x52\xd9\xd2\x59\x68\xfa\x52\x7a\x71\xe3\xa5\x94\x1d\xf3\xc8\x66\x80\x1b\xde\x76\x6d\x44\xf2\x67\x94\x46\x9c\x37\x5e\x19\x77\xac\xdc\xce\xb2\x18\x40\xba\x7f\xee\xfe\x9a\x78\xcf\xdc\x61\x41\xfc\x8d\xf1\x68\xf5\xfe\xa9\x9a\xf1\xf3\x6d\xb3\x7f\x63\x3c\x32\x76\xed\xc4\x8e\xf9\xbb\x8b\xe1\x8e\xf0\xda\x53\x09\x4f\xe5\x15\xd3\x66\x6f\x70\xc6\xbe\x03\x39\x87\x76\xff\xb9\x11\x9a\xd5\xf7\x4f\xd7\x9e\x73\x77\x9a\x3e\x77\x6d\x53\xd1\x4f\xdc\x03\x72\x34\x0d\xeb\x1d\x41\x7d\x50\x4c\x53\x8f\xb0\x1d\xd4\x44\xff\x8c\xa8\xa6\x89\x5d\x58\x27\x0b\x15\xa1\xda\xb3\x32\xa1\xbf\x0c\x3b\xa4\x8b\xa9\xfb\xfb\xcb\x11\xee\xb1\xfe\x32\x41\x39\x86\xeb\x3d\x40\x13\x4a\x67\xec\x23\x86\x67\xa5\x19\xb6\x60\xa5\x0c\x88\x38\xa3\x08\x16\x7c\xa2\x36\x71\x6d\x95\x2e\xa6\x7c\x23\xab\x30\x33\x63\xb5\x5f\xe9\x1f\x22\x47\xee\x63\xfe\x76\x3a\x91\x0c\xe6\x0b\x38\x76\xaf\xb7\xd3\x89\x4b\xc9\x79\x8c\x24\x56\x17\x6f\xca\xe5\xcc\x91\x6f\xd6\x6c\x9e\x92\x5d\x2e\x4f\x27\x58\x39\x52\xba\x7b\x77\x74\x82\x7e\x1e\xe9\xf4\xee\x46\x7c\xfc\xcf\xc3\x88\x7f\x77\x43\x21\xb6\xe7\x7e\x28\xbc\xd3\x18\xef\x74\xe1\x18\x0f\xba\x3a\x68\xe7\x38\xd4\xbd\xbb\xd1\x24\x5a\xe7\xd0\x96\x97\x2c\x1b\x8f\xd9\x7c\x36\x9d\xdc\x4d\x27\x5c\x69\x38\x9f\x41\x69\x1d\x2a\xba\x94\x4b\xe6\x44\xa6\x21\xef\x50\x92\x2c\x25\xbd\x2b\x2d\x3a\x9e\xe5\xef\x61\x01\xa5\x45\x41\x82\x83\x66\xdc\x49\x21\x6b\x5f\xe0\xeb\xb3\x05\x48\xd1\x04\x19\xae\x08\x2d\xe2\x3a\x69\xc6\x73\xa2\x27\xc1\xb2\x00\xe2\x4b\x68\x28\x5e\x33\xbb\xd1\x12\x24\xeb\xc2\x84\xfa\xdd\xed\x38\xc1\x70\xa4\x40\xa1\xc7\xb1\x48\xc1\xc9\x19\xaf\x43\x63\x9b\xc6\x4a\x46\x27\xa8\x19\x30\xad\xdd\xfb\x2d\x7a\xc7\xb4\x76\xde\xf1\xba\x78\xa9\x75\x96\xbf\x40\x42\xe2\x5f\xb0\x50\x34\x33\xe0\xad\x75\x5c\x4a\xf3\x8c\xb2\x03\xbe\xf8\x73\x0e\x5f\x5c\x1d\xcd\xdc\x7c\x5c\x48\x37\x3d\x47\xd7\x0c\xa2\x76\x8c\x3a\x6f\x87\x31\x06\x71\x02\xc6\x12\x57\xfd\x11\x47\x99\x0d\xc3\x18\x47\x7c\x20\x63\x27\x3c\x4f\x07\x90\xb2\x15\xb3\x38\xd4\x45\x6d\xe8\x5f\xe7\x9d\x0d\xa1\x49\x9d\x4e\x62\x6b\xda\x8d\x06\x8a\x1b\xf5\x5d\xde\xbc\x93\x1b\xfa\x3e\x42\x0b\x75\xa7\xfd\xe0\x1c\x75\xf7\x3a\xc4\x8e\x33\x36\x7c\xf3\xe8\x73\xec\xea\x86\xc9\x80\xc3\xfd\x74\xe8\x7a\x3d\x1c\x6f\x98\xcc\x78\x5d\x74\xd4\x1c\x85\x84\xae\x28\xea\x88\x14\x1c\x8e\xdd\x51\xd4\x11\x29\x5b\x29\x07\x9f\x97\x74\x7c\x3b\xe9\x0c\x3f\x24\xe9\x0c\xc7\x20\x80\xc5\xfd\x91\xd8\x0a\x63\x5c\x25\xc6\xcd\x43\xb8\x49\xce\x90\x10\x9f\x47\x33\x27\xcb\xa9\xc8\xa3\x6c\x77\x0a\x9b\x2f\x00\x8f\x5f\x0e\x37\x77\x2c\xcb\x5f\x10\xfd\xd9\x02\xbe\x0a\x76\xe2\x71\x6d\x01\xc7\x6e\x00\x27\xbb\xed\x8e\xce\xce\xbe\x8b\x07\x3c\x14\x40\x55\x4a\xb8\x60\x80\x97\x4a\xac\x06\xab\x90\x67\xc9\x24\xd3\x25\xe6\xa7\x9b\xf9\xa3\xd2\xc0\xae\xcb\x76\xdd\xb0\x19\x48\x65\xa1\x04\x97\xb6\xd8\x18\x37\xe2\x92\x81\x15\x2d\x2b\xce\xd4\xc7\x02\xad\x3c\x9f\x85\xdc\x74\xfb\x4b\xf1\x4b\xa9\xcd\xaa\x6c\xb2\x2e\xee\x7c\xae\x26\x08\x19\x5e\xf4\x4e\x36\x8b\x24\x4a\xd3\x72\x63\xf8\xcc\xcd\xe9\x6a\x0e\x6d\xb9\xdb\x35\x87\xce\xfa\x58\x73\xe8\x71\xac\xe6\xe0\xe4\x4c\xd4\xd7\xee\x40\x5b\xb3\xeb\xfe\x06\x45\xa2\x6f\xa3\xee\x63\x24\x38\x6b\x71\xa3\xf6\xe9\x24\xea\x6b\xec\x82\x31\x83\x69\x4f\x9e\xc7\x01\x7a\x1f\xe6\xb6\x1b\xe9\x32\x3b\x4d\x18\x37\xd2\x4b\x97\x3b\xef\xa9\xc7\xd0\x5f\x77\xd1\x6a\xe1\x4a\x25\xd7\x67\x31\xac\xdd\x93\x82\x12\xfe\xfb\xfa\xd7\x33\x37\x19\x3b\x19\xbf\xd0\x35\xa3\x85\x46\x16\x27\xe0\x75\xef\x2a\x85\xfe\x78\x84\x7a\x4a\x33\x13\x74\xbb\x06\xc9\x6b\xca\x21\xbb\x80\x77\xef\x2f\x6e\x2c\xd5\xcf\xa4\x40\x1b\xac\xa1\x34\xd7\x61\x46\xf7\x6b\xf3\x70\x61\x44\xaf\x59\x9e\xee\xe1\x42\xd2\xa5\x68\xe6\xaf\x32\x71\x93\xff\x95\x7b\xcd\x79\xee\xd3\x6d\x16\xb2\xc1\x07\x99\x29\xdc\x9a\xe3\x15\x4f\x60\x3d\x78\x2f\xf0\x4e\xc5\xcd\xc0\x0c\xf7\x82\xa1\x1a\x5a\xd1\xc7\xd7\x43\x0d\x5e\xd4\x55\x72\x86\x41\x15\x14\x45\x43\x1e\x43\x97\xaf\x76\x2c\xed\x30\x5c\xe7\x89\x89\x48\xc1\xec\x2a\xda\x7a\xcd\x64\x9d\x79\xc2\xac\xeb\xe6\x92\x2c\xc9\xf2\xdc\xc3\xe4\xaf\x28\x53\x07\xfc\x8d\xe6\x53\xba\xe0\x52\x37\x3a\xe1\x6d\xf0\x6e\x84\xfb\xd4\xc4\x91\xd3\x60\x64\x9a\xfa\xa3\xde\x0c\x16\x1d\xef\x5a\x9f\x3e\xb6\xe8\x92\xf6\xf1\xf5\xf8\x89\xbd\x62\x6c\x72\x5f\x59\xde\xca\xb6\x57\x5b\xa8\x40\x18\xda\x06\xc4\x15\x93\x70\xb1\xe1\x9c\x69\xc0\x92\xe2\xab\x6b\xb8\xf5\xc5\x32\x31\x90\x90\x5d\x6c\xb8\xaf\x09\xae\x73\x23\xe2\x6c\x57\x65\xe8\xc1\x80\x16\x46\x71\x4e\xd0\x0c\xcc\x7e\x20\x98\xd6\x69\x40\xf0\x2e\x1c\x8c\xaf\xbe\x38\x25\x69\x17\x0b\xbf\x01\x9a\x91\x96\x71\x5b\xb4\x93\x9d\x6c\x3f\xe9\xee\x13\xab\x0e\x3e\x19\x7f\xa3\x6c\x55\xb8\x9d\xa6\x93\x51\x5a\x2e\x3d\x60\x99\x01\x0f\x4b\x0e\xc3\xd2\x35\xac\xaf\x08\x9b\xb3\x0d\xa5\xf7\xf2\xab\x57\xf1\xf6\x64\x57\x0a\x91\x98\x41\x9b\xa4\x0c\x99\x8c\x87\x81\xb2\xf5\x9d\xc5\x78\x0d\x6e\xaf\x63\xfd\x9d\x4e\x26\xfe\x80\x99\x5a\xe3\x0b\x63\x7b\x9d\x77\x70\x8f\x20\xdb\x6f\x7f\x9c\xf6\x18\xb7\x32\x89\x5a\x67\x2f\x1a\xfc\xa1\xb7\xa6\xbc\x5b\xd1\x89\x6b\x05\xbc\xfe\xee\xf8\xd0\xcf\x66\xc7\x36\x62\xca\xa7\xda\x82\xc6\xb8\x16\x25\xde\x06\x2e\xe0\x38\x3c\x93\x44\x2c\x27\xbe\x23\xf8\x30\x43\x92\xff\x8e\x81\x44\xab\x69\xaf\x9f\x24\x1f\x27\xe6\x20\x66\x9d\xf0\x10\xac\x49\xb9\xf2\xcd\x03\x18\x1e\x00\xd9\xb5\x49\x3c\x36\xe8\xbb\x36\x87\xcf\xda\x1d\x50\xea\xbe\xfd\xe1\x09\xac\xdf\xb9\x2f\x3c\x64\x63\x40\x05\xf4\x6d\x2d\x75\x83\x36\x87\x47\x8f\xfb\xce\x7e\x54\x19\xac\xa7\xcf\x7e\x89\xed\x3f\x91\x41\x8f\x18\x8f\xc1\x0c\xff\x61\x30\xf5\xd5\xef\x50\x8f\xe9\xac\xe0\x40\x8a\x7a\x82\x4c\xe1\xbf\x58\x26\x9e\xbe\xf2\xf6\x0c\x5c\xfd\x54\xbf\x92\xbb\x8b\x5e\x29\xf7\x09\x48\xb5\x9c\xce\x60\x9f\x51\xcb\x7b\xfd\xe1\xce\x62\xbe\xbb\x7e\x7e\x72\x39\x1f\xaf\x8e\x87\x15\xc7\xdd\x2b\x18\xf7\xbe\x9d\x65\x2f\x60\x8b\x3c\xf7\x55\xaf\x2d\xcc\x47\xb1\x4b\xdb\xac\x9d\xd0\xed\x4a\xc0\x4f\x04\x6e\x2c\xbd\x0e\xcd\xae\x98\x5c\x14\x58\x31\x00\x79\xd9\xd0\x3d\xe2\xdd\xc1\x2e\xf7\x5a\xbe\x9d\x3e\xef\xce\xc4\xc3\xbd\x1e\xcd\xb3\xc3\xd2\x6c\x8f\x3b\x1c\xe8\x6a\x2e\x87\xae\xab\xea\xec\x16\x1c\x9e\xc5\x83\x3d\xfc\xfd\xb7\x7b\x3b\x95\x5c\x15\x67\x9b\x96\x69\x51\x65\xf9\xa0\x9f\x43\x3d\x72\x06\xea\x92\x5a\xb5\xf4\x4e\xa0\xc8\x78\xa3\x4a\xfb\xed\x37\xe4\xed\x33\x75\x99\x4e\x4e\x4b\xce\x46\xb2\xeb\x35\xab\x2c\xab\x07\x97\x1d\x78\xcf\x12\xaf\x58\xe6\x74\xc7\x92\x5e\xb1\x98\x8f\xc2\x56\x2b\xb0\xa4\x1d\x4d\x75\xfd\xcf\x0b\xa7\xa9\x2a\x0d\x03\x0b\xff\x59\x40\xfa\xe5\xdd\xfe\x13\x8e\x8f\xc1\xc2\xbf\x07\xe4\x6f\xbf\x99\xbb\x4a\x3e\xbc\xd5\xa0\x8b\x1b\x99\x8f\x8b\x7b\x2b\xc6\xe5\xbd\x15\x3b\x05\x6e\x3a\x89\x63\x85\xad\xab\x2c\xf0\x51\x97\x6b\x93\xfe\xcf\x86\xa7\x97\xb2\xa6\x3e\x30\x10\x5a\x66\x57\xaa\x86\x8f\xc2\xae\x40\xb3\x4a\x5d\x51\xf3\xcf\xa4\xd9\x68\x06\x52\xc1\xba\x94\xa2\x32\x20\x24\xf8\x4e\x5d\xc8\xa5\x2f\x87\x49\x25\xe3\x75\xf2\x79\x1a\x3c\x31\x87\x77\xef\xbb\xff\xa9\xb8\xcb\x21\xf3\x45\x2b\x21\x0f\x6f\x12\x6a\xe6\x8e\x1f\x4e\xbc\x8f\x17\xc1\xe1\x0a\xf3\x97\x8c\x73\x7d\xfc\x55\xaf\x88\xe1\xe5\x52\x2f\x24\xbe\x78\x13\xbc\x23\xe3\xe3\xdd\xef\x0c\xae\xb0\xc5\xe3\xa1\x80\x61\x14\xe2\x3e\xe1\x3a\xdd\x10\x5d\x75\x11\x1c\x98\x0d\xd0\xa5\x86\x68\x0b\x5c\x22\x3f\x14\xca\xf4\x0e\x20\x45\x93\xe8\x01\x4c\xfc\x92\xe2\xb0\xa4\x4e\xad\x23\x3e\x05\x92\x3d\xff\x7a\x60\x12\x90\xcc\x37\x88\xa3\x38\xa6\x93\xb7\xa1\x0c\x9d\xd9\x16\x98\x61\xe0\xa1\x70\xf6\x6f\x24\x52\x40\xc3\x48\x80\x94\xee\xfe\x1c\xa6\x22\xfe\x5f\x56\xa4\x3f\x21\xac\xc1\xd3\x11\x60\x45\xec\x5b\xf7\x41\x1b\x1d\x19\x82\x4b\x27\xd5\x2d\x68\x89\xfc\x50\x60\xf7\x9d\x60\x33\xea\x00\x09\xbf\x5f\xba\x53\xec\x93\xe0\x47\xee\x8c\xa0\x47\x46\xec\xc7\x8e\xbc\xd8\x42\x8e\x9a\x82\x2d\xe4\x88\xfc\x50\xe4\x7a\x3d\x4f\x12\x90\x44\x0f\xe1\xe8\xde\x30\x1a\xa9\x59\xe9\x88\x4f\x08\x25\xf9\x37\x02\xe5\xca\x37\x49\xfb\xa0\xf4\xe6\x0f\xa1\xf4\xdd\xc6\x16\x96\x9e\xfe\x50\x30\xfb\xdd\x54\x82\xa6\x1f\xc8\x31\x36\xbd\x32\x07\xa7\xef\x88\x3a\xea\x13\xe2\xe9\xd5\x8e\x00\xba\x0e\x3d\xd8\x3e\x44\x83\x0b\xb3\x5e\x03\x16\xaf\x6b\x2c\xa4\x17\x36\x79\xef\x0d\x4f\x1c\x4a\x83\x2d\x7e\x16\xb2\xce\x72\x58\x2c\xe2\xf8\x2b\x8b\x9d\xda\xc4\xc2\x02\x6c\xf1\xb2\x61\x6d\xd6\x6b\x25\xec\xf4\x6e\xfa\xff\x00\x00\x00\xff\xff\x2f\xcd\xaf\xc0\x4e\x2d\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 11598, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
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
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
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
