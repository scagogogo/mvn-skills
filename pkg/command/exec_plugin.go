package command

// Exec plugin-related commands
// The exec plugin is used to execute Java programs or system commands during the Maven build process

// ExecJava runs a Java program (mvn exec:java)
// Requires mainClass to be configured in the POM or specified via -Dexec.mainClass
func ExecJava(executable string) (string, error) {
	return ExecForStdout(executable, "exec:java")
}

// ExecJavaWithMainClass runs a Java program with the specified main class
func ExecJavaWithMainClass(executable, mainClass string) (string, error) {
	return ExecForStdout(executable, "exec:java", "-Dexec.mainClass="+mainClass)
}

// ExecExec executes a system command (mvn exec:exec)
// Requires an executable to be configured in the POM or specified via -Dexec.executable
func ExecExec(executable string) (string, error) {
	return ExecForStdout(executable, "exec:exec")
}