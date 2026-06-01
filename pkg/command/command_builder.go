package command

import (
	"fmt"
	"io"
	"strings"
)

// CommandBuilder 使用 builder 模式构建和执行 Maven 命令
// 提供流式 API 来设置命令选项，比直接拼接字符串更安全、更易用
type CommandBuilder struct {
	executable       string
	workingDirectory string
	pomFile          string
	settingsFile     string
	globalSettings   string
	profiles         []string
	properties       map[string]string
	projects         []string
	alsoMake         bool
	alsoMakeDependents bool
	offline          bool
	batchMode        bool
	updateSnapshots  bool
	skipTests        bool
	mavenTestSkip    bool
	showErrors       bool
	debug            bool
	quiet            bool
	threads          int
	nonRecursive     bool
	resumeFrom       string
	failAtEnd        bool
	failNever        bool
	failFast         bool
	noTransferProgress bool
	strictChecksums  bool
	laxChecksums     bool
	showVersion      bool
	toolchains       string
	stdin            io.Reader
	stdout           io.Writer
	stderr           io.Writer
	goals            []string
}

// NewCommandBuilder 创建一个新的 Maven 命令构建器
func NewCommandBuilder() *CommandBuilder {
	return &CommandBuilder{
		executable: "mvn",
		properties: make(map[string]string),
	}
}

// WithExecutable 设置 Maven 可执行文件路径
func (b *CommandBuilder) WithExecutable(executable string) *CommandBuilder {
	b.executable = executable
	return b
}

// WithWorkingDirectory 设置命令的工作目录
func (b *CommandBuilder) WithWorkingDirectory(dir string) *CommandBuilder {
	b.workingDirectory = dir
	return b
}

// WithPomFile 指定要使用的 POM 文件路径（-f / --file）
func (b *CommandBuilder) WithPomFile(pomPath string) *CommandBuilder {
	b.pomFile = pomPath
	return b
}

// WithSettingsFile 指定用户 settings.xml 路径（-s / --settings）
func (b *CommandBuilder) WithSettingsFile(settingsPath string) *CommandBuilder {
	b.settingsFile = settingsPath
	return b
}

// WithGlobalSettings 指定全局 settings.xml 路径（-gs / --global-settings）
func (b *CommandBuilder) WithGlobalSettings(path string) *CommandBuilder {
	b.globalSettings = path
	return b
}

// WithToolchains 指定 toolchains 文件路径（-t / --toolchains）
func (b *CommandBuilder) WithToolchains(path string) *CommandBuilder {
	b.toolchains = path
	return b
}

// WithProfiles 激活指定的 Maven profile（-P）
func (b *CommandBuilder) WithProfiles(profiles ...string) *CommandBuilder {
	b.profiles = append(b.profiles, profiles...)
	return b
}

// WithProperty 设置系统属性（-Dkey=value）
func (b *CommandBuilder) WithProperty(key, value string) *CommandBuilder {
	b.properties[key] = value
	return b
}

// WithProperties 批量设置系统属性
func (b *CommandBuilder) WithProperties(props map[string]string) *CommandBuilder {
	for k, v := range props {
		b.properties[k] = v
	}
	return b
}

// WithProjects 指定要构建的模块（-pl / --projects）
func (b *CommandBuilder) WithProjects(modules ...string) *CommandBuilder {
	b.projects = append(b.projects, modules...)
	return b
}

// WithAlsoMake 同时构建被选中模块所依赖的模块（-am / --also-make）
func (b *CommandBuilder) WithAlsoMake() *CommandBuilder {
	b.alsoMake = true
	return b
}

// WithAlsoMakeDependents 同时构建依赖于选中模块的模块（-amd / --also-make-dependents）
func (b *CommandBuilder) WithAlsoMakeDependents() *CommandBuilder {
	b.alsoMakeDependents = true
	return b
}

// WithOffline 启用离线模式（-o / --offline）
func (b *CommandBuilder) WithOffline() *CommandBuilder {
	b.offline = true
	return b
}

// WithBatchMode 启用批处理/非交互模式（-B / --batch-mode），CI/CD 环境必需
func (b *CommandBuilder) WithBatchMode() *CommandBuilder {
	b.batchMode = true
	return b
}

// WithUpdateSnapshots 强制更新 SNAPSHOT 依赖（-U / --update-snapshots）
func (b *CommandBuilder) WithUpdateSnapshots() *CommandBuilder {
	b.updateSnapshots = true
	return b
}

// WithSkipTests 跳过测试执行但仍然编译测试代码（-DskipTests）
func (b *CommandBuilder) WithSkipTests() *CommandBuilder {
	b.skipTests = true
	return b
}

