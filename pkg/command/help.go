package command

// EffectivePom displays the effective POM configuration for the current project, merging all active profiles
func EffectivePom(executable string) (string, error) {
	return ExecForStdout(executable, "help:effective-pom")
}

// EffectiveSettings displays the effective Maven settings, merging global and user-level settings.xml
func EffectiveSettings(executable string) (string, error) {
	return ExecForStdout(executable, "help:effective-settings")
}

// ActiveProfiles displays all active Maven profiles in the current project
func ActiveProfiles(executable string) (string, error) {
	return ExecForStdout(executable, "help:active-profiles")
}

// DescribePlugin describes the goal details of the specified Maven plugin
func DescribePlugin(executable string, plugin string) (string, error) {
	return ExecForStdout(executable, "help:describe", "-Dplugin="+plugin)
}