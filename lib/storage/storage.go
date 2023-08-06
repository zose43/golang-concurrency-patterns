package storage

import (
	"fmt"
	"os"
	"path"
)

const Dir = "xxx"

func OpenData(filename string) ([]byte, error) {
	bytes, err := os.ReadFile(path.Join(Dir, filename))
	if err != nil {
		return nil, fmt.Errorf("can't read file %w", err)
	}
	return bytes, nil
}
