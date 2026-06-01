package command

// 扩展生命周期阶段命令
// Maven 有 3 个内置生命周期（clean/default/site），共 28 个阶段
// lifecycle.go 已实现最常用的 9 个阶段，本文件补全剩余阶段

// --- Default Lifecycle（剩余阶段）---

// Initialize 初始化构建状态（mvn initialize），常用于在多模块构建中做早期设置
func Initialize(executable string) (string, error) {
	return ExecForStdout(executable, "initialize")
}

// GenerateSources 生成源码（mvn generate-sources），常用于 protobuf、XSD 等代码生成
func GenerateSources(executable string) (string, error) {
	return ExecForStdout(executable, "generate-sources")
}

// ProcessSources 处理源码（mvn process-sources）
func ProcessSources(executable string) (string, error) {
	return ExecForStdout(executable, "process-sources")
}

// GenerateResources 生成资源文件（mvn generate-resources）
func GenerateResources(executable string) (string, error) {
	return ExecForStdout(executable, "generate-resources")
}

// ProcessResources 处理资源文件（mvn process-resources），如变量替换等
func ProcessResources(executable string) (string, error) {
	return ExecForStdout(executable, "process-resources")
}

// ProcessClasses 处理编译后的字节码（mvn process-classes），如字节码增强
func ProcessClasses(executable string) (string, error) {
	return ExecForStdout(executable, "process-classes")
}

// GenerateTestSources 生成测试源码（mvn generate-test-sources）
func GenerateTestSources(executable string) (string, error) {
	return ExecForStdout(executable, "generate-test-sources")
}

// ProcessTestSources 处理测试源码（mvn process-test-sources）
func ProcessTestSources(executable string) (string, error) {
	return ExecForStdout(executable, "process-test-sources")
}

// GenerateTestResources 生成测试资源文件（mvn generate-test-resources）
func GenerateTestResources(executable string) (string, error) {
	return ExecForStdout(executable, "generate-test-resources")
}

// ProcessTestResources 处理测试资源文件（mvn process-test-resources）
func ProcessTestResources(executable string) (string, error) {
	return ExecForStdout(executable, "process-test-resources")
}

// ProcessTestClasses 处理编译后的测试字节码（mvn process-test-classes）
func ProcessTestClasses(executable string) (string, error) {
	return ExecForStdout(executable, "process-test-classes")
}

// PreparePackage 打包前的准备工作（mvn prepare-package），常用于 CI 流水线
func PreparePackage(executable string) (string, error) {
	return ExecForStdout(executable, "prepare-package")
}

// PreIntegrationTest 集成测试前的准备工作（mvn pre-integration-test），如启动服务器
func PreIntegrationTest(executable string) (string, error) {
	return ExecForStdout(executable, "pre-integration-test")
}

// IntegrationTest 运行集成测试（mvn integration-test），通常配合 failsafe 插件使用
func IntegrationTest(executable string) (string, error) {
	return ExecForStdout(executable, "integration-test")
}

// PostIntegrationTest 集成测试后的清理工作（mvn post-integration-test），如关闭服务器
func PostIntegrationTest(executable string) (string, error) {
	return ExecForStdout(executable, "post-integration-test")
}

// StandaloneInstall 安装到本地仓库（mvn install），不先执行 clean
// 注意：lifecycle.go 中的 Install() 执行的是 "clean install"
func StandaloneInstall(executable string) (string, error) {
	return ExecForStdout(executable, "install")
}

// --- Clean Lifecycle（剩余阶段）---

// PreClean 清理前的准备工作（mvn pre-clean）
func PreClean(executable string) (string, error) {
	return ExecForStdout(executable, "pre-clean")
}

// PostClean 清理后的善后工作（mvn post-clean）
func PostClean(executable string) (string, error) {
	return ExecForStdout(executable, "post-clean")
}

// --- Site Lifecycle（剩余阶段）---

// PreSite 站点生成前的准备工作（mvn pre-site）
func PreSite(executable string) (string, error) {
	return ExecForStdout(executable, "pre-site")
}

// PostSite 站点生成后的善后工作（mvn post-site）
func PostSite(executable string) (string, error) {
	return ExecForStdout(executable, "post-site")
}

// SiteDeploy 部署生成的站点到 Web 服务器（mvn site-deploy）
func SiteDeploy(executable string) (string, error) {
	return ExecForStdout(executable, "site-deploy")
}
