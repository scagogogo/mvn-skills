package command

import (
	"path/filepath"
	"runtime"
)

// BuildExecutable builds the path to the mvn executable based on the Maven home directory.
// On Windows, it returns the path to mvn.cmd; on other platforms, it returns the path to mvn.
func BuildExecutable(mavenHomeDirectory string) string {
	if runtime.GOOS == "windows" {
		return filepath.Join(mavenHomeDirectory, "bin", "mvn.cmd")
	}
	return filepath.Join(mavenHomeDirectory, "bin", "mvn")
}
