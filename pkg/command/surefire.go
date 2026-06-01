package command

// Surefire 相关命令
// surefire 插件是 Maven 执行单元测试的标准插件

// SurefireTest 直接调用 surefire:test 运行单元测试
// 相比 Test() 生命周期阶段，直接调用 surefire 插件可以更精细地控制测试执行
func SurefireTest(executable string) (string, error) {
	return ExecForStdout(executable, "surefire:test")
}

// SurefireTestSingleClass 运行单个测试类
// className 格式为完全限定名，如 "com.example.MyTest"
func SurefireTestSingleClass(executable, className string) (string, error) {
	return ExecForStdout(executable, "surefire:test", "-Dtest="+className)
}

// SurefireTestMethod 运行单个测试方法
// methodSpec 格式为 "ClassName#methodName"
func SurefireTestMethod(executable, methodSpec string) (string, error) {
	return ExecForStdout(executable, "surefire:test", "-Dtest="+methodSpec)
}
