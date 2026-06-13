// Package command provides a comprehensive API for executing Maven commands.
//
// There are three ways to use this package:
//
// 1. Standalone functions (simple, one-off commands):
//
//	output, err := command.Clean("mvn")
//
// 2. Option structs (complex commands with many parameters):
//
//	opts := &command.DependencyGetOption{
//	    GroupId:    "joda-time",
//	    ArtifactId: "joda-time",
//	    Version:    "2.10.10",
//	}
//	output, err := command.DependencyGetWithOptions("mvn", opts)
//
// 3. CommandBuilder (recommended for complex builds):
//
//	output, err := command.NewCommandBuilder().
//	    WithExecutable("mvn").
//	    WithWorkingDirectory("/path/to/project").
//	    WithBatchMode().
//	    WithSkipTests().
//	    CleanInstall()
//
// The CommandBuilder provides a fluent, composable API with 30+ Maven CLI
// options. Convenience methods do NOT mutate the builder — they create copies,
// so the original builder remains reusable.
//
// Context and Cancellation:
//
// The builder supports context.Context for cancellation and timeouts:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
//	defer cancel()
//	output, err := command.NewCommandBuilder().
//	    WithContext(ctx).
//	    CleanInstall()
//
// Error Handling:
//
// When a Maven command fails, ExecForStdout returns a *MavenError containing
// the command, arguments, exit code, and stderr output for diagnosis:
//
//	output, err := command.Clean("mvn")
//	if err != nil {
//	    var me *command.MavenError
//	    if errors.As(err, &me) {
//	        log.Printf("Exit code: %d, stderr: %s", me.ExitCode, me.Stderr)
//	    }
//	}
//
// Version Parsing:
//
// The ParseVersion function can parse Maven version output into a structured type:
//
//	output, _ := command.Version("mvn")
//	v, err := command.ParseVersion(output)
//	if err == nil {
//	    fmt.Printf("Maven %d.%d.%d\n", v.Major, v.Minor, v.Patch)
//	    if v.IsAtLeast(3, 8, 0) { /* ... */ }
//	}
package command
