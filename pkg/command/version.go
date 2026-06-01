package command

// Version gets the Maven version
func Version(executable string) (string, error) {
	return ExecForStdout(executable, "-v")
}