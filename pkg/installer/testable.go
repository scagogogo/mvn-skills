package installer

import (
	"fmt"
	"os"
	"path/filepath"
)

// InstallOptions holds configurable options for Maven installation
type InstallOptions struct {
	// MavenURL is the download URL for the Maven binary archive
	MavenURL string
	// HomeDir is the user home directory (defaults to real home directory if empty)
	HomeDir string
	// SkipEnvSetup skips environment variable configuration (useful for testing)
	SkipEnvSetup bool
}

// DefaultInstallOptions returns the default installation options
func DefaultInstallOptions() InstallOptions {
	return InstallOptions{
		MavenURL:     "https://archive.apache.org/dist/maven/maven-3/3.9.11/binaries/apache-maven-3.9.11-bin.tar.gz",
		HomeDir:      "",
		SkipEnvSetup: false,
	}
}

// InstallMacOSWithOptions installs Maven on macOS with configurable options
// This function is designed for testing, allowing dependency injection
func InstallMacOSWithOptions(options InstallOptions) (string, error) {
	// Determine home directory
	homeDir := options.HomeDir
	if homeDir == "" {
		var err error
		homeDir, err = os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get user home directory: %w", err)
		}
	}

	// Create installation directory
	mavenDir := filepath.Join(homeDir, ".m2", "maven")
	if err := os.MkdirAll(mavenDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create Maven installation directory: %w", err)
	}

	// Download Maven
	tarPath := filepath.Join(mavenDir, "maven.tar.gz")
	if err := downloadFile(options.MavenURL, tarPath); err != nil {
		return "", fmt.Errorf("failed to download Maven: %w", err)
	}

	// Extract Maven
	extractDir := filepath.Join(mavenDir, "maven-install")
	if err := untar(tarPath, extractDir); err != nil {
		return "", fmt.Errorf("failed to extract Maven: %w", err)
	}

	// Find the extracted directory
	mavenHome := ""
	err := filepath.Walk(extractDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && filepath.Base(path) == "apache-maven-3.9.11" {
			mavenHome = path
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to find Maven directory: %w", err)
	}

	if mavenHome == "" {
		return "", fmt.Errorf("Maven installation directory not found")
	}

	// Ensure bin directory exists
	binDir := filepath.Join(mavenHome, "bin")
	if _, err := os.Stat(binDir); os.IsNotExist(err) {
		return "", fmt.Errorf("Maven installation incomplete, bin directory not found")
	}

	// Set environment variables (can be skipped)
	if !options.SkipEnvSetup {
		if err := setMacOSEnvironmentVars(mavenHome); err != nil {
			return "", fmt.Errorf("failed to set environment variables: %w", err)
		}
	}

	// Clean up temporary file
	os.Remove(tarPath)

	return mavenHome, nil
}
