package command

// Exec 插件相关命令
// exec 插件用于在 Maven 构建过程中执行 Java 程序或系统命令

// ExecJava 运行 Java 程序（mvn exec:java）
// 需要在 POM 中配置 mainClass 或通过 -Dexec.mainClass 指定
func ExecJava(executable string) (string, error) {
	return ExecForStdout(executable, "exec:java")
}

// ExecJavaWithMainClass 运行指定主类的 Java 程序
func ExecJavaWithMainClass(executable, mainClass string) (string, error) {
	return ExecForStdout(executable, "exec:java", "-Dexec.mainClass="+mainClass)
}

// ExecExec 执行系统命令（mvn exec:exec）
// 需要在 POM 中配置可执行文件或通过 -Dexec.executable 指定
func ExecExec(executable string) (string, error) {
	return ExecForStdout(executable, "exec:exec")
}
