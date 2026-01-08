package src

import (
	"errors"
	"fmt"
	"io/fs"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strconv"
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

		fileName := d.Name()
		if matched, _ := filepath.Match(config.Pattern, fileName); !matched {
			return nil
		}

		return SafeMoveFile(config, fileName, oldFullPath)
	})
	return err
}

func SafeMoveFile(config *Config, fileName string, oldFullPath string) error {
	attempt := 0
	extension := filepath.Ext(fileName)
	prefix := fileName[:len(fileName)-len(extension)]

	for attempt < int(config.MaxAttempts) {
		randomSuffix := ""
		// Attempting to move the file previously would have resulted in overriding another file so we are retrying
		if attempt > 0 {
			randomSuffix = "-" + strconv.FormatInt(rand.Int64(), 10)
		}
		newFullPath := filepath.Join(config.Destination, prefix, randomSuffix, extension)

		newFileInfo, err := os.Stat(newFullPath)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}

		// File exists in Destination directory so don't move it as that would override the file. Try again
		if newFileInfo != nil {
			attempt++
			continue
		}

		// err = os.Rename(oldFullPath, newFullPath)
		// return err
		fmt.Println("Moving from " + oldFullPath + " to " + newFullPath)
	}
	return errors.New("Failed to move " + fileName + " after " + strconv.FormatUint(uint64(config.MaxAttempts), 10) + " max attempts")
}
