package command

// Wrapper 在项目中生成 Maven Wrapper 文件，使项目可以在没有安装 Maven 的情况下构建
func Wrapper(executable string) (string, error) {
	return ExecForStdout(executable, "wrapper:wrapper")
}