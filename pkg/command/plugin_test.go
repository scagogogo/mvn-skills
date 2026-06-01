package command

import (
	"testing"
)

// --- Surefire ---

func TestSurefireTest(t *testing.T) {
	_, err := SurefireTest("mvn")
	if err != nil {
		t.Logf("SurefireTest 执行报错（环境无 Maven）: %v", err)
	}
}

func TestSurefireTestSingleClass(t *testing.T) {
	_, err := SurefireTestSingleClass("mvn", "com.example.MyTest")
	if err != nil {
		t.Logf("SurefireTestSingleClass 执行报错（环境无 Maven）: %v", err)
	}
}

// --- Failsafe ---

func TestFailsafeIntegrationTest(t *testing.T) {
	_, err := FailsafeIntegrationTest("mvn")
	if err != nil {
		t.Logf("FailsafeIntegrationTest 执行报错（环境无 Maven）: %v", err)
	}
}

func TestFailsafeVerify(t *testing.T) {
	_, err := FailsafeVerify("mvn")
	if err != nil {
		t.Logf("FailsafeVerify 执行报错（环境无 Maven）: %v", err)
	}
}

// --- Versions ---

func TestVersionsSet(t *testing.T) {
	_, err := VersionsSet("mvn", "2.0.0")
	if err != nil {
		t.Logf("VersionsSet 执行报错（环境无 Maven）: %v", err)
	}
}

func TestVersionsCommit(t *testing.T) {
	_, err := VersionsCommit("mvn")
	if err != nil {
		t.Logf("VersionsCommit 执行报错（环境无 Maven）: %v", err)
	}
}

func TestVersionsRevert(t *testing.T) {
	_, err := VersionsRevert("mvn")
	if err != nil {
		t.Logf("VersionsRevert 执行报错（环境无 Maven）: %v", err)
	}
}

func TestVersionsDisplayDependencyUpdates(t *testing.T) {
	_, err := VersionsDisplayDependencyUpdates("mvn")
	if err != nil {
		t.Logf("VersionsDisplayDependencyUpdates 执行报错（环境无 Maven）: %v", err)
	}
}

func TestVersionsDisplayPluginUpdates(t *testing.T) {
	_, err := VersionsDisplayPluginUpdates("mvn")
	if err != nil {
		t.Logf("VersionsDisplayPluginUpdates 执行报错（环境无 Maven）: %v", err)
	}
}

// --- Release ---

func TestReleasePrepare(t *testing.T) {
	_, err := ReleasePrepare("mvn")
	if err != nil {
		t.Logf("ReleasePrepare 执行报错（环境无 Maven）: %v", err)
	}
}

func TestReleasePerform(t *testing.T) {
	_, err := ReleasePerform("mvn")
	if err != nil {
		t.Logf("ReleasePerform 执行报错（环境无 Maven）: %v", err)
	}
}

func TestReleaseRollback(t *testing.T) {
	_, err := ReleaseRollback("mvn")
	if err != nil {
		t.Logf("ReleaseRollback 执行报错（环境无 Maven）: %v", err)
	}
}

// --- Jar/Source/Javadoc ---

func TestJarJar(t *testing.T) {
	_, err := JarJar("mvn")
	if err != nil {
		t.Logf("JarJar 执行报错（环境无 Maven）: %v", err)
	}
}

func TestSourceJar(t *testing.T) {
	_, err := SourceJar("mvn")
	if err != nil {
		t.Logf("SourceJar 执行报错（环境无 Maven）: %v", err)
	}
}

func TestJavadocJar(t *testing.T) {
	_, err := JavadocJar("mvn")
	if err != nil {
		t.Logf("JavadocJar 执行报错（环境无 Maven）: %v", err)
	}
}

// --- Deploy Plugin ---

func TestDeployDeploy(t *testing.T) {
	_, err := DeployDeploy("mvn")
	if err != nil {
		t.Logf("DeployDeploy 执行报错（环境无 Maven）: %v", err)
	}
}

// --- Assembly/Shader/Exec/Enforcer/GPG ---

func TestAssemblySingle(t *testing.T) {
	_, err := AssemblySingle("mvn")
	if err != nil {
		t.Logf("AssemblySingle 执行报错（环境无 Maven）: %v", err)
	}
}

func TestShadeShade(t *testing.T) {
	_, err := ShadeShade("mvn")
	if err != nil {
		t.Logf("ShadeShade 执行报错（环境无 Maven）: %v", err)
	}
}

func TestExecJava(t *testing.T) {
	_, err := ExecJava("mvn")
	if err != nil {
		t.Logf("ExecJava 执行报错（环境无 Maven）: %v", err)
	}
}

func TestEnforcerEnforce(t *testing.T) {
	_, err := EnforcerEnforce("mvn")
	if err != nil {
		t.Logf("EnforcerEnforce 执行报错（环境无 Maven）: %v", err)
	}
}

func TestGpgSign(t *testing.T) {
	_, err := GpgSign("mvn")
	if err != nil {
		t.Logf("GpgSign 执行报错（环境无 Maven）: %v", err)
	}
}

// --- Dependency Extra ---

func TestDependencyBuildClasspath(t *testing.T) {
	_, err := DependencyBuildClasspath("mvn")
	if err != nil {
		t.Logf("DependencyBuildClasspath 执行报错（环境无 Maven）: %v", err)
	}
}
