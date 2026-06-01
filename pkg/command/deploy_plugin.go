package command

// Deploy 插件相关命令
// deploy 插件用于将构建产物部署到远程 Maven 仓库

// DeployDeploy 直接调用 deploy:deploy 部署构件（mvn deploy:deploy）
// 相比 Deploy() 生命周期阶段，只执行部署操作而不执行前置阶段
func DeployDeploy(executable string) (string, error) {
	return ExecForStdout(executable, "deploy:deploy")
}

// DeployDeployFile 部署指定文件到远程仓库（mvn deploy:deploy-file）
// 常用于部署第三方构件到内部仓库
func DeployDeployFile(executable, file, groupId, artifactId, version, repositoryId, url string) (string, error) {
	return ExecForStdout(executable,
		"deploy:deploy-file",
		"-Dfile="+file,
		"-DgroupId="+groupId,
		"-DartifactId="+artifactId,
		"-Dversion="+version,
		"-Dpackaging=jar",
		"-DrepositoryId="+repositoryId,
		"-Durl="+url,
	)
}
