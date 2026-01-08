package src

import (
	"errors"
	"flag"
	"math"
	"path/filepath"
	"strconv"
)

func ParseArgs() (*Config, error) {
	maxAttemptPtr := flag.Uint("a", 3, "Number of attempts to try to move a file before skipping")
	flag.Parse()

	args := flag.Args()
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

	maxAttempts := *maxAttemptPtr
	if maxAttempts <= 0 || maxAttempts > math.MaxUint8 {
		return nil, errors.New("Attempts must be greater than 0 and less than " + strconv.FormatUint(math.MaxUint8, 10) + " but got " + strconv.FormatInt(int64(maxAttempts), 10))
	}

	config := Config{
		Source:      source,
		Destination: fullDstPath,
		Pattern:     pattern,
		MaxAttempts: uint8(maxAttempts),
	}
	return &config, nil
}
