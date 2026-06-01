# API Reference

This document provides detailed API reference for Maven SDK Go.

## Package Overview

### Finder

The Finder package is used to locate a locally-installed Maven executable.

```go
package finder

// FindMaven searches for a locally installed Maven executable.
// It checks the system PATH first, then M2_HOME and MAVEN_HOME environment variables.
// Returns the executable path or ErrNotFoundMaven.
func FindMaven() (string, error)

// Check validates that the given directory contains a valid Maven installation
// by checking for the presence of bin/mvn executable.
func Check(mavenHomeDirectory string) bool

// ErrNotFoundMaven is returned when no Maven installation can be found.
var ErrNotFoundMaven error
```

### Command

The Command package provides functionality to execute Maven commands.

#### Generic Execution

```go
package command

// Options holds configuration for Maven command execution.
type Options struct {
    Executable       string      // Path to mvn executable (defaults to "mvn")
    Args             []string    // Arguments to pass to mvn
    WorkingDirectory string      // Working directory for the command
    Stdin            io.Reader   // Standard input
    Stdout           io.Writer   // Standard output
    Stderr           io.Writer   // Standard error
}

// Exec executes a Maven command with the given options.
func Exec(options *Options) error

// ExecForStdout executes a Maven command and returns its stdout as a string.
func ExecForStdout(executable string, args ...string) (string, error)

// BuildExecutable constructs the path to the mvn executable from a MAVEN_HOME directory.
func BuildExecutable(mavenHomeDirectory string) string
```

#### Lifecycle Phases

```go
// Clean removes the target directory (mvn clean).
func Clean(executable string) (string, error)

// Compile compiles the project source code (mvn compile).
func Compile(executable string) (string, error)

// Test runs the project unit tests (mvn test).
func Test(executable string) (string, error)

// TestCompile compiles the project test source code (mvn test-compile).
func TestCompile(executable string) (string, error)

// Package packages the compiled code (mvn package).
func Package(executable string) (string, error)

// Verify runs integration tests and checks (mvn verify).
func Verify(executable string) (string, error)

// Deploy deploys the built artifact to a remote repository (mvn deploy).
func Deploy(executable string) (string, error)

// Site generates project site documentation (mvn site).
func Site(executable string) (string, error)

// Validate validates the project structure (mvn validate).
func Validate(executable string) (string, error)

// Install runs mvn clean install.
func Install(executable string) (string, error)

// InstallJar installs a JAR file into the local repository (mvn install:install-file).
func InstallJar(executable, jarPath, groupId, artifactId, version string) (string, error)
```

#### Dependency Commands

```go
// DependencyGet downloads a specific artifact to the local repository (mvn dependency:get).
func DependencyGet(executable, groupId, artifactId, version string) (string, error)

// DependencyTree displays the dependency tree (mvn dependency:tree).
func DependencyTree(executable string) (string, error)

// DependencyResolve resolves all dependencies (mvn dependency:resolve).
func DependencyResolve(executable string) (string, error)

// DependencyAnalyze analyzes dependency usage (mvn dependency:analyze).
func DependencyAnalyze(executable string) (string, error)

// DependencyList lists all resolved dependencies (mvn dependency:list).
func DependencyList(executable string) (string, error)

// DependencyPurgeLocalRepository purges the local repository (mvn dependency:purge-local-repository).
func DependencyPurgeLocalRepository(executable string) (string, error)
```

#### Help Commands

```go
// Version returns the Maven version (mvn -v).
func Version(executable string) (string, error)

// GetLocalRepositoryDirectory returns the local repository path (mvn help:evaluate).
func GetLocalRepositoryDirectory(executable string) (string, error)

// EffectivePom displays the effective POM (mvn help:effective-pom).
func EffectivePom(executable string) (string, error)

// EffectiveSettings displays the effective settings (mvn help:effective-settings).
func EffectiveSettings(executable string) (string, error)

// ActiveProfiles displays active profiles (mvn help:active-profiles).
func ActiveProfiles(executable string) (string, error)

// DescribePlugin describes a plugin's goals (mvn help:describe -Dplugin=...).
func DescribePlugin(executable, plugin string) (string, error)
```

