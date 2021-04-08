package dryer

import (
	"io"
	"os"
	"path/filepath"
)

type file struct {
	src          io.Reader
	path         string
	absolutePath string
}

func open(path string) (*file, error) {
	src, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	return &file{
		src:          src,
		path:         path,
		absolutePath: abs,
	}, nil
}
