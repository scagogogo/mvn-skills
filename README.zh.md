# mvn-skills

[![CI](https://github.com/scagogogo/mvn-skills/actions/workflows/ci.yml/badge.svg)](https://github.com/scagogogo/mvn-skills/actions/workflows/ci.yml)
[![Release](https://github.com/scagogogo/mvn-skills/actions/workflows/release.yml/badge.svg)](https://github.com/scagogogo/mvn-skills/actions/workflows/release.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/mvn-skills.svg)](https://pkg.go.dev/github.com/scagogogo/mvn-skills)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/mvn-skills)](https://goreportcard.com/report/github.com/scagogogo/mvn-skills)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**语言**: [English](README.md) | 简体中文

Maven 操作工具包，面向 AI 智能体和 Go 应用 —— 执行构建、解析 POM 文件、管理依赖、安装 Maven、自动化 Java 项目工作流。

---

## 安装

### 🤖 AI 智能体（Claude Code / OpenCode）

```bash
# 1. 添加 marketplace
claude plugin marketplace add scagogogo/mvn-skills

# 2. 安装插件
claude plugin install maven-skills@mvn-skills
```

或使用社区 CLI（[skills](https://github.com/vercel-labs/skills)）：

```bash
npx skills add scagogogo/mvn-skills@maven-operations
```

搞定！AI 智能体现在可以直接执行 Maven 命令、解析 POM 文件、管理依赖、安装 Maven。

<details>
<summary>📋 手动安装（marketplace 不可用时）</summary>

```bash
# 方式 A：仅当前会话加载（无需安装）
claude --plugin-dir /path/to/mvn-skills/plugins/maven-skills

# 方式 B：复制到用户技能目录（持久化）
mkdir -p ~/.claude/skills/maven-operations
git clone https://github.com/scagogogo/mvn-skills.git /tmp/mvn-skills-clone
cp -r /tmp/mvn-skills-clone/plugins/maven-skills/skills/maven-operations/* ~/.claude/skills/maven-operations/
rm -rf /tmp/mvn-skills-clone
```

</details>

### 📦 Go 应用

```bash
go get github.com/scagogogo/mvn-skills@latest
```

```go
package main

import (
    "fmt"
    "github.com/scagogogo/mvn-skills/pkg/command"
    "github.com/scagogogo/mvn-skills/pkg/finder"
    "github.com/scagogogo/mvn-skills/pkg/pom"
)

func main() {
    // 查找 Maven
    mvn, _ := finder.FindMaven()

    // 执行构建
    output, _ := command.NewCommandBuilder().
        WithExecutable(mvn).       // 使用找到的 Maven
        WithBatchMode().
        WithSkipTests().
        CleanInstall()

    // 解析 POM 文件
    project, _ := pom.ParseFile("pom.xml")
    fmt.Printf("%s:%s:%s\n", project.GroupId, project.ArtifactId, project.Version)
}
```

### 🖥️ CLI（独立二进制）

```bash
# 下载最新 Release
curl -sL https://github.com/scagogogo/mvn-skills/releases/latest/download/mvn-skills-latest.tar.gz | tar -xz
```

### 🔌 MCP 服务器

将 Go SDK 封装为 MCP 服务器，为任何兼容 MCP 的 AI 工具提供 Maven 操作。详见[文档站](https://scagogogo.github.io/mvn-skills/)。

---

## Go SDK 示例

### 查找 Maven

```go
// 查找系统 Maven
maven, err := finder.FindMaven()

// 在项目中查找 Maven Wrapper
maven, err := finder.FindMavenWrapper("/path/to/project")

// 查找最佳可用（优先 Wrapper，回退系统 Maven）
maven, err := finder.FindBestMaven("/path/to/project")
```

### 执行 Maven 命令

**简单调用：**

```go
output, err := command.Clean("mvn")
output, err := command.Version("mvn")
output, err := command.DependencyGet("mvn", "joda-time", "joda-time", "2.10.10")
```

**构建器模式（推荐用于复杂构建）：**

```go
output, err := command.NewCommandBuilder().
    WithWorkingDirectory("/path/to/project").
    WithBatchMode().
    WithSkipTests().
    CleanInstall()   // mvn clean install
```

### 解析 POM 文件

```go
project, err := pom.ParseFile("pom.xml")
fmt.Printf("GAV: %s:%s:%s\n", project.GroupId, project.ArtifactId, project.Version)
```

### 超时与取消

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()

output, err := command.NewCommandBuilder().
    WithContext(ctx).
    CleanDeploy()
```

### 检查 Maven 版本

```go
output, _ := command.Version("mvn")
v, _ := command.ParseVersion(output)
if v.IsAtLeast(3, 8, 0) {
    fmt.Println("Maven 3.8+ 特性可用")
}
```

---

## 特性

| 特性 | 说明 |
|------|------|
| 🔍 **Maven 查找器** | 自动从 PATH、M2_HOME 或 Maven Wrapper 检测 Maven |
| 📦 **命令构建器** | 支持 30+ Maven CLI 选项的流式 API |
| 📄 **POM 解析器** | 解析和分析 pom.xml 文件 |
| ⚙️ **配置解析器** | 解析 Maven settings.xml |
| 🗂️ **本地仓库** | 导航和搜索本地 Maven 仓库 |
| 📥 **Maven 安装器** | 在 Linux/macOS/Windows 上下载和安装 Maven |
| 🏗️ **Context 支持** | 通过 context.Context 取消和超时 Maven 命令 |
| 🖥️ **跨平台** | 完整支持 Windows、macOS 和 Linux |

---

## Go API 参考

### 命令构建器选项

| 选项 | 标志 | 说明 |
|------|------|------|
| `WithBatchMode()` | `-B` | 非交互模式（CI/CD） |
| `WithSkipTests()` | `-DskipTests` | 跳过测试执行 |
| `WithSkipTestsCompletely()` | `-Dmaven.test.skip=true` | 完全跳过测试 |
| `WithOffline()` | `-o` | 离线模式 |
| `WithUpdateSnapshots()` | `-U` | 强制更新 SNAPSHOT |
| `WithProfiles(...)` | `-P` | 激活 profile |
| `WithProperty(k, v)` | `-Dk=v` | 设置系统属性 |
| `WithPomFile(path)` | `-f` | 指定 POM 文件 |
| `WithSettingsFile(path)` | `-s` | 指定 settings.xml |
| `WithProjects(...)` | `-pl` | 构建指定模块 |
| `WithAlsoMake()` | `-am` | 同时构建依赖模块 |
| `WithDebug()` | `-X` | 调试输出 |
| `WithQuiet()` | `-q` | 静默模式 |
| `WithThreads(n)` | `-T` | 并行线程数 |
| `WithFailAtEnd()` | `-fae` | 最后才失败 |
| `WithNoTransferProgress()` | `-ntp` | 不显示下载进度 |
| `WithEnv(...)` | — | 设置环境变量 |
| `WithContext(ctx)` | — | 取消/超时支持 |

### 生命周期方法

```go
builder.Clean()          // mvn clean
builder.Compile()        // mvn compile
builder.Test()           // mvn test
builder.Package()        // mvn package
builder.Verify()         // mvn verify
builder.Install()        // mvn install
builder.Deploy()         // mvn deploy
```

### 多阶段快捷方法

```go
builder.CleanInstall()   // mvn clean install  ← 最常见的 CI 构建
builder.CleanPackage()   // mvn clean package
builder.CleanDeploy()    // mvn clean deploy
builder.CleanVerify()    // mvn clean verify
builder.CleanTest()      // mvn clean test
```

<details>
<summary>📖 结构化选项与错误处理</summary>

### 结构化选项

用于有多个参数的命令：

```go
// 下载构件
opts := &command.DependencyGetOption{
    GroupId:    "joda-time",
    ArtifactId: "joda-time",
    Version:    "2.10.10",
    Classifier: "sources",
}
output, err := command.DependencyGetWithOptions("mvn", opts)

// 部署文件
deployOpts := &command.DeployDeployFileOption{
    File:         "target/my-app.jar",
    PomFile:      "pom.xml",
    RepositoryId: "internal-repo",
    URL:          "https://repo.example.com/maven2",
}
output, err := command.DeployDeployFileWithOptions("mvn", deployOpts)

// 安装本地构件
installOpts := &command.InstallFileOption{
    File:       "my-app.war",
    GroupId:    "com.example",
    ArtifactId: "my-app",
    Version:    "1.0.0",
    Packaging:  "war",
}
output, err := command.InstallFile("mvn", installOpts)
```

### 错误处理

```go
output, err := command.Clean("mvn")
if err != nil {
    var me *command.MavenError
    if errors.As(err, &me) {
        log.Printf("退出码: %d", me.ExitCode)
        log.Printf("命令: %s %s", me.Command, strings.Join(me.Args, " "))
        log.Printf("标准错误: %s", me.Stderr)
    }
}
```

</details>

### 包列表

| 包 | 说明 |
|----|------|
| `pkg/command` | Maven 命令执行（构建器模式 + 独立函数） |
| `pkg/finder` | 在系统中查找 Maven 可执行文件 |
| `pkg/installer` | 下载和安装 Maven |
| `pkg/pom` | 解析和分析 pom.xml 文件 |
| `pkg/settings` | 解析 Maven settings.xml 文件 |
| `pkg/local_repository` | 导航和搜索本地 Maven 仓库 |

---

## Release 发布

Release 通过 [GoReleaser](https://goreleaser.com/) 自动发布到 [GitHub Releases](https://github.com/scagogogo/mvn-skills/releases)。

```bash
# 指定版本
go get github.com/scagogogo/mvn-skills@v0.1.0

# 下载并验证 Release 二进制
curl -sL https://github.com/scagogogo/mvn-skills/releases/latest/download/mvn-skills-latest.tar.gz -o mvn-skills.tar.gz
curl -sL https://github.com/scagogogo/mvn-skills/releases/latest/download/checksums.txt -o checksums.txt
sha256sum -c checksums.txt --ignore-missing
```

---

## 文档

- 🌐 [VitePress 文档](https://scagogogo.github.io/mvn-skills/)
- 📦 [Go 包参考](https://pkg.go.dev/github.com/scagogogo/mvn-skills)

## 贡献

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'feat: add amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 许可证

MIT — 详见 [LICENSE](LICENSE) 文件。
