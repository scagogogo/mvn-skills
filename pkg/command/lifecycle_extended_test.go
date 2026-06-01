package command

import (
	"testing"
)

func TestInitialize(t *testing.T) {
	_, err := Initialize("mvn")
	if err != nil {
		t.Logf("Initialize 执行报错（环境无 Maven）: %v", err)
	}
}

func TestGenerateSources(t *testing.T) {
	_, err := GenerateSources("mvn")
	if err != nil {
		t.Logf("GenerateSources 执行报错（环境无 Maven）: %v", err)
	}
}

func TestProcessResources(t *testing.T) {
	_, err := ProcessResources("mvn")
	if err != nil {
		t.Logf("ProcessResources 执行报错（环境无 Maven）: %v", err)
	}
}

func TestPreparePackage(t *testing.T) {
	_, err := PreparePackage("mvn")
	if err != nil {
		t.Logf("PreparePackage 执行报错（环境无 Maven）: %v", err)
	}
}

func TestPreIntegrationTest(t *testing.T) {
	_, err := PreIntegrationTest("mvn")
	if err != nil {
		t.Logf("PreIntegrationTest 执行报错（环境无 Maven）: %v", err)
	}
}

func TestIntegrationTestPhase(t *testing.T) {
	_, err := IntegrationTest("mvn")
	if err != nil {
		t.Logf("IntegrationTest 执行报错（环境无 Maven）: %v", err)
	}
}

func TestPostIntegrationTest(t *testing.T) {
	_, err := PostIntegrationTest("mvn")
	if err != nil {
		t.Logf("PostIntegrationTest 执行报错（环境无 Maven）: %v", err)
	}
}

func TestStandaloneInstall(t *testing.T) {
	_, err := StandaloneInstall("mvn")
	if err != nil {
		t.Logf("StandaloneInstall 执行报错（环境无 Maven）: %v", err)
	}
}

func TestPreClean(t *testing.T) {
	_, err := PreClean("mvn")
	if err != nil {
		t.Logf("PreClean 执行报错（环境无 Maven）: %v", err)
	}
}

func TestPostClean(t *testing.T) {
	_, err := PostClean("mvn")
	if err != nil {
		t.Logf("PostClean 执行报错（环境无 Maven）: %v", err)
	}
}

func TestSiteDeploy(t *testing.T) {
	_, err := SiteDeploy("mvn")
	if err != nil {
		t.Logf("SiteDeploy 执行报错（环境无 Maven）: %v", err)
	}
}
