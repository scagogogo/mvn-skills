package command

// JAR/Source/Javadoc plugin-related commands
// These commands are required steps for publishing to Maven Central

// JarJar directly invokes jar:jar to create a JAR package
// Compared to the Package() lifecycle phase, it provides finer-grained control
func JarJar(executable string) (string, error) {
	return ExecForStdout(executable, "jar:jar")
}

// SourceJar generates a source JAR package (mvn source:jar)
// A source package is required when publishing to Maven Central
func SourceJar(executable string) (string, error) {
	return ExecForStdout(executable, "source:jar")
}

// SourceJarNoFork generates a source JAR package without forking the lifecycle (mvn source:jar-no-fork)
// Used within an existing build process; does not re-run the lifecycle
func SourceJarNoFork(executable string) (string, error) {
	return ExecForStdout(executable, "source:jar-no-fork")
}

// JavadocJavadoc generates Javadoc documentation (mvn javadoc:javadoc)
func JavadocJavadoc(executable string) (string, error) {
	return ExecForStdout(executable, "javadoc:javadoc")
}

// JavadocJar generates a Javadoc JAR package (mvn javadoc:jar)
// A Javadoc package is required when publishing to Maven Central
func JavadocJar(executable string) (string, error) {
	return ExecForStdout(executable, "javadoc:jar")
}
