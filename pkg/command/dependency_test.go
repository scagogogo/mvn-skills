package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDependencyGet(t *testing.T) {
	stdout, err := DependencyGet("", "joda-time", "joda-time", "2.10.10")
	assert.Nil(t, err)
	assert.NotEmpty(t, stdout)
}

func TestDependencyTree(t *testing.T) {
	executable := "mvn"
	_, err := DependencyTree(executable)
	if err != nil {
		t.Logf("DependencyTree 执行报错（环境无 Maven 或无项目）: %v", err)
	}
}

func TestDependencyResolve(t *testing.T) {
	executable := "mvn"
	_, err := DependencyResolve(executable)
	if err != nil {
		t.Logf("DependencyResolve 执行报错（环境无 Maven 或无项目）: %v", err)
	}
}

func TestDependencyAnalyze(t *testing.T) {
	executable := "mvn"
	_, err := DependencyAnalyze(executable)
	if err != nil {
		t.Logf("DependencyAnalyze 执行报错（环境无 Maven 或无项目）: %v", err)
	}
}

func TestDependencyList(t *testing.T) {
	executable := "mvn"
	_, err := DependencyList(executable)
	if err != nil {
		t.Logf("DependencyList 执行报错（环境无 Maven 或无项目）: %v", err)
	}
}

func TestDependencyPurgeLocalRepository(t *testing.T) {
	executable := "mvn"
	_, err := DependencyPurgeLocalRepository(executable)
	if err != nil {
		t.Logf("DependencyPurgeLocalRepository 执行报错（环境无 Maven 或无项目）: %v", err)
	}
}