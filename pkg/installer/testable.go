package installer

// InstallMacOSWithOptions installs Maven on macOS with configurable options.
//
// Deprecated: use InstallWithOptions, which handles all platforms uniformly.
// Retained for backward compatibility with existing callers and tests.
func InstallMacOSWithOptions(options InstallOptions) (string, error) {
	// Force the binary path on macOS (skip Homebrew) to preserve the original
	// semantics of this function: always download & extract to ~/.m2/maven.
	// The shared InstallWithOptions would otherwise short-circuit when Homebrew
	// Maven is present, which the original tests do not expect.
	options.Force = true
	// Ensure environment setup is honored per the caller's SkipEnvSetup flag
	// (Force only affects the idempotency check, not env setup).
	return installMacOSBinary(options)
}

// installMacOSBinary installs Maven from a binary tar.gz on macOS, bypassing
// Homebrew. Mirrors and version come from opts.
func installMacOSBinary(opts InstallOptions) (string, error) {
	version := opts.resolvedVersion()
	mavenHome, err := installBinary(opts, version)
	if err != nil {
		return "", err
	}
	return mavenHome, nil
}
