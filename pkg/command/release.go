package command

// Release 插件相关命令
// release 插件是 Maven 标准的发布流程工具

// ReleasePrepare 准备发布（mvn release:prepare）
// 执行版本检查、打标签、更新到下一个开发版本
func ReleasePrepare(executable string) (string, error) {
	return ExecForStdout(executable, "release:prepare")
}

// ReleasePrepareWithArgs 带参数准备发布
// 常用参数如 -Darguments="-DskipTests" 跳过测试
func ReleasePrepareWithArgs(executable string, args ...string) (string, error) {
	allArgs := append([]string{"release:prepare"}, args...)
	return ExecForStdout(executable, allArgs...)
}

// ReleasePerform 执行发布（mvn release:perform）
// 从标签检出代码并执行 deploy
func ReleasePerform(executable string) (string, error) {
	return ExecForStdout(executable, "release:perform")
}

// ReleaseRollback 回滚发布准备（mvn release:rollback）
// 当 release:prepare 失败或发现问题时执行
func ReleaseRollback(executable string) (string, error) {
	return ExecForStdout(executable, "release:rollback")
}

// ReleaseClean 清理发布状态（mvn release:clean）
// 清理 release.properties 和其他发布临时文件
func ReleaseClean(executable string) (string, error) {
	return ExecForStdout(executable, "release:clean")
}
