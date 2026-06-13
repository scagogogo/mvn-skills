package command

import (
	"testing"
)

func TestEffectivePom(t *testing.T) {
	executable := "mvn"
	_, err := EffectivePom(executable)
	if err != nil {
		t.Logf("EffectivePom 执行报错（环境无 Maven 或无项目）: %v", err)
	}
}

func TestEffectiveSettings(t *testing.T) {
	executable := "mvn"
	_, err := EffectiveSettings(executable)
	if err != nil {
		t.Logf("EffectiveSettings 执行报错（环境无 Maven）: %v", err)
	}
}

func TestActiveProfiles(t *testing.T) {
	executable := "mvn"
	_, err := ActiveProfiles(executable)
	if err != nil {
		t.Logf("ActiveProfiles 执行报错（环境无 Maven 或无项目）: %v", err)
	}
}

func TestDescribePlugin(t *testing.T) {
	executable := "mvn"
	_, err := DescribePlugin(executable, "compiler")
	if err != nil {
		t.Logf("DescribePlugin 执行报错（环境无 Maven）: %v", err)
	}
}
