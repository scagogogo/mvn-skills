package pom

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testPomXML = `<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <parent>
        <groupId>com.example</groupId>
        <artifactId>parent-project</artifactId>
        <version>1.0.0</version>
    </parent>

    <artifactId>child-module</artifactId>
    <packaging>jar</packaging>
    <name>Child Module</name>
    <description>A child module for testing</description>

    <properties>
        <java.version>17</java.version>
        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
    </properties>

    <dependencies>
        <dependency>
            <groupId>org.springframework</groupId>
            <artifactId>spring-core</artifactId>
            <version>5.3.21</version>
        </dependency>
        <dependency>
            <groupId>junit</groupId>
            <artifactId>junit</artifactId>
            <version>4.13.2</version>
            <scope>test</scope>
        </dependency>
    </dependencies>

    <modules>
        <module>module-a</module>
        <module>module-b</module>
    </modules>

    <build>
        <plugins>
            <plugin>
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-compiler-plugin</artifactId>
                <version>3.10.1</version>
            </plugin>
        </plugins>
    </build>

    <profiles>
        <profile>
            <id>ci</id>
            <activation>
                <activeByDefault>false</activeByDefault>
            </activation>
        </profile>
    </profiles>

    <repositories>
        <repository>
            <id>central</id>
            <url>https://repo.maven.apache.org/maven2</url>
        </repository>
    </repositories>

    <scm>
        <connection>scm:git:https://github.com/example/project.git</connection>
    </scm>
</project>
`

func TestParseBytes(t *testing.T) {
	project, err := ParseBytes([]byte(testPomXML))
	assert.Nil(t, err)
	assert.NotNil(t, project)

	// 验证基础字段
	assert.Equal(t, "4.0.0", project.ModelVersion)
	assert.Equal(t, "child-module", project.ArtifactId)
	assert.Equal(t, "jar", project.Packaging)
	assert.Equal(t, "Child Module", project.Name)
}

func TestGetGAV(t *testing.T) {
	project, err := ParseBytes([]byte(testPomXML))
	assert.Nil(t, err)

	groupId, artifactId, version := project.GetGAV()
	// groupId 和 version 从 parent 继承
	assert.Equal(t, "com.example", groupId)
	assert.Equal(t, "child-module", artifactId)
	assert.Equal(t, "1.0.0", version)
}

func TestGetGAVWithoutParent(t *testing.T) {
	pomXML := `<?xml version="1.0" encoding="UTF-8"?>
<project>
    <modelVersion>4.0.0</modelVersion>
    <groupId>com.standalone</groupId>
    <artifactId>standalone-project</artifactId>
    <version>2.0.0</version>
</project>`

	project, err := ParseBytes([]byte(pomXML))
	assert.Nil(t, err)

	groupId, artifactId, version := project.GetGAV()
	assert.Equal(t, "com.standalone", groupId)
	assert.Equal(t, "standalone-project", artifactId)
	assert.Equal(t, "2.0.0", version)
}

func TestGetDependencies(t *testing.T) {
	project, err := ParseBytes([]byte(testPomXML))
	assert.Nil(t, err)

	deps := project.GetDependencies()
	assert.Len(t, deps, 2)

	assert.Equal(t, "org.springframework", deps[0].GroupId)
	assert.Equal(t, "spring-core", deps[0].ArtifactId)
	assert.Equal(t, "5.3.21", deps[0].Version)

	assert.Equal(t, "junit", deps[1].GroupId)
	assert.Equal(t, "junit", deps[1].ArtifactId)
	assert.Equal(t, "test", deps[1].Scope)
}

func TestGetModules(t *testing.T) {
	project, err := ParseBytes([]byte(testPomXML))
	assert.Nil(t, err)

	modules := project.GetModules()
	assert.Len(t, modules, 2)
	assert.Contains(t, modules, "module-a")
	assert.Contains(t, modules, "module-b")
}

func TestIsMultiModule(t *testing.T) {
	project, err := ParseBytes([]byte(testPomXML))
	assert.Nil(t, err)
	assert.True(t, project.IsMultiModule())

	singlePom := `<?xml version="1.0" encoding="UTF-8"?><project><artifactId>single</artifactId></project>`
	singleProject, err := ParseBytes([]byte(singlePom))
	assert.Nil(t, err)
	assert.False(t, singleProject.IsMultiModule())
}

func TestHasParent(t *testing.T) {
	project, err := ParseBytes([]byte(testPomXML))
	assert.Nil(t, err)
	assert.True(t, project.HasParent())

	singlePom := `<?xml version="1.0" encoding="UTF-8"?><project><artifactId>single</artifactId></project>`
	singleProject, err := ParseBytes([]byte(singlePom))
	assert.Nil(t, err)
	assert.False(t, singleProject.HasParent())
}

func TestFindDependency(t *testing.T) {
	project, err := ParseBytes([]byte(testPomXML))
	assert.Nil(t, err)

	dep := project.FindDependency("org.springframework", "spring-core")
	assert.NotNil(t, dep)
	assert.Equal(t, "5.3.21", dep.Version)

	dep = project.FindDependency("nonexistent", "nonexistent")
	assert.Nil(t, dep)
}

func TestGetPlugins(t *testing.T) {
	project, err := ParseBytes([]byte(testPomXML))
	assert.Nil(t, err)

	plugins := project.GetPlugins()
	assert.Len(t, plugins, 1)
	assert.Equal(t, "maven-compiler-plugin", plugins[0].ArtifactId)
}

