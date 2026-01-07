package files

import (
	"errors"
	"os"
	"path/filepath"
)

/**
* Initializes the destination directory and returns the absolute path to it.
 */
func InitDestinationDir(dst string) (string, error) {
	path, err := filepath.Abs(dst)
	if err != nil {
		return path, err
	}

	dir, err := os.Stat(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return path, err
	}

	if dir != nil && !dir.IsDir() {
		return path, errors.New("Error: Destination is not a directory")
	}

	const directoryPerms = 0755
	err = os.Mkdir(dst, directoryPerms)
	if errors.Is(err, os.ErrExist) {
		return path, nil
	}
	return path, err
}
