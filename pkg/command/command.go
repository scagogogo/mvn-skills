package command

import (
	"bytes"
	"os/exec"
)

// Exec 执行mvn命令，比较底层的API
func Exec(options *Options) error {

	// 如果没有设置可执行文件的位置的话，默认是从PATH中获取执行
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

// ExecForStdout 执行 Maven 命令并返回标准输出
// 当命令执行失败时，返回包含 stderr 信息的 MavenError
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
