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