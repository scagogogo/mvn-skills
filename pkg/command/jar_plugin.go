package command

// JAR/Source/Javadoc 插件相关命令
// 这些命令是发布到 Maven Central 的必要步骤

// JarJar 直接调用 jar:jar 创建 JAR 包
// 相比 Package() 生命周期阶段，提供更细粒度的控制
func JarJar(executable string) (string, error) {
	return ExecForStdout(executable, "jar:jar")
}

// SourceJar 生成源码 JAR 包（mvn source:jar）
// 发布到 Maven Central 时必须提供源码包
func SourceJar(executable string) (string, error) {
	return ExecForStdout(executable, "source:jar")
}

// SourceJarNoFork 生成源码 JAR 包但不 fork 生命周期（mvn source:jar-no-fork）
// 在已有构建过程中使用，不会重新运行生命周期
func SourceJarNoFork(executable string) (string, error) {
	return ExecForStdout(executable, "source:jar-no-fork")
}

// JavadocJavadoc 生成 Javadoc 文档（mvn javadoc:javadoc）
func JavadocJavadoc(executable string) (string, error) {
	return ExecForStdout(executable, "javadoc:javadoc")
}

// JavadocJar 生成 Javadoc JAR 包（mvn javadoc:jar）
// 发布到 Maven Central 时必须提供 Javadoc 包
func JavadocJar(executable string) (string, error) {
	return ExecForStdout(executable, "javadoc:jar")
}
