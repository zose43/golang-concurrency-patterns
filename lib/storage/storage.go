package storage

import (
	"fmt"
	"os"
	"path"
)

const Dir = "xxx"

func OpenData(filename string) ([]byte, error) {
	bytes, err := os.ReadFile(dir(filename))
	if err != nil {
		return nil, fmt.Errorf("can't read file %w", err)
	}
	return bytes, nil
}

func OpenFile(filename string) (*os.File, error) {
	f, err := os.Open(dir(filename))
	if err != nil {
		return nil, err
	}
	return f, nil
}

func dir(p string) string {
	return path.Join(Dir, p)
}
