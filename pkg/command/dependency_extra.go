package command

// Additional dependency plugin commands
// Supplements the commands already in dependency.go

// DependencyCopy copies the specified artifact to the target directory (mvn dependency:copy)
func DependencyCopy(executable, groupId, artifactId, version, outputDirectory string) (string, error) {
	return ExecForStdout(executable,
		"dependency:copy",
		"-Dartifact="+groupId+":"+artifactId+":"+version,
		"-DoutputDirectory="+outputDirectory,
	)
}

// DependencyCopyDependencies copies all dependencies to the target directory (mvn dependency:copy-dependencies)
// Commonly used for creating distribution packages
func DependencyCopyDependencies(executable, outputDirectory string) (string, error) {
	return ExecForStdout(executable, "dependency:copy-dependencies", "-DoutputDirectory="+outputDirectory)
}

// DependencyUnpack unpacks the specified artifact to the target directory (mvn dependency:unpack)
func DependencyUnpack(executable, groupId, artifactId, version, outputDirectory string) (string, error) {
	return ExecForStdout(executable,
		"dependency:unpack",
		"-Dartifact="+groupId+":"+artifactId+":"+version,
		"-DoutputDirectory="+outputDirectory,
	)
}

// DependencyBuildClasspath generates a classpath string (mvn dependency:build-classpath)
// Returns the project's complete classpath, commonly used for scripts and IDE integration
func DependencyBuildClasspath(executable string) (string, error) {
	return ExecForStdout(executable, "dependency:build-classpath")
}

// DependencyGetWithOptions downloads an artifact using structured options
func DependencyGetWithOptions(executable string, opts *DependencyGetOption) (string, error) {
	args := opts.ToArgs()
	return ExecForStdout(executable, args...)
}

// DependencyResolvePlugins resolves all plugin dependencies (mvn dependency:resolve-plugins)
func DependencyResolvePlugins(executable string) (string, error) {
	return ExecForStdout(executable, "dependency:resolve-plugins")
}

// DependencyGoOffline resolves all dependencies needed for offline builds (mvn dependency:go-offline)
func DependencyGoOffline(executable string) (string, error) {
	return ExecForStdout(executable, "dependency:go-offline")
}

// DependencyUnpackDependencies unpacks all dependencies to the target directory (mvn dependency:unpack-dependencies)
func DependencyUnpackDependencies(executable, outputDirectory string) (string, error) {
	return ExecForStdout(executable, "dependency:unpack-dependencies", "-DoutputDirectory="+outputDirectory)
}
