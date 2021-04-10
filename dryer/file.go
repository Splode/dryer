package dryer

import (
	"io"
	"os"
	"path/filepath"
)

type file struct {
	src          io.Reader // src represents the contents of the file.
	path         string    // path is original path used to access the file.
	absolutePath string    // absolutePath is the absolute path on the system to the file.
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
