package installer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// configureEnvironment sets up MAVEN_HOME and PATH for the user's shell/platform.
// On Linux/macOS it appends exports to the appropriate shell rc file.
// On Windows it sets user-level environment variables via setx (safely).
// Returns a warning string (non-nil) if env setup could not be completed fully,
// or an error if a hard failure occurred.
func configureEnvironment(mavenHome string, opts InstallOptions) (warning string, err error) {
	if opts.SkipEnvSetup {
		return "", nil
	}

	if runtime.GOOS == "windows" {
		return configureWindowsEnv(mavenHome, opts)
	}
	return configureUnixEnv(mavenHome, opts, runtime.GOOS)
}

// configureUnixEnv writes MAVEN_HOME/PATH exports to the user's shell rc file.
func configureUnixEnv(mavenHome string, opts InstallOptions, goos string) (string, error) {
	homeDir := opts.HomeDir
	if homeDir == "" {
		var err error
		homeDir, err = os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get user home directory: %w", err)
		}
	}

	rcFile := shellRCFileForGOOS(homeDir, goos)

	content, err := os.ReadFile(rcFile)
	if err != nil && !os.IsNotExist(err) {
		return "", fmt.Errorf("failed to read shell config file %s: %w", rcFile, err)
	}

	if strings.Contains(string(content), "MAVEN_HOME=") {
		// Already configured; nothing to do.
		return "", nil
	}

	envVars := fmt.Sprintf(
		"\n# Maven environment variables\nexport MAVEN_HOME=%s\nexport PATH=$PATH:$MAVEN_HOME/bin\n",
		mavenHome,
	)

	f, err := os.OpenFile(rcFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to open shell config file %s: %w", rcFile, err)
	}
	defer f.Close()

	if _, err := f.WriteString(envVars); err != nil {
		return "", fmt.Errorf("failed to write environment variables: %w", err)
	}

	fmt.Printf("Maven environment variables configured in %s\n", rcFile)
	fmt.Printf("Run 'source %s' or restart your terminal to apply.\n", rcFile)
	return "", nil
}

// shellRCFile returns the rc file path for the user's shell on the current OS.
func shellRCFile(homeDir string) string {
	return shellRCFileForGOOS(homeDir, runtime.GOOS)
}

// shellRCFileForGOOS returns the rc file path for the user's shell on the given OS.
// On macOS defaults to .zshrc (default shell since Catalina), on Linux to .bashrc.
func shellRCFileForGOOS(homeDir, goos string) string {
	shell := os.Getenv("SHELL")
	base := filepath.Base(shell)
	switch base {
	case "zsh":
		return filepath.Join(homeDir, ".zshrc")
	case "bash":
		// On macOS bash uses .bash_profile for login shells; on Linux .bashrc.
		if goos == "darwin" {
			return filepath.Join(homeDir, ".bash_profile")
		}
		return filepath.Join(homeDir, ".bashrc")
	case "fish":
		return filepath.Join(homeDir, ".config", "fish", "config.fish")
	default:
		// Unknown shell — fall back to .bashrc on Linux, .zshrc on macOS.
		if goos == "darwin" {
			return filepath.Join(homeDir, ".zshrc")
		}
		return filepath.Join(homeDir, ".bashrc")
	}
}

// commandExists reports whether the given executable is on PATH.
func commandExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
