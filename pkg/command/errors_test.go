package command

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMavenError_Error(t *testing.T) {
	inner := errors.New("exit status 1")
	me := NewMavenError("mvn", []string{"clean", "install"}, "BUILD FAILURE\nCompilation error", inner)

	errMsg := me.Error()
	assert.Contains(t, errMsg, "maven command failed")
	assert.Contains(t, errMsg, "mvn")
	assert.Contains(t, errMsg, "clean install")
	assert.Contains(t, errMsg, "BUILD FAILURE")
	assert.Contains(t, errMsg, "exit status 1")
}

func TestMavenError_Unwrap(t *testing.T) {
	inner := errors.New("exit status 1")
	me := NewMavenError("mvn", []string{"test"}, "", inner)
	assert.Equal(t, inner, me.Unwrap())
}

func TestMavenError_StderrTruncation(t *testing.T) {
	longStderr := strings.Repeat("x", 600)
	me := NewMavenError("mvn", []string{"compile"}, longStderr, nil)

	errMsg := me.Error()
	assert.Contains(t, errMsg, "truncated")
	// 确保错误消息不会过长
	assert.Less(t, len(errMsg), 1000)
}

func TestMavenError_EmptyStderr(t *testing.T) {
	me := NewMavenError("mvn", []string{"clean"}, "", nil)
	errMsg := me.Error()
	assert.Contains(t, errMsg, "maven command failed")
	assert.NotContains(t, errMsg, "stderr")
}

func TestExecForStdout_InvalidCommand(t *testing.T) {
	// 执行不存在的命令应该返回 MavenError
	_, err := ExecForStdout("/nonexistent/mvn", "clean")
	assert.NotNil(t, err)

	// 应该是 MavenError 类型
	var me *MavenError
	assert.True(t, errors.As(err, &me))
	assert.Equal(t, "/nonexistent/mvn", me.Command)
}

func TestExecForStdout_DefaultExecutable(t *testing.T) {
	// 空字符串应该默认为 "mvn"
	// 这个测试验证空字符串被替换为 "mvn"
	// 在没有 Maven 的环境中会失败，但错误应该是 MavenError
	_, err := ExecForStdout("", "clean")
	if err != nil {
		var me *MavenError
		assert.True(t, errors.As(err, &me))
		assert.Equal(t, "mvn", me.Command)
	}
}
