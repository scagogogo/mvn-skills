package command

// Surefire-related commands
// The surefire plugin is the standard plugin for running unit tests in Maven

// SurefireTest directly invokes surefire:test to run unit tests
// Compared to the Test() lifecycle phase, directly invoking the surefire plugin provides finer control over test execution
func SurefireTest(executable string) (string, error) {
	return ExecForStdout(executable, "surefire:test")
}

// SurefireTestSingleClass runs a single test class
// The className format is a fully qualified name, e.g. "com.example.MyTest"
func SurefireTestSingleClass(executable, className string) (string, error) {
	return ExecForStdout(executable, "surefire:test", "-Dtest="+className)
}

// SurefireTestMethod runs a single test method
// The methodSpec format is "ClassName#methodName"
func SurefireTestMethod(executable, methodSpec string) (string, error) {
	return ExecForStdout(executable, "surefire:test", "-Dtest="+methodSpec)
}