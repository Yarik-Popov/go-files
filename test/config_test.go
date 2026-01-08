package src

import (
	"Yarik-Popov/go-files/src"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var source string = "source"
var destination string = "destination"
var pattern string = ".go"

func TestCreateConfigMaxAttempts0(t *testing.T) {
	config, err := src.CreateConfig(source, destination, pattern, 0)

	assert.Nil(t, config)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "greater than 0 and less than 255 but got 0")
}

func TestCreateConfigMaxAttempts256(t *testing.T) {
	config, err := src.CreateConfig(source, destination, pattern, 256)

	assert.Nil(t, config)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "greater than 0 and less than 255 but got 256")
}

func TestCreateConfigSuccess(t *testing.T) {
	config, err := src.CreateConfig(source, destination, pattern, 5)

	assert.Nil(t, err)
	assert.Equal(t, config.Source, source)
	assert.Equal(t, config.Pattern, pattern)
	assert.Equal(t, config.MaxAttempts, uint8(5))

	fullDstPath, _ := filepath.Abs(destination)
	assert.Equal(t, config.Destination, fullDstPath)
	assert.Contains(t, config.Destination, filepath.Join("test", destination))
}
