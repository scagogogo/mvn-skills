package command

// Assembly plugin-related commands
// The assembly plugin is used to create custom distribution packages (zip, tar.gz, etc.)

// AssemblySingle creates a distribution package (mvn assembly:single)
func AssemblySingle(executable string) (string, error) {
	return ExecForStdout(executable, "assembly:single")
}
