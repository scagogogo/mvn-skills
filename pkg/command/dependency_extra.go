package command

// 依赖插件的额外命令
// 补充 dependency.go 中已有的命令

// DependencyCopy 复制指定构件到目标目录（mvn dependency:copy）
func DependencyCopy(executable, groupId, artifactId, version, outputDirectory string) (string, error) {
	return ExecForStdout(executable,
		"dependency:copy",
		"-Dartifact="+groupId+":"+artifactId+":"+version,
		"-DoutputDirectory="+outputDirectory,
	)
}

// DependencyCopyDependencies 复制所有依赖到目标目录（mvn dependency:copy-dependencies）
// 常用于创建分发包
func DependencyCopyDependencies(executable, outputDirectory string) (string, error) {
	return ExecForStdout(executable, "dependency:copy-dependencies", "-DoutputDirectory="+outputDirectory)
}

// DependencyUnpack 解压指定构件到目标目录（mvn dependency:unpack）
func DependencyUnpack(executable, groupId, artifactId, version, outputDirectory string) (string, error) {
	return ExecForStdout(executable,
		"dependency:unpack",
		"-Dartifact="+groupId+":"+artifactId+":"+version,
		"-DoutputDirectory="+outputDirectory,
	)
}

// DependencyBuildClasspath 生成 classpath 字符串（mvn dependency:build-classpath）
// 返回项目的完整 classpath，常用于脚本和 IDE 集成
func DependencyBuildClasspath(executable string) (string, error) {
	return ExecForStdout(executable, "dependency:build-classpath")
}
