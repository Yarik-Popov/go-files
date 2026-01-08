package src

import (
	"Yarik-Popov/go-files/src"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

var source string = "source"
var destination string = "destination"
var pattern string = ".go"

func TestCreateConfigMaxAttempts0(t *testing.T) {
	config, err := src.CreateConfig(source, destination, pattern, 0)

	require.Nil(t, config)
	require.NotNil(t, err)
	require.ErrorContains(t, err, "greater than 0 and less than 255 but got 0")
}

func TestCreateConfigMaxAttempts256(t *testing.T) {
	config, err := src.CreateConfig(source, destination, pattern, 256)

	require.Nil(t, config)
	require.NotNil(t, err)
	require.ErrorContains(t, err, "greater than 0 and less than 255 but got 256")
}

func TestCreateConfigSuccess(t *testing.T) {
	tmpDir := t.TempDir()
	t.Chdir(tmpDir)
	config, err := src.CreateConfig(source, destination, pattern, 5)

	require.Nil(t, err)
	require.Equal(t, config.Source, source)
	require.Equal(t, config.Pattern, pattern)
	require.Equal(t, config.MaxAttempts, uint8(5))

	fullDstPath, _ := filepath.Abs(destination)
	require.Equal(t, config.Destination, fullDstPath)
	require.Contains(t, config.Destination, filepath.Join(tmpDir, destination))
}