// WithSkipTestsCompletely 完全跳过测试（不编译也不执行）（-Dmaven.test.skip=true）
func (b *CommandBuilder) WithSkipTestsCompletely() *CommandBuilder {
	b.mavenTestSkip = true
	return b
}

// WithErrors 显示完整的错误堆栈信息（-e / --errors）
func (b *CommandBuilder) WithErrors() *CommandBuilder {
	b.showErrors = true
	return b
}

// WithDebug 启用调试输出（-X / --debug）
func (b *CommandBuilder) WithDebug() *CommandBuilder {
	b.debug = true
	return b
}

// WithQuiet 启用安静模式，只输出错误（-q / --quiet）
func (b *CommandBuilder) WithQuiet() *CommandBuilder {
	b.quiet = true
	return b
}

// WithThreads 设置并行构建线程数（-T / --threads）
func (b *CommandBuilder) WithThreads(n int) *CommandBuilder {
	b.threads = n
	return b
}

// WithThreadSpec 设置并行构建线程规格，如 "2C" 表示每核心 2 线程
func (b *CommandBuilder) WithThreadSpec(spec string) *CommandBuilder {
	b.properties["maven.threads"] = spec
	return b
}

// WithNonRecursive 不递归构建子模块（-N / --non-recursive）
func (b *CommandBuilder) WithNonRecursive() *CommandBuilder {
	b.nonRecursive = true
	return b
}

// WithResumeFrom 从指定模块恢复构建（-rf / --resume-from）
func (b *CommandBuilder) WithResumeFrom(module string) *CommandBuilder {
	b.resumeFrom = module
	return b
}

// WithFailAtEnd 构建失败时继续其他模块，最后再失败（-fae / --fail-at-end）
func (b *CommandBuilder) WithFailAtEnd() *CommandBuilder {
	b.failAtEnd = true
	return b
}

// WithFailNever 永不因构建失败而停止（-fn / --fail-never）
func (b *CommandBuilder) WithFailNever() *CommandBuilder {
	b.failNever = true
	return b
}

// WithFailFast 遇到第一个失败就停止（-ff / --fail-fast），这是默认行为
func (b *CommandBuilder) WithFailFast() *CommandBuilder {
	b.failFast = true
	return b
}

// WithNoTransferProgress 不显示下载/上传进度（-ntp / --no-transfer-progress），CI 日志更干净
func (b *CommandBuilder) WithNoTransferProgress() *CommandBuilder {
	b.noTransferProgress = true
	return b
}

// WithStrictChecksums 校验和不匹配时构建失败（-C / --strict-checksums）
func (b *CommandBuilder) WithStrictChecksums() *CommandBuilder {
	b.strictChecksums = true
	return b
}

// WithLaxChecksums 校验和不匹配时发出警告（-c / --lax-checksums）
func (b *CommandBuilder) WithLaxChecksums() *CommandBuilder {
	b.laxChecksums = true
	return b
}

// WithShowVersion 显示版本信息但不停止构建（-V / --show-version）
func (b *CommandBuilder) WithShowVersion() *CommandBuilder {
	b.showVersion = true
	return b
}

// WithStdin 设置标准输入
func (b *CommandBuilder) WithStdin(r io.Reader) *CommandBuilder {
	b.stdin = r
	return b
}

// WithStdout 设置标准输出
func (b *CommandBuilder) WithStdout(w io.Writer) *CommandBuilder {
	b.stdout = w
	return b
}

// WithStderr 设置标准错误
func (b *CommandBuilder) WithStderr(w io.Writer) *CommandBuilder {
	b.stderr = w
	return b
}

// WithGoal 添加要执行的 Maven 目标/阶段
func (b *CommandBuilder) WithGoal(goal string) *CommandBuilder {
	b.goals = append(b.goals, goal)
	return b
}

// WithGoals 批量添加要执行的 Maven 目标/阶段
func (b *CommandBuilder) WithGoals(goals ...string) *CommandBuilder {
	b.goals = append(b.goals, goals...)
	return b
}

