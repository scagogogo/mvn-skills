package finder

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

// ErrNotFoundMavenWrapper indicates Maven Wrapper was not found
var ErrNotFoundMavenWrapper = errors.New("not found maven wrapper")

// FindMavenWrapper locates the Maven Wrapper (mvnw/mvnw.cmd) in the specified project directory
// Maven Wrapper is the recommended practice for modern Maven projects, the project directory contains the mvnw script
// which allows building the project without having Maven installed
func FindMavenWrapper(projectDir string) (string, error) {
	wrapperName := getWrapperScriptName()
	wrapperPath := filepath.Join(projectDir, wrapperName)

	stat, err := os.Stat(wrapperPath)
	if err != nil {
		return "", ErrNotFoundMavenWrapper
	}
	if stat.IsDir() {
		return "", ErrNotFoundMavenWrapper
	}

	return wrapperPath, nil
}

// FindBestMaven locates the most suitable Maven executable in the specified project directory
// It prioritizes the Maven Wrapper in the project directory, and falls back to the system-installed Maven if not found
func FindBestMaven(projectDir string) (string, error) {
	// Prioritize locating the Maven Wrapper
	wrapper, err := FindMavenWrapper(projectDir)
	if err == nil {
		return wrapper, nil
	}

	// Fall back to system Maven
	return FindMaven()
}

// HasMavenWrapper checks whether the Maven Wrapper exists in the specified project directory
func HasMavenWrapper(projectDir string) bool {
	wrapperName := getWrapperScriptName()
	wrapperPath := filepath.Join(projectDir, wrapperName)

	stat, err := os.Stat(wrapperPath)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

// getWrapperScriptName returns the Maven Wrapper script filename based on the operating system
func getWrapperScriptName() string {
	if runtime.GOOS == "windows" {
		return "mvnw.cmd"
	}
	return "mvnw"
}