---
name: maven-operations
version: 1.0.0
description: |
  Executes Maven commands, parses POM files and settings.xml, manages
  dependencies, installs Maven, and automates Java project builds.
  Use when the user asks to run Maven builds, parse pom.xml, find
  Maven artifacts, manage local repository, install Maven, or perform
  any mvn-related operation. Also use when working with Java/Maven
  projects that need build automation, dependency resolution, or
  POM analysis.
license: MIT
compatibility: claude-code opencode
allowed-tools:
  - Read
  - Write
  - Edit
  - Grep
  - Glob
  - Bash
  - AskUserQuestion
---

# Maven Operations Skill

This skill provides comprehensive Maven operations through the `mvn-skills` Go SDK, enabling Claude Code to execute Maven commands, parse project files, and automate Java project workflows.

## When to Use

- User asks to build, test, package, or deploy a Maven project
- User needs to parse or analyze a `pom.xml` file
- User wants to find Maven dependencies or artifacts
- User needs to install Maven on the system
- User asks about Maven settings or configuration
- User wants to navigate the local Maven repository
- User needs to run any `mvn` command with specific options

## When NOT to Use

- The task is purely about Gradle or other build tools (not Maven)
- The user is asking about Java code without Maven context
- The task only involves running a simple shell command unrelated to Maven

## Installation

This skill uses the `mvn-skills` Go SDK. Install it first:

```bash
go get github.com/scagogogo/mvn-skills@latest
```

## Core Capabilities

### 1. Find Maven

Locate the Maven executable on the system:

```go
maven, err := finder.FindMaven()
// Or find Maven Wrapper in a project directory:
maven, err := finder.FindMavenWrapper("/path/to/project")
// Or find the best available (wrapper preferred, system Maven fallback):
maven, err := finder.FindBestMaven("/path/to/project")
```

### 2. Execute Maven Commands

**Simple commands:**

```go
output, err := command.Clean("mvn")
output, err := command.Compile("mvn")
output, err := command.Test("mvn")
output, err := command.Package("mvn")
output, err := command.Install("mvn")
output, err := command.Deploy("mvn")
output, err := command.Version("mvn")
```

**Builder pattern for complex builds:**

```go
output, err := command.NewCommandBuilder().
    WithWorkingDirectory("/path/to/project").
    WithBatchMode().
    WithSkipTests().
    WithProfiles("ci", "release").
    WithProperty("maven.test.skip", "true").
    CleanInstall()
```

**With context for cancellation/timeout:**

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
defer cancel()
output, err := command.NewCommandBuilder().
    WithContext(ctx).
    CleanDeploy()
```

### 3. Parse POM Files

```go
project, err := pom.ParseFile("pom.xml")
// Access GAV coordinates
fmt.Printf("GAV: %s:%s:%s\n", project.GroupId, project.ArtifactId, project.Version)
// Check for parent POM
if project.HasParent() { ... }
// Get dependencies
deps := project.GetDependencies()
// Get modules (multi-module project)
modules := project.GetModules()
```

### 4. Parse Maven Settings

```go
settings, err := settings.ParseFile("~/.m2/settings.xml")
// Access mirrors, profiles, servers, localRepository
mirrors := settings.Mirrors
localRepo := settings.LocalRepository
```

### 5. Dependency Operations

```go
// Download an artifact
output, err := command.DependencyGet("mvn", "joda-time", "joda-time", "2.10.10")

// With structured options
opts := &command.DependencyGetOption{
    GroupId:    "com.example",
    ArtifactId: "my-lib",
    Version:    "1.0.0",
    Classifier: "sources",
}
output, err := command.DependencyGetWithOptions("mvn", opts)

// Print dependency tree
output, err := command.DependencyTree("mvn")

// Analyze dependency usage
output, err := command.DependencyAnalyze("mvn")
```

### 6. Local Repository Navigation

```go
// Find local repository directory
directory := local_repository.ParseLocalRepositoryDirectory("mvn")

// Find a JAR file
jar, err := local_repository.FindJar(directory, "com.alibaba", "fastjson", "2.0.2")

// Find JAR with classifier
jar, err := local_repository.FindJarWithClassifier(directory, "com.example", "lib", "1.0", "sources")
```

### 7. Install Maven

```go
mavenHome, err := installer.Install()
// Platform-specific:
mavenHome, err := installer.InstallMacOS()
mavenHome, err := installer.InstallLinux()
mavenHome, err := installer.InstallWindows()
```

## Command Builder Options Reference

| Method | Flag | Description |
|--------|------|-------------|
| `WithBatchMode()` | `-B` | Non-interactive mode (essential for CI/CD) |
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
| `WithThreads(n)` | `-T` | Parallel build threads |
| `WithFailAtEnd()` | `-fae` | Fail at end |
| `WithNoTransferProgress()` | `-ntp` | No download progress |
| `WithEnv(...)` | — | Set environment variables |
| `WithContext(ctx)` | — | Cancellation/timeout support |

## Multi-Phase Convenience Methods

```go
builder.CleanInstall()   // mvn clean install — most common CI build
builder.CleanPackage()   // mvn clean package
builder.CleanDeploy()    // mvn clean deploy
builder.CleanVerify()    // mvn clean verify
builder.CleanTest()      // mvn clean test
```

## Error Handling

All Maven errors return `*command.MavenError` with structured information:

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

## Version Parsing

```go
output, _ := command.Version("mvn")
v, err := command.ParseVersion(output)
if v.IsAtLeast(3, 8, 0) {
    fmt.Println("Maven 3.8+ features available")
}
```
