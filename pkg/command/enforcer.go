package command

// Enforcer 插件相关命令
// enforcer 插件用于执行构建规则（依赖收敛、Java 版本等）

// EnforcerEnforce 执行构建规则检查（mvn enforcer:enforce）
// 常用于 CI 流水线中强制执行项目规范
func EnforcerEnforce(executable string) (string, error) {
	return ExecForStdout(executable, "enforcer:enforce")
}
