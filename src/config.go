package src

import (
	"errors"
	"math"
	"path/filepath"
	"strconv"
)

type Config struct {
	Source      string
	Destination string
	Pattern     string
	MaxAttempts uint8
}

func CreateConfig(source string, destination string, pattern string, maxAttempts uint) (*Config, error) {
	fullDstPath, err := filepath.Abs(destination)
	if err != nil {
		return nil, err
	}

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
