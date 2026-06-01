package command

// DependencyTree prints the entire project dependency tree; must be run from the project root directory
func DependencyTree(executable string) (string, error) {
	return ExecForStdout(executable, "dependency:tree")
}

// DependencyGet downloads the specified artifact to the local repository
func DependencyGet(executable string, groupId, artifactId, version string) (string, error) {
	return ExecForStdout(executable, "dependency:get", "-DgroupId="+groupId, "-DartifactId="+artifactId, "-Dversion="+version)
}

// DependencyResolve resolves all dependencies and downloads them to the local repository
func DependencyResolve(executable string) (string, error) {
	return ExecForStdout(executable, "dependency:resolve")
}

// DependencyAnalyze analyzes project dependency usage, detecting unused and undeclared dependencies
func DependencyAnalyze(executable string) (string, error) {
	return ExecForStdout(executable, "dependency:analyze")
}

// DependencyList lists all resolved dependencies in the project
func DependencyList(executable string) (string, error) {
	return ExecForStdout(executable, "dependency:list")
}

// DependencyPurgeLocalRepository purges dependencies from the local repository and then re-resolves them
func DependencyPurgeLocalRepository(executable string) (string, error) {
	return ExecForStdout(executable, "dependency:purge-local-repository")
}