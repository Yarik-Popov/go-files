package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

/**
* Initializes the destination directory and returns the absolute path to it.
 */
func initDestinationDir(dst string) (string, error) {
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

func main() {
	fmt.Println("Go files")

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("You must use 2 arguments. src dst")
		os.Exit(1)
	}

	source := args[0]
	destination := args[1]

	fullDst, err := initDestinationDir(destination)
	fmt.Printf("src: %s, dst: %s \n", source, fullDst)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
