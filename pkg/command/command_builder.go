package command

import (
	"fmt"
	"io"
	"strings"
)

// CommandBuilder uses the builder pattern to construct and execute Maven commands
// Provides a fluent API to set command options, safer and easier to use than directly concatenating strings
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

// NewCommandBuilder creates a new Maven command builder
func NewCommandBuilder() *CommandBuilder {
	return &CommandBuilder{
		executable: "mvn",
		properties: make(map[string]string),
	}
}

// WithExecutable sets the Maven executable path
func (b *CommandBuilder) WithExecutable(executable string) *CommandBuilder {
	b.executable = executable
	return b
}

// WithWorkingDirectory sets the working directory for the command
func (b *CommandBuilder) WithWorkingDirectory(dir string) *CommandBuilder {
	b.workingDirectory = dir
	return b
}

// WithPomFile specifies the POM file path to use (-f / --file)
func (b *CommandBuilder) WithPomFile(pomPath string) *CommandBuilder {
	b.pomFile = pomPath
	return b
}

// WithSettingsFile specifies the user settings.xml path (-s / --settings)
func (b *CommandBuilder) WithSettingsFile(settingsPath string) *CommandBuilder {
	b.settingsFile = settingsPath
	return b
}

// WithGlobalSettings specifies the global settings.xml path (-gs / --global-settings)
func (b *CommandBuilder) WithGlobalSettings(path string) *CommandBuilder {
	b.globalSettings = path
	return b
}

// WithToolchains specifies the toolchains file path (-t / --toolchains)
func (b *CommandBuilder) WithToolchains(path string) *CommandBuilder {
	b.toolchains = path
	return b
}

// WithProfiles activates the specified Maven profiles (-P)
func (b *CommandBuilder) WithProfiles(profiles ...string) *CommandBuilder {
	b.profiles = append(b.profiles, profiles...)
	return b
}

// WithProperty sets a system property (-Dkey=value)
func (b *CommandBuilder) WithProperty(key, value string) *CommandBuilder {
	b.properties[key] = value
	return b
}

// WithProperties batch-sets system properties
func (b *CommandBuilder) WithProperties(props map[string]string) *CommandBuilder {
	for k, v := range props {
		b.properties[k] = v
	}
	return b
}

// WithProjects specifies the modules to build (-pl / --projects)
func (b *CommandBuilder) WithProjects(modules ...string) *CommandBuilder {
	b.projects = append(b.projects, modules...)
	return b
}

// WithAlsoMake also builds modules that the selected modules depend on (-am / --also-make)
func (b *CommandBuilder) WithAlsoMake() *CommandBuilder {
	b.alsoMake = true
	return b
}

// WithAlsoMakeDependents also builds modules that depend on the selected modules (-amd / --also-make-dependents)
func (b *CommandBuilder) WithAlsoMakeDependents() *CommandBuilder {
	b.alsoMakeDependents = true
	return b
}

// WithOffline enables offline mode (-o / --offline)
func (b *CommandBuilder) WithOffline() *CommandBuilder {
	b.offline = true
	return b
}

// WithBatchMode enables batch/non-interactive mode (-B / --batch-mode), required for CI/CD environments
func (b *CommandBuilder) WithBatchMode() *CommandBuilder {
	b.batchMode = true
	return b
}

// WithUpdateSnapshots forces updating of SNAPSHOT dependencies (-U / --update-snapshots)
func (b *CommandBuilder) WithUpdateSnapshots() *CommandBuilder {
	b.updateSnapshots = true
	return b
}

// WithSkipTests skips test execution but still compiles test code (-DskipTests)
func (b *CommandBuilder) WithSkipTests() *CommandBuilder {
	b.skipTests = true
	return b
}

// WithSkipTestsCompletely skips tests entirely (no compilation or execution) (-Dmaven.test.skip=true)
func (b *CommandBuilder) WithSkipTestsCompletely() *CommandBuilder {
	b.mavenTestSkip = true
	return b
}

// WithErrors displays full error stack traces (-e / --errors)
func (b *CommandBuilder) WithErrors() *CommandBuilder {
	b.showErrors = true
	return b
}

// WithDebug enables debug output (-X / --debug)
func (b *CommandBuilder) WithDebug() *CommandBuilder {
	b.debug = true
	return b
}

// WithQuiet enables quiet mode, only outputting errors (-q / --quiet)
func (b *CommandBuilder) WithQuiet() *CommandBuilder {
	b.quiet = true
	return b
}

// WithThreads sets the number of parallel build threads (-T / --threads)
func (b *CommandBuilder) WithThreads(n int) *CommandBuilder {
	b.threads = n
	return b
}

// WithThreadSpec sets the parallel build thread specification, e.g. "2C" means 2 threads per core
func (b *CommandBuilder) WithThreadSpec(spec string) *CommandBuilder {
	b.properties["maven.threads"] = spec
	return b
}

// WithNonRecursive does not recursively build submodules (-N / --non-recursive)
func (b *CommandBuilder) WithNonRecursive() *CommandBuilder {
	b.nonRecursive = true
	return b
}

// WithResumeFrom resumes the build from the specified module (-rf / --resume-from)
func (b *CommandBuilder) WithResumeFrom(module string) *CommandBuilder {
	b.resumeFrom = module
	return b
}

