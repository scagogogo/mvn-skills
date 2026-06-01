package command

// Failsafe 相关命令
// failsafe 插件是 Maven 执行集成测试的标准插件

// FailsafeIntegrationTest 运行集成测试（mvn failsafe:integration-test）
func FailsafeIntegrationTest(executable string) (string, error) {
	return ExecForStdout(executable, "failsafe:integration-test")
}

// FailsafeVerify 验证集成测试结果（mvn failsafe:verify）
// 通常在 integration-test 阶段之后执行
func FailsafeVerify(executable string) (string, error) {
	return ExecForStdout(executable, "failsafe:verify")
}
