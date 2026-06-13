package command

// Install runs clean install on the current project
func Install(executable string) (string, error) {
	return ExecForStdout(executable, "clean", "install")
}

// InstallJar installs a JAR file to the local repository
func InstallJar(executable string, jarPath string, groupId, artifactId, version string) (string, error) {
	return ExecForStdout(executable, "install:install-file", "-Dfile="+jarPath, "-DgroupId="+groupId, "-DartifactId="+artifactId, "-Dversion="+version, "-Dpackaging=jar")
}

// InstallFile installs an artifact to the local repository using structured options.
// This supports all packaging types (jar, war, pom, ear, aar, etc.) and optional
// classifiers, source/javadoc JARs, and POM files.
func InstallFile(executable string, opts *InstallFileOption) (string, error) {
	args := opts.ToArgs()
	return ExecForStdout(executable, args...)
}

// DeployDeployFileWithOptions deploys an artifact to a remote repository using structured options.
// This supports all packaging types, repository selection, and optional classifiers.
func DeployDeployFileWithOptions(executable string, opts *DeployDeployFileOption) (string, error) {
	args := opts.ToArgs()
	return ExecForStdout(executable, args...)
}

// ArchetypeGenerate creates a new project from an archetype using structured options.
func ArchetypeGenerate(executable string, opts *ArchetypeGenerateOption) (string, error) {
	args := opts.ToArgs()
	return ExecForStdout(executable, args...)
}
