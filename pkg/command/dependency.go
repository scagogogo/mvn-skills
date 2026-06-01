package command

// DependencyTree 打印出整个项目的依赖树，需要现在处于项目的根目录下
func DependencyTree(executable string) (string, error) {
	return ExecForStdout(executable, "dependency:tree")
}

// DependencyGet 下载指定的包到本地
func DependencyGet(executable string, groupId, artifactId, version string) (string, error) {
	return ExecForStdout(executable, "dependency:get", "-DgroupId="+groupId, "-DartifactId="+artifactId, "-Dversion="+version)
}

// DependencyResolve 解析所有依赖，将它们下载到本地仓库
func DependencyResolve(executable string) (string, error) {
	return ExecForStdout(executable, "dependency:resolve")
}

// DependencyAnalyze 分析项目的依赖使用情况，检测未使用的依赖和未声明的依赖
func DependencyAnalyze(executable string) (string, error) {
	return ExecForStdout(executable, "dependency:analyze")
}

// DependencyList 列出项目所有已解析的依赖
func DependencyList(executable string) (string, error) {
	return ExecForStdout(executable, "dependency:list")
}

// DependencyPurgeLocalRepository 清除本地仓库中的依赖，然后重新解析
func DependencyPurgeLocalRepository(executable string) (string, error) {
	return ExecForStdout(executable, "dependency:purge-local-repository")
}
