package web

import (
	"embed"
	"fmt"
	"io/fs"
	"path"
	"strings"
)

//go:embed all:dist
var distFS embed.FS

type neuteredFS struct{ fs fs.FS }

func (n *neuteredFS) Open(name string) (fs.File, error) {
	name = strings.TrimPrefix(name, "/")
	if name == "" {
		name = "."
	}

	f, err := n.fs.Open(name)
	if err != nil {
		return nil, err
	}

	if stat, err := f.Stat(); err != nil {
		f.Close()
		return nil, err
	} else if stat.IsDir() {
		_, err := fs.Stat(n.fs, path.Join(name, "index.html"))
		if err != nil {
			f.Close()
			return nil, fs.ErrNotExist
		}
	}

	return f, nil
}

var FS fs.FS

func init() {
	fs, err := fs.Sub(distFS, "dist")
	if err != nil {
		panic(fmt.Sprintf("embedded frontend could not be mounted: %v", err))
	}

	FS = &neuteredFS{fs}
}
