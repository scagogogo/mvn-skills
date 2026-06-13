package command

import (
	"context"
	"strings"
	"testing"
	"time"

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

	// Verify key arguments exist
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
	// Simulate a typical CI build scenario
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

func TestCommandBuilder_NoMutation(t *testing.T) {
	// Verify convenience methods do not modify the original builder
	builder := NewCommandBuilder().
		WithExecutable("mvn").
		WithBatchMode().
		WithGoal("existing-goal")

	// Record original goals
	originalArgs := builder.buildArgs()
	originalStr := strings.Join(originalArgs, " ")
	assert.Contains(t, originalStr, "existing-goal")
	assert.NotContains(t, originalStr, "clean")

	// Call Clean convenience method (will fail since no mvn, but we only care about no mutation)
	builder.Clean()

	// Verify original builder was not modified
	afterArgs := builder.buildArgs()
	afterStr := strings.Join(afterArgs, " ")
	assert.Equal(t, originalStr, afterStr, "builder should not be mutated by convenience methods")
	assert.NotContains(t, afterStr, "clean")
}

func TestCommandBuilder_ConsecutiveCalls(t *testing.T) {
	// Verify consecutive convenience method calls do not accumulate goals
	builder := NewCommandBuilder().WithExecutable("mvn").WithGoal("base")

	// First call
	args1 := builder.withGoal("clean").buildArgs()
	assert.Contains(t, strings.Join(args1, " "), "clean")
	assert.NotContains(t, strings.Join(args1, " "), "compile")

	// Second call
	args2 := builder.withGoal("compile").buildArgs()
	assert.Contains(t, strings.Join(args2, " "), "compile")
	assert.NotContains(t, strings.Join(args2, " "), "clean")

	// Original builder unchanged
	originalArgs := builder.buildArgs()
	originalStr := strings.Join(originalArgs, " ")
	assert.Contains(t, originalStr, "base")
	assert.NotContains(t, originalStr, "clean")
	assert.NotContains(t, originalStr, "compile")
}

// --- New feature tests ---

func TestCommandBuilder_WithEnv(t *testing.T) {
	builder := NewCommandBuilder().
		WithEnv("JAVA_HOME=/usr/lib/jvm/java-17", "MAVEN_OPTS=-Xmx2g").
		WithGoal("compile")

	opts := builder.Build()
	assert.Contains(t, opts.Env, "JAVA_HOME=/usr/lib/jvm/java-17")
	assert.Contains(t, opts.Env, "MAVEN_OPTS=-Xmx2g")
}

func TestCommandBuilder_WithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	builder := NewCommandBuilder().
		WithContext(ctx).
		WithGoal("compile")

	opts := builder.Build()
	assert.Equal(t, ctx, opts.Context)
}

func TestCommandBuilder_CleanInstall(t *testing.T) {
	// CleanInstall adds "clean" and "install" goals
	args := NewCommandBuilder().withGoals("clean", "install").buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "clean")
	assert.Contains(t, argsStr, "install")
}

func TestCommandBuilder_CleanPackage(t *testing.T) {
	args := NewCommandBuilder().WithGoals("clean", "package").buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "clean")
	assert.Contains(t, argsStr, "package")
}

func TestCommandBuilder_CleanDeploy(t *testing.T) {
	args := NewCommandBuilder().WithGoals("clean", "deploy").buildArgs()
	argsStr := strings.Join(args, " ")
	assert.Contains(t, argsStr, "clean")
	assert.Contains(t, argsStr, "deploy")
}

func TestCommandBuilder_MultiPhaseConvenienceMethods(t *testing.T) {
	// Verify multi-phase methods don't mutate original builder
	builder := NewCommandBuilder().WithBatchMode().WithGoal("base")

	// Build a copy with clean+install
	copy := builder.withGoals("clean", "install")
	copyArgs := strings.Join(copy.buildArgs(), " ")
	assert.Contains(t, copyArgs, "clean")
	assert.Contains(t, copyArgs, "install")

	// Original unchanged
	origArgs := strings.Join(builder.buildArgs(), " ")
	assert.NotContains(t, origArgs, "clean")
	assert.NotContains(t, origArgs, "install")
}

func TestCommandBuilder_FailModeMutualExclusion(t *testing.T) {
	// WithFailAtEnd should clear failNever and failFast
	builder := NewCommandBuilder().WithFailNever().WithFailAtEnd()
	args := strings.Join(builder.buildArgs(), " ")
	assert.Contains(t, args, "-fae")
	assert.NotContains(t, args, "-fn")
	assert.NotContains(t, args, "-ff")

	// WithFailNever should clear failAtEnd and failFast
	builder = NewCommandBuilder().WithFailAtEnd().WithFailNever()
	args = strings.Join(builder.buildArgs(), " ")
	assert.Contains(t, args, "-fn")
	assert.NotContains(t, args, "-fae")
	assert.NotContains(t, args, "-ff")

	// WithFailFast should clear failAtEnd and failNever
	builder = NewCommandBuilder().WithFailNever().WithFailFast()
	args = strings.Join(builder.buildArgs(), " ")
	assert.Contains(t, args, "-ff")
	assert.NotContains(t, args, "-fae")
	assert.NotContains(t, args, "-fn")
}

func TestCommandBuilder_ChecksumMutualExclusion(t *testing.T) {
	// WithStrictChecksums should clear laxChecksums
	builder := NewCommandBuilder().WithLaxChecksums().WithStrictChecksums()
	args := strings.Join(builder.buildArgs(), " ")
	assert.Contains(t, args, "-C")
	assert.NotContains(t, args, "-c")

	// WithLaxChecksums should clear strictChecksums
	builder = NewCommandBuilder().WithStrictChecksums().WithLaxChecksums()
	args = strings.Join(builder.buildArgs(), " ")
	assert.Contains(t, args, "-c")
	assert.NotContains(t, args, "-C")
}

func TestCommandBuilder_PropertiesSorted(t *testing.T) {
	// Properties should be sorted for deterministic output
	builder := NewCommandBuilder().
		WithProperty("zebra", "z").
		WithProperty("alpha", "a").
		WithProperty("middle", "m").
		WithGoal("compile")

	args := builder.buildArgs()

	// Find the property arguments
	var propArgs []string
	for _, arg := range args {
		if strings.HasPrefix(arg, "-D") && !strings.HasPrefix(arg, "-Dskip") && !strings.HasPrefix(arg, "-Dmaven.test") {
			propArgs = append(propArgs, arg)
		}
	}
	assert.Equal(t, []string{"-Dalpha=a", "-Dmiddle=m", "-Dzebra=z"}, propArgs)
}

func TestCommandBuilder_Validate(t *testing.T) {
	args := NewCommandBuilder().withGoal("validate").buildArgs()
	assert.Contains(t, strings.Join(args, " "), "validate")
}

func TestCommandBuilder_Site(t *testing.T) {
	args := NewCommandBuilder().withGoal("site").buildArgs()
	assert.Contains(t, strings.Join(args, " "), "site")
}

func TestCommandBuilder_DependencyTree(t *testing.T) {
	args := NewCommandBuilder().withGoal("dependency:tree").buildArgs()
	assert.Contains(t, strings.Join(args, " "), "dependency:tree")
}

func TestCommandBuilder_HelpEffectivePom(t *testing.T) {
	args := NewCommandBuilder().withGoal("help:effective-pom").buildArgs()
	assert.Contains(t, strings.Join(args, " "), "help:effective-pom")
}