func TestFindPlugin(t *testing.T) {
	project, err := ParseBytes([]byte(testPomXML))
	assert.Nil(t, err)

	plugin := project.FindPlugin("org.apache.maven.plugins", "maven-compiler-plugin")
	assert.NotNil(t, plugin)
	assert.Equal(t, "3.10.1", plugin.Version)
}

func TestGetProfiles(t *testing.T) {
	project, err := ParseBytes([]byte(testPomXML))
	assert.Nil(t, err)

	profiles := project.GetProfiles()
	assert.Len(t, profiles, 1)
	assert.Equal(t, "ci", profiles[0].Id)
}

func TestGetRepositories(t *testing.T) {
	project, err := ParseBytes([]byte(testPomXML))
	assert.Nil(t, err)

	repos := project.GetRepositories()
	assert.Len(t, repos, 1)
	assert.Equal(t, "central", repos[0].Id)
}

func TestParseFile(t *testing.T) {
	// 创建临时 POM 文件
	tmpDir := t.TempDir()
	pomPath := filepath.Join(tmpDir, "pom.xml")
	err := os.WriteFile(pomPath, []byte(testPomXML), 0644)
	assert.Nil(t, err)

	project, err := ParseFile(pomPath)
	assert.Nil(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, "child-module", project.ArtifactId)
}

func TestParseFileNotFound(t *testing.T) {
	_, err := ParseFile("/nonexistent/path/pom.xml")
	assert.NotNil(t, err)
}

func TestParseInvalidXML(t *testing.T) {
	_, err := ParseBytes([]byte("not xml at all"))
	assert.NotNil(t, err)
}

func TestGetDependenciesEmpty(t *testing.T) {
	pomXML := `<?xml version="1.0" encoding="UTF-8"?><project><artifactId>empty</artifactId></project>`
	project, err := ParseBytes([]byte(pomXML))
	assert.Nil(t, err)

	deps := project.GetDependencies()
	assert.NotNil(t, deps)
	assert.Empty(t, deps)
}

func TestGetModulesEmpty(t *testing.T) {
	pomXML := `<?xml version="1.0" encoding="UTF-8"?><project><artifactId>empty</artifactId></project>`
	project, err := ParseBytes([]byte(pomXML))
	assert.Nil(t, err)

	modules := project.GetModules()
	assert.NotNil(t, modules)
	assert.Empty(t, modules)
}

func TestScmInfo(t *testing.T) {
	project, err := ParseBytes([]byte(testPomXML))
	assert.Nil(t, err)

	assert.NotNil(t, project.Scm)
	assert.Equal(t, "scm:git:https://github.com/example/project.git", project.Scm.Connection)
}

func TestGetScm(t *testing.T) {
	project, err := ParseBytes([]byte(testPomXML))
	assert.Nil(t, err)

	scm := project.GetScm()
	assert.NotNil(t, scm)
	assert.Equal(t, "scm:git:https://github.com/example/project.git", scm.Connection)
}

func TestGetScmNil(t *testing.T) {
	pomXML := `<?xml version="1.0" encoding="UTF-8"?><project><artifactId>no-scm</artifactId></project>`
	project, err := ParseBytes([]byte(pomXML))
	assert.Nil(t, err)
	assert.Nil(t, project.GetScm())
}

func TestGetPackaging(t *testing.T) {
	project, err := ParseBytes([]byte(testPomXML))
	assert.Nil(t, err)
	assert.Equal(t, "jar", project.GetPackaging())

	// No packaging specified → defaults to "jar"
	pomXML := `<?xml version="1.0" encoding="UTF-8"?><project><artifactId>no-pack</artifactId></project>`
	project2, err := ParseBytes([]byte(pomXML))
	assert.Nil(t, err)
	assert.Equal(t, "jar", project2.GetPackaging())

	// WAR packaging
	warPom := `<?xml version="1.0" encoding="UTF-8"?><project><artifactId>war-proj</artifactId><packaging>war</packaging></project>`
	project3, err := ParseBytes([]byte(warPom))
	assert.Nil(t, err)
	assert.Equal(t, "war", project3.GetPackaging())
}

func TestGetBuild(t *testing.T) {
	project, err := ParseBytes([]byte(testPomXML))
	assert.Nil(t, err)

	build := project.GetBuild()
	assert.NotNil(t, build)

	pomXML := `<?xml version="1.0" encoding="UTF-8"?><project><artifactId>no-build</artifactId></project>`
	project2, err := ParseBytes([]byte(pomXML))
	assert.Nil(t, err)
	assert.Nil(t, project2.GetBuild())
}

func TestGetLicensesEmpty(t *testing.T) {
	pomXML := `<?xml version="1.0" encoding="UTF-8"?><project><artifactId>no-lic</artifactId></project>`
	project, err := ParseBytes([]byte(pomXML))
	assert.Nil(t, err)

	licenses := project.GetLicenses()
	assert.NotNil(t, licenses)
	assert.Empty(t, licenses)
}

func TestGetDevelopersEmpty(t *testing.T) {
	pomXML := `<?xml version="1.0" encoding="UTF-8"?><project><artifactId>no-dev</artifactId></project>`
	project, err := ParseBytes([]byte(pomXML))
	assert.Nil(t, err)

	developers := project.GetDevelopers()
	assert.NotNil(t, developers)
	assert.Empty(t, developers)
}
