package src

import (
	"errors"
	"os"
	"path/filepath"
)

func ParseArgs() (*Config, error) {
	args := os.Args[1:]
	if len(args) != 3 {
		return nil, errors.New("'src' 'dst' 'pattern' must be provided")
	}

	source := args[0]
	destination := args[1]
	pattern := args[2]

	fullDstPath, err := filepath.Abs(destination)
	if err != nil {
		return nil, err
	}

	config := Config{Source: source, Destination: fullDstPath, Pattern: pattern}
	return &config, nil
}
