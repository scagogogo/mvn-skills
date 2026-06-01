// Package command provides a comprehensive API for executing Maven commands.
//
// There are two ways to use this package:
//
// 1. Standalone functions (simple, one-off commands):
//
//	output, err := command.Clean("mvn")
//
// 2. CommandBuilder (recommended for complex builds):
//
//	output, err := command.NewCommandBuilder().
//	    WithExecutable("mvn").
//	    WithWorkingDirectory("/path/to/project").
//	    WithBatchMode().
//	    WithSkipTests().
//	    Clean()
//
// The CommandBuilder provides a fluent, composable API with 27+ Maven CLI
// options. Convenience methods (Clean/Compile/Test/Package/Install/Deploy/Verify)
// do NOT mutate the builder — they create copies, so the original builder
// remains reusable.
//
// Error Handling:
//
// When a Maven command fails, ExecForStdout returns a *MavenError containing
// the command, arguments, and stderr output for diagnosis:
//
//	output, err := command.Clean("mvn")
//	if err != nil {
//	    var me *command.MavenError
//	    if errors.As(err, &me) {
//	        log.Printf("Maven stderr: %s", me.Stderr)
//	    }
//	}
package command
