package command

import (
	"context"
	"io"
)

// Options holds the configuration for executing a Maven command
type Options struct {
	Executable      string
	Args            []string
	WorkingDirectory string
	Stdin           io.Reader
	Stdout          io.Writer
	Stderr          io.Writer
	// Env specifies optional environment variables for the Maven process.
	// If nil, the current process environment is used.
	// If set, these are appended to the current process environment.
	Env []string
	// Context allows cancellation of the Maven command.
	// If nil, context.Background() is used.
	Context context.Context
}


