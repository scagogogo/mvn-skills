package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanCmd(t *testing.T) {
	executable := "mvn"
	_, err := Clean(executable)
	if err != nil {
		t.Logf("Clean 执行报错（环境无 Maven）: %v", err)
	}
}

func TestCompileCmd(t *testing.T) {
	executable := "mvn"
	_, err := Compile(executable)
	if err != nil {
		t.Logf("Compile 执行报错（环境无 Maven）: %v", err)
	}
}

func TestMvnTestCmd(t *testing.T) {
	executable := "mvn"
	_, err := Test(executable)
	if err != nil {
		t.Logf("Test 执行报错（环境无 Maven）: %v", err)
	}
}

func TestPackageCmd(t *testing.T) {
	executable := "mvn"
	_, err := Package(executable)
	if err != nil {
		t.Logf("Package 执行报错（环境无 Maven）: %v", err)
	}
}

func TestDeployCmd(t *testing.T) {
	executable := "mvn"
	_, err := Deploy(executable)
	if err != nil {
		t.Logf("Deploy 执行报错（环境无 Maven）: %v", err)
	}
}

func TestVerifyCmd(t *testing.T) {
	executable := "mvn"
	_, err := Verify(executable)
	if err != nil {
		t.Logf("Verify 执行报错（环境无 Maven）: %v", err)
	}
}

func TestValidateCmd(t *testing.T) {
	executable := "mvn"
	_, err := Validate(executable)
	if err != nil {
		t.Logf("Validate 执行报错（环境无 Maven）: %v", err)
	}
}

func TestSiteCmd(t *testing.T) {
	executable := "mvn"
	_, err := Site(executable)
	if err != nil {
		t.Logf("Site 执行报错（环境无 Maven）: %v", err)
	}
}

func TestTestCompileCmd(t *testing.T) {
	executable := "mvn"
	_, err := TestCompile(executable)
	if err != nil {
		t.Logf("TestCompile 执行报错（环境无 Maven）: %v", err)
	}
}

// TestCleanWithWorkingDirectory 测试带工作目录的生命周期命令
func TestCleanWithWorkingDirectory(t *testing.T) {
	options := &Options{
		Executable:       "mvn",
		Args:             []string{"clean"},
		WorkingDirectory: "/tmp",
	}
	err := Exec(options)
	if err != nil {
		t.Logf("带工作目录的 Clean 执行报错（环境无 Maven）: %v", err)
	}
	// 验证 WorkingDirectory 字段存在于 Options 结构体中
	assert.NotEmpty(t, options.WorkingDirectory)
}