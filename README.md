# mvn-skills

[![CI](https://github.com/scagogogo/mvn-skills/actions/workflows/ci.yml/badge.svg)](https://github.com/scagogogo/mvn-skills/actions/workflows/ci.yml)
[![Release](https://github.com/scagogogo/mvn-skills/actions/workflows/release.yml/badge.svg)](https://github.com/scagogogo/mvn-skills/actions/workflows/release.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/mvn-skills.svg)](https://pkg.go.dev/github.com/scagogogo/mvn-skills)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/mvn-skills)](https://goreportcard.com/report/github.com/scagogogo/mvn-skills)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Language**: [English](README.md) | [简体中文](README.zh.md)

Maven operations toolkit for AI agents and Go applications — execute builds, parse POM files, manage dependencies, install Maven, and automate Java project workflows.

---

## Install

### 🤖 AI Agents (Claude Code / OpenCode)

```bash
# 1. Add the marketplace
claude plugin marketplace add scagogogo/mvn-skills

# 2. Install the plugin
claude plugin install maven-skills@mvn-skills
```

Or use the community CLI ([skills](https://github.com/vercel-labs/skills)):

```bash
npx skills add scagogogo/mvn-skills@maven-operations
```

Done! Your AI agent can now run Maven commands, parse POM files, manage dependencies, and install Maven.

<details>
<summary>📋 Manual install (if marketplace is unavailable)</summary>

```bash
# Option A: Load for current session only (no install)
claude --plugin-dir /path/to/mvn-skills/plugins/maven-skills

# Option B: Copy to user skills directory (persistent)
mkdir -p ~/.claude/skills/maven-operations
git clone https://github.com/scagogogo/mvn-skills.git /tmp/mvn-skills-clone
cp -r /tmp/mvn-skills-clone/plugins/maven-skills/skills/maven-operations/* ~/.claude/skills/maven-operations/
rm -rf /tmp/mvn-skills-clone
```

</details>

### 📦 Go Applications

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
    // Find Maven
    mvn, _ := finder.FindMaven()

    // Run a build
    output, _ := command.NewCommandBuilder().
        WithExecutable(mvn).       // use the found Maven
        WithBatchMode().
        WithSkipTests().
        CleanInstall()

    // Parse a POM file
    project, _ := pom.ParseFile("pom.xml")
    fmt.Printf("%s:%s:%s\n", project.GroupId, project.ArtifactId, project.Version)
}
```

### 🖥️ CLI (Standalone Binary)

```bash
# Download the latest release
curl -sL https://github.com/scagogogo/mvn-skills/releases/latest/download/mvn-skills-latest.tar.gz | tar -xz
```

### 🔌 MCP Server

Wrap the Go SDK as an MCP server to provide Maven operations to any MCP-compatible AI tool. See the [documentation site](https://scagogogo.github.io/mvn-skills/) for setup details.

---

## Go SDK Examples

### Find Maven

```go
// Find system Maven
maven, err := finder.FindMaven()

// Find Maven Wrapper in a project
maven, err := finder.FindMavenWrapper("/path/to/project")

// Best available (wrapper preferred, system Maven fallback)
maven, err := finder.FindBestMaven("/path/to/project")
```

### Execute Maven Commands

**Simple one-liners:**

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
    CleanInstall()   // mvn clean install
```

### Parse POM Files

```go
project, err := pom.ParseFile("pom.xml")
fmt.Printf("GAV: %s:%s:%s\n", project.GroupId, project.ArtifactId, project.Version)
```

### Timeout & Cancellation

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()

output, err := command.NewCommandBuilder().
    WithContext(ctx).
    CleanDeploy()
```

### Check Maven Version

```go
output, _ := command.Version("mvn")
v, _ := command.ParseVersion(output)
if v.IsAtLeast(3, 8, 0) {
    fmt.Println("Maven 3.8+ features available")
}
```

---

## Features

| Feature | Description |
|---------|-------------|
| 🔍 **Maven Finder** | Auto-detect Maven from PATH, M2_HOME, or Maven Wrapper |
| 📦 **Command Builder** | Fluent API with 30+ Maven CLI options |
| 📄 **POM Parser** | Parse and analyze pom.xml files |
| ⚙️ **Settings Parser** | Parse Maven settings.xml |
| 🗂️ **Local Repository** | Navigate and search the local Maven repository |
| 📥 **Maven Installer** | Download and install Maven on Linux/macOS/Windows |
| 🏗️ **Context Support** | Cancel and timeout Maven commands via context.Context |
| 🖥️ **Cross-Platform** | Full Windows, macOS, and Linux support |

---

## Go API Reference

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

### Lifecycle Methods

```go
builder.Clean()          // mvn clean
builder.Compile()        // mvn compile
builder.Test()           // mvn test
builder.Package()        // mvn package
builder.Verify()         // mvn verify
builder.Install()        // mvn install
builder.Deploy()         // mvn deploy
```

### Multi-Phase Shortcuts

```go
builder.CleanInstall()   // mvn clean install  ← most common CI build
builder.CleanPackage()   // mvn clean package
builder.CleanDeploy()    // mvn clean deploy
builder.CleanVerify()    // mvn clean verify
builder.CleanTest()      // mvn clean test
```

<details>
<summary>📖 Structured Options & Error Handling</summary>

### Structured Options

For commands with many parameters:

```go
// Download an artifact
opts := &command.DependencyGetOption{
    GroupId:    "joda-time",
    ArtifactId: "joda-time",
    Version:    "2.10.10",
    Classifier: "sources",
}
output, err := command.DependencyGetWithOptions("mvn", opts)

// Deploy a file
deployOpts := &command.DeployDeployFileOption{
    File:         "target/my-app.jar",
    PomFile:      "pom.xml",
    RepositoryId: "internal-repo",
    URL:          "https://repo.example.com/maven2",
}
output, err := command.DeployDeployFileWithOptions("mvn", deployOpts)

// Install a local artifact
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

</details>

### Packages

| Package | Description |
|---------|-------------|
| `pkg/command` | Maven command execution (builder pattern + standalone functions) |
| `pkg/finder` | Find Maven executable on the system |
| `pkg/installer` | Download and install Maven |
| `pkg/pom` | Parse and analyze pom.xml files |
| `pkg/settings` | Parse Maven settings.xml files |
| `pkg/local_repository` | Navigate and search local Maven repository |

---

## Releases

Releases are automated via [GoReleaser](https://goreleaser.com/) and published to [GitHub Releases](https://github.com/scagogogo/mvn-skills/releases).

```bash
# Specific version
go get github.com/scagogogo/mvn-skills@v0.1.0

# Download and verify a release binary
curl -sL https://github.com/scagogogo/mvn-skills/releases/latest/download/mvn-skills-latest.tar.gz -o mvn-skills.tar.gz
curl -sL https://github.com/scagogogo/mvn-skills/releases/latest/download/checksums.txt -o checksums.txt
sha256sum -c checksums.txt --ignore-missing
```

---

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

MIT — see the [LICENSE](LICENSE) file for details.
