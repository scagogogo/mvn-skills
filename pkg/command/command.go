package command

import (
	"bytes"
	"context"
	"os/exec"
)

// Exec executes a Maven command; this is a low-level API
func Exec(options *Options) error {

	// If the executable path is not set, default to finding it from PATH
	if options.Executable == "" {
		options.Executable = "mvn"
	}

	ctx := options.Context
	if ctx == nil {
		ctx = context.Background()
	}

	var command *exec.Cmd
	if len(options.Env) > 0 {
		command = exec.CommandContext(ctx, options.Executable, options.Args...)
		command.Env = append(command.Environ(), options.Env...)
	} else {
		command = exec.CommandContext(ctx, options.Executable, options.Args...)
	}
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

// ExecWithContext executes a Maven command with context support for cancellation and timeouts
func ExecWithContext(ctx context.Context, options *Options) error {
	options.Context = ctx
	return Exec(options)
}

// ExecForStdoutWithContext executes a Maven command with context and returns stdout
func ExecForStdoutWithContext(ctx context.Context, executable string, args ...string) (string, error) {
	if executable == "" {
		executable = "mvn"
	}

	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}
	err := Exec(&Options{
		Context:    ctx,
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
