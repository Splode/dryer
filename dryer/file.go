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

func matchPattern(dir string, cfg *Config) error {
	g := filepath.Join(dir, cfg.Pattern)
	m, err := filepath.Glob(g)
	if err != nil {
		return err
	}
	cfg.Paths = append(cfg.Paths, m...)
	return nil
}

func walkPattern(cfg *Config) error {
	d := cfg.Dir
	return filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return matchPattern(path, cfg)
		}
		return nil
	})
}
