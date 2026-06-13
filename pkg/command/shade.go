package command

// Shade plugin-related commands
// The shade plugin is used to create an uber-jar / fat JAR (bundling all dependencies into a single JAR)

// ShadeShade creates an uber-jar (mvn shade:shade)
func ShadeShade(executable string) (string, error) {
	return ExecForStdout(executable, "shade:shade")
}
