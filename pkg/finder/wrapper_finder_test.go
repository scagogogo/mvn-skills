package finder

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMavenWrapper(t *testing.T) {
	// 创建临时目录模拟项目
	tmpDir := t.TempDir()

	// 目录中没有 mvnw
	_, err := FindMavenWrapper(tmpDir)
	assert.Equal(t, ErrNotFoundMavenWrapper, err)

	// 创建 mvnw 文件
	wrapperPath := filepath.Join(tmpDir, "mvnw")
	err = os.WriteFile(wrapperPath, []byte("#!/bin/sh\n"), 0644)
	assert.Nil(t, err)

	// 现在应该能找到
	found, err := FindMavenWrapper(tmpDir)
	assert.Nil(t, err)
	assert.Equal(t, wrapperPath, found)
}

func TestHasMavenWrapper(t *testing.T) {
	tmpDir := t.TempDir()

	// 没有 mvnw
	assert.False(t, HasMavenWrapper(tmpDir))

	// 创建 mvnw
	wrapperPath := filepath.Join(tmpDir, "mvnw")
	err := os.WriteFile(wrapperPath, []byte("#!/bin/sh\n"), 0644)
	assert.Nil(t, err)

	// 有 mvnw
	assert.True(t, HasMavenWrapper(tmpDir))
}

func TestHasMavenWrapperWithDirectory(t *testing.T) {
	tmpDir := t.TempDir()

	// 创建 mvnw 目录（不是文件）
	wrapperDir := filepath.Join(tmpDir, "mvnw")
	err := os.MkdirAll(wrapperDir, 0755)
	assert.Nil(t, err)

	// 目录不是有效的 wrapper
	assert.False(t, HasMavenWrapper(tmpDir))
}

func TestFindBestMavenWithWrapper(t *testing.T) {
	tmpDir := t.TempDir()

	// 创建 mvnw
	wrapperPath := filepath.Join(tmpDir, "mvnw")
	err := os.WriteFile(wrapperPath, []byte("#!/bin/sh\n"), 0644)
	assert.Nil(t, err)

	// 应该返回 wrapper（优先于系统 Maven）
	maven, err := FindBestMaven(tmpDir)
	assert.Nil(t, err)
	assert.Equal(t, wrapperPath, maven)
}

func TestFindBestMavenWithoutWrapper(t *testing.T) {
	tmpDir := t.TempDir()

	// 没有 mvnw，回退到系统 Maven
	// 这个测试在没有 Maven 的环境中会失败
	maven, err := FindBestMaven(tmpDir)
	if err != nil {
		// System has no Maven — should be a NotFoundError
		var nfe *NotFoundError
		assert.True(t, errors.As(err, &nfe))
	} else {
		assert.NotEmpty(t, maven)
	}
}

func TestGetWrapperScriptName(t *testing.T) {
	name := getWrapperScriptName()
	if runtime.GOOS == "windows" {
		assert.Equal(t, "mvnw.cmd", name)
	} else {
		assert.Equal(t, "mvnw", name)
	}
}