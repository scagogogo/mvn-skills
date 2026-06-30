package installer

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// InstallMacOS installs Maven on macOS.
// It first tries Homebrew (which handles arm64/intel differences automatically),
// then falls back to a binary tar.gz installation.
func InstallMacOS() (string, error) {
	return InstallWithOptions(DefaultInstallOptions())
}

// tryHomebrewInstall attempts to install Maven via Homebrew.
// Returns (true, mavenHome) on success. Maven's tar.gz is architecture-agnostic
// (pure Java), so no GOARCH-specific handling is needed for the binary fallback.
func tryHomebrewInstall() (bool, string) {
	if !commandExists("brew") {
		return false, ""
	}
	cmd := exec.Command("brew", "install", "maven")
	if err := cmd.Run(); err != nil {
		return false, ""
	}
	mvnPath, err := exec.LookPath("mvn")
	if err != nil || mvnPath == "" {
		return false, ""
	}
	mvnPath = strings.TrimSpace(mvnPath)
	mavenHome := filepath.Dir(filepath.Dir(mvnPath))
	return true, mavenHome
}

// configureMacOSEnv is a thin alias for the shared Unix env configuration,
// kept for backward compatibility with existing tests that call it.
func configureMacOSEnv(mavenHome string, opts InstallOptions) (string, error) {
	return configureUnixEnv(mavenHome, opts, "darwin")
}

// formatArchInfo returns a human-readable architecture label for diagnostics.
func formatArchInfo() string {
	return fmt.Sprintf("macOS (%s/%s)", currentGOOS(), runtime.GOARCH)
}
