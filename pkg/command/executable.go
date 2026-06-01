package command

import "path/filepath"

// BuildExecutable builds the path to the mvn executable based on the Maven home directory
func BuildExecutable(mavenHomeDirectory string) string {
	return filepath.Join(mavenHomeDirectory, "bin/mvn")
}