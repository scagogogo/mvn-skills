package finder

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

// ErrNotFoundMavenWrapper 未找到 Maven Wrapper
var ErrNotFoundMavenWrapper = errors.New("not found maven wrapper")

// FindMavenWrapper 在指定项目目录中查找 Maven Wrapper（mvnw/mvnw.cmd）
// Maven Wrapper 是现代 Maven 项目的推荐做法，项目目录中包含 mvnw 脚本
// 可以在没有安装 Maven 的情况下构建项目
func FindMavenWrapper(projectDir string) (string, error) {
	wrapperName := getWrapperScriptName()
	wrapperPath := filepath.Join(projectDir, wrapperName)

	stat, err := os.Stat(wrapperPath)
	if err != nil {
		return "", ErrNotFoundMavenWrapper
	}
	if stat.IsDir() {
		return "", ErrNotFoundMavenWrapper
	}

	return wrapperPath, nil
}

// FindBestMaven 在指定项目目录中查找最合适的 Maven 可执行文件
// 优先使用项目目录中的 Maven Wrapper，如果没有则使用系统安装的 Maven
func FindBestMaven(projectDir string) (string, error) {
	// 优先查找 Maven Wrapper
	wrapper, err := FindMavenWrapper(projectDir)
	if err == nil {
		return wrapper, nil
	}

	// 回退到系统 Maven
	return FindMaven()
}

// HasMavenWrapper 检查指定项目目录中是否存在 Maven Wrapper
func HasMavenWrapper(projectDir string) bool {
	wrapperName := getWrapperScriptName()
	wrapperPath := filepath.Join(projectDir, wrapperName)

	stat, err := os.Stat(wrapperPath)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

// getWrapperScriptName 根据操作系统返回 Maven Wrapper 脚本文件名
func getWrapperScriptName() string {
	if runtime.GOOS == "windows" {
		return "mvnw.cmd"
	}
	return "mvnw"
}