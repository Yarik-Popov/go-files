package src

import (
	"Yarik-Popov/go-files/src"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

var testDestinationDir string = "test"

/**
* Creates a temp directory and switches into it.
 */
func switchToTempDir(t *testing.T) string {
	tmpDir := t.TempDir()
	t.Chdir(tmpDir)
	return tmpDir
}

func TestInitDestinationDirSuccess(t *testing.T) {
	tmpDir := switchToTempDir(t)

	err := src.InitDestinationDir(testDestinationDir)
	require.Nil(t, err)
	require.DirExists(t, filepath.Join(tmpDir, testDestinationDir))
}

func TestInitDestinationDirExists(t *testing.T) {
	tmpDir := switchToTempDir(t)

	dirToUse := filepath.Join(tmpDir, testDestinationDir)
	err := os.Mkdir(dirToUse, 0755)

	t.Cleanup(func() {
		os.Remove(dirToUse)
	})

	// Should not error if directory exists
	err = src.InitDestinationDir(testDestinationDir)
	require.Nil(t, err)
	require.DirExists(t, dirToUse)
}

func TestInitDestinationDirExistsAsFile(t *testing.T) {
	tmpDir := switchToTempDir(t)

	dirToUse := filepath.Join(tmpDir, testDestinationDir)

	// Create a file so the init function should error
	_, err := os.Create(dirToUse)
	require.Nil(t, err)
	require.FileExists(t, dirToUse)

	t.Cleanup(func() {
		os.Remove(dirToUse)
	})

	err = src.InitDestinationDir(testDestinationDir)
	require.NotNil(t, err)
	require.NoDirExists(t, dirToUse)
	require.ErrorContains(t, err, "Destination is not a directory")
}

var safeMoveFileName string = "test.txt"

func defaultConfig(tmpDir string) *src.Config {
	config := src.Config{
		Source:      "source",
		Destination: tmpDir,
		Pattern:     "*.txt",
		MaxAttempts: 3,
	}
	return &config
}

func TestSafeMoveFileInstantSuccess(t *testing.T) {
	oldFullPath, err := filepath.Abs(filepath.Join("safeMove", safeMoveFileName))
	require.Nil(t, err)

	tmpDir := switchToTempDir(t)
	config := defaultConfig(tmpDir)

	oldFileData, err := os.ReadFile(oldFullPath)
	require.Nil(t, err)

	expectedNewPath := filepath.Join(tmpDir, safeMoveFileName)
	t.Cleanup(func() {
		os.Remove(expectedNewPath)
		os.WriteFile(oldFullPath, oldFileData, 0664) // Restore the file as it was being moved
	})

	err = src.SafeMoveFile(config, safeMoveFileName, oldFullPath)
	require.Nil(t, err)
	require.FileExists(t, expectedNewPath)

	newFileData, err := os.ReadFile(expectedNewPath)
	require.Nil(t, err)
	require.Equal(t, newFileData, oldFileData)
}
