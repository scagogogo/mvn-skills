package command

import (
	"fmt"
	"strings"
)

// MavenError represents an error from a failed Maven command execution
// It contains command information and stderr output for easier diagnosis
type MavenError struct {
	Command   string   // The full command executed (e.g. "mvn clean install")
	Args      []string // Command arguments
	Stderr    string   // Maven's stderr output
	ExitCode  int      // Process exit code (if available)
	Inner     error    // The original error
}

func (e *MavenError) Error() string {
	var parts []string
	parts = append(parts, fmt.Sprintf("maven command failed: %s %s", e.Command, strings.Join(e.Args, " ")))
	if e.Stderr != "" {
		// Only take the first 500 characters of stderr to avoid overly long error messages
		stderr := e.Stderr
		if len(stderr) > 500 {
			stderr = stderr[:500] + "... (truncated)"
		}
		parts = append(parts, fmt.Sprintf("stderr:\n%s", stderr))
	}
	if e.Inner != nil {
		parts = append(parts, fmt.Sprintf("cause: %v", e.Inner))
	}
	return strings.Join(parts, "\n")
}

func (e *MavenError) Unwrap() error {
	return e.Inner
}

// NewMavenError creates a new MavenError
func NewMavenError(command string, args []string, stderr string, inner error) *MavenError {
	return &MavenError{
		Command: command,
		Args:    args,
		Stderr:  stderr,
		Inner:   inner,
	}
}