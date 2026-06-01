package command

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandBuilder_BuildArgs(t *testing.T) {
	builder := NewCommandBuilder().
		WithProfiles("ci", "release").
		WithProperty("skipTests", "true").
		WithProjects("module-a", "module-b").
		WithAlsoMake().
		WithBatchMode().
		WithOffline().
		WithPomFile("/path/to/pom.xml").
		WithSettingsFile("/path/to/settings.xml").
		WithGoal("clean").
		WithGoal("install")

	args := builder.buildArgs()

	// 验证关键参数存在
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-P ci,release")
	assert.Contains(t, argsStr, "-DskipTests=true")
	assert.Contains(t, argsStr, "-pl module-a,module-b")
	assert.Contains(t, argsStr, "-am")
	assert.Contains(t, argsStr, "-B")
	assert.Contains(t, argsStr, "-o")
	assert.Contains(t, argsStr, "-f /path/to/pom.xml")
	assert.Contains(t, argsStr, "-s /path/to/settings.xml")
	assert.Contains(t, argsStr, "clean")
	assert.Contains(t, argsStr, "install")
}

func TestCommandBuilder_WithSkipTests(t *testing.T) {
	builder := NewCommandBuilder().WithSkipTests().WithGoal("install")
	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-DskipTests")
}

func TestCommandBuilder_WithSkipTestsCompletely(t *testing.T) {
	builder := NewCommandBuilder().WithSkipTestsCompletely().WithGoal("install")
	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-Dmaven.test.skip=true")
}

func TestCommandBuilder_WithUpdateSnapshots(t *testing.T) {
	builder := NewCommandBuilder().WithUpdateSnapshots().WithGoal("compile")
	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-U")
}

func TestCommandBuilder_WithThreads(t *testing.T) {
	builder := NewCommandBuilder().WithThreads(4).WithGoal("install")
	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-T 4")
}

func TestCommandBuilder_WithFailAtEnd(t *testing.T) {
	builder := NewCommandBuilder().WithFailAtEnd().WithGoal("install")
	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-fae")
}

func TestCommandBuilder_WithNoTransferProgress(t *testing.T) {
	builder := NewCommandBuilder().WithNoTransferProgress().WithGoal("install")
	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-ntp")
}

func TestCommandBuilder_WithDebug(t *testing.T) {
	builder := NewCommandBuilder().WithDebug().WithGoal("compile")
	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-X")
}

func TestCommandBuilder_WithQuiet(t *testing.T) {
	builder := NewCommandBuilder().WithQuiet().WithGoal("compile")
	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-q")
}

func TestCommandBuilder_WithNonRecursive(t *testing.T) {
	builder := NewCommandBuilder().WithNonRecursive().WithGoal("install")
	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-N")
}

func TestCommandBuilder_WithResumeFrom(t *testing.T) {
	builder := NewCommandBuilder().WithResumeFrom("module-b").WithGoal("install")
	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-rf module-b")
}

func TestCommandBuilder_WithGlobalSettings(t *testing.T) {
	builder := NewCommandBuilder().WithGlobalSettings("/global/settings.xml").WithGoal("install")
	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-gs /global/settings.xml")
}

func TestCommandBuilder_WithToolchains(t *testing.T) {
	builder := NewCommandBuilder().WithToolchains("/path/to/toolchains.xml").WithGoal("compile")
	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-t /path/to/toolchains.xml")
}

func TestCommandBuilder_WithShowVersion(t *testing.T) {
	builder := NewCommandBuilder().WithShowVersion().WithGoal("compile")
	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-V")
}

func TestCommandBuilder_WithStrictChecksums(t *testing.T) {
	builder := NewCommandBuilder().WithStrictChecksums().WithGoal("install")
	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-C")
}

func TestCommandBuilder_WithLaxChecksums(t *testing.T) {
	builder := NewCommandBuilder().WithLaxChecksums().WithGoal("install")
	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-c")
}

func TestCommandBuilder_WithAlsoMakeDependents(t *testing.T) {
	builder := NewCommandBuilder().WithAlsoMakeDependents().WithGoal("install")
	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "-amd")
}

func TestCommandBuilder_Build(t *testing.T) {
	builder := NewCommandBuilder().
		WithExecutable("mvn").
		WithWorkingDirectory("/tmp").
		WithBatchMode().
		WithGoals("clean", "install")

	opts := builder.Build()
	assert.Equal(t, "mvn", opts.Executable)
	assert.Equal(t, "/tmp", opts.WorkingDirectory)
	assert.Contains(t, opts.Args, "-B")
	assert.Contains(t, opts.Args, "clean")
	assert.Contains(t, opts.Args, "install")
}

func TestCommandBuilder_ChainedUsage(t *testing.T) {
	// 模拟典型的 CI 构建场景
	builder := NewCommandBuilder().
		WithExecutable("mvn").
		WithWorkingDirectory("/home/user/project").
		WithBatchMode().
		WithNoTransferProgress().
		WithProfiles("ci").
		WithSkipTests().
		WithUpdateSnapshots().
		WithGoals("clean", "deploy")

	args := builder.buildArgs()
	argsStr := strings.Join(args, " ")

	assert.Contains(t, argsStr, "-B")
	assert.Contains(t, argsStr, "-ntp")
	assert.Contains(t, argsStr, "-P ci")
	assert.Contains(t, argsStr, "-DskipTests")
	assert.Contains(t, argsStr, "-U")
	assert.Contains(t, argsStr, "clean")
	assert.Contains(t, argsStr, "deploy")
}
