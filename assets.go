// Code generated by Lorca. DO NOT EDIT.
package main

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"time"
)

var assets = map[string][]byte{}

var FS = &fs{}

type fs struct {}

func (fs *fs) Open(name string) (http.File, error) {
	if name == "/" {
		return fs, nil;
	}
	b, ok := assets[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	return &file{name: name, size: len(b), Reader: bytes.NewReader(b)}, nil
}

func (fs *fs) Close() error { return nil }
func (fs *fs) Read(p []byte) (int, error) { return 0, nil }
func (fs *fs) Seek(offset int64, whence int) (int64, error) { return 0, nil }
func (fs *fs) Stat() (os.FileInfo, error) { return fs, nil }
func (fs *fs) Name() string { return "/" }
func (fs *fs) Size() int64 { return 0 }
func (fs *fs) Mode() os.FileMode { return 0755}
func (fs *fs) ModTime() time.Time{ return time.Time{} }
func (fs *fs) IsDir() bool { return true }
func (fs *fs) Sys() interface{} { return nil }
func (fs *fs) Readdir(count int) ([]os.FileInfo, error) {
	files := []os.FileInfo{}
	for name, data := range assets {
		files = append(files, &file{name: name, size: len(data), Reader: bytes.NewReader(data)})
	}
	return files, nil
}

type file struct {
	name string
	size int
	*bytes.Reader 
}

func (f *file) Close() error { return nil }
func (f *file) Readdir(count int) ([]os.FileInfo, error) { return nil, errors.New("not supported") }
func (f *file) Stat() (os.FileInfo, error) { return f, nil }
func (f *file) Name() string { return f.name }
func (f *file) Size() int64 { return int64(f.size) }
func (f *file) Mode() os.FileMode { return 0644 }
func (f *file) ModTime() time.Time{ return time.Time{} }
func (f *file) IsDir() bool { return false }
func (f *file) Sys() interface{} { return nil }

func init() {
	assets["/assets/favicon.ico"] = []byte{0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x10, 0x10, 0x00, 0x00, 0x01, 0x00, 0x20, 0x00, 0x68, 0x04, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x28, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x01, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4c, 0x5c, 0xe0, 0xff, 0x3f, 0x4d, 0xa5, 0xff, 0x3e, 0x51, 0xa9, 0xff, 0x3c, 0x54, 0xa8, 0xff, 0x3a, 0x57, 0xa8, 0xff, 0x38, 0x5a, 0xa8, 0xff, 0x36, 0x5d, 0xa8, 0xff, 0x34, 0x60, 0xa8, 0xff, 0x32, 0x64, 0xa8, 0xff, 0x31, 0x67, 0xa8, 0xff, 0x2f, 0x6a, 0xa8, 0xff, 0x2d, 0x6d, 0xa8, 0xff, 0x2b, 0x70, 0xa8, 0xff, 0x29, 0x74, 0xa9, 0xff, 0x27, 0x74, 0xa5, 0xff, 0x27, 0x9c, 0xe0, 0xff, 0x42, 0x46, 0xa5, 0xff, 0x20, 0x11, 0x00, 0xff, 0x23, 0x15, 0x04, 0xff, 0x23, 0x14, 0x02, 0xff, 0x23, 0x14, 0x02, 0xff, 0x23, 0x14, 0x02, 0xff, 0x23, 0x14, 0x02, 0xff, 0x23, 0x13, 0x02, 0xff, 0x23, 0x13, 0x02, 0xff, 0x23, 0x13, 0x02, 0xff, 0x23, 0x13, 0x02, 0xff, 0x23, 0x13, 0x02, 0xff, 0x24, 0x12, 0x02, 0xff, 0x24, 0x14, 0x04, 0xff, 0x24, 0x0c, 0x00, 0xff, 0x27, 0x75, 0xa5, 0xff, 0x45, 0x45, 0xa9, 0xff, 0x22, 0x16, 0x04, 0xff, 0x25, 0x19, 0x10, 0xff, 0x24, 0x19, 0x0e, 0xff, 0x24, 0x19, 0x0e, 0xff, 0x24, 0x19, 0x0e, 0xff, 0x25, 0x1a, 0x10, 0xff, 0x25, 0x1c, 0x12, 0xff, 0x25, 0x1c, 0x12, 0xff, 0x24, 0x1a, 0x0f, 0xff, 0x24, 0x19, 0x0e, 0xff, 0x24, 0x19, 0x0e, 0xff, 0x24, 0x19, 0x0e, 0xff, 0x24, 0x1a, 0x10, 0xff, 0x24, 0x13, 0x04, 0xff, 0x29, 0x75, 0xa9, 0xff, 0x46, 0x42, 0xa8, 0xff, 0x22, 0x15, 0x02, 0xff, 0x24, 0x19, 0x0e, 0xff, 0x24, 0x18, 0x0c, 0xff, 0x24, 0x18, 0x0c, 0xff, 0x25, 0x1a, 0x0e, 0xff, 0x24, 0x12, 0x05, 0xff, 0x20, 0x0b, 0x00, 0xff, 0x20, 0x0b, 0x00, 0xff, 0x24, 0x14, 0x07, 0xff, 0x24, 0x19, 0x0d, 0xff, 0x24, 0x18, 0x0c, 0xff, 0x24, 0x18, 0x0c, 0xff, 0x24, 0x19, 0x0e, 0xff, 0x24, 0x12, 0x02, 0xff, 0x2a, 0x72, 0xa8, 0xff, 0x47, 0x40, 0xa8, 0xff, 0x22, 0x16, 0x02, 0xff, 0x24, 0x19, 0x0e, 0xff, 0x24, 0x18, 0x0c, 0xff, 0x25, 0x1a, 0x0f, 0xff, 0x23, 0x11, 0x02, 0xff, 0x24, 0x37, 0x3d, 0xff, 0x34, 0x5c, 0x6d, 0xff, 0x34, 0x5f, 0x74, 0xff, 0x24, 0x31, 0x34, 0xff, 0x24, 0x15, 0x07, 0xff, 0x24, 0x19, 0x0e, 0xff, 0x24, 0x18, 0x0c, 0xff, 0x24, 0x19, 0x0e, 0xff, 0x24, 0x12, 0x02, 0xff, 0x2c, 0x6f, 0xa8, 0xff, 0x49, 0x3d, 0xa8, 0xff, 0x22, 0x16, 0x02, 0xff, 0x24, 0x18, 0x0e, 0xff, 0x24, 0x19, 0x0f, 0xff, 0x25, 0x14, 0x04, 0xff, 0x1e, 0x25, 0x2a, 0xff, 0x7b, 0xbb, 0xec, 0xff, 0xba, 0xec, 0xff, 0xff, 0x69, 0xb1, 0xdc, 0xff, 0x1c, 0x1f, 0x1a, 0xff, 0x25, 0x13, 0x00, 0xff, 0x25, 0x18, 0x0b, 0xff, 0x24, 0x19, 0x0d, 0xff, 0x24, 0x19, 0x0e, 0xff, 0x23, 0x13, 0x02, 0xff, 0x2d, 0x6d, 0xa8, 0xff, 0x4a, 0x3a, 0xa8, 0xff, 0x21, 0x16, 0x02, 0xff, 0x25, 0x19, 0x0f, 0xff, 0x26, 0x18, 0x0a, 0xff, 0x1a, 0x12, 0x0b, 0xff, 0x64, 0x89, 0xc8, 0xff, 0xc4, 0xd0, 0xff, 0xff, 0x92, 0xae, 0xf1, 0xff, 0x27, 0x35, 0x51, 0xff, 0x26, 0x21, 0x25, 0xff, 0x23, 0x27, 0x36, 0xff, 0x21, 0x1a, 0x18, 0xff, 0x23, 0x16, 0x07, 0xff, 0x24, 0x1a, 0x0f, 0xff, 0x23, 0x13, 0x02, 0xff, 0x2f, 0x6a, 0xa8, 0xff, 0x4c, 0x38, 0xa8, 0xff, 0x21, 0x16, 0x02, 0xff, 0x26, 0x1a, 0x10, 0xff, 0x1e, 0x11, 0x02, 0xff, 0x4f, 0x5c, 0x99, 0xff, 0xc0, 0xbd, 0xff, 0xff, 0xb2, 0xb3, 0xff, 0xff, 0x35, 0x3d, 0x61, 0xff, 0x23, 0x26, 0x3d, 0xff, 0x81, 0xae, 0xff, 0xff, 0x9f, 0xca, 0xff, 0xff, 0x64, 0x83, 0xca, 0xff, 0x20, 0x15, 0x11, 0xff, 0x25, 0x1a, 0x0d, 0xff, 0x23, 0x13, 0x03, 0xff, 0x30, 0x67, 0xa8, 0xff, 0x4d, 0x35, 0xa8, 0xff, 0x22, 0x17, 0x04, 0xff, 0x25, 0x17, 0x07, 0xff, 0x22, 0x1c, 0x28, 0xff, 0xa1, 0x91, 0xf3, 0xff, 0xcf, 0xaf, 0xff, 0xff, 0x58, 0x54, 0xa1, 0xff, 0x1c, 0x11, 0x01, 0xff, 0x20, 0x11, 0x00, 0xff, 0x49, 0x4a, 0x84, 0xff, 0xc0, 0xc5, 0xff, 0xff, 0xa5, 0xae, 0xfd, 0xff, 0x27, 0x21, 0x35, 0xff, 0x23, 0x17, 0x05, 0xff, 0x24, 0x14, 0x04, 0xff, 0x32, 0x65, 0xa8, 0xff, 0x4f, 0x33, 0xa8, 0xff, 0x22, 0x17, 0x04, 0xff, 0x24, 0x17, 0x07, 0xff, 0x25, 0x1b, 0x28, 0xff, 0xb6, 0x87, 0xff, 0xff, 0x8a, 0x68, 0xd3, 0xff, 0x1b, 0x12, 0x0f, 0xff, 0x26, 0x19, 0x0b, 0xff, 0x26, 0x1a, 0x0e, 0xff, 0x18, 0x0d, 0x03, 0xff, 0x72, 0x5f, 0xb9, 0xff, 0xb6, 0x99, 0xff, 0xff, 0x2a, 0x1e, 0x31, 0xff, 0x22, 0x17, 0x05, 0xff, 0x24, 0x14, 0x04, 0xff, 0x33, 0x62, 0xa8, 0xff, 0x50, 0x30, 0xa8, 0xff, 0x22, 0x17, 0x04, 0xff, 0x22, 0x17, 0x06, 0xff, 0x30, 0x1d, 0x2d, 0xff, 0x93, 0x55, 0xee, 0xff, 0x2f, 0x1d, 0x33, 0xff, 0x21, 0x16, 0x02, 0xff, 0x25, 0x19, 0x0f, 0xff, 0x24, 0x18, 0x0e, 0xff, 0x23, 0x18, 0x06, 0xff, 0x26, 0x18, 0x1f, 0xff, 0x88, 0x53, 0xe6, 0xff, 0x36, 0x21, 0x3b, 0xff, 0x20, 0x17, 0x03, 0xff, 0x24, 0x14, 0x04, 0xff, 0x35, 0x60, 0xa8, 0xff, 0x52, 0x2e, 0xa8, 0xff, 0x21, 0x17, 0x03, 0xff, 0x22, 0x18, 0x07, 0xff, 0x2f, 0x1a, 0x2c, 0xff, 0x3b, 0x19, 0x56, 0xff, 0x1f, 0x17, 0x00, 0xff, 0x25, 0x19, 0x10, 0xff, 0x24, 0x18, 0x0c, 0xff, 0x24, 0x18, 0x0c, 0xff, 0x25, 0x19, 0x0f, 0xff, 0x20, 0x17, 0x00, 0xff, 0x35, 0x18, 0x46, 0xff, 0x30, 0x1a, 0x30, 0xff, 0x22, 0x18, 0x06, 0xff, 0x23, 0x14, 0x03, 0xff, 0x36, 0x5d, 0xa8, 0xff, 0x53, 0x2b, 0xa8, 0xff, 0x21, 0x17, 0x02, 0xff, 0x25, 0x18, 0x0f, 0xff, 0x23, 0x18, 0x0a, 0xff, 0x22, 0x18, 0x05, 0xff, 0x25, 0x19, 0x0e, 0xff, 0x24, 0x18, 0x0c, 0xff, 0x24, 0x18, 0x0c, 0xff, 0x24, 0x18, 0x0c, 0xff, 0x24, 0x18, 0x0c, 0xff, 0x25, 0x18, 0x0e, 0xff, 0x22, 0x18, 0x06, 0xff, 0x23, 0x18, 0x09, 0xff, 0x25, 0x19, 0x0f, 0xff, 0x23, 0x14, 0x02, 0xff, 0x38, 0x5b, 0xa8, 0xff, 0x55, 0x29, 0xa9, 0xff, 0x22, 0x17, 0x04, 0xff, 0x25, 0x18, 0x10, 0xff, 0x25, 0x18, 0x0f, 0xff, 0x25, 0x19, 0x10, 0xff, 0x25, 0x18, 0x0e, 0xff, 0x25, 0x18, 0x0e, 0xff, 0x25, 0x18, 0x0e, 0xff, 0x25, 0x18, 0x0e, 0xff, 0x24, 0x19, 0x0e, 0xff, 0x24, 0x19, 0x0e, 0xff, 0x25, 0x19, 0x10, 0xff, 0x25, 0x19, 0x0f, 0xff, 0x25, 0x1a, 0x10, 0xff, 0x23, 0x15, 0x04, 0xff, 0x39, 0x58, 0xa9, 0xff, 0x55, 0x26, 0xa5, 0xff, 0x1d, 0x16, 0x00, 0xff, 0x22, 0x17, 0x04, 0xff, 0x21, 0x17, 0x02, 0xff, 0x21, 0x17, 0x02, 0xff, 0x21, 0x16, 0x02, 0xff, 0x21, 0x16, 0x02, 0xff, 0x21, 0x16, 0x02, 0xff, 0x22, 0x16, 0x02, 0xff, 0x22, 0x16, 0x02, 0xff, 0x22, 0x15, 0x02, 0xff, 0x22, 0x15, 0x02, 0xff, 0x22, 0x15, 0x02, 0xff, 0x22, 0x16, 0x04, 0xff, 0x21, 0x10, 0x00, 0xff, 0x3a, 0x54, 0xa5, 0xff, 0x6a, 0x28, 0xe0, 0xff, 0x55, 0x26, 0xa5, 0xff, 0x54, 0x2a, 0xa9, 0xff, 0x52, 0x2d, 0xa8, 0xff, 0x51, 0x30, 0xa8, 0xff, 0x4f, 0x33, 0xa8, 0xff, 0x4d, 0x36, 0xa8, 0xff, 0x4b, 0x39, 0xa8, 0xff, 0x49, 0x3c, 0xa8, 0xff, 0x47, 0x40, 0xa8, 0xff, 0x46, 0x43, 0xa8, 0xff, 0x44, 0x46, 0xa8, 0xff, 0x42, 0x49, 0xa8, 0xff, 0x40, 0x4c, 0xa9, 0xff, 0x3e, 0x4e, 0xa5, 0xff, 0x45, 0x67, 0xe0, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	assets["/index.html"] = []byte{0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0x0a, 0x20, 0x20, 0x3c, 0x68, 0x65, 0x61, 0x64, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x56, 0x41, 0x4c, 0x4f, 0x4c, 0x49, 0x4e, 0x45, 0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x22, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x22, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x2f, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x63, 0x73, 0x73, 0x22, 0x20, 0x2f, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x22, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x22, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x63, 0x73, 0x73, 0x22, 0x20, 0x2f, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x22, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x78, 0x2d, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x69, 0x63, 0x6f, 0x22, 0x20, 0x2f, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x64, 0x6e, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x73, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x61, 0x6e, 0x74, 0x22, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x22, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x22, 0x20, 0x2f, 0x3e, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x63, 0x6b, 0x28, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x29, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x28, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x20, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x7e, 0x20, 0x5b, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5d, 0x20, 0x3d, 0x3e, 0x20, 0x22, 0x20, 0x2b, 0x20, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x29, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x0a, 0x20, 0x20, 0x3c, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x3e, 0x0a, 0x0a, 0x20, 0x20, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x20, 0x6f, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x3d, 0x22, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x28, 0x28, 0x6e, 0x65, 0x77, 0x20, 0x44, 0x61, 0x74, 0x65, 0x28, 0x29, 0x29, 0x2e, 0x74, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x28, 0x29, 0x29, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3e, 0x56, 0x41, 0x4c, 0x4f, 0x4c, 0x49, 0x4e, 0x45, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x62, 0x74, 0x6e, 0x20, 0x62, 0x74, 0x6e, 0x2d, 0x2d, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x62, 0x74, 0x6e, 0x5f, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x62, 0x74, 0x6e, 0x5f, 0x5f, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x62, 0x74, 0x6e, 0x5f, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3e, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x20, 0x6d, 0x6f, 0x64, 0x65, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x62, 0x74, 0x6e, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x62, 0x74, 0x6e, 0x5f, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x62, 0x74, 0x6e, 0x5f, 0x5f, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x62, 0x74, 0x6e, 0x5f, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x20, 0x6d, 0x6f, 0x64, 0x65, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0x0a, 0x20, 0x20, 0x3c, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0x0a, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0x0a}
	assets["/style/button_style.css"] = []byte{0x3a, 0x72, 0x6f, 0x6f, 0x74, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x23, 0x30, 0x66, 0x31, 0x39, 0x32, 0x33, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2f, 0x2a, 0x20, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x20, 0x2a, 0x2f, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x6d, 0x6f, 0x7a, 0x2d, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x20, 0x6e, 0x6f, 0x6e, 0x65, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x77, 0x65, 0x62, 0x6b, 0x69, 0x74, 0x2d, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x20, 0x6e, 0x6f, 0x6e, 0x65, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x20, 0x6e, 0x6f, 0x6e, 0x65, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x3a, 0x20, 0x6e, 0x6f, 0x6e, 0x65, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x20, 0x6e, 0x6f, 0x6e, 0x65, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x30, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x3a, 0x20, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2f, 0x2a, 0x20, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x20, 0x2a, 0x2f, 0x0d, 0x0a, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x2d, 0x68, 0x6f, 0x76, 0x65, 0x72, 0x3a, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x2d, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x23, 0x37, 0x64, 0x38, 0x30, 0x38, 0x32, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x23, 0x65, 0x63, 0x65, 0x38, 0x65, 0x31, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x2d, 0x68, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x23, 0x66, 0x66, 0x34, 0x36, 0x35, 0x35, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2d, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x62, 0x69, 0x74, 0x73, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x62, 0x69, 0x74, 0x73, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x2d, 0x68, 0x6f, 0x76, 0x65, 0x72, 0x3a, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x38, 0x70, 0x78, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x2d, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x3a, 0x20, 0x32, 0x30, 0x70, 0x78, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x20, 0x75, 0x70, 0x70, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x62, 0x6f, 0x6c, 0x64, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x20, 0x31, 0x34, 0x70, 0x78, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x30, 0x2e, 0x31, 0x35, 0x73, 0x20, 0x65, 0x61, 0x73, 0x65, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x32, 0x35, 0x30, 0x70, 0x78, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x3a, 0x3a, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x2c, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x3a, 0x3a, 0x61, 0x66, 0x74, 0x65, 0x72, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3a, 0x20, 0x22, 0x22, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x3a, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x72, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x30, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x6c, 0x65, 0x66, 0x74, 0x3a, 0x20, 0x30, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x63, 0x61, 0x6c, 0x63, 0x28, 0x35, 0x30, 0x25, 0x20, 0x2d, 0x20, 0x35, 0x70, 0x78, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x3a, 0x20, 0x31, 0x70, 0x78, 0x20, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x30, 0x2e, 0x31, 0x35, 0x73, 0x20, 0x65, 0x61, 0x73, 0x65, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x3a, 0x3a, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x74, 0x6f, 0x70, 0x3a, 0x20, 0x30, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x2d, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x30, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x3a, 0x3a, 0x61, 0x66, 0x74, 0x65, 0x72, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x3a, 0x20, 0x30, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x74, 0x6f, 0x70, 0x2d, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x30, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x3a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2c, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x3a, 0x66, 0x6f, 0x63, 0x75, 0x73, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x6f, 0x75, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x3a, 0x20, 0x6e, 0x6f, 0x6e, 0x65, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x3a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x3a, 0x3a, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x2c, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x3a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x3a, 0x3a, 0x61, 0x66, 0x74, 0x65, 0x72, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x72, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x33, 0x70, 0x78, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x6c, 0x65, 0x66, 0x74, 0x3a, 0x20, 0x33, 0x70, 0x78, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x3a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x3a, 0x3a, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x74, 0x6f, 0x70, 0x3a, 0x20, 0x33, 0x70, 0x78, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x3a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x3a, 0x3a, 0x61, 0x66, 0x74, 0x65, 0x72, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x3a, 0x20, 0x33, 0x70, 0x78, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x5f, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x3a, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x32, 0x30, 0x70, 0x78, 0x20, 0x33, 0x30, 0x70, 0x78, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x6f, 0x76, 0x65, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x20, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x6f, 0x78, 0x2d, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x3a, 0x20, 0x69, 0x6e, 0x73, 0x65, 0x74, 0x20, 0x30, 0x70, 0x78, 0x20, 0x30, 0x70, 0x78, 0x20, 0x30, 0x70, 0x78, 0x20, 0x31, 0x70, 0x78, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2d, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x5f, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3a, 0x20, 0x22, 0x22, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x3a, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x74, 0x6f, 0x70, 0x3a, 0x20, 0x30, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x6c, 0x65, 0x66, 0x74, 0x3a, 0x20, 0x30, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x32, 0x70, 0x78, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x32, 0x70, 0x78, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x62, 0x69, 0x74, 0x73, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x5f, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x61, 0x66, 0x74, 0x65, 0x72, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3a, 0x20, 0x22, 0x22, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x3a, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x72, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x30, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x3a, 0x20, 0x30, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x34, 0x70, 0x78, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x34, 0x70, 0x78, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x62, 0x69, 0x74, 0x73, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x30, 0x2e, 0x32, 0x73, 0x20, 0x65, 0x61, 0x73, 0x65, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x5f, 0x5f, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x3a, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x74, 0x6f, 0x70, 0x3a, 0x20, 0x30, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x3a, 0x20, 0x2d, 0x31, 0x70, 0x78, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x6c, 0x65, 0x66, 0x74, 0x3a, 0x20, 0x2d, 0x38, 0x70, 0x78, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x30, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x68, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x20, 0x73, 0x6b, 0x65, 0x77, 0x28, 0x2d, 0x31, 0x35, 0x64, 0x65, 0x67, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x30, 0x2e, 0x32, 0x73, 0x20, 0x65, 0x61, 0x73, 0x65, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x5f, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x3a, 0x68, 0x6f, 0x76, 0x65, 0x72, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x2d, 0x68, 0x6f, 0x76, 0x65, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x3a, 0x68, 0x6f, 0x76, 0x65, 0x72, 0x20, 0x2e, 0x62, 0x74, 0x6e, 0x5f, 0x5f, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x63, 0x61, 0x6c, 0x63, 0x28, 0x31, 0x30, 0x30, 0x25, 0x20, 0x2b, 0x20, 0x31, 0x35, 0x70, 0x78, 0x29, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x3a, 0x68, 0x6f, 0x76, 0x65, 0x72, 0x20, 0x2e, 0x62, 0x74, 0x6e, 0x5f, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x61, 0x66, 0x74, 0x65, 0x72, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x62, 0x69, 0x74, 0x73, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x2d, 0x68, 0x6f, 0x76, 0x65, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x62, 0x74, 0x6e, 0x2d, 0x2d, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x68, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2d, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x68, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x2d, 0x68, 0x6f, 0x76, 0x65, 0x72, 0x3a, 0x20, 0x23, 0x65, 0x63, 0x65, 0x38, 0x65, 0x31, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x62, 0x69, 0x74, 0x73, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x2d, 0x68, 0x6f, 0x76, 0x65, 0x72, 0x3a, 0x20, 0x23, 0x65, 0x63, 0x65, 0x38, 0x65, 0x31, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a}
	assets["/style/general_style.css"] = []byte{0x62, 0x6f, 0x64, 0x79, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x6d, 0x69, 0x6e, 0x2d, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x39, 0x30, 0x76, 0x68, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x76, 0x61, 0x72, 0x28, 0x2d, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x3a, 0x20, 0x75, 0x72, 0x6c, 0x28, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x64, 0x75, 0x63, 0x6b, 0x64, 0x75, 0x63, 0x6b, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x75, 0x2f, 0x3f, 0x75, 0x3d, 0x68, 0x74, 0x74, 0x70, 0x73, 0x25, 0x33, 0x41, 0x25, 0x32, 0x46, 0x25, 0x32, 0x46, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x73, 0x64, 0x65, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x25, 0x32, 0x46, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x25, 0x32, 0x46, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x25, 0x32, 0x46, 0x76, 0x69, 0x70, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x6f, 0x6c, 0x2d, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x61, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6e, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x47, 0x68, 0x6f, 0x61, 0x32, 0x57, 0x55, 0x6d, 0x5a, 0x71, 0x61, 0x72, 0x61, 0x57, 0x6b, 0x70, 0x4a, 0x52, 0x6d, 0x61, 0x47, 0x74, 0x72, 0x72, 0x57, 0x78, 0x72, 0x62, 0x51, 0x2e, 0x6a, 0x70, 0x67, 0x26, 0x66, 0x3d, 0x31, 0x26, 0x6e, 0x6f, 0x66, 0x62, 0x3d, 0x31, 0x22, 0x29, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x3a, 0x20, 0x6e, 0x6f, 0x2d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x20, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x34, 0x30, 0x25, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x3a, 0x20, 0x66, 0x6c, 0x65, 0x78, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66, 0x6c, 0x65, 0x78, 0x2d, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x2d, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x3a, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3a, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x74, 0x6f, 0x70, 0x3a, 0x20, 0x34, 0x30, 0x25, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x28, 0x31, 0x30, 0x25, 0x2c, 0x20, 0x2d, 0x35, 0x30, 0x25, 0x29, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x2d, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x3a, 0x20, 0x32, 0x30, 0x25, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a, 0x0d, 0x0a, 0x2e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x20, 0x7b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x3a, 0x20, 0x22, 0x56, 0x41, 0x4c, 0x4f, 0x52, 0x41, 0x4e, 0x54, 0x22, 0x2c, 0x20, 0x73, 0x61, 0x6e, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x69, 0x66, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x77, 0x68, 0x69, 0x74, 0x65, 0x3b, 0x0d, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x20, 0x32, 0x65, 0x6d, 0x3b, 0x0d, 0x0a, 0x7d, 0x0d, 0x0a}
}
