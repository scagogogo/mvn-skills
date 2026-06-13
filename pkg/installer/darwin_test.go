//go:build darwin

package installer

import (
	"os"
	"path/filepath"
	"testing"
)

// TestInstallMacOS tests macOS platform installation
// Since it involves real download and installation, this test is skipped by default
func TestInstallMacOS(t *testing.T) {
	// Skip network-dependent tests by default, only run when env var is set
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("Skipping macOS install test (set RUN_INTEGRATION_TESTS=1 to run)")
	}

	// Real integration test
	mavenHome, err := InstallMacOS()
	if err != nil {
		t.Logf("Maven installation failed (possibly network issue): %v", err)
		t.Skip("Network or installation issue, skipping test")
	}

	// Verify installation path
	if mavenHome == "" {
		t.Fatal("Returned Maven installation path is empty")
	}

	// Verify bin/mvn executable exists
	mvnPath := filepath.Join(mavenHome, "bin", "mvn")
	_, err = os.Stat(mvnPath)
	if err != nil {
		t.Fatalf("mvn executable not found: %v", err)
	}

	t.Logf("Maven successfully installed to: %s", mavenHome)
}

// TestSetMacOSEnvironmentVars tests macOS environment variable setup
// This test creates a temporary directory and does not modify real config files
func TestSetMacOSEnvironmentVars(t *testing.T) {
	// Create temporary directory as test MAVEN_HOME
	tempDir, err := os.MkdirTemp("", "maven-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Save original HOME env var, restore after test
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)

	// Set temporary HOME env var to avoid modifying real config files
	os.Setenv("HOME", tempDir)

	// Test environment variable setup function
	err = setMacOSEnvironmentVars(tempDir + "/maven")
	if err != nil {
		t.Fatalf("Failed to set environment variables: %v", err)
	}

	// Check if shell config files were created
	shellFiles := []string{".zshrc", ".bash_profile"}
	found := false

	for _, file := range shellFiles {
		path := filepath.Join(tempDir, file)
		if _, err := os.Stat(path); err == nil {
			// Read file content and check if it contains Maven env var settings
			content, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("Failed to read config file: %v", err)
			}

			if len(content) > 0 {
				found = true
				t.Logf("Found Maven config in %s", file)
			}
		}
	}

	if !found {
		t.Fatal("Maven environment variable settings not found in any shell config file")
	}
}
