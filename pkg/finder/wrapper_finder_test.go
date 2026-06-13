package finder

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// wrapperScriptName returns the platform-specific wrapper script name
func wrapperScriptName() string {
	if runtime.GOOS == "windows" {
		return "mvnw.cmd"
	}
	return "mvnw"
}

func TestFindMavenWrapper(t *testing.T) {
	tmpDir := t.TempDir()

	// No wrapper in the directory
	_, err := FindMavenWrapper(tmpDir)
	assert.Equal(t, ErrNotFoundMavenWrapper, err)

	// Create the wrapper file with platform-specific name
	scriptName := wrapperScriptName()
	wrapperPath := filepath.Join(tmpDir, scriptName)
	err = os.WriteFile(wrapperPath, []byte("#!/bin/sh\n"), 0644)
	assert.Nil(t, err)

	// Now it should be found
	found, err := FindMavenWrapper(tmpDir)
	assert.Nil(t, err)
	assert.Equal(t, wrapperPath, found)
}

func TestHasMavenWrapper(t *testing.T) {
	tmpDir := t.TempDir()

	// No wrapper
	assert.False(t, HasMavenWrapper(tmpDir))

	// Create wrapper with platform-specific name
	scriptName := wrapperScriptName()
	wrapperPath := filepath.Join(tmpDir, scriptName)
	err := os.WriteFile(wrapperPath, []byte("#!/bin/sh\n"), 0644)
	assert.Nil(t, err)

	// Has wrapper
	assert.True(t, HasMavenWrapper(tmpDir))
}

func TestHasMavenWrapperWithDirectory(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a directory with the wrapper name (not a file)
	scriptName := wrapperScriptName()
	wrapperDir := filepath.Join(tmpDir, scriptName)
	err := os.MkdirAll(wrapperDir, 0755)
	assert.Nil(t, err)

	// A directory is not a valid wrapper
	assert.False(t, HasMavenWrapper(tmpDir))
}

func TestFindBestMavenWithWrapper(t *testing.T) {
	tmpDir := t.TempDir()

	// Create wrapper with platform-specific name
	scriptName := wrapperScriptName()
	wrapperPath := filepath.Join(tmpDir, scriptName)
	err := os.WriteFile(wrapperPath, []byte("#!/bin/sh\n"), 0644)
	assert.Nil(t, err)

	// Should return the wrapper (preferred over system Maven)
	maven, err := FindBestMaven(tmpDir)
	assert.Nil(t, err)
	assert.Equal(t, wrapperPath, maven)
}

func TestFindBestMavenWithoutWrapper(t *testing.T) {
	tmpDir := t.TempDir()

	// No wrapper, fall back to system Maven
	maven, err := FindBestMaven(tmpDir)
	if err != nil {
		// System has no Maven — should be a NotFoundError
		var nfe *NotFoundError
		assert.True(t, errors.As(err, &nfe))
	} else {
		assert.NotEmpty(t, maven)
	}
}

func TestGetWrapperScriptName(t *testing.T) {
	name := getWrapperScriptName()
	if runtime.GOOS == "windows" {
		assert.Equal(t, "mvnw.cmd", name)
	} else {
		assert.Equal(t, "mvnw", name)
	}
}