#### Archetype & Wrapper

```go
// ArchetypeCreate generates a new project from an archetype (mvn archetype:generate).
func ArchetypeCreate(executable, directory, groupId, artifactId, version string) (string, error)

// Wrapper generates Maven Wrapper files (mvn wrapper:wrapper).
func Wrapper(executable string) (string, error)
```

### Local Repository

The Local Repository package is used to parse and navigate the Maven local repository.

```go
package local_repository

// DefaultLocalRepositoryDirectory is the default path (~/.m2/repository/).
var DefaultLocalRepositoryDirectory string

// ParseLocalRepositoryDirectory resolves the local repository path from Maven,
// falling back to the default if Maven is not available.
func ParseLocalRepositoryDirectory(executable string) string

// BuildDirectory constructs the GAV-relative path within the repository.
func BuildDirectory(groupId, artifactId, version string) string

// FindDirectory finds the GAV directory in the local repository.
func FindDirectory(localRepositoryDirectory, groupId, artifactId, version string) (string, error)

// FindJar finds a JAR file by GAV coordinates in the local repository.
func FindJar(localRepositoryDirectory, groupId, artifactId, version string) (string, error)

// FindJarWithClassifier finds a JAR file with a classifier (e.g., "sources", "javadoc").
func FindJarWithClassifier(localRepositoryDirectory, groupId, artifactId, version, classifier string) (string, error)
```

### Installer

The Installer package provides automatic Maven installation functionality.

```go
package installer

// Install downloads and installs Maven for the current platform.
// Returns the Maven home directory path.
func Install() (string, error)

// InstallLinux installs Maven on Linux (apt-get/yum or binary fallback).
func InstallLinux() (string, error)

// InstallMacOS installs Maven on macOS (Homebrew or binary fallback).
func InstallMacOS() (string, error)

// InstallWindows installs Maven on Windows (zip download + env setup).
func InstallWindows() (string, error)

// InstallOptions holds configuration for macOS installation.
type InstallOptions struct {
    MavenURL    string
    HomeDir     string
    SkipEnvSetup bool
}

// InstallMacOSWithOptions installs Maven on macOS with configurable options.
func InstallMacOSWithOptions(options InstallOptions) (string, error)
```

## Usage Examples

### Finding Maven

```go
maven, err := finder.FindMaven()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Maven executable: %s\n", maven)
```

### Executing Maven Commands

```go
maven, _ := finder.FindMaven()

// Using specific lifecycle commands
output, err := command.Clean(maven)
output, err = command.Compile(maven)
output, err = command.Test(maven)
output, err = command.Package(maven)

// With working directory
err = command.Exec(&command.Options{
    Executable:       maven,
    Args:             []string{"clean", "install"},
    WorkingDirectory: "/path/to/project",
})
```

### Finding JAR Files

```go
maven, _ := finder.FindMaven()
repoDir := local_repository.ParseLocalRepositoryDirectory(maven)

// Find main artifact
jarPath, err := local_repository.FindJar(repoDir, "org.springframework", "spring-core", "5.3.21")

// Find sources JAR
sourcesPath, err := local_repository.FindJarWithClassifier(repoDir, "org.springframework", "spring-core", "5.3.21", "sources")

// Find javadoc JAR
javadocPath, err := local_repository.FindJarWithClassifier(repoDir, "org.springframework", "spring-core", "5.3.21", "javadoc")
```

### Getting Local Repository Path

```go
maven, _ := finder.FindMaven()
repoDir := local_repository.ParseLocalRepositoryDirectory(maven)
fmt.Printf("Local repository: %s\n", repoDir)
```

### Dependency Operations

```go
maven, _ := finder.FindMaven()

// Download a dependency
output, err := command.DependencyGet(maven, "com.alibaba", "fastjson", "2.0.2")

// View dependency tree
tree, err := command.DependencyTree(maven)

// Analyze dependencies
analysis, err := command.DependencyAnalyze(maven)
```

### Installing Maven

```go
mavenHome, err := installer.Install()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Maven installed at: %s\n", mavenHome)
```
