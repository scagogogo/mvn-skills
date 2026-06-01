package installer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// InstallMacOS installs Maven on macOS
func InstallMacOS() (string, error) {
	// Try installing via Homebrew first
	if installed, path := tryHomebrewInstall(); installed {
		return path, nil
	}

	// If Homebrew installation fails, fall back to binary package installation (similar to Linux)
	return installFromBinaryMacOS()
}

// tryHomebrewInstall attempts to install Maven via Homebrew
func tryHomebrewInstall() (bool, string) {
	// Check if brew command is available
	cmd := exec.Command("which", "brew")
	if err := cmd.Run(); err == nil {
		// Install maven via brew
		cmd = exec.Command("brew", "install", "maven")
		if err := cmd.Run(); err == nil {
			// Installation succeeded, find the installation path
			cmd = exec.Command("which", "mvn")
			output, err := cmd.Output()
			if err == nil && len(output) > 0 {
				mvnPath := strings.TrimSpace(string(output))
				// Get MAVEN_HOME
				mavenHome := filepath.Dir(filepath.Dir(mvnPath))
				return true, mavenHome
			}
		}
	}
	return false, ""
}

// installFromBinaryMacOS installs Maven from a binary package on macOS
func installFromBinaryMacOS() (string, error) {
	// Use the same tar.gz package as Linux
	mavenURL := "https://archive.apache.org/dist/maven/maven-3/3.9.11/binaries/apache-maven-3.9.11-bin.tar.gz"

	mavenHome, err := installFromTarGz(mavenURL)
	if err != nil {
		return "", err
	}

	// Set environment variables for macOS
	if err := setMacOSEnvironmentVars(mavenHome); err != nil {
		return "", fmt.Errorf("failed to set environment variables: %w", err)
	}

	return mavenHome, nil
}

// setMacOSEnvironmentVars configures Maven environment variables on macOS
func setMacOSEnvironmentVars(mavenHome string) error {
	// Write environment variables to .zshrc or .bash_profile
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}

	// Check which shell the user is using
	var rcFile string
	shell := os.Getenv("SHELL")
	if strings.Contains(shell, "zsh") {
		rcFile = filepath.Join(homeDir, ".zshrc")
	} else {
		// Default to bash
		rcFile = filepath.Join(homeDir, ".bash_profile")
	}

	// Read existing file content
	content, err := os.ReadFile(rcFile)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read shell config file: %w", err)
	}

	// Build new environment variable settings
	envVars := fmt.Sprintf("\n# Maven environment variables\nexport MAVEN_HOME=%s\nexport PATH=$PATH:$MAVEN_HOME/bin\n", mavenHome)

	// Check if these settings already exist
	if !strings.Contains(string(content), "MAVEN_HOME=") {
		// Append to file
		f, err := os.OpenFile(rcFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to open shell config file: %w", err)
		}
		defer f.Close()

		if _, err := f.WriteString(envVars); err != nil {
			return fmt.Errorf("failed to write environment variables: %w", err)
		}
	}

	fmt.Println("Maven environment variables have been set. Run 'source " + rcFile + "' or restart your terminal for changes to take effect")

	return nil
}
