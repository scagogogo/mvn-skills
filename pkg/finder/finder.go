package finder

import (
	"errors"
	"github.com/scagogogo/mvn-skills/pkg/command"
	"os"
	"strings"
)

// ErrNotFoundMaven indicates Maven was not found
var ErrNotFoundMaven = errors.New("not found maven")

// FindMaven locates the locally installed Maven
func FindMaven() (string, error) {

	// Try to find an executable Maven from PATH
	stdout, err := command.ExecForStdout("mvn", "--help")
	if err == nil && strings.Contains(stdout, "usage: mvn [options] [<goal(s)>] [<phase(s)>]") {
		return "mvn", nil
	}

	// Search from several environment variables
	envNameSlice := []string{"M2_HOME", "MAVEN_HOME"}
	for _, envName := range envNameSlice {
		getenv := os.Getenv(envName)
		if getenv == "" {
			continue
		}
		if Check(getenv) {
			return command.BuildExecutable(getenv), nil
		}
	}

	return "", ErrNotFoundMaven
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