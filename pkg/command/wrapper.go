package command

// Wrapper generates Maven Wrapper files in the project, allowing the project to be built without Maven installed
func Wrapper(executable string) (string, error) {
	return ExecForStdout(executable, "wrapper:wrapper")
}