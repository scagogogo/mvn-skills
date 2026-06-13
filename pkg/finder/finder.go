package finder

import (
	"fmt"
	"github.com/scagogogo/mvn-skills/pkg/command"
	"os"
	"os/exec"
	"strings"
)

// NotFoundError indicates Maven was not found, with details about which paths were searched
type NotFoundError struct {
	SearchPaths []string // Paths that were checked
	Message     string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("maven not found: %s (searched: %v)", e.Message, e.SearchPaths)
}

// ErrNotFoundMaven is a sentinel error for Maven not found (backward compatible)
var ErrNotFoundMaven = &NotFoundError{Message: "not found maven"}

// FindMaven locates the locally installed Maven.
// It first uses exec.LookPath for a fast check, then falls back to running mvn --help,
// and finally checks M2_HOME and MAVEN_HOME environment variables.
func FindMaven() (string, error) {
	searchPaths := []string{}

	// Fast check: use exec.LookPath first (much faster than running mvn --help)
	path, err := exec.LookPath("mvn")
	if err == nil && path != "" {
		searchPaths = append(searchPaths, "PATH:"+path)
		// Verify it actually works
		stdout, execErr := command.ExecForStdout("mvn", "--help")
		if execErr == nil && strings.Contains(stdout, "usage: mvn") {
			return "mvn", nil
		}
	}

	// Search from several environment variables
	envNameSlice := []string{"M2_HOME", "MAVEN_HOME"}
	for _, envName := range envNameSlice {
		getenv := os.Getenv(envName)
		if getenv == "" {
			continue
		}
		searchPaths = append(searchPaths, envName+"="+getenv)
		if Check(getenv) {
			return command.BuildExecutable(getenv), nil
		}
	}

	return "", &NotFoundError{
		SearchPaths: searchPaths,
		Message:     "maven executable not found in PATH or M2_HOME/MAVEN_HOME",
	}
}

// Check verifies whether the directory is a valid Maven directory, based on whether the mvn executable exists
func Check(mavenHomeDirectory string) bool {
	executable := command.BuildExecutable(mavenHomeDirectory)
	stat, err := os.Stat(executable)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}
