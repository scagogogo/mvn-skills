# API 参考

本文档提供了 Maven SDK Go 的详细 API 参考。

## 包概览

### Finder

Finder 包用于查找本机已安装的 Maven 可执行文件。

```go
package finder

// FindMaven 查找本机已安装的 Maven 可执行文件。
// 先检查系统 PATH，然后检查 M2_HOME 和 MAVEN_HOME 环境变量。
// 找不到时返回 ErrNotFoundMaven。
func FindMaven() (string, error)

// Check 检查给定目录是否包含有效的 Maven 安装
// 通过检查是否存在 bin/mvn 可执行文件来判断。
func Check(mavenHomeDirectory string) bool

// ErrNotFoundMaven 未找到 Maven 时返回的错误
var ErrNotFoundMaven error
```

### Command

Command 包提供了执行 Maven 命令的功能。

#### 通用执行

```go
package command

// Options 保存 Maven 命令执行的配置
type Options struct {
    Executable       string      // mvn 可执行文件路径（默认为 "mvn"）
    Args             []string    // 传递给 mvn 的参数
    WorkingDirectory string      // 命令的工作目录
    Stdin            io.Reader   // 标准输入
    Stdout           io.Writer   // 标准输出
    Stderr           io.Writer   // 标准错误
}

// Exec 使用给定的选项执行 Maven 命令
func Exec(options *Options) error

// ExecForStdout 执行 Maven 命令并返回其标准输出
func ExecForStdout(executable string, args ...string) (string, error)

// BuildExecutable 根据 MAVEN_HOME 目录构造 mvn 可执行文件路径
func BuildExecutable(mavenHomeDirectory string) string
```

#### 生命周期阶段

```go
// Clean 清理项目构建产物，删除 target 目录（mvn clean）
func Clean(executable string) (string, error)

// Compile 编译项目源码（mvn compile）
func Compile(executable string) (string, error)

// Test 运行项目单元测试（mvn test）
func Test(executable string) (string, error)

// TestCompile 编译项目测试源码（mvn test-compile）
func TestCompile(executable string) (string, error)

// Package 将编译后的代码打包（mvn package）
func Package(executable string) (string, error)

// Verify 运行集成测试和验证（mvn verify）
func Verify(executable string) (string, error)

// Deploy 将构建产物部署到远程仓库（mvn deploy）
func Deploy(executable string) (string, error)

// Site 生成项目站点文档（mvn site）
func Site(executable string) (string, error)

// Validate 验证项目结构（mvn validate）
func Validate(executable string) (string, error)

// Install 执行 mvn clean install
func Install(executable string) (string, error)

// InstallJar 安装 JAR 文件到本地仓库（mvn install:install-file）
func InstallJar(executable, jarPath, groupId, artifactId, version string) (string, error)
```

#### 依赖命令

```go
// DependencyGet 下载指定构件到本地仓库（mvn dependency:get）
func DependencyGet(executable, groupId, artifactId, version string) (string, error)

// DependencyTree 显示依赖树（mvn dependency:tree）
func DependencyTree(executable string) (string, error)

// DependencyResolve 解析所有依赖（mvn dependency:resolve）
func DependencyResolve(executable string) (string, error)

// DependencyAnalyze 分析依赖使用情况（mvn dependency:analyze）
func DependencyAnalyze(executable string) (string, error)

// DependencyList 列出所有已解析的依赖（mvn dependency:list）
func DependencyList(executable string) (string, error)

// DependencyPurgeLocalRepository 清理本地仓库（mvn dependency:purge-local-repository）
func DependencyPurgeLocalRepository(executable string) (string, error)
```

#### Help 命令

```go
// Version 获取 Maven 版本（mvn -v）
func Version(executable string) (string, error)

// GetLocalRepositoryDirectory 获取本地仓库路径（mvn help:evaluate）
func GetLocalRepositoryDirectory(executable string) (string, error)

// EffectivePom 显示有效的 POM 配置（mvn help:effective-pom）
func EffectivePom(executable string) (string, error)

// EffectiveSettings 显示有效的 Maven 设置（mvn help:effective-settings）
func EffectiveSettings(executable string) (string, error)

// ActiveProfiles 显示激活的 Maven profile（mvn help:active-profiles）
func ActiveProfiles(executable string) (string, error)

// DescribePlugin 描述插件的目标信息（mvn help:describe -Dplugin=...）
func DescribePlugin(executable, plugin string) (string, error)
```

