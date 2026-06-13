package command

// Release plugin-related commands
// The release plugin is the standard release workflow tool for Maven

// ReleasePrepare prepares a release (mvn release:prepare)
// Performs version checking, creates a tag, and updates to the next development version
func ReleasePrepare(executable string) (string, error) {
	return ExecForStdout(executable, "release:prepare")
}

// ReleasePrepareWithArgs prepares a release with additional arguments
// Common arguments include -Darguments="-DskipTests" to skip tests
func ReleasePrepareWithArgs(executable string, args ...string) (string, error) {
	allArgs := append([]string{"release:prepare"}, args...)
	return ExecForStdout(executable, allArgs...)
}

// ReleasePerform performs the release (mvn release:perform)
// Checks out code from the tag and executes deploy
func ReleasePerform(executable string) (string, error) {
	return ExecForStdout(executable, "release:perform")
}

// ReleaseRollback rolls back the release preparation (mvn release:rollback)
// Executed when release:prepare fails or issues are found
func ReleaseRollback(executable string) (string, error) {
	return ExecForStdout(executable, "release:rollback")
}

// ReleaseClean cleans up the release state (mvn release:clean)
// Cleans up release.properties and other release temporary files
func ReleaseClean(executable string) (string, error) {
	return ExecForStdout(executable, "release:clean")
}
