package command

// Extended lifecycle phase commands
// Maven has 3 built-in lifecycles (clean/default/site) with a total of 28 phases
// lifecycle.go implements the 9 most commonly used phases; this file adds the remaining phases

// --- Default Lifecycle (remaining phases) ---

// Initialize initializes the build state (mvn initialize), commonly used for early setup in multi-module builds
func Initialize(executable string) (string, error) {
	return ExecForStdout(executable, "initialize")
}

// GenerateSources generates source code (mvn generate-sources), commonly used for protobuf, XSD, etc. code generation
func GenerateSources(executable string) (string, error) {
	return ExecForStdout(executable, "generate-sources")
}

// ProcessSources processes source code (mvn process-sources)
func ProcessSources(executable string) (string, error) {
	return ExecForStdout(executable, "process-sources")
}

// GenerateResources generates resource files (mvn generate-resources)
func GenerateResources(executable string) (string, error) {
	return ExecForStdout(executable, "generate-resources")
}

// ProcessResources processes resource files (mvn process-resources), such as variable substitution
func ProcessResources(executable string) (string, error) {
	return ExecForStdout(executable, "process-resources")
}

// ProcessClasses processes compiled bytecode (mvn process-classes), such as bytecode enhancement
func ProcessClasses(executable string) (string, error) {
	return ExecForStdout(executable, "process-classes")
}

// GenerateTestSources generates test source code (mvn generate-test-sources)
func GenerateTestSources(executable string) (string, error) {
	return ExecForStdout(executable, "generate-test-sources")
}

// ProcessTestSources processes test source code (mvn process-test-sources)
func ProcessTestSources(executable string) (string, error) {
	return ExecForStdout(executable, "process-test-sources")
}

// GenerateTestResources generates test resource files (mvn generate-test-resources)
func GenerateTestResources(executable string) (string, error) {
	return ExecForStdout(executable, "generate-test-resources")
}

// ProcessTestResources processes test resource files (mvn process-test-resources)
func ProcessTestResources(executable string) (string, error) {
	return ExecForStdout(executable, "process-test-resources")
}

// ProcessTestClasses processes compiled test bytecode (mvn process-test-classes)
func ProcessTestClasses(executable string) (string, error) {
	return ExecForStdout(executable, "process-test-classes")
}

// PreparePackage performs preparation work before packaging (mvn prepare-package), commonly used in CI pipelines
func PreparePackage(executable string) (string, error) {
	return ExecForStdout(executable, "prepare-package")
}

// PreIntegrationTest performs preparation work before integration tests (mvn pre-integration-test), such as starting servers
func PreIntegrationTest(executable string) (string, error) {
	return ExecForStdout(executable, "pre-integration-test")
}

// IntegrationTest runs integration tests (mvn integration-test), typically used with the failsafe plugin
func IntegrationTest(executable string) (string, error) {
	return ExecForStdout(executable, "integration-test")
}

// PostIntegrationTest performs cleanup after integration tests (mvn post-integration-test), such as shutting down servers
func PostIntegrationTest(executable string) (string, error) {
	return ExecForStdout(executable, "post-integration-test")
}

// StandaloneInstall installs to the local repository (mvn install) without running clean first
// Note: The Install() in lifecycle.go executes "clean install"
func StandaloneInstall(executable string) (string, error) {
	return ExecForStdout(executable, "install")
}

// --- Clean Lifecycle (remaining phases) ---

// PreClean performs preparation work before cleaning (mvn pre-clean)
func PreClean(executable string) (string, error) {
	return ExecForStdout(executable, "pre-clean")
}

// PostClean performs cleanup work after cleaning (mvn post-clean)
func PostClean(executable string) (string, error) {
	return ExecForStdout(executable, "post-clean")
}

// --- Site Lifecycle (remaining phases) ---

// PreSite performs preparation work before site generation (mvn pre-site)
func PreSite(executable string) (string, error) {
	return ExecForStdout(executable, "pre-site")
}

// PostSite performs cleanup work after site generation (mvn post-site)
func PostSite(executable string) (string, error) {
	return ExecForStdout(executable, "post-site")
}

// SiteDeploy deploys the generated site to a web server (mvn site-deploy)
func SiteDeploy(executable string) (string, error) {
	return ExecForStdout(executable, "site-deploy")
}