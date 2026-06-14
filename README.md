# mvn-skills

[![CI](https://github.com/scagogogo/mvn-skills/actions/workflows/ci.yml/badge.svg)](https://github.com/scagogogo/mvn-skills/actions/workflows/ci.yml)
[![Release](https://github.com/scagogogo/mvn-skills/actions/workflows/release.yml/badge.svg)](https://github.com/scagogogo/mvn-skills/actions/workflows/release.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/mvn-skills.svg)](https://pkg.go.dev/github.com/scagogogo/mvn-skills)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/mvn-skills)](https://goreportcard.com/report/github.com/scagogogo/mvn-skills)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Language**: [English](README.md) | [简体中文](README.zh.md)

A comprehensive toolkit for operating Maven (`mvn`) — supporting multiple integration methods: **Skills**, Go SDK, CLI, and MCP.

## Integration Methods

### 🎯 Skills (Claude Code Plugin)

Add mvn-skills as a Claude Code skill with one click:

```bash
claude plugin marketplace add scagogogo/mvn-skills
```

Or use the community CLI:

```bash
npx skills add scagogogo/mvn-skills@maven-operations
```

Once installed, Claude Code can directly execute Maven commands, parse POM files, manage dependencies, install Maven, and automate Java project builds.

<details>
<summary>📋 One-click copy</summary>

```bash
# Add the marketplace
claude plugin marketplace add scagogogo/mvn-skills

# Then install the plugin
claude plugin install maven-skills@mvn-skills
```

Or manual installation:

```bash
mkdir -p ~/.claude/skills/maven-operations
git clone https://github.com/scagogogo/mvn-skills.git /tmp/mvn-skills-clone
cp -r /tmp/mvn-skills-clone/plugins/maven-skills/skills/maven-operations/* ~/.claude/skills/maven-operations/
rm -rf /tmp/mvn-skills-clone
```

</details>

### 📦 Go SDK

Use as a Go library in your applications:

```bash
go get github.com/scagogogo/mvn-skills@latest
```

```go
import (
    "github.com/scagogogo/mvn-skills/pkg/command"
    "github.com/scagogogo/mvn-skills/pkg/finder"
    "github.com/scagogogo/mvn-skills/pkg/pom"
    "github.com/scagogogo/mvn-skills/pkg/settings"
    "github.com/scagogogo/mvn-skills/pkg/installer"
    "github.com/scagogogo/mvn-skills/pkg/local_repository"
)
```

### 🖥️ CLI

Use directly from the command line via releases:

```bash
# Download latest release
curl -sL https://github.com/scagogogo/mvn-skills/releases/latest/download/mvn-skills-latest.tar.gz | tar -xz
```

### 🔌 MCP

Integrate with AI assistants via the Model Context Protocol. The Go SDK can be wrapped as an MCP server to provide Maven operations to any MCP-compatible AI tool.

## Features

- 🔍 **Maven Finder** — Auto-detect Maven from PATH, M2_HOME, or Maven Wrapper
- 📦 **Command Builder** — Fluent API with 30+ Maven CLI options
- 📄 **POM Parser** — Parse and analyze pom.xml files
- ⚙️ **Settings Parser** — Parse Maven settings.xml
- 🗂️ **Local Repository** — Navigate and search the local Maven repository
- 📥 **Maven Installer** — Download and install Maven on Linux/macOS/Windows
- 🏗️ **Context Support** — Cancel and timeout Maven commands via context.Context
- 🖥️ **Cross-Platform** — Full Windows, macOS, and Linux support

## Quick Start

### Find Maven

```go
maven, err := finder.FindMaven()
// Or find Maven Wrapper in a project directory:
maven, err := finder.FindMavenWrapper("/path/to/project")
// Or find the best available (wrapper preferred, system Maven fallback):
maven, err := finder.FindBestMaven("/path/to/project")
```

### Execute Maven Commands

**Simple (standalone functions):**

```go
output, err := command.Clean("mvn")
output, err := command.Version("mvn")
output, err := command.DependencyGet("mvn", "joda-time", "joda-time", "2.10.10")
```

**Builder pattern (recommended for complex builds):**

```go
output, err := command.NewCommandBuilder().
    WithWorkingDirectory("/path/to/project").
    WithBatchMode().
    WithSkipTests().
    CleanInstall()
```

### Context & Cancellation

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()

output, err := command.NewCommandBuilder().
    WithContext(ctx).
    CleanDeploy()
```

### Parse POM Files

```go
project, err := pom.ParseFile("pom.xml")
fmt.Printf("GAV: %s:%s:%s\n", project.GroupId, project.ArtifactId, project.Version)
```

### Parse Maven Version

```go
output, _ := command.Version("mvn")
v, err := command.ParseVersion(output)
if v.IsAtLeast(3, 8, 0) {
    fmt.Println("Maven 3.8+ features available")
}
```

## API Overview

### Command Builder Options

| Option | Flag | Description |
|--------|------|-------------|
| `WithBatchMode()` | `-B` | Non-interactive mode (CI/CD) |
| `WithSkipTests()` | `-DskipTests` | Skip test execution |
| `WithSkipTestsCompletely()` | `-Dmaven.test.skip=true` | Skip tests entirely |
| `WithOffline()` | `-o` | Offline mode |
| `WithUpdateSnapshots()` | `-U` | Force update SNAPSHOTs |
| `WithProfiles(...)` | `-P` | Activate profiles |
| `WithProperty(k, v)` | `-Dk=v` | Set system property |
| `WithPomFile(path)` | `-f` | Specify POM file |
| `WithSettingsFile(path)` | `-s` | Specify settings.xml |
| `WithProjects(...)` | `-pl` | Build specific modules |
| `WithAlsoMake()` | `-am` | Also build dependencies |
| `WithDebug()` | `-X` | Debug output |
| `WithQuiet()` | `-q` | Quiet mode |
| `WithThreads(n)` | `-T` | Parallel threads |
| `WithFailAtEnd()` | `-fae` | Fail at end |
| `WithNoTransferProgress()` | `-ntp` | No download progress |
| `WithEnv(...)` | — | Set environment variables |
| `WithContext(ctx)` | — | Cancellation/timeout support |

### Lifecycle Convenience Methods

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

### Multi-Phase Convenience Methods

```go
builder.CleanInstall()   // mvn clean install — most common CI build
builder.CleanPackage()   // mvn clean package
builder.CleanDeploy()    // mvn clean deploy
builder.CleanVerify()    // mvn clean verify
builder.CleanTest()      // mvn clean test
```

### Structured Option Types

For commands with many parameters:

```go
// Dependency get with options
opts := &command.DependencyGetOption{
    GroupId:    "joda-time",
    ArtifactId: "joda-time",
    Version:    "2.10.10",
    Classifier: "sources",
}
output, err := command.DependencyGetWithOptions("mvn", opts)

// Deploy file with options
deployOpts := &command.DeployDeployFileOption{
    File:         "target/my-app.jar",
    PomFile:      "pom.xml",
    RepositoryId: "internal-repo",
    URL:          "https://repo.example.com/maven2",
}
output, err := command.DeployDeployFileWithOptions("mvn", deployOpts)

// Install artifact with flexible packaging
installOpts := &command.InstallFileOption{
    File:       "my-app.war",
    GroupId:    "com.example",
    ArtifactId: "my-app",
    Version:    "1.0.0",
    Packaging:  "war",
}
output, err := command.InstallFile("mvn", installOpts)
```

### Error Handling

```go
output, err := command.Clean("mvn")
if err != nil {
    var me *command.MavenError
    if errors.As(err, &me) {
        log.Printf("Exit code: %d", me.ExitCode)
        log.Printf("Command: %s %s", me.Command, strings.Join(me.Args, " "))
        log.Printf("Stderr: %s", me.Stderr)
    }
}
```

## Packages

| Package | Description |
|---------|-------------|
| `pkg/command` | Maven command execution (builder pattern + standalone functions) |
| `pkg/finder` | Find Maven executable on the system |
| `pkg/installer` | Download and install Maven |
| `pkg/pom` | Parse and analyze pom.xml files |
| `pkg/settings` | Parse Maven settings.xml files |
| `pkg/local_repository` | Navigate and search local Maven repository |

## Releases

Releases are automated via [GoReleaser](https://goreleaser.com/) and published to [GitHub Releases](https://github.com/scagogogo/mvn-skills/releases).

### Use as Go Module

```bash
# Latest version
go get github.com/scagogogo/mvn-skills@latest

# Specific version
go get github.com/scagogogo/mvn-skills@v0.1.0
```

### Download a Release

Each release includes source archives and SHA256 checksums for verification:

```bash
# Download latest source archive
curl -sL https://github.com/scagogogo/mvn-skills/releases/latest/download/mvn-skills-latest.tar.gz -o mvn-skills.tar.gz

# Verify checksum
curl -sL https://github.com/scagogogo/mvn-skills/releases/latest/download/checksums.txt -o checksums.txt
sha256sum -c checksums.txt --ignore-missing
```

## Documentation

- 🌐 [VitePress Documentation](https://scagogogo.github.io/mvn-skills/)
- 📦 [Go Package Reference](https://pkg.go.dev/github.com/scagogogo/mvn-skills)

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.
