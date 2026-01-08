package src

import (
	"errors"
	"flag"
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
	maxAttempts := *maxAttemptPtr

	return CreateConfig(source, destination, pattern, maxAttempts)
}
