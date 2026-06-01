package command

// Enforcer plugin-related commands
// The enforcer plugin is used to enforce build rules (dependency convergence, Java version, etc.)

// EnforcerEnforce executes build rule checks (mvn enforcer:enforce)
// Commonly used in CI pipelines to enforce project standards
func EnforcerEnforce(executable string) (string, error) {
	return ExecForStdout(executable, "enforcer:enforce")
}