// WithFailAtEnd continues building other modules on failure, failing at the end (-fae / --fail-at-end)
func (b *CommandBuilder) WithFailAtEnd() *CommandBuilder {
	b.failAtEnd = true
	return b
}

// WithFailNever never stops the build due to failures (-fn / --fail-never)
func (b *CommandBuilder) WithFailNever() *CommandBuilder {
	b.failNever = true
	return b
}

// WithFailFast stops on the first failure (-ff / --fail-fast), this is the default behavior
func (b *CommandBuilder) WithFailFast() *CommandBuilder {
	b.failFast = true
	return b
}

// WithNoTransferProgress does not show download/upload progress (-ntp / --no-transfer-progress), cleaner CI logs
func (b *CommandBuilder) WithNoTransferProgress() *CommandBuilder {
	b.noTransferProgress = true
	return b
}

// WithStrictChecksums fails the build on checksum mismatch (-C / --strict-checksums)
func (b *CommandBuilder) WithStrictChecksums() *CommandBuilder {
	b.strictChecksums = true
	return b
}

// WithLaxChecksums issues a warning on checksum mismatch (-c / --lax-checksums)
func (b *CommandBuilder) WithLaxChecksums() *CommandBuilder {
	b.laxChecksums = true
	return b
}

// WithShowVersion displays version information without stopping the build (-V / --show-version)
func (b *CommandBuilder) WithShowVersion() *CommandBuilder {
	b.showVersion = true
	return b
}

// WithStdin sets the standard input
func (b *CommandBuilder) WithStdin(r io.Reader) *CommandBuilder {
	b.stdin = r
	return b
}

// WithStdout sets the standard output
func (b *CommandBuilder) WithStdout(w io.Writer) *CommandBuilder {
	b.stdout = w
	return b
}

// WithStderr sets the standard error
func (b *CommandBuilder) WithStderr(w io.Writer) *CommandBuilder {
	b.stderr = w
	return b
}

// WithGoal adds a Maven goal/phase to execute
func (b *CommandBuilder) WithGoal(goal string) *CommandBuilder {
	b.goals = append(b.goals, goal)
	return b
}

// WithGoals batch-adds Maven goals/phases to execute
func (b *CommandBuilder) WithGoals(goals ...string) *CommandBuilder {
	b.goals = append(b.goals, goals...)
	return b
}

// buildArgs converts all builder options to Maven command-line arguments
func (b *CommandBuilder) buildArgs() []string {
	var args []string

	// POM file
	if b.pomFile != "" {
		args = append(args, "-f", b.pomFile)
	}

	// Settings file
	if b.settingsFile != "" {
		args = append(args, "-s", b.settingsFile)
	}

	// Global settings
	if b.globalSettings != "" {
		args = append(args, "-gs", b.globalSettings)
	}

	// Toolchains
	if b.toolchains != "" {
		args = append(args, "-t", b.toolchains)
	}

	// Profiles
	if len(b.profiles) > 0 {
		args = append(args, "-P", strings.Join(b.profiles, ","))
	}

	// Projects/modules
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

	// System properties
	for key, value := range b.properties {
		args = append(args, fmt.Sprintf("-D%s=%s", key, value))
	}

	// Goals / phases
	args = append(args, b.goals...)

	return args
}

// Build constructs the command options without executing, returning *Options for later use
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

// Run executes the constructed Maven command
func (b *CommandBuilder) Run() error {
	return Exec(b.Build())
}

// RunForStdout executes the constructed Maven command and returns the standard output
func (b *CommandBuilder) RunForStdout() (string, error) {
	builder := *b // shallow copy
	builder.stdout = nil
	builder.stderr = nil
	builder.stdin = nil

	args := builder.buildArgs()
	return ExecForStdout(builder.executable, args...)
}

// Convenience methods: using the builder to execute common lifecycle phases
// These methods do not modify the original builder; instead they create copies and add goals

// withGoal creates a copy and adds a goal without modifying the original builder
func (b *CommandBuilder) withGoal(goal string) *CommandBuilder {
	copy := *b // shallow copy
	copy.goals = append([]string{}, b.goals...) // deep copy goals slice
	copy.goals = append(copy.goals, goal)
	return &copy
}

// Clean cleans up build artifacts
func (b *CommandBuilder) Clean() (string, error) {
	return b.withGoal("clean").RunForStdout()
}

// Compile compiles source code
func (b *CommandBuilder) Compile() (string, error) {
	return b.withGoal("compile").RunForStdout()
}

// Test runs tests
func (b *CommandBuilder) Test() (string, error) {
	return b.withGoal("test").RunForStdout()
}

// Package packages the project
func (b *CommandBuilder) Package() (string, error) {
	return b.withGoal("package").RunForStdout()
}

// Install installs to the local repository (executes mvn install, without clean)
// Note: If you need to run clean install, use WithGoals("clean", "install").RunForStdout()
func (b *CommandBuilder) Install() (string, error) {
	return b.withGoal("install").RunForStdout()
}

// Deploy deploys to a remote repository
func (b *CommandBuilder) Deploy() (string, error) {
	return b.withGoal("deploy").RunForStdout()
}

// Verify verifies the project
func (b *CommandBuilder) Verify() (string, error) {
	return b.withGoal("verify").RunForStdout()
}