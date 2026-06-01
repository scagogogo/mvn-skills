package command

// Shade 插件相关命令
// shade 插件用于创建 uber-jar / fat JAR（将所有依赖打入一个 JAR）

// ShadeShade 创建 uber-jar（mvn shade:shade）
func ShadeShade(executable string) (string, error) {
	return ExecForStdout(executable, "shade:shade")
}
