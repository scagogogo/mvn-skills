package command

import (
	"bytes"
	"os/exec"
)

// Exec executes a Maven command; this is a low-level API
func Exec(options *Options) error {

	// If the executable path is not set, default to finding it from PATH
	if options.Executable == "" {
		options.Executable = "mvn"
	}

	command := exec.Command(options.Executable, options.Args...)
	if options.WorkingDirectory != "" {
		command.Dir = options.WorkingDirectory
	}
	if options.Stdout != nil {
		command.Stdout = options.Stdout
	}
	if options.Stdin != nil {
		command.Stdin = options.Stdin
	}
	if options.Stderr != nil {
		command.Stderr = options.Stderr
	}
	return command.Run()
}

// ExecForStdout executes a Maven command and returns the standard output
// When the command fails, it returns a MavenError containing stderr information
func ExecForStdout(executable string, args ...string) (string, error) {
	if executable == "" {
		executable = "mvn"
	}

	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}
	err := Exec(&Options{
		Executable: executable,
		Stdout:     &stdout,
		Stderr:     &stderr,
		Args:       args,
	})
	if err != nil {
		return "", NewMavenError(executable, args, stderr.String(), err)
	}
	return stdout.String(), nil
}