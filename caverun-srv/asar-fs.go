package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"layeh.com/asar"
)

// AsarFileSystem implements http.FileSystem on top of a path that may contain one ".asar" archive file,
// and treats it like a directory.
type AsarFileSystem struct {
	RegularBaseDir http.FileSystem
	AsarPrefix     string
	AsarFp         *os.File
	AsarRootEntry  *asar.Entry
}

// NewAsarFileSystem returns a new instance.  The path
// provided can be a regular filesystem path, or a
// path that contains a .asar file which is opened
// and used like a directory.
func NewAsarFileSystem(basePath string) (*AsarFileSystem, error) {

	ret := &AsarFileSystem{}

	// check for normal directory (no .asar file)
	st, err := os.Stat(basePath)
	if err == nil && st.IsDir() {
		ret.RegularBaseDir = http.Dir(basePath)
		return ret, nil
	}

	asarPrefix := ""
	p := basePath
	for {
		st, err := os.Stat(p)
		if err == nil {
			if !st.IsDir() {
				// proceed with asar file
				ret.AsarPrefix = asarPrefix

				f, err := os.Open(p)
				if err != nil {
					return nil, err
				}
				ret.AsarFp = f

				archive, err := asar.Decode(f)
				if err != nil {
					return nil, err
				}
				ret.AsarRootEntry = archive

				return ret, nil

			} else {

				// this is just odd - directory with some other path components that don't exist, will never work, send back error
				return nil, fmt.Errorf("path provided contains extra path components after a directory, nonsensical path")
			}
		}

		asarPrefix = filepath.Base(p) + asarPrefix
		newp := filepath.Dir(p)
		// check for being done
		if newp == p {
			break
		}
		p = newp

	}

	return nil, fmt.Errorf("failed to find valid directory or .asar file in base path")

}

// Open implements http.FileSystem
func (fs *AsarFileSystem) Open(name string) (http.File, error) {

	if fs.AsarFp == nil {
		return fs.RegularBaseDir.Open(name)
	}

	parts := strings.Split(strings.Trim(path.Clean("/"+fs.AsarPrefix+"/"+name), "/"), "/")

	entry := fs.AsarRootEntry.Find(parts...)
	if entry == nil {
		return nil, os.ErrNotExist
	}

	isr := entry.Open()
	if isr == nil {
		return nil, fmt.Errorf("unable to open entry %q, is it a directory?", name)
	}

	return &AsarFile{entry: entry, SectionReader: isr}, nil
}

func (fs *AsarFileSystem) Close() error {
	if fs.AsarFp != nil {
		return fs.AsarFp.Close()
	}
	return nil
}

type AsarFile struct {
	entry *asar.Entry
	*io.SectionReader
}

func (f *AsarFile) Close() error {
	return nil
}

func (f *AsarFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("directory operations not implemented on asar files")
}

func (f *AsarFile) Stat() (os.FileInfo, error) {
	return f.entry.FileInfo(), nil
}
