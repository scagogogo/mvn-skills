package command

import (
	"io"
)

type Options struct {
	Executable      string
	Args            []string
	WorkingDirectory string
	Stdin           io.Reader
	Stdout          io.Writer
	Stderr          io.Writer
}


