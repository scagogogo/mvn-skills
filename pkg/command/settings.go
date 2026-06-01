package command

// GetLocalRepositoryDirectory gets the local repository location from the command
func GetLocalRepositoryDirectory(executable string) (string, error) {
	return ExecForStdout(executable, "help:evaluate", "-Dexpression=settings.localRepository", "-q", "-DforceStdout")
}