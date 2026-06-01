package command

// Clean 清理项目构建产物，删除 target 目录
func Clean(executable string) (string, error) {
	return ExecForStdout(executable, "clean")
}

// Compile 编译项目源码
func Compile(executable string) (string, error) {
	return ExecForStdout(executable, "compile")
}

// Test 运行项目单元测试
func Test(executable string) (string, error) {
	return ExecForStdout(executable, "test")
}

// TestCompile 编译项目测试代码
func TestCompile(executable string) (string, error) {
	return ExecForStdout(executable, "test-compile")
}

// Package 将编译后的代码打包为分发格式（如 jar/war）
func Package(executable string) (string, error) {
	return ExecForStdout(executable, "package")
}

// Verify 对集成测试结果进行检查和验证
func Verify(executable string) (string, error) {
	return ExecForStdout(executable, "verify")
}

// Deploy 将构建产物部署到远程仓库
func Deploy(executable string) (string, error) {
	return ExecForStdout(executable, "deploy")
}

// Site 生成项目站点文档
func Site(executable string) (string, error) {
	return ExecForStdout(executable, "site")
}

// Validate 验证项目结构是否正确且所有必要信息可用
func Validate(executable string) (string, error) {
	return ExecForStdout(executable, "validate")
}
