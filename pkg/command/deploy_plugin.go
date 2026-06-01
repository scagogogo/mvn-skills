package command

// Deploy plugin-related commands
// The deploy plugin is used to deploy build artifacts to a remote Maven repository

// DeployDeploy directly invokes deploy:deploy to deploy an artifact (mvn deploy:deploy)
// Compared to the Deploy() lifecycle phase, it only performs the deploy operation without running prior phases
func DeployDeploy(executable string) (string, error) {
	return ExecForStdout(executable, "deploy:deploy")
}

// DeployDeployFile deploys a specified file to a remote repository (mvn deploy:deploy-file)
// Commonly used to deploy third-party artifacts to an internal repository
func DeployDeployFile(executable, file, groupId, artifactId, version, repositoryId, url string) (string, error) {
	return ExecForStdout(executable,
		"deploy:deploy-file",
		"-Dfile="+file,
		"-DgroupId="+groupId,
		"-DartifactId="+artifactId,
		"-Dversion="+version,
		"-Dpackaging=jar",
		"-DrepositoryId="+repositoryId,
		"-Durl="+url,
	)
}