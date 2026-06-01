package command

// Assembly 插件相关命令
// assembly 插件用于创建自定义的分发包（zip、tar.gz 等）

// AssemblySingle 创建分发包（mvn assembly:single）
func AssemblySingle(executable string) (string, error) {
	return ExecForStdout(executable, "assembly:single")
}
