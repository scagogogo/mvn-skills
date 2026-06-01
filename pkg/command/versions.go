package command

// Versions 插件相关命令
// versions 插件用于管理项目版本和检查依赖/插件更新

// VersionsSet 设置项目版本（mvn versions:set -DnewVersion=...）
func VersionsSet(executable, newVersion string) (string, error) {
	return ExecForStdout(executable, "versions:set", "-DnewVersion="+newVersion)
}

// VersionsCommit 提交版本变更（mvn versions:commit）
// 在 versions:set 确认无误后执行，删除备份 POM
func VersionsCommit(executable string) (string, error) {
	return ExecForStdout(executable, "versions:commit")
}

// VersionsRevert 回滚版本变更（mvn versions:revert）
// 在 versions:set 后发现问题时执行，恢复备份 POM
func VersionsRevert(executable string) (string, error) {
	return ExecForStdout(executable, "versions:revert")
}

// VersionsDisplayDependencyUpdates 检查可用的依赖更新（mvn versions:display-dependency-updates）
func VersionsDisplayDependencyUpdates(executable string) (string, error) {
	return ExecForStdout(executable, "versions:display-dependency-updates")
}

// VersionsDisplayPluginUpdates 检查可用的插件更新（mvn versions:display-plugin-updates）
func VersionsDisplayPluginUpdates(executable string) (string, error) {
	return ExecForStdout(executable, "versions:display-plugin-updates")
}

// VersionsUseLatestReleases 自动更新到最新的发布版本（mvn versions:use-latest-releases）
func VersionsUseLatestReleases(executable string) (string, error) {
	return ExecForStdout(executable, "versions:use-latest-releases")
}

// VersionsUseNextReleases 自动更新到下一个发布版本（mvn versions:use-next-releases）
func VersionsUseNextReleases(executable string) (string, error) {
	return ExecForStdout(executable, "versions:use-next-releases")
}
