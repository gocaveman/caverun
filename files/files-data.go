// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package files

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// EmbeddedAssets statically implements the virtual filesystem provided to vfsgen.
var EmbeddedAssets = func() http.FileSystem {
	mustUnmarshalTextTime := func(text string) time.Time {
		var t time.Time
		err := t.UnmarshalText([]byte(text))
		if err != nil {
			panic(err)
		}
		return t
	}

	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: mustUnmarshalTextTime("2018-08-24T01:44:15Z"),
		},
		"/static": &vfsgen۰DirInfo{
			name:    "static",
			modTime: mustUnmarshalTextTime("2018-08-24T01:44:22Z"),
		},
		"/static/index.html": &vfsgen۰CompressedFileInfo{
			name:             "index.html",
			modTime:          mustUnmarshalTextTime("2018-08-24T21:53:48Z"),
			uncompressedSize: 19025,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x7c\xeb\x92\xdb\x36\xb2\xf0\x6f\xe9\x29\x3a\xca\xd6\x4a\x53\x16\xa9\x19\x5f\x92\xac\x46\xa3\xc4\x76\x9c\x64\xb7\x92\x8d\xcb\xce\xb7\x5b\x5b\x5b\x5b\x53\x20\xd9\x14\x61\x81\x00\x03\x80\xd2\xe8\xf3\xf1\x53\x9c\xbf\xe7\xe9\xce\x93\x9c\xc2\x8d\x37\x69\xc6\x93\xc4\x71\xfc\xc3\x23\x01\x60\xa3\xef\xdd\xe8\x06\xb5\xfa\xe4\xeb\x1f\x9f\xff\xf4\xaf\x97\x2f\xa0\xd0\x25\x5b\x8f\x57\xe1\x0f\x92\x6c\x3d\x1e\xad\x34\xd5\x0c\xd7\x29\xd9\xa1\xac\xf9\x6a\xe1\xbe\x8e\xc7\x00\x00\x2b\x46\xf9\x16\x24\xb2\xab\x89\xd2\x07\x86\xaa\x40\xd4\x13\xd0\x87\x0a\xaf\x26\x1a\x6f\xf4\x22\x55\x6a\x02\x85\xc4\xfc\x6a\x52\x68\x5d\xa9\xe5\x62\x91\x0b\xae\x55\xbc\x11\x62\xc3\x90\x54\x54\xc5\xa9\x28\xcd\xba\x2f\x73\x52\x52\x76\xb8\x7a\x25\x12\xa1\xc5\xf2\xd1\xf9\xf9\xfc\xf1\xf9\xf9\xfc\xc9\xf9\xf9\xfc\xf3\xf3\xf3\xff\xfa\x81\x68\x94\x94\xb0\x07\x7f\x4d\x05\x57\x93\xf5\xaf\x47\xa0\xe6\xd5\x76\x63\x77\xdd\xd5\xa8\x69\x7e\xf8\xea\x22\xbe\x88\x2f\x1e\x2f\x32\xaa\x74\x18\x8b\x4b\xca\x63\xf3\xf0\x7a\x3c\x5e\x59\xd8\xeb\xf1\xa7\xa4\xaa\x20\x26\x55\xc5\x68\x4a\x34\x15\x3c\x8a\xf6\x92\x54\xb0\x86\x78\x17\x69\x21\x58\x42\x64\xef\xcb\xf5\x75\x2a\xb8\x46\xae\xe1\xed\x78\x94\x90\x74\xbb\x91\xa2\xe6\xd9\x12\x3e\x7d\xfc\xe2\xe2\xe9\xc3\xaf\x2f\xc7\xef\x3c\xd0\x5d\x94\x0b\xa1\x51\xbe\x77\xa1\x97\x43\x64\x98\x20\x18\x0e\xd7\x27\x8c\xa4\xdb\xcb\xf1\x28\x15\x4c\xc8\x25\xec\x0b\xaa\xf1\x72\x3c\xaa\x48\x96\x51\xbe\x59\xc2\x93\xea\xe6\x72\x3c\x32\x22\x88\x1c\xbb\x97\x50\x0a\x2e\x54\x45\x52\xb3\xce\xae\x8f\xec\xb7\x25\x54\xd2\x0c\x89\x1d\xca\x9c\x89\xfd\x12\x54\x2a\x05\x63\x06\x95\xd5\xc2\x73\x64\xbc\x5a\x38\x35\x59\x25\x22\x3b\x98\xef\x19\xdd\x01\xcd\xae\x26\xa4\xaa\xac\x8c\x56\xbb\xc8\xe0\x6d\x86\x28\x57\x15\x95\x18\x44\xb7\x8b\x38\xd9\xd1\x8d\x63\x64\x26\xc9\x1e\xa5\x9d\x00\xc8\xe9\x0d\x66\xfe\xf3\x2e\x2a\x45\x66\xe4\xeb\x56\x4c\xfc\x70\x29\x12\xca\x30\x4a\x24\x92\x6d\x54\x09\xca\xf5\xd5\xe4\xf1\xf9\x79\x98\x4e\x19\xad\xaa\x06\x06\xa9\x2a\xfb\x69\xed\xbf\xaf\x3e\x89\x22\x28\x29\xa7\xd1\x8e\x48\x4a\xb8\x86\x28\x6a\xe6\x76\x11\xa3\x4a\x7b\x05\xef\x8c\x44\x9a\x32\x84\xaf\x52\x46\xd3\xed\xd5\x44\xa5\x12\x91\x5f\x4d\x93\x9a\xb2\x6c\x3a\x59\x37\xab\xfb\xeb\x23\x92\x1a\xf2\xba\xd3\x76\x01\x4d\x05\x07\x46\xe4\x06\xd7\x16\xc2\x6a\xe1\xc6\x7a\x70\x16\x77\x03\xea\xed\xe3\xf5\xec\x68\xa3\x76\x85\xb3\xdb\x67\x66\x37\xf8\x33\x29\xab\x4b\x78\x65\xcc\xf9\x78\xc9\xad\x28\x1c\xed\xd1\x9b\xbe\x2f\xc7\x36\xe2\x7a\x83\x1c\x25\xd1\xf8\x9b\xf8\x46\x79\x55\xeb\x8f\xc6\xb7\x6f\x05\x7c\xeb\xb1\xfe\x03\x78\x66\x8c\x7e\x83\xfc\x37\xf1\x4b\xe4\x39\xa3\x1c\xaf\x13\xc1\x3e\x1e\xdb\x9e\x3b\xc4\xff\x00\x96\x65\x22\xdd\xa2\xfc\x4d\x1c\x4b\x99\xa8\x3f\x9e\x65\x7e\x6d\x11\xfe\x23\x0c\x92\xea\xdf\xc4\x26\x89\x4a\x0b\x89\x1f\xcf\x14\xa9\xfe\x03\xb8\xa4\x50\x6b\xca\x37\xea\x37\xb1\x2a\x00\xf9\x68\xbc\x7a\xdd\xd9\xf0\x43\x33\x2c\x8c\xfa\x70\xbe\x38\x11\xcf\x9b\x48\x1f\x92\x23\x9b\x96\x98\x4c\x20\xa3\x1b\x31\x81\x8c\xc8\xad\x8b\xf6\x26\x46\x87\xa8\x1d\x31\xcc\xdb\x00\xdc\x3e\x1d\x29\x9a\xa1\x63\xa7\x93\x4e\xac\xb4\xa8\x42\x66\x00\x57\xf0\x89\xcf\x11\xd6\x06\x99\xa3\x87\xd6\xc7\x00\xdb\xd4\x76\x85\xe5\xda\xa6\xb7\x58\xf6\x9e\x0e\xd9\xee\xc8\x3e\x67\x13\x23\x69\x17\xf8\x8f\x76\xc6\xfc\x1b\xad\x76\x51\xa2\x39\x74\xd0\x6b\x95\x67\x8f\x2c\x15\xa5\x8d\x77\x61\xf9\xc8\x6b\xc6\x9a\x64\xd9\x75\x22\x6e\x5a\x95\x68\x00\x2e\x2c\xc4\xb0\x45\xd8\x20\x67\x44\xaf\xc7\xa3\x91\x85\xa0\x90\x61\xaa\xc7\xa3\x51\x2b\xcb\x25\xd5\x58\xaa\xab\x89\xfd\x33\xe9\x4d\x31\x92\x98\x4c\xea\xa5\x14\x6f\x30\xd5\xdd\x39\x47\x92\x05\xb6\x0e\xdb\xf5\xb7\x7f\x0f\x85\x3e\x2d\x35\xf9\x6e\x4e\x37\x27\x09\x0d\xca\x7f\xdd\x49\xa0\xd5\x5d\x64\x07\xb5\xf2\xb2\x68\x94\xa9\xd1\xcf\x56\xa0\x66\x88\x50\x8e\x12\x72\x56\xd3\x0c\x36\x92\x66\x4e\x65\xcb\x0c\xcc\x39\x20\xba\x51\x51\x8a\xdc\x64\xd8\xbb\x88\xe6\x0d\xe2\x3d\xd9\x8c\x2d\xb2\x8c\x1c\x44\xad\x41\x8a\x3d\x98\xec\xde\xf2\x7a\xb5\x8b\x72\x86\x37\x70\xa3\x2e\x1e\xda\x75\x76\x28\x25\x32\x73\x3a\xec\xf5\xba\x92\xb4\x24\xf2\x30\x71\xf2\x09\x4b\x22\x83\x00\xa4\x8c\x28\x75\x35\xa9\x6e\xa2\xf3\x30\x3f\xf2\x7f\x56\xc5\xc5\xfa\x9f\x0e\x0d\xd0\x02\x9e\x87\xa3\x56\x71\xe1\xf7\x1a\x8d\x4c\x5e\xbd\xfe\x16\x35\x28\x4d\xa4\xc6\x0c\x92\x03\x38\x79\x51\xbe\x01\xc2\x01\x6f\xa8\x72\x9f\x5b\xee\xfa\x67\x85\x04\x9f\x71\xd9\x79\xe0\xb8\x07\xc1\x31\x5e\x2d\x0c\x50\xbf\x83\x61\x74\x83\xac\xd7\x2f\x3f\xe4\x96\x98\x6f\x86\x07\xfe\x5b\xc3\x90\xcf\xd6\xbf\x86\x1d\x76\x28\x68\x5f\x50\x11\xb8\x89\x9c\x9f\xcc\x05\xcb\x50\x5e\x8b\xca\xa5\x0e\x5e\x41\xba\x0f\x18\xcc\x7f\xac\x90\xc3\x8b\x40\xb7\xd7\x6a\x4f\xd4\xfb\x69\xfa\xc8\x24\xa5\x12\x89\xc6\x6b\x8e\xfb\x6b\x47\xdd\x1d\x84\x3d\xb7\x6b\xe1\xef\xb8\xff\x90\x54\x19\xcd\xed\xe0\xec\xa6\x07\xb1\x03\x54\x9d\x98\x03\x9d\xf3\x6c\x5d\xdf\x61\x9c\x4d\x98\x03\xca\x15\xea\xf5\x2b\x34\x16\x15\x50\xb4\xa6\x7c\xd7\xe3\x4d\x04\xe9\xcf\x98\x23\x5e\xee\x58\x6b\xe1\x00\xe5\x20\x2d\xe4\x6b\x3f\x62\xbd\x58\xef\x89\xe5\x16\x0f\xcd\x03\x71\x46\x55\xc5\xc8\xe1\x9a\x93\x12\x8f\x96\x92\x1d\xd1\x44\x0e\x47\x83\xfb\x1a\x2c\x5f\x0f\xd7\x9d\x8c\xef\x83\x35\xc3\x20\xff\x3e\xe5\x85\x5b\x83\xfd\x9d\xbb\x37\x5e\xef\xd4\xf6\x83\xd8\x0e\x3b\xab\x1c\xb7\xb0\x68\x7d\x32\x1b\xb8\x1b\xac\xaa\x13\xbf\xf0\x07\xc1\xe1\xe1\x39\x3c\xad\x37\xf0\xf0\xfc\xe2\x0b\x78\xb4\x7c\xfc\xa4\x2a\xfb\x20\xdb\xd5\x77\x13\xde\x7a\xf2\x3b\xf9\x6e\x45\x78\x0a\x43\x73\x8e\x5f\xd1\x72\x03\x4b\x25\xd3\x96\x5a\xf7\xc0\x64\x6d\x8f\xf5\x77\x73\xfe\x14\xe8\x41\xba\xd3\xb5\x92\x26\xe5\xb9\xcb\x3f\xba\xcf\x2e\x8c\x34\x41\x6a\xd4\xc9\x9a\x9a\x68\xb5\x1e\x4e\x1e\x47\xb2\x9c\x32\x16\x15\x48\x37\x85\xee\x47\xb5\x41\x18\x6b\x0a\x11\x66\xfb\x4f\xa2\x28\x40\xb6\x0e\xc5\x4e\x2e\x64\xcd\xc1\xad\xf7\xee\x24\x2c\xa9\x59\xf3\x79\xb4\x62\x74\xbd\x11\xb0\x41\x0d\x49\xad\xb5\xe0\x30\x13\x9c\x1d\x80\x23\x66\x98\x01\xcd\x81\x0b\xd8\x88\xb8\x14\xd9\xd9\x6a\xc1\x68\xff\x49\xbb\x91\x7f\xf0\x78\xd6\x20\x70\xdb\x9c\x7d\xf2\xc1\x5d\x2b\x42\xd5\xe0\xc1\xfb\x97\x86\xe2\x98\xa8\x75\x55\x6b\x88\x80\xf2\x94\xd5\x99\x09\x13\x4a\x4b\x24\xa5\xf9\x94\x90\x74\x0b\x44\x01\xd5\x20\x6b\x93\x89\x74\xa0\xac\x16\x86\x27\xa3\x91\x51\xa0\x5b\xb2\x82\x66\xbf\x61\x6a\xd0\x61\x7c\x08\xfa\x9d\xa2\x9d\xd1\xf7\x09\x2c\x6d\xfd\xec\x6a\xa2\x2a\x46\xf5\x73\x37\xfa\xda\x0c\xd9\xc0\x12\x40\x77\x3e\x9a\x14\xc1\xd7\x6e\x5c\xdd\xc6\xe7\x06\xed\xbc\x5c\xf4\x07\x1a\x94\x09\xa3\x1b\xee\x73\x6a\x37\x37\x08\x0b\xd0\xdd\x73\x14\xd2\xbc\x61\xc0\x7b\x16\xea\x54\x36\x39\xbb\xc7\x03\xbe\xbc\x74\xdf\xe5\x83\x0a\xcb\x2f\x7a\x4c\x9f\x7a\xa2\x6b\x90\x0d\x57\x3a\x86\xf9\xe1\x79\x45\x79\x2e\x26\x7d\x31\xfd\x02\x42\x3c\xf1\x73\x78\x0f\x80\x8f\x4f\x97\xaa\xd3\x14\x95\x9a\xac\x9f\xd7\x4a\x8b\x12\x9e\xda\x78\x05\x17\xf7\x20\xef\x96\x47\x1f\xfe\x1a\xc2\x3a\xb6\xd5\x64\xae\x61\xc0\x44\x82\xbf\xff\xf8\xd3\x8b\x25\xe4\x42\x02\xa9\xb5\xf0\xe5\x6a\x5b\xf5\x0e\x95\xff\x12\x33\x5a\x97\xb6\xf4\xff\x55\x81\x44\x17\x28\x13\x21\xb6\x28\x17\x85\xd8\x47\x5a\x44\xe6\xc1\xc8\x3d\x68\xbe\xea\x02\xa3\x44\x68\x2d\xca\x48\xe4\x11\x89\x32\xba\x8b\x1e\x5f\x3c\xc1\xbf\x7c\xf6\x39\x7e\x4e\x1e\x3e\xf6\x01\xe6\x18\x09\x5d\x50\x05\x12\x09\x63\x07\x50\x85\xa8\x8d\x53\x44\x20\x90\x8a\xb2\x12\x1c\xb9\x8e\xe3\x78\xf0\xf0\x69\x77\x61\xdc\xd8\xb1\xbb\x30\xa3\xde\x5d\x3c\x23\x26\xf3\xa0\x95\xaa\x4b\xc8\x0c\xcf\x81\x94\xe6\x98\x50\x08\xa9\x41\xd2\x44\x81\x26\x94\xc1\x1b\x94\xdb\x03\xe8\x28\x11\x1c\xe7\xb0\xa5\xc8\x12\xa2\x08\xe8\x5a\x6e\xf1\x00\x05\x29\xa1\x22\x4a\x4b\x52\x52\xd8\xe2\x8e\x72\x48\x24\x2a\x22\x18\x31\x30\xf0\x80\x50\xd1\x8d\x39\x7a\xf2\x2d\x24\x84\x31\xd0\xb4\x8a\xe1\x6b\x61\xe2\x54\x25\xe4\x16\x98\xa0\x1c\x4a\x24\xda\xce\x26\x94\x69\x61\x9c\x6d\x45\x24\x3a\x2c\x1c\x02\x01\x36\x55\x5b\xd4\x90\x92\x8a\xa6\x82\x91\x71\x2e\x09\xdf\xe6\xb5\xd4\x36\xec\x75\x37\x71\xd0\xd3\x42\x54\xc0\x30\x41\xb9\x25\x0a\xde\x88\x3d\x8b\xe1\x1b\xb7\x0e\x31\xf7\x64\xd6\x32\xab\xd3\x2d\xf2\x39\x64\x16\x2f\x43\x54\x21\xd2\x2d\x74\xa1\x1b\x14\x99\x20\xb9\x83\xab\xf6\x94\x23\x10\x9e\x89\x9a\x32\x86\x5e\x54\x28\x03\x01\x31\x3c\x65\x29\xd1\xb2\xe1\x42\x26\xeb\x52\x69\x9a\x6e\x3d\x83\x2d\xd9\xcd\x46\x15\xe1\x29\x6a\x4d\x20\x37\x7c\x68\x16\x59\xf4\xcc\xa2\x2e\x22\x56\x2a\xe3\x06\xfd\x18\x7e\x08\xcc\x6b\x04\xe1\xa4\x05\x95\x14\x2a\xa5\x26\xd2\x75\xf9\x69\xf1\x4f\xd0\x28\x98\x17\x62\x8b\x5b\x23\xde\xce\xfe\x7d\xc2\x2d\x43\x2d\x0a\x05\x29\x93\x5a\x6e\x50\xc6\xf0\x93\xe0\x9b\x1a\x3d\x53\x12\x61\xe2\xe5\x1c\xc6\x8d\xb0\x0a\x6a\xb8\xeb\x45\xd8\x3c\x06\xa9\xd8\x83\x46\x9e\xa1\xb4\xcc\xe8\x6c\xd9\xee\xa4\x08\x23\x25\x8d\xc7\xe3\xef\x8c\xa2\x35\xc3\x0d\xa5\x2d\x34\xb7\xad\xc9\x7a\x50\x43\x49\x37\x5c\xf0\x79\x47\xc8\x2d\xd7\x63\x78\xd6\x15\xbd\x61\x80\xc1\x24\xa9\xf3\x9c\x30\x01\x06\x5c\x26\x36\x1b\xe4\xad\x78\x63\xf8\x3e\xa8\x50\x87\x7b\x81\x2e\x2d\x69\x64\x14\xae\x01\x26\xb9\x39\x66\x23\xe6\x31\x3c\x0b\xc6\xa0\xb4\xa4\x15\x28\x8d\x64\xdb\x18\xd3\xb1\xf8\x03\x12\x96\xc1\x0d\xe3\x3b\x2c\xab\x2b\x5a\xc6\xf0\x37\x6b\x12\x46\x9b\x5b\xc5\x1b\x57\x34\x25\xbc\x20\xf3\x80\x47\xa3\xd6\x3d\x9e\xb4\x74\x7a\x4b\x6b\x59\xd4\x18\x61\x20\x88\x38\x15\x8e\x61\xfc\xd2\x52\xed\x9f\xb4\x72\x9c\x5b\xa6\x39\x89\xab\xc2\x18\x54\x6f\x97\x8e\x92\x7b\xc1\x84\xfd\x1a\xd1\xa5\x45\x9d\x6e\x63\x78\x6e\x78\xef\x35\xa5\x2b\x2e\x03\xd2\xb2\xa1\xcb\x4f\xaf\x07\x81\x5d\x86\x03\xf3\xae\x0e\x19\x5f\x93\x58\xbf\xd6\x6c\x13\x2c\xaa\xe1\x66\xcb\x17\x03\xde\x99\x8f\x9d\x68\x77\xb7\x6c\x9e\x77\xfc\x93\x73\x2d\x8d\xbc\x1c\xd9\x27\x6c\xb5\xcb\x83\x8e\x61\x06\x77\x05\x63\xa7\xc2\x2d\x45\x1e\xe3\xd6\x65\xa5\x85\xa5\xac\x71\x59\x6e\xe7\x20\x91\x16\x23\x6f\x18\xf0\x37\xa3\x04\x0d\x6d\x6e\x74\xee\xa0\x06\xd5\xf3\x20\x1a\xb2\x3b\xbe\xc0\x35\x7f\xc1\xfd\xdf\x71\x29\x3f\xb5\x8b\x49\xad\xc8\x06\xbd\x67\x1c\x0f\xfc\x92\xd3\xde\x80\x54\x57\xdd\x8d\xfc\x18\xb6\x92\x0d\x24\x06\x54\xe7\x1d\xad\xee\xb9\xd1\xa0\x36\x81\xe4\x9e\x5a\x35\x2a\xea\xd1\x8a\xc7\x4e\x35\xad\x53\x70\xfe\x25\x2c\x99\xf7\x9d\xa6\xf3\x88\x8d\xb5\x38\x46\x75\xad\xd9\x59\x93\x07\x0b\xe3\xe7\x41\x64\x56\x4d\x5b\x5b\xb7\xb6\x15\xc3\x33\xaf\xcd\x01\xcb\xae\x79\x07\xc6\xf7\x34\xd3\x70\xab\x55\x30\x6f\xab\xc7\xfb\x53\xaf\xc7\x96\xa5\x84\x67\x6f\x08\x6e\x42\x68\xf3\xc2\x0e\xcf\xba\xb8\x18\x10\xb3\x5a\x3f\xef\x30\x28\x50\xda\x06\x81\x8e\xaa\x5b\xc3\x72\x91\x69\x7e\xe4\x78\x2d\x06\x3b\xe4\x54\x09\x8f\x49\x2b\x2a\x6f\x97\x8e\x9b\x6d\xb4\xf0\x68\x5b\x09\xc4\xe3\xf1\x6b\x6b\x1f\x0e\xc1\x56\xa4\x9d\xec\xc1\xa9\x64\x83\xa3\x61\x8e\x77\x36\xa0\x5d\x2c\x69\x63\x92\xa7\x70\xdc\x93\x67\x87\xdf\xde\x9f\x07\x4d\xeb\x21\x6e\xb1\xea\x66\x30\x27\xfc\x77\x0c\xe3\x57\x75\x59\xf5\xa2\xa5\xa5\xd2\x2c\x9d\x87\xe7\x1a\x6b\x6a\x79\xe1\x85\x12\x02\x7d\x27\x70\x34\xce\x67\x6c\xa5\xda\xd0\x12\xc3\x33\xef\x43\x7b\x86\x17\x1c\x76\xab\x2d\x5d\x0f\xd1\x11\x4f\x67\x8b\xb1\x31\xf1\x60\x65\xad\xa6\xc4\xe0\x78\x6f\x3d\x58\x50\xe7\x79\x93\x8a\x35\xfe\xa5\x09\x34\x8d\x5d\x04\xbb\x1b\x37\x52\x69\xbc\x4a\xe3\xbe\x3a\x21\xc5\x0a\xb0\x17\xea\x5c\x5a\x35\x88\x35\x0d\xbe\x56\xe3\xac\x13\x69\xaa\xc7\xdd\x4c\xdd\x65\xef\x4d\x92\x7e\x54\x04\x19\x14\x3f\xee\x57\xf5\x18\x14\x3a\x06\xf7\x07\xc6\x9d\xe4\xbb\x57\xf3\xb0\x45\x0c\xb7\xae\x5f\xf5\xb8\xb5\xec\xe1\xeb\x1d\xf6\xe8\x40\x4e\x54\x15\x6c\x81\x54\xe4\x50\x91\x74\x4b\x36\xa8\x60\x4f\x75\x01\x8b\xc5\x46\x2c\x37\xcd\xc9\x75\xf8\x10\x92\xb4\x70\x0b\x8d\x13\xc2\x74\x9b\x88\x1b\xe3\x2a\x81\x80\xc2\x8a\x58\xf4\xc2\xd1\xcf\xd7\x33\xe6\xbe\xb2\x0f\x84\xb1\x05\x17\x1c\x2f\xdb\x15\x4f\x4d\x76\x7d\xaf\xaa\xc7\xf1\x3c\x47\xcc\x0c\x75\x12\x4b\x2c\x13\x94\xb0\x2f\x68\x5a\x80\xe0\x38\x35\x71\x60\x87\x46\x05\xb8\xc3\x12\x33\x98\x75\x8a\x26\x64\x67\xfe\x68\x61\x15\xe3\x4b\x28\xca\x39\x94\xe4\x90\x20\xec\xd1\x56\x86\xba\xd4\x28\x4d\x74\xed\x82\x68\x3b\x98\x4b\x51\x82\x2e\x10\x5c\x83\x08\x22\x60\x74\x8b\xa0\xea\x84\xd1\x12\xcf\x40\x09\xd0\x05\xd1\x70\xa2\xc8\x93\x12\x0e\x5b\x2e\xf6\xb0\x37\x0b\xb4\x00\x51\x39\x5e\x0c\x78\xe0\x6a\x36\xe0\xff\xf9\xca\x8d\x2b\x25\x36\xb5\x9b\x70\xe0\x1a\x56\xc5\x5b\x45\xf4\x55\xf4\x00\xe6\xb8\x77\xc9\x8c\x52\x46\x09\xab\xd1\xf5\x2f\xd7\xbd\x16\x9b\x2b\x49\x9e\xea\x3e\x9e\xee\x49\xb6\x78\x1e\xef\x78\x7c\xf5\x04\x5e\x7a\xd5\x3b\xd9\xa2\x1c\x54\xeb\x4f\xf5\x2a\xfb\x4b\x42\x37\xaf\x4f\x42\x53\xcf\x5e\x2b\x24\x32\x2d\xba\x45\xec\xce\x9a\xb6\x3d\x78\x0c\xb4\x57\x0f\xe9\x2d\x18\x35\xa4\xfc\xef\xff\xfc\x37\x66\xb7\x80\xb4\x5c\x0c\x35\x0a\x43\xa7\x09\xce\xee\x20\x3e\x00\xd7\x6d\x4a\x0e\x66\x46\x8a\x09\x7d\x35\x21\xa9\xa6\x3b\xa2\x85\x9c\x0c\xe7\x87\x64\xb7\xad\xca\xc1\xdd\xa2\xc1\x8e\x6d\xe9\xc2\x62\xa8\x2a\xc2\x9b\xea\x0d\xbc\xb6\xa6\x8b\x59\x47\x54\x76\x41\xdb\xbb\xf1\xf4\xac\xc7\xa1\x3c\xed\x0a\xb6\x70\x4c\xea\x71\xfd\x7a\x34\x90\x5a\x1f\xff\x80\xbe\xc4\x5c\xa2\x2a\x6e\xe9\x3e\x1c\x11\x60\x37\xb7\x48\xbe\x72\x0f\x06\xdc\xc1\x78\x3c\x8f\xff\x10\xd9\x0e\x21\x43\x1d\xee\xf5\x6c\xc7\x77\xf6\x98\xde\xd3\x61\xea\xaa\x7b\xbf\xb5\xd4\x7b\xac\xdb\x59\xea\x32\xc4\x76\x89\x94\x28\x51\x17\x94\x6f\x26\xe3\xdb\x7b\x40\xed\xe8\x91\x2d\x18\x6a\x5d\x9b\xe0\xd8\x56\x6f\x6d\x54\x00\x1c\xc1\x08\x4d\xa2\xa5\x2f\xf2\xfc\x9b\x6a\x2c\x63\x33\xf4\xdc\x0c\xfc\x67\xb2\x7e\xfb\x16\x9a\x31\x78\xf7\xae\x11\x5f\xd8\x37\x34\xef\x7c\x14\xf1\x23\xa3\xe6\x86\x64\x98\x98\xf8\x19\x6b\xfb\x61\xf0\x18\xf3\x93\x8d\x90\xbb\xc8\xeb\x76\x6b\x86\x7a\x77\xd4\x4e\xf2\x38\x54\x44\x17\x0b\x91\x2f\xc8\xc2\x47\xcd\x45\x81\x12\x1b\x6a\x4e\x35\xa2\x86\xe0\x5b\xfe\x0d\x97\x06\x86\xb9\x9e\x97\xe3\xd8\x70\x4d\xaf\xca\x76\x3b\xc0\xb6\x63\x15\x80\xaa\x3a\x39\x0d\xb7\x5d\x7a\x02\xf6\xed\xfd\xad\xdb\xb5\xa6\x69\x2b\x0e\x90\x1c\x54\x48\x1b\x76\x07\x5f\x73\xfc\xc4\xb0\xa8\x3a\xa4\xb7\xb9\xbe\x21\x69\x55\xdd\xc6\x9c\x81\xbe\x7a\x5f\xbe\x91\x78\x00\x1b\xfb\x90\x47\x17\x93\xb5\xf1\xed\x7d\x0d\x3d\xbd\xa9\x47\xe9\xbd\xac\x3a\xc5\x85\xe1\xbd\xa3\x10\x17\x7a\x3c\xca\xe8\x8e\xb6\x1e\xc3\x3c\xe2\x47\x4e\x84\xc4\x81\x7b\xf9\x86\xb2\x5f\xed\x5b\x7c\xd3\xda\xa8\x0a\x50\x6e\x55\x46\x3d\x1c\xba\x18\xd7\xa7\x6e\x75\x74\x38\xdf\x34\xa7\xbb\x83\xf7\xf4\x4b\xf7\x72\x3c\xbf\xda\xe7\x7c\x38\x67\xf1\xfe\x6e\xf5\xbd\xec\xf8\x6e\x38\xbf\xca\x7c\x7f\x6f\xd3\xed\x99\xda\x29\x33\xfb\x05\x16\x76\x5f\x5b\xff\xc5\x46\x75\x2a\x78\x87\x16\x77\x3b\xd2\xb4\xb9\x3b\x70\xcc\x61\x6a\x5f\x20\x87\x83\xa8\xdb\xa3\x16\x81\x70\x0c\xf1\x7d\x8f\x4a\x8a\x84\x24\xae\x0f\xb2\x87\xba\x9a\x83\xbd\x9a\x6b\x4f\x22\xe6\x40\xe0\x72\x1d\x7b\x24\x4a\x99\x50\x24\x61\xae\x98\xb1\xb7\x27\x8e\xa2\x3d\xce\xf8\x1b\x30\x6e\xf7\x90\xc7\xf7\x5a\x46\xbf\xd3\x89\xf3\xe8\xf6\xb5\x3b\x60\xfa\xe1\xfb\xb5\xd4\x45\x65\xaf\xbb\x2d\x8c\x30\x17\xe1\x20\x19\x0e\x34\xef\x39\x3d\x9a\xa3\x7d\x41\x94\x39\x63\x95\xb6\x0e\xe9\x8e\x50\x6e\x5f\x73\x72\xa2\x7a\xaa\x00\x89\x3a\x18\x60\x55\xcd\xd3\xc2\xb8\x24\xcb\x3b\xb7\xed\xbc\x01\xaa\x90\x2b\x9a\x30\x84\x0c\x73\x52\x33\xad\x4e\x75\xe2\xc3\x31\xd8\x9c\xbe\x8c\x5c\x8c\x34\xd4\xad\xc7\xca\x3d\x91\x1c\x68\x6e\xf7\xab\x95\x71\xae\x0a\x48\x62\x0e\x5c\x5a\x18\x91\x26\x89\x2d\x0e\xb6\x37\xd4\xcc\xc9\xf0\xc4\xc9\xed\x03\x4a\x6c\x78\xf9\xdb\x09\xcc\x8d\xde\x4f\x5e\x8e\xe4\x84\x28\x9a\x42\x78\xce\x9e\x6b\x8f\xe9\xf7\xe7\x5d\x91\x43\x83\xa1\x02\x89\x8c\x68\xa7\xe4\x55\xb8\xc2\x75\xe2\xbc\xae\xa5\x60\xca\x56\x1c\xb4\xa8\x16\xf6\x6e\xdf\x42\xa2\xfd\xdb\x85\x36\x33\xd3\x60\x64\x4f\x98\x12\x67\x73\x2b\x15\x2f\x20\x8f\x9d\xed\x2c\xaa\x53\x7c\xfd\x70\x6c\xed\xdd\x14\xf7\x55\x16\xaa\xef\x79\xa7\x84\x6a\xcf\xc1\xb9\x6b\xcf\xfa\x43\x0a\xaa\x2f\x1d\x9b\x43\xaf\x74\x0f\x69\x2d\x25\x72\x0d\x89\x24\x46\x95\x0d\xad\xc6\x52\xa6\x0a\xd2\x82\xf0\x0d\x66\xb7\x94\x36\x08\xec\x89\x35\x01\xfb\x52\x98\x2b\x74\x08\x8d\x30\xdb\x50\x5d\xd4\xc9\x1c\x50\xa7\xf1\x89\x4b\x2b\xa9\x28\x4b\xc1\x7d\x59\xc1\x98\xcb\xd2\xe1\x91\xd1\x3c\x87\x99\xad\xe3\x59\x8f\xf6\x25\x08\xe9\xcb\x1e\x6f\x6a\xa5\xa1\x14\x19\x61\x73\xa3\xfa\x54\x43\x41\x32\xe7\xc8\x95\x2d\x64\x4c\x6d\xc3\x37\x15\x82\x9d\x19\xd1\x94\x54\x03\xee\x50\x1e\xec\x49\xc4\xd5\x83\x4a\x54\xae\xac\x57\xf3\x4c\x58\xe1\x62\x46\xb5\x9a\x43\x55\xab\xc2\xfc\xcf\xd8\x25\x94\xc4\x56\x49\x24\x9a\x3d\x32\x81\x8a\x4f\x75\x63\x54\x4a\xd7\x79\xee\xf6\x9f\x2a\x63\xf2\x04\xf2\x9a\x6f\x0f\x96\xd1\xbf\xaf\x2a\x1c\x5f\x87\x77\xfa\x10\x2e\xf1\x85\xf9\x7b\x96\xde\xfc\x95\xd7\x60\x2b\xae\x78\x34\x93\x48\x32\x10\x9c\x1d\x4e\x08\xcd\xdf\x3a\x84\x8c\x4a\xb3\x90\x11\x4d\x77\xd8\x31\x37\x07\x62\x6f\x4e\x1b\x50\x09\x65\x9d\xde\x09\x30\xdf\xfe\xf8\xf2\xe9\x4f\xdf\xc1\xcc\x07\x2b\x5b\x05\xdb\x08\x91\x59\x1d\x2d\x89\xa6\x29\xec\x08\xab\x4f\x3d\xfa\x92\xe8\x22\x18\xa1\x2d\x6f\xfd\x06\x34\x7e\x20\x94\x87\xba\x22\x70\x52\x22\xcc\x5a\x04\xde\xfb\xf4\x8b\x1b\x4c\x6b\x7b\xaf\xe2\x07\x91\x21\x44\xc0\x85\x2c\x09\x33\xea\xea\xbc\x43\xe4\xbd\x03\xcc\xfa\xde\x02\xf6\xd4\xd6\x13\x29\xcb\xfc\x33\xec\x00\x49\xad\x21\x95\x42\x29\xbb\xca\xa0\xef\x4c\xaa\xbe\x01\x52\x66\x9f\x3d\x76\xce\xc7\x04\x7e\xb4\xfb\x1a\xaf\x1f\x62\xf5\x00\x3c\xe5\x4a\x5b\x19\xe6\x86\x4b\xdc\xac\xca\xa8\xc4\x54\xb3\xc3\x31\x15\x4f\xb3\x8c\x1a\x22\x08\x6b\x2b\xac\x26\x38\x0a\x0b\xc6\xd8\xc7\xac\x56\x98\xd7\xcc\x5e\xf1\xf0\x95\x43\x2b\x27\xb3\x1d\x43\x8d\xb7\x5a\xb8\xbf\x7c\xe2\x08\x35\xe6\x48\x78\x66\x81\x3b\x3b\x36\x36\x46\xb9\x16\xbe\xb2\xaa\x40\xb8\xe8\xe9\xd6\xfb\x38\x3b\x73\xb1\xd4\x3a\x6b\x72\xf0\xb5\xd3\x8a\x11\x9d\x0b\x59\xce\xd4\x19\xd8\x58\xed\x83\xa7\x8f\x98\x1c\x22\x13\x9f\x4d\x82\x94\x12\xee\xb4\x8b\xc0\x9e\xf2\xcc\x4c\xa7\x0e\x2b\xe7\x12\x5c\x49\xd8\xb2\x79\x51\x92\xb4\x37\xf9\xe5\x09\x92\xdc\x25\x21\x4d\x36\x27\x22\xd2\xab\x9a\x43\x5a\x66\xcc\xb6\x37\xd9\xc9\x25\xcf\xa4\xd8\x2b\x74\x39\x18\x55\xf0\xff\x5e\x7d\xef\xfd\x98\xa5\x59\xc1\x9f\xed\x9d\x38\x48\x85\xbc\x4d\x5c\xcf\x9a\x34\xc1\xa4\x6d\x68\x19\x16\x94\x7e\x66\x3c\xef\xce\xfa\x2c\x6b\x19\x99\x14\x15\x64\x62\xcf\xcf\x7e\xdf\xb8\x7f\xfc\x16\x42\xd7\x37\x89\x1d\x4a\xd7\xc7\x75\xce\xc6\xad\xaa\x9d\xd3\xbf\xa7\xa3\xda\x08\xaf\xf5\x26\x4f\x85\x8a\xe8\xe2\x44\x42\x74\xa2\x26\xee\x74\xc8\x29\x4d\x41\x78\xc6\x50\x59\x97\x6f\xb8\x16\xd4\xdb\xba\xf3\x06\x12\x40\x04\x1b\x31\x87\x0d\xd5\x73\x48\x83\x2d\x4a\x98\xa1\xaa\x30\xa5\xd6\x54\x05\x0f\xca\x74\x36\x87\x8d\xa0\x65\x25\xa4\x89\x20\xce\x0c\x2f\x8d\x52\x35\xf0\xa8\xb6\x58\x59\xb5\xb7\x6f\x1a\xd4\x15\x50\x4e\xb5\x83\x64\x13\xf9\x26\x83\xcb\xa9\x54\xee\x56\xa4\xd5\x4b\x2e\xac\x1a\xb6\xa0\x54\xc0\x19\x33\x98\x09\x09\x04\xdc\x95\x09\x7f\xb9\xc9\x64\x80\x25\x55\xc6\x2b\x9c\xd0\xdc\xd7\xce\xd5\x8a\x3c\x47\x69\x90\xa9\xab\xcc\x56\xf3\x59\x06\x3b\x94\x8a\x5a\x03\xcc\x9d\xe6\x2b\x88\xa0\x43\xef\xb7\xe2\x18\x5c\xe7\x7d\x07\x60\x62\xe3\x1d\xae\x09\xb0\x52\x68\x37\xec\xb3\xe0\xdf\x57\xf7\x38\xee\xc3\x85\xf5\xa1\xe2\x71\xdc\x07\xd3\x98\x03\x96\x95\x3e\x18\xd7\xec\x5a\x22\x58\x1a\x1f\x62\x5c\x97\x71\x1f\xe1\x6b\xeb\x2e\x94\x11\x95\x27\x00\x66\xd6\x1e\x55\x25\x78\xe6\xfb\x31\xb6\xa5\xe2\x0f\x22\xde\xb1\x9d\xf5\x32\xb1\x0f\x41\x9a\x61\xe9\x75\x62\x3d\xc6\x51\x52\x6d\xd9\xed\xe7\xac\x63\x16\x15\x1a\x47\xbf\x70\x3d\x23\xeb\x92\xef\x42\xc8\x96\x74\xcc\x81\xee\x3e\x88\x05\xfc\x9b\x9e\x4e\xe7\x1d\x36\xe3\xca\x69\x7e\xf0\x6f\xf9\x74\x26\xdc\xdd\xc7\xc1\xf0\xe0\xcd\x3a\xdb\x0c\xea\xbf\x27\x74\xf4\x6e\x5d\x28\xce\xb7\xb7\x28\x7b\xd3\xcd\x81\xdf\xbe\x05\x00\x4b\xf7\x9b\x03\x4a\xd4\x32\xc5\x09\x68\x33\xa8\xaf\x26\xd7\x09\x23\x7c\x3b\x81\x61\x6f\x62\x08\xf0\xe8\x5d\x54\x91\x9d\x7c\xc3\x12\xba\xa5\x80\xc1\xb0\xad\xdb\xbf\xb6\x08\x84\x26\xc4\xf0\xb1\xb6\x17\xd1\x1b\xee\xf6\x71\xfb\xc7\xea\x5b\x94\xca\x16\x11\xba\x33\xcd\xfb\x83\xab\xe6\x97\x0d\x48\xd5\xec\x63\x51\x0b\xb7\x1d\xdd\x6f\x0e\xd8\x17\x16\x26\xeb\x3f\xf3\x44\x55\x97\xfe\xff\x54\x54\x87\x4b\xf8\xd3\x5b\x7f\x20\xb8\x3e\x20\x91\xef\xba\xa4\x58\x5c\x2d\x74\xf3\xdd\x7c\xb3\x9b\x04\x27\xbe\x52\xa9\xa4\x95\x06\xfb\x72\xc0\xc9\x5f\x7f\xf8\xea\x61\xfc\x24\xbe\xf8\xbc\xf9\xe5\x87\xf8\x8d\x9a\xac\x57\x0b\xf7\xdc\xfa\x1e\x00\x6e\xbe\x7a\x14\x9f\xc7\x17\x0d\x80\x9b\x5f\x0c\xe1\xce\x1f\xa0\xe8\x03\x0b\xd0\xd6\x63\x9a\xc3\xcc\x79\xfe\x98\x09\xe7\xfc\xe2\x82\xa8\x02\xae\xae\x60\xf2\x69\x4e\x25\x26\xf5\x66\x72\x06\x6f\xc7\xa3\x4c\xa4\x75\x89\x5c\xc7\x7b\x49\x35\x32\x3e\x9b\x06\x9c\x3a\xbf\x92\xf1\x86\xec\x88\x1b\x9d\xf4\x51\xdd\xa0\xf6\xd0\x2c\xbe\xfe\x73\xc4\xa8\xf6\xbc\x9a\x3e\x98\x2e\xa6\x0f\xa6\x2a\x9d\x3e\x98\x5a\xdc\xa6\x67\xfe\x17\x22\x8e\x90\x1e\x2f\x16\xcd\x11\x66\x2f\xe4\x16\x22\x90\xf8\x73\x4d\xa5\x49\xf4\x0a\x04\x85\x72\x67\x4e\x35\x34\x43\x17\x0b\xe3\x38\x86\xc4\xe8\xbf\x79\xf2\x1f\x35\xc6\xb5\xc2\xd9\x3f\x1c\x7b\xe6\xf0\xd6\x8c\x8e\x74\x81\x25\x2e\xfd\x97\x91\xbf\x4d\xbd\x84\xc9\xa7\x5f\x7c\x71\xfe\xe2\xf1\x37\x93\xb9\x9b\x50\x98\x0a\x9e\xf9\xa9\xa7\x5f\x5f\x3c\x7e\xf2\x79\x98\x22\xa9\x31\x78\x33\xfe\xfc\xb3\x87\x5f\x3c\xfc\x22\x8c\xa3\x94\x42\x9a\xe1\xfc\xf1\xe3\x47\x8f\x3e\x0b\xc3\x7b\x22\xb9\xbd\x54\x3c\xf9\x34\xcf\x31\x79\x94\x84\x09\xca\x73\x61\x46\x1f\x5e\xfc\xe5\xb3\xfc\x51\xb3\xb1\xbb\x00\x6d\x26\x1e\xa7\x24\x7f\x72\x3e\xb1\xe3\xef\xec\xf4\xbb\xb3\xcb\xf1\xd8\x9f\x2e\x63\x26\x36\xb3\xc9\x6b\x73\xf4\xf7\xaf\xf1\xc5\x71\x3c\x39\x1b\x8f\x77\xc4\xda\xce\xf5\xae\x46\xb8\xb2\x2f\xef\xfd\xa3\xc6\x99\x91\x2c\x32\x5a\x52\x8d\x52\x2d\xe1\xdf\xd3\x3f\xbd\x9d\xce\xa7\xef\xa6\xff\x99\x8f\x47\xc8\x96\x30\xfd\x94\x54\xd5\x74\x3e\x1e\x65\x44\x93\x25\xcc\xce\xe0\x6a\x0d\xb3\xb7\xee\xf6\xb7\x7b\x4f\x76\x09\xbc\x66\xcc\xd5\x86\x46\x5d\x1b\x5b\xda\x5d\xbe\x26\x1a\x67\x67\xf1\x06\xf5\x37\x35\x63\xff\x42\x22\x67\x67\x73\x67\xe5\x23\x17\x1a\x96\xd0\xbc\x3d\x19\x26\x06\x6f\x6e\x2d\xe1\xdf\xfe\xbe\xf9\xdb\xee\x2b\x48\x4b\x98\x24\xac\xc6\x68\x4f\xb3\x0d\x6a\x15\x51\x9e\x4e\xde\xcd\x6f\x59\x59\x0a\xbe\xc5\x43\x54\x10\x1d\xe5\x24\xd5\x42\x1e\x6e\x5f\xab\x49\x2a\x22\x34\x69\x10\xad\x4b\xf3\x81\x89\x03\x62\xa4\x69\x89\x91\x96\xc4\x64\x45\xcd\xc3\xff\x69\x88\x19\xbe\x84\xb1\x84\xe9\x74\xde\x9f\x6b\x6e\x5c\xbb\x39\x37\xb9\x58\xd8\x80\xdd\x7c\x61\x42\x62\xb9\x84\xe9\xf7\xe6\x6f\xef\x4e\xb6\xa2\xda\xde\xcb\x9e\x43\x89\x0c\x88\x39\xc2\x53\x4d\xe0\xe7\x9a\xf0\x4c\xc4\xf0\x13\xda\x15\xc2\xa6\xa2\xb0\xa3\xba\xb6\xb5\x08\x21\x69\x52\xab\xb9\x39\x10\x02\xc9\x3c\xc0\x4a\x28\x9a\x09\x4e\xeb\x12\x4a\xcc\xa8\x48\x25\x35\xa7\xfd\x72\x0e\x78\x63\x33\x21\x49\x21\xa3\x4a\x21\xd7\xb4\xe6\x1a\xd2\x1a\x4a\x24\x31\xbc\xc2\xaa\xce\x28\x91\x68\x66\xab\xda\xa5\x46\x14\x76\xb4\x04\xca\xe7\x50\x0a\xc6\x4c\x5e\x27\x69\x2d\x4d\xfe\x6a\x4e\x20\x73\x13\x3b\x91\xd7\x25\x10\xb9\xb1\x4e\xa4\x2e\x81\xd6\x0a\x48\x16\xc3\x4b\x49\x01\x6b\x1b\x76\x05\x10\xfc\xb9\x36\xa7\x09\xf9\x73\x4d\xb4\x50\xb1\x63\xd1\xe8\xdd\xd9\xd8\xaa\xb8\xd5\x61\x97\xf6\xfe\xd3\x7a\xad\x57\xa8\xe8\xff\x37\xea\x9c\xd7\xdc\xd6\xa8\x67\xd6\x57\x99\x65\xbe\x73\xfc\x1d\x5c\x41\xe3\xba\x7e\xae\x51\x1e\x5c\x67\x5d\xc8\xd9\xf4\x3e\xbf\xc3\x33\x3d\x8b\x45\x9e\x2b\xd4\xdf\xd9\xcc\xe1\xd2\x01\x77\xc1\xe2\x2e\xd8\xcd\x2f\xf1\x9c\x06\x50\xdb\x02\xb5\x01\xe0\xdd\x2f\xe5\x1c\xa5\x5b\x02\x11\xcc\x02\xf6\x0f\xfc\x4e\x67\x10\xc1\x93\xf3\xcb\xf1\xc8\x5b\x70\x7c\xa4\x6c\x70\x05\x53\x97\xdd\x2c\xa7\x0f\x2a\x22\x15\xfe\x95\xeb\x99\xdf\x67\xf1\xf0\xec\xc1\xb4\xba\x99\x9e\x00\xd0\x68\xe4\x3d\x01\xbc\x1b\x1f\x0b\x60\x76\x76\x39\xf6\x74\x90\x2c\x7b\xb1\x43\xae\xbf\xa7\x4a\x23\x47\x39\x9b\x4a\xbb\x64\x3a\x3f\x21\x38\x23\xd3\xae\x87\x5f\xb8\x5f\xff\x59\x2d\xec\x4f\x47\xfd\x5f\x00\x00\x00\xff\xff\x7d\x78\x53\xcb\x51\x4a\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/static"].(os.FileInfo),
	}
	fs["/static"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/static/index.html"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr: gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
