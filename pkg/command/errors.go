package command

import (
	"fmt"
	"os/exec"
	"strings"
)

// MavenError represents an error from a failed Maven command execution
// It contains command information and stderr output for easier diagnosis
type MavenError struct {
	Command  string   // The full command executed (e.g. "mvn clean install")
	Args     []string // Command arguments
	Stderr   string   // Maven's stderr output
	ExitCode int      // Process exit code (0 if not available)
	Inner    error    // The original error
}

func (e *MavenError) Error() string {
	var parts []string
	parts = append(parts, fmt.Sprintf("maven command failed: %s %s", e.Command, strings.Join(e.Args, " ")))
	if e.ExitCode > 0 {
		parts = append(parts, fmt.Sprintf("exit code: %d", e.ExitCode))
	}
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

// NewMavenError creates a new MavenError, extracting the exit code from the inner error if available
func NewMavenError(command string, args []string, stderr string, inner error) *MavenError {
	me := &MavenError{
		Command: command,
		Args:    args,
		Stderr:  stderr,
		Inner:   inner,
	}

	// Extract exit code from exec.ExitError
	if inner != nil {
		var exitErr *exec.ExitError
		if exitErr, me.Inner = extractExitError(inner); exitErr != nil {
			me.ExitCode = exitErr.ExitCode()
		}
	}

	return me
}

// extractExitError attempts to find an *exec.ExitError in the error chain
func extractExitError(err error) (*exec.ExitError, error) {
	if exitErr, ok := err.(*exec.ExitError); ok {
		return exitErr, nil
	}
	// Try unwrapping for wrapped errors (e.g. from context cancellation)
	if unwrapper, ok := err.(interface{ Unwrap() error }); ok {
		if inner := unwrapper.Unwrap(); inner != nil {
			return extractExitError(inner)
		}
	}
	return nil, err
}
