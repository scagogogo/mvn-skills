package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDependencyGet(t *testing.T) {
	stdout, err := DependencyGet("", "joda-time", "joda-time", "2.10.10")
	if err != nil {
		t.Skip("Maven not installed, skipping DependencyGet test")
	}
	assert.NotEmpty(t, stdout)
}

func TestDependencyTree(t *testing.T) {
	executable := "mvn"
	_, err := DependencyTree(executable)
	if err != nil {
		t.Skip("Maven not installed or no project, skipping DependencyTree test")
	}
}

func TestDependencyResolve(t *testing.T) {
	executable := "mvn"
	_, err := DependencyResolve(executable)
	if err != nil {
		t.Skip("Maven not installed or no project, skipping DependencyResolve test")
	}
}

func TestDependencyAnalyze(t *testing.T) {
	executable := "mvn"
	_, err := DependencyAnalyze(executable)
	if err != nil {
		t.Skip("Maven not installed or no project, skipping DependencyAnalyze test")
	}
}

func TestDependencyList(t *testing.T) {
	executable := "mvn"
	_, err := DependencyList(executable)
	if err != nil {
		t.Skip("Maven not installed or no project, skipping DependencyList test")
	}
}

func TestDependencyPurgeLocalRepository(t *testing.T) {
	executable := "mvn"
	_, err := DependencyPurgeLocalRepository(executable)
	if err != nil {
		t.Skip("Maven not installed or no project, skipping DependencyPurgeLocalRepository test")
	}
}
