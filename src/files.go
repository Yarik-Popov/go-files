package src

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

/**
* Initializes the destination directory, returning any errors if they exist.
 */
func InitDestinationDir(dst string) error {
	dir, err := os.Stat(dst)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if dir != nil && !dir.IsDir() {
		return errors.New("Error: Destination is not a directory")
	}

	const directoryPerms = 0755
	err = os.Mkdir(dst, directoryPerms)
	if errors.Is(err, os.ErrExist) {
		return nil
	}
	return err
}

func OperateOnFiles(config *Config) error {
	err := filepath.WalkDir(config.Source, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil // Skip problematic directories
		}

		oldFullPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		// Don't process the destination directory as this was already handled
		parentDirectory := filepath.Dir(oldFullPath)
		if parentDirectory == config.Destination {
			return nil
		}

		if matched, _ := filepath.Match(config.Pattern, d.Name()); !matched {
			return nil
		}

		newFullPath := filepath.Join(config.Destination, d.Name())

		newFileInfo, err := os.Stat(newFullPath)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}

		if newFileInfo != nil {
			fmt.Println("We will handle the duplicate later")
		}

		// err = os.Rename(oldFullPath, newFullPath)
		// return err
		fmt.Println("Moving from " + oldFullPath + " to " + newFullPath)
		return nil
	})
	return err
}