#### Archetype 与 Wrapper

```go
// ArchetypeCreate 从原型生成新项目（mvn archetype:generate）
func ArchetypeCreate(executable, directory, groupId, artifactId, version string) (string, error)

// Wrapper 生成 Maven Wrapper 文件（mvn wrapper:wrapper）
func Wrapper(executable string) (string, error)
```

### Local Repository

Local Repository 包用于解析和管理 Maven 本地仓库。

```go
package local_repository

// DefaultLocalRepositoryDirectory 默认本地仓库路径（~/.m2/repository/）
var DefaultLocalRepositoryDirectory string

// ParseLocalRepositoryDirectory 从 Maven 解析本地仓库路径，
// 如果 Maven 不可用则返回默认路径。
func ParseLocalRepositoryDirectory(executable string) string

// BuildDirectory 构造 GAV 坐标的相对路径
func BuildDirectory(groupId, artifactId, version string) string

// FindDirectory 在本地仓库中查找 GAV 目录
func FindDirectory(localRepositoryDirectory, groupId, artifactId, version string) (string, error)

// FindJar 在本地仓库中根据 GAV 坐标查找 JAR 文件
func FindJar(localRepositoryDirectory, groupId, artifactId, version string) (string, error)

// FindJarWithClassifier 查找带分类器的 JAR 文件（如 "sources"、"javadoc"）
func FindJarWithClassifier(localRepositoryDirectory, groupId, artifactId, version, classifier string) (string, error)
```

### Installer

Installer 包提供了 Maven 的自动安装功能。

```go
package installer

// Install 下载并安装 Maven 到当前平台
// 返回 Maven 主目录路径
func Install() (string, error)

// InstallLinux 在 Linux 上安装 Maven（apt-get/yum 或二进制包回退）
func InstallLinux() (string, error)

// InstallMacOS 在 macOS 上安装 Maven（Homebrew 或二进制包回退）
func InstallMacOS() (string, error)

// InstallWindows 在 Windows 上安装 Maven（zip 下载 + 环境变量配置）
func InstallWindows() (string, error)

// InstallOptions macOS 安装的配置选项
type InstallOptions struct {
    MavenURL     string
    HomeDir      string
    SkipEnvSetup bool
}

// InstallMacOSWithOptions 使用可配置选项在 macOS 上安装 Maven
func InstallMacOSWithOptions(options InstallOptions) (string, error)
```

## 使用示例

### 查找 Maven

```go
maven, err := finder.FindMaven()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Maven 可执行文件路径: %s\n", maven)
```

### 执行 Maven 命令

```go
maven, _ := finder.FindMaven()

// 使用特定生命周期命令
output, err := command.Clean(maven)
output, err = command.Compile(maven)
output, err = command.Test(maven)
output, err = command.Package(maven)

// 带工作目录
err = command.Exec(&command.Options{
    Executable:       maven,
    Args:             []string{"clean", "install"},
    WorkingDirectory: "/path/to/project",
})
```

### 查找 JAR 文件

```go
maven, _ := finder.FindMaven()
repoDir := local_repository.ParseLocalRepositoryDirectory(maven)

// 查找主构件
jarPath, err := local_repository.FindJar(repoDir, "org.springframework", "spring-core", "5.3.21")

// 查找源码 JAR
sourcesPath, err := local_repository.FindJarWithClassifier(repoDir, "org.springframework", "spring-core", "5.3.21", "sources")

// 查找文档 JAR
javadocPath, err := local_repository.FindJarWithClassifier(repoDir, "org.springframework", "spring-core", "5.3.21", "javadoc")
```

### 获取本地仓库路径

```go
maven, _ := finder.FindMaven()
repoDir := local_repository.ParseLocalRepositoryDirectory(maven)
fmt.Printf("本地仓库路径: %s\n", repoDir)
```

### 依赖操作

```go
maven, _ := finder.FindMaven()

// 下载依赖
output, err := command.DependencyGet(maven, "com.alibaba", "fastjson", "2.0.2")

// 查看依赖树
tree, err := command.DependencyTree(maven)

// 分析依赖
analysis, err := command.DependencyAnalyze(maven)
```

### 安装 Maven

```go
mavenHome, err := installer.Install()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Maven 安装路径: %s\n", mavenHome)
```
