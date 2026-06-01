package command

import (
	"fmt"
	"strings"
)

// MavenError 表示 Maven 命令执行失败的错误
// 包含命令信息和 stderr 输出，便于诊断问题
type MavenError struct {
	Command   string   // 执行的完整命令（如 "mvn clean install"）
	Args      []string // 命令参数
	Stderr    string   // Maven 的 stderr 输出
	ExitCode  int      // 进程退出码（如果可用）
	Inner     error    // 原始错误
}

func (e *MavenError) Error() string {
	var parts []string
	parts = append(parts, fmt.Sprintf("maven command failed: %s %s", e.Command, strings.Join(e.Args, " ")))
	if e.Stderr != "" {
		// 只取 stderr 的前 500 字符，避免错误信息过长
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

// NewMavenError 创建一个新的 MavenError
func NewMavenError(command string, args []string, stderr string, inner error) *MavenError {
	return &MavenError{
		Command: command,
		Args:    args,
		Stderr:  stderr,
		Inner:   inner,
	}
}