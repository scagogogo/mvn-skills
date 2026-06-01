package command

// Install installs the current project's dependencies
func Install(executable string) (string, error) {
	// mvn -Dmaven.repo.local=/my/local/repository/path clean install
	//return ExecForStdout(executable, "-Dmaven.repo.local=/my/local/repository/path clean install")
	return ExecForStdout(executable, "clean", "install")
}

// InstallJar installs a JAR file to the local repository
func InstallJar(executable string, jarPath string, groupId, artifactId, version string) (string, error) {
	// mvn install:install-file -Dfile=D:\jaxen-1.1-beta-6.jar -DgroupId=org.jaxen -DartifactId=jaxen1.0 -Dversion=1.1-beta-6 -Dpackaging=jar
	return ExecForStdout(executable, "install:install-file", "-Dfile="+jarPath, "-DgroupId="+groupId, "-DartifactId="+artifactId, "-Dversion="+version, "-Dpackaging=jar")
}