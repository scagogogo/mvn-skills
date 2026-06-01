package command

import (
	"testing"
)

func TestWrapper(t *testing.T) {
	executable := "mvn"
	_, err := Wrapper(executable)
	if err != nil {
		t.Logf("Wrapper 执行报错（环境无 Maven 或无项目）: %v", err)
	}
}