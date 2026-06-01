package command

// GPG plugin-related commands
// The gpg plugin is used to GPG-sign artifacts; a required step for publishing to Maven Central

// GpgSign GPG-signs the artifacts (mvn gpg:sign)
func GpgSign(executable string) (string, error) {
	return ExecForStdout(executable, "gpg:sign")
}