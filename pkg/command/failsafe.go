package command

// Failsafe-related commands
// The failsafe plugin is the standard plugin for running integration tests in Maven

// FailsafeIntegrationTest runs integration tests (mvn failsafe:integration-test)
func FailsafeIntegrationTest(executable string) (string, error) {
	return ExecForStdout(executable, "failsafe:integration-test")
}

// FailsafeVerify verifies integration test results (mvn failsafe:verify)
// Typically executed after the integration-test phase
func FailsafeVerify(executable string) (string, error) {
	return ExecForStdout(executable, "failsafe:verify")
}