package command

// Clean cleans up project build artifacts, deleting the target directory
func Clean(executable string) (string, error) {
	return ExecForStdout(executable, "clean")
}

// Compile compiles the project source code
func Compile(executable string) (string, error) {
	return ExecForStdout(executable, "compile")
}

// Test runs the project unit tests
func Test(executable string) (string, error) {
	return ExecForStdout(executable, "test")
}

// TestCompile compiles the project test code
func TestCompile(executable string) (string, error) {
	return ExecForStdout(executable, "test-compile")
}

// Package packages the compiled code into a distribution format (e.g. jar/war)
func Package(executable string) (string, error) {
	return ExecForStdout(executable, "package")
}

// Verify checks and verifies integration test results
func Verify(executable string) (string, error) {
	return ExecForStdout(executable, "verify")
}

// Deploy deploys the build artifacts to a remote repository
func Deploy(executable string) (string, error) {
	return ExecForStdout(executable, "deploy")
}

// Site generates project site documentation
func Site(executable string) (string, error) {
	return ExecForStdout(executable, "site")
}

// Validate validates that the project structure is correct and all necessary information is available
func Validate(executable string) (string, error) {
	return ExecForStdout(executable, "validate")
}
