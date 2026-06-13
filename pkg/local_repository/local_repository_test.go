package local_repository

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildDirectory(t *testing.T) {
	result := BuildDirectory("com.alibaba", "fastjson", "2.0.2")
	assert.Equal(t, filepath.Join("com/alibaba", "fastjson", "2.0.2"), result)
}

func TestFindJarWithClassifier(t *testing.T) {
	// 创建临时目录模拟本地仓库结构
	tmpDir := t.TempDir()
	gavDir := filepath.Join(tmpDir, "com", "example", "lib", "1.0.0")
	err := os.MkdirAll(gavDir, 0755)
	assert.Nil(t, err)

	// 创建主构件 JAR
	mainJar := filepath.Join(gavDir, "lib-1.0.0.jar")
	err = os.WriteFile(mainJar, []byte("main jar"), 0644)
	assert.Nil(t, err)

	// 创建 sources JAR
	sourcesJar := filepath.Join(gavDir, "lib-1.0.0-sources.jar")
	err = os.WriteFile(sourcesJar, []byte("sources jar"), 0644)
	assert.Nil(t, err)

	// 创建 javadoc JAR
	javadocJar := filepath.Join(gavDir, "lib-1.0.0-javadoc.jar")
	err = os.WriteFile(javadocJar, []byte("javadoc jar"), 0644)
	assert.Nil(t, err)

	// 测试无 classifier（主构件）
	jarPath, err := FindJarWithClassifier(tmpDir, "com.example", "lib", "1.0.0", "")
	assert.Nil(t, err)
	assert.Equal(t, mainJar, jarPath)

	// 测试 classifier = sources
	jarPath, err = FindJarWithClassifier(tmpDir, "com.example", "lib", "1.0.0", "sources")
	assert.Nil(t, err)
	assert.Equal(t, sourcesJar, jarPath)

	// 测试 classifier = javadoc
	jarPath, err = FindJarWithClassifier(tmpDir, "com.example", "lib", "1.0.0", "javadoc")
	assert.Nil(t, err)
	assert.Equal(t, javadocJar, jarPath)

	// 测试不存在的 classifier
	_, err = FindJarWithClassifier(tmpDir, "com.example", "lib", "1.0.0", "nonexistent")
	assert.NotNil(t, err)

	// 测试 FindJar（无 classifier 版本）
	jarPath, err = FindJar(tmpDir, "com.example", "lib", "1.0.0")
	assert.Nil(t, err)
	assert.Equal(t, mainJar, jarPath)
}

func TestParseLocalRepositoryDirectory(t *testing.T) {
	// 无 Maven 环境时返回默认目录
	directory := ParseLocalRepositoryDirectory("mvn")
	assert.NotEmpty(t, directory)
	assert.Contains(t, directory, ".m2")
}
