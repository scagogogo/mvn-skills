# mvn-skills — Maven SDK for Go

[![CI](https://github.com/scagogogo/mvn-skills/actions/workflows/ci.yml/badge.svg)](https://github.com/scagogogo/mvn-skills/actions/workflows/ci.yml)
[![Release](https://github.com/scagogogo/mvn-skills/actions/workflows/release.yml/badge.svg)](https://github.com/scagogogo/mvn-skills/actions/workflows/release.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/mvn-skills.svg)](https://pkg.go.dev/github.com/scagogogo/mvn-skills)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/mvn-skills.svg)](https://goreportcard.com/report/github.com/scagogogo/mvn-skills)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**语言**: [English](README.md) | [中文](README.zh.md)

用于方便的在Go中操作Maven（`mvn`）的SDK。会检测使用本地已经安装的Maven来执行命令、解析POM文件、管理配置等。

## 特性

- 🔍 **Maven查找器** — 自动从PATH、M2_HOME或Maven Wrapper检测Maven
- 📦 **命令构建器** — 支持30+ Maven CLI选项的流式API
- 📄 **POM解析器** — 解析和分析pom.xml文件
- ⚙️ **配置解析器** — 解析Maven settings.xml
- 🗂️ **本地仓库** — 导航和搜索本地Maven仓库
- 📥 **Maven安装器** — 在Linux/macOS/Windows上下载和安装Maven
- 🏗️ **Context支持** — 通过context.Context取消和超时Maven命令
- 🖥️ **跨平台** — 完整支持Windows、macOS和Linux

## 安装

```bash
go get github.com/scagogogo/mvn-skills
```

### 指定版本

```bash
go get github.com/scagogogo/mvn-skills@v0.2.0
```

### 从Release下载

从[最新Release](https://github.com/scagogogo/mvn-skills/releases/latest)下载源代码归档：

```bash
# 下载并解压
curl -sL https://github.com/scagogogo/mvn-skills/releases/latest/download/mvn-skills-latest.tar.gz | tar -xz
cd mvn-skills-*/
go mod download
```

## 快速开始

### 查找Maven

```go
package main

import (
    "fmt"
    "github.com/scagogogo/mvn-skills/pkg/finder"
)

func main() {
    maven, err := finder.FindMaven()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Maven可执行文件: %s\n", maven)
}
```

### 执行Maven命令

**简单方式（独立函数）：**

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
    CleanInstall()
```

### Context与取消

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()

output, err := command.NewCommandBuilder().
    WithContext(ctx).
    CleanDeploy()
```

### 解析POM文件

```go
project, err := pom.ParseFile("pom.xml")
if err != nil {
    panic(err)
}
fmt.Printf("GAV: %s:%s:%s\n", project.GroupId, project.ArtifactId, project.Version)
```

### 解析Maven版本

```go
output, _ := command.Version("mvn")
v, err := command.ParseVersion(output)
if err != nil {
    panic(err)
}
fmt.Printf("Maven %d.%d.%d\n", v.Major, v.Minor, v.Patch)
if v.IsAtLeast(3, 8, 0) {
    fmt.Println("Maven 3.8+ 特性可用")
}
```

## API概览

### 命令构建器选项

| 选项 | 标志 | 描述 |
|------|------|------|
| `WithBatchMode()` | `-B` | 非交互模式（CI/CD） |
| `WithSkipTests()` | `-DskipTests` | 跳过测试执行 |
| `WithSkipTestsCompletely()` | `-Dmaven.test.skip=true` | 完全跳过测试 |
| `WithOffline()` | `-o` | 离线模式 |
| `WithUpdateSnapshots()` | `-U` | 强制更新SNAPSHOT |
| `WithProfiles(...)` | `-P` | 激活profile |
| `WithProperty(k, v)` | `-Dk=v` | 设置系统属性 |
| `WithPomFile(path)` | `-f` | 指定POM文件 |
| `WithSettingsFile(path)` | `-s` | 指定settings.xml |
| `WithProjects(...)` | `-pl` | 构建指定模块 |
| `WithAlsoMake()` | `-am` | 同时构建依赖模块 |
| `WithDebug()` | `-X` | 调试输出 |
| `WithQuiet()` | `-q` | 静默模式 |
| `WithThreads(n)` | `-T` | 并行线程数 |
| `WithFailAtEnd()` | `-fae` | 最后才失败 |
| `WithFailNever()` | `-fn` | 永不失败 |
| `WithNoTransferProgress()` | `-ntp` | 不显示下载进度 |
| `WithEnv(...)` | — | 设置环境变量 |
| `WithContext(ctx)` | — | 取消/超时支持 |

### 生命周期便捷方法

```go
builder.Clean()          // mvn clean
builder.Compile()        // mvn compile
builder.Test()           // mvn test
builder.Package()        // mvn package
builder.Verify()         // mvn verify
builder.Install()        // mvn install
builder.Deploy()         // mvn deploy
builder.Site()           // mvn site
builder.Validate()       // mvn validate
```

### 多阶段便捷方法

```go
builder.CleanInstall()   // mvn clean install
builder.CleanPackage()   // mvn clean package
builder.CleanDeploy()    // mvn clean deploy
builder.CleanVerify()    // mvn clean verify
builder.CleanTest()      // mvn clean test
```

### 结构化选项类型

对于有多个参数的命令，使用结构化选项类型：

```go
// 带选项的依赖获取
opts := &command.DependencyGetOption{
    GroupId:    "joda-time",
    ArtifactId: "joda-time",
    Version:    "2.10.10",
    Classifier: "sources",
}
output, err := command.DependencyGetWithOptions("mvn", opts)

// 带选项的文件部署
deployOpts := &command.DeployDeployFileOption{
    File:         "target/my-app.jar",
    PomFile:      "pom.xml",
    RepositoryId: "internal-repo",
    URL:          "https://repo.example.com/maven2",
}
output, err := command.DeployDeployFileWithOptions("mvn", deployOpts)

// 灵活打包的构件安装
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

## 包列表

| 包 | 描述 |
|----|------|
| `pkg/command` | Maven命令执行（构建器模式 + 独立函数） |
| `pkg/finder` | 在系统中查找Maven可执行文件 |
| `pkg/installer` | 下载和安装Maven |
| `pkg/pom` | 解析和分析pom.xml文件 |
| `pkg/settings` | 解析Maven settings.xml文件 |
| `pkg/local_repository` | 导航和搜索本地Maven仓库 |

## Release发布

Release通过[GoReleaser](https://goreleaser.com/)自动发布到[GitHub Releases](https://github.com/scagogogo/mvn-skills/releases)。

### 下载Release

每个Release包含：
- **源代码归档** — 完整源代码tarball
- **校验和** — SHA256校验和用于验证

```bash
# 下载最新源代码归档
curl -sL https://github.com/scagogogo/mvn-skills/releases/latest/download/mvn-skills-latest.tar.gz -o mvn-skills.tar.gz

# 验证校验和
curl -sL https://github.com/scagogogo/mvn-skills/releases/latest/download/checksums.txt -o checksums.txt
sha256sum -c checksums.txt --ignore-missing
```

### 作为Go模块使用

推荐使用Go模块方式：

```bash
# 最新版本
go get github.com/scagogogo/mvn-skills@latest

# 指定版本
go get github.com/scagogogo/mvn-skills@v0.2.0

# 指定commit
go get github.com/scagogogo/mvn-skills@abc1234
```

### 导入路径

```go
import (
    "github.com/scagogogo/mvn-skills/pkg/command"
    "github.com/scagogogo/mvn-skills/pkg/finder"
    "github.com/scagogogo/mvn-skills/pkg/installer"
    "github.com/scagogogo/mvn-skills/pkg/pom"
    "github.com/scagogogo/mvn-skills/pkg/settings"
    "github.com/scagogogo/mvn-skills/pkg/local_repository"
)
```

## 文档

完整API文档可在以下地址访问：
- 🌐 [https://scagogogo.github.io/mvn-skills/](https://scagogogo.github.io/mvn-skills/)
- 📦 [pkg.go.dev/github.com/scagogogo/mvn-skills](https://pkg.go.dev/github.com/scagogogo/mvn-skills)

## 贡献

1. Fork本仓库
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'feat: add amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建Pull Request

## 许可证

本项目基于MIT许可证 — 详见[LICENSE](LICENSE)文件。
