package command

// Versions plugin-related commands
// The versions plugin is used to manage project versions and check for dependency/plugin updates

// VersionsSet sets the project version (mvn versions:set -DnewVersion=...)
func VersionsSet(executable, newVersion string) (string, error) {
	return ExecForStdout(executable, "versions:set", "-DnewVersion="+newVersion)
}

// VersionsCommit commits the version change (mvn versions:commit)
// Executed after versions:set is confirmed to be correct; deletes the backup POM
func VersionsCommit(executable string) (string, error) {
	return ExecForStdout(executable, "versions:commit")
}

// VersionsRevert reverts the version change (mvn versions:revert)
// Executed when issues are found after versions:set; restores the backup POM
func VersionsRevert(executable string) (string, error) {
	return ExecForStdout(executable, "versions:revert")
}

// VersionsDisplayDependencyUpdates checks for available dependency updates (mvn versions:display-dependency-updates)
func VersionsDisplayDependencyUpdates(executable string) (string, error) {
	return ExecForStdout(executable, "versions:display-dependency-updates")
}

// VersionsDisplayPluginUpdates checks for available plugin updates (mvn versions:display-plugin-updates)
func VersionsDisplayPluginUpdates(executable string) (string, error) {
	return ExecForStdout(executable, "versions:display-plugin-updates")
}

// VersionsUseLatestReleases automatically updates to the latest release version (mvn versions:use-latest-releases)
func VersionsUseLatestReleases(executable string) (string, error) {
	return ExecForStdout(executable, "versions:use-latest-releases")
}

// VersionsUseNextReleases automatically updates to the next release version (mvn versions:use-next-releases)
func VersionsUseNextReleases(executable string) (string, error) {
	return ExecForStdout(executable, "versions:use-next-releases")
}
