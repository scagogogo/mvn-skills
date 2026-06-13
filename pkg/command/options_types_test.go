package command

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDependencyGetOption_ToArgs(t *testing.T) {
	opts := &DependencyGetOption{
		GroupId:    "joda-time",
		ArtifactId: "joda-time",
		Version:    "2.10.10",
	}
	args := opts.ToArgs()
	argsStr := strings.Join(args, " ")

	assert.Contains(t, argsStr, "dependency:get")
	assert.Contains(t, argsStr, "-DgroupId=joda-time")
	assert.Contains(t, argsStr, "-DartifactId=joda-time")
	assert.Contains(t, argsStr, "-Dversion=2.10.10")
	// No classifier or packaging by default
	assert.NotContains(t, argsStr, "-Dclassifier=")
	assert.NotContains(t, argsStr, "-Dpackaging=")
}

func TestDependencyGetOption_ToArgsWithAll(t *testing.T) {
	opts := &DependencyGetOption{
		GroupId:            "com.example",
		ArtifactId:         "my-artifact",
		Version:            "1.0.0",
		Classifier:         "sources",
		Type:               "war",
		RemoteRepositories: []string{"https://repo.example.com/maven2"},
	}
	args := opts.ToArgs()
	argsStr := strings.Join(args, " ")

	assert.Contains(t, argsStr, "-Dclassifier=sources")
	assert.Contains(t, argsStr, "-Dpackaging=war")
	assert.Contains(t, argsStr, "-DremoteRepositories=")
}

func TestDeployDeployFileOption_ToArgs(t *testing.T) {
	opts := &DeployDeployFileOption{
		File:         "/path/to/artifact.jar",
		GroupId:      "com.example",
		ArtifactId:   "my-artifact",
		Version:      "1.0.0",
		Packaging:    "jar",
		RepositoryId: "internal-repo",
		URL:          "https://repo.example.com/maven2",
	}
	args := opts.ToArgs()
	argsStr := strings.Join(args, " ")

	assert.Contains(t, argsStr, "deploy:deploy-file")
	assert.Contains(t, argsStr, "-Dfile=/path/to/artifact.jar")
	assert.Contains(t, argsStr, "-DgroupId=com.example")
	assert.Contains(t, argsStr, "-DartifactId=my-artifact")
	assert.Contains(t, argsStr, "-Dversion=1.0.0")
	assert.Contains(t, argsStr, "-Dpackaging=jar")
	assert.Contains(t, argsStr, "-DrepositoryId=internal-repo")
	assert.Contains(t, argsStr, "-Durl=https://repo.example.com/maven2")
}

func TestDeployDeployFileOption_WithPomFile(t *testing.T) {
	opts := &DeployDeployFileOption{
		File:    "/path/to/artifact.jar",
		PomFile: "/path/to/pom.xml",
	}
	args := opts.ToArgs()
	argsStr := strings.Join(args, " ")

	assert.Contains(t, argsStr, "-DpomFile=/path/to/pom.xml")
	// When PomFile is set, groupId/artifactId/version/packaging should NOT be set
	assert.NotContains(t, argsStr, "-DgroupId=")
	assert.NotContains(t, argsStr, "-DartifactId=")
}

func TestInstallFileOption_ToArgs(t *testing.T) {
	opts := &InstallFileOption{
		File:       "/path/to/artifact.war",
		GroupId:    "com.example",
		ArtifactId: "my-webapp",
		Version:    "2.0.0",
		Packaging:  "war",
	}
	args := opts.ToArgs()
	argsStr := strings.Join(args, " ")

	assert.Contains(t, argsStr, "install:install-file")
	assert.Contains(t, argsStr, "-Dfile=/path/to/artifact.war")
	assert.Contains(t, argsStr, "-Dpackaging=war")
}

func TestInstallFileOption_DefaultPackaging(t *testing.T) {
	opts := &InstallFileOption{
		File:       "/path/to/artifact.jar",
		GroupId:    "com.example",
		ArtifactId: "my-lib",
		Version:    "1.0.0",
	}
	args := opts.ToArgs()
	argsStr := strings.Join(args, " ")

	// Default packaging should be "jar"
	assert.Contains(t, argsStr, "-Dpackaging=jar")
}

func TestInstallFileOption_WithPomFile(t *testing.T) {
	opts := &InstallFileOption{
		File:    "/path/to/artifact.jar",
		PomFile: "/path/to/pom.xml",
	}
	args := opts.ToArgs()
	argsStr := strings.Join(args, " ")

	assert.Contains(t, argsStr, "-DpomFile=/path/to/pom.xml")
	assert.NotContains(t, argsStr, "-Dpackaging=")
}

func TestArchetypeGenerateOption_ToArgs(t *testing.T) {
	opts := &ArchetypeGenerateOption{
		ArchetypeGroupId:    "org.apache.maven.archetypes",
		ArchetypeArtifactId: "maven-archetype-quickstart",
		ArchetypeVersion:    "1.4",
		GroupId:             "com.example",
		ArtifactId:          "my-app",
		Version:             "1.0-SNAPSHOT",
		InteractiveMode:     false,
	}
	args := opts.ToArgs()
	argsStr := strings.Join(args, " ")

	assert.Contains(t, argsStr, "archetype:generate")
	assert.Contains(t, argsStr, "-DarchetypeGroupId=org.apache.maven.archetypes")
	assert.Contains(t, argsStr, "-DarchetypeArtifactId=maven-archetype-quickstart")
	assert.Contains(t, argsStr, "-DarchetypeVersion=1.4")
	assert.Contains(t, argsStr, "-DgroupId=com.example")
	assert.Contains(t, argsStr, "-DartifactId=my-app")
	assert.Contains(t, argsStr, "-Dversion=1.0-SNAPSHOT")
	assert.Contains(t, argsStr, "-DinteractiveMode=false")
}