// buildArgs 将所有构建器选项转换为 Maven 命令行参数
func (b *CommandBuilder) buildArgs() []string {
	var args []string

	// POM 文件
	if b.pomFile != "" {
		args = append(args, "-f", b.pomFile)
	}

	// Settings 文件
	if b.settingsFile != "" {
		args = append(args, "-s", b.settingsFile)
	}

	// 全局 Settings
	if b.globalSettings != "" {
		args = append(args, "-gs", b.globalSettings)
	}

	// Toolchains
	if b.toolchains != "" {
		args = append(args, "-t", b.toolchains)
	}

	// Profile
	if len(b.profiles) > 0 {
		args = append(args, "-P", strings.Join(b.profiles, ","))
	}

	// 项目/模块
	if len(b.projects) > 0 {
		args = append(args, "-pl", strings.Join(b.projects, ","))
	}

	// Also make
	if b.alsoMake {
		args = append(args, "-am")
	}

	// Also make dependents
	if b.alsoMakeDependents {
		args = append(args, "-amd")
	}

	// Offline
	if b.offline {
		args = append(args, "-o")
	}

	// Batch mode
	if b.batchMode {
		args = append(args, "-B")
	}

	// Update snapshots
	if b.updateSnapshots {
		args = append(args, "-U")
	}

	// Skip tests
	if b.skipTests {
		args = append(args, "-DskipTests")
	}

	// Maven test skip
	if b.mavenTestSkip {
		args = append(args, "-Dmaven.test.skip=true")
	}

	// Show errors
	if b.showErrors {
		args = append(args, "-e")
	}

	// Debug
	if b.debug {
		args = append(args, "-X")
	}

	// Quiet
	if b.quiet {
		args = append(args, "-q")
	}

	// Threads
	if b.threads > 0 {
		args = append(args, "-T", fmt.Sprintf("%d", b.threads))
	}

	// Non-recursive
	if b.nonRecursive {
		args = append(args, "-N")
	}

	// Resume from
	if b.resumeFrom != "" {
		args = append(args, "-rf", b.resumeFrom)
	}

	// Fail at end
	if b.failAtEnd {
		args = append(args, "-fae")
	}

	// Fail never
	if b.failNever {
		args = append(args, "-fn")
	}

	// Fail fast
	if b.failFast {
		args = append(args, "-ff")
	}

	// No transfer progress
	if b.noTransferProgress {
		args = append(args, "-ntp")
	}

	// Strict checksums
	if b.strictChecksums {
		args = append(args, "-C")
	}

	// Lax checksums
	if b.laxChecksums {
		args = append(args, "-c")
	}

	// Show version
	if b.showVersion {
		args = append(args, "-V")
	}

	// 系统属性
	for key, value := range b.properties {
		args = append(args, fmt.Sprintf("-D%s=%s", key, value))
	}

	// Goals / phases
	args = append(args, b.goals...)

	return args
}

// Build 构建命令选项但不执行，返回 *Options 供后续使用
func (b *CommandBuilder) Build() *Options {
	return &Options{
		Executable:       b.executable,
		Args:             b.buildArgs(),
		WorkingDirectory: b.workingDirectory,
		Stdin:            b.stdin,
		Stdout:           b.stdout,
		Stderr:           b.stderr,
	}
}

// Run 执行构建好的 Maven 命令
func (b *CommandBuilder) Run() error {
	return Exec(b.Build())
}

// RunForStdout 执行构建好的 Maven 命令并返回标准输出
func (b *CommandBuilder) RunForStdout() (string, error) {
	builder := *b // 浅拷贝
	builder.stdout = nil
	builder.stderr = nil
	builder.stdin = nil

	args := builder.buildArgs()
	return ExecForStdout(builder.executable, args...)
}

// 便捷方法：使用 builder 执行常用生命周期阶段
// 这些方法不会修改原始 builder，而是创建副本并添加目标

// withGoal 创建一个副本并添加目标，不修改原始 builder
func (b *CommandBuilder) withGoal(goal string) *CommandBuilder {
	copy := *b // 浅拷贝
	copy.goals = append([]string{}, b.goals...) // 深拷贝 goals 列表
	copy.goals = append(copy.goals, goal)
	return &copy
}

// Clean 清理构建产物
func (b *CommandBuilder) Clean() (string, error) {
	return b.withGoal("clean").RunForStdout()
}

// Compile 编译源码
func (b *CommandBuilder) Compile() (string, error) {
	return b.withGoal("compile").RunForStdout()
}

// Test 运行测试
func (b *CommandBuilder) Test() (string, error) {
	return b.withGoal("test").RunForStdout()
}

// Package 打包
func (b *CommandBuilder) Package() (string, error) {
	return b.withGoal("package").RunForStdout()
}

// Install 安装到本地仓库（执行 mvn install，不带 clean）
// 注意：如果需要执行 clean install，使用 WithGoals("clean", "install").RunForStdout()
func (b *CommandBuilder) Install() (string, error) {
	return b.withGoal("install").RunForStdout()
}

// Deploy 部署到远程仓库
func (b *CommandBuilder) Deploy() (string, error) {
	return b.withGoal("deploy").RunForStdout()
}

// Verify 验证
func (b *CommandBuilder) Verify() (string, error) {
	return b.withGoal("verify").RunForStdout()
}
