package installer

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/scagogogo/mvn-skills/pkg/command"
	"github.com/scagogogo/mvn-skills/pkg/finder"
)

// Install selects the appropriate installation method based on the current
// operating system. It is the simple, zero-configuration entry point.
// For configurable installs, use InstallWithOptions.
func Install() (string, error) {
	return InstallWithOptions(DefaultInstallOptions())
}

// InstallWithOptions installs Maven according to the supplied options.
//
// Workflow:
//  1. Unless Force is set, check whether a usable Maven already exists; if so,
//     skip the download (idempotency).
//  2. On Linux, try the system package manager first.
//  3. On macOS, try Homebrew first.
//  4. Otherwise (or as a fallback), download the binary archive from the
//     configured mirrors with SHA512 verification, extract it, and configure
//     the environment.
func InstallWithOptions(opts InstallOptions) (string, error) {
	version := opts.resolvedVersion()

	// 1. Idempotency: skip if a usable Maven is already installed.
	if !opts.Force {
		if existing := findUsableMaven(version); existing != "" {
			return existing, nil
		}
	}

	// 2/3. Platform-specific fast paths (package manager / Homebrew).
	switch runtime.GOOS {
	case "linux":
		if ok, path := tryPackageManagerLinux(); ok {
			return path, nil
		}
	case "darwin":
		if ok, path := tryHomebrewInstall(); ok {
			return path, nil
		}
	}

	// 4. Binary archive fallback for all platforms.
	return installBinary(opts, version)
}

// installBinary downloads, verifies, and extracts the Maven binary archive,
// then configures the environment.
func installBinary(opts InstallOptions, version string) (string, error) {
	return installBinaryForGOOS(opts, version, currentGOOS())
}

// installBinaryForGOOS is the testable form of installBinary, allowing the
// target OS to be injected so the zip path can be exercised on any host.
func installBinaryForGOOS(opts InstallOptions, version, goos string) (string, error) {

	baseDir, err := opts.installBaseDir()
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create Maven installation directory: %w", err)
	}

	// Download (with mirror fallback + SHA512 verification).
	archivePath, err := downloadFromMirrors(opts, version, goos, baseDir)
	if err != nil {
		return "", fmt.Errorf("failed to download Maven: %w", err)
	}
	defer os.Remove(archivePath)

	// Extract.
	extractDir := filepath.Join(baseDir, "maven-install")
	// Clear any previous extraction to avoid stale leftovers mixing in.
	os.RemoveAll(extractDir)
	if err := extractArchive(archivePath, extractDir); err != nil {
		return "", fmt.Errorf("failed to extract Maven: %w", err)
	}

	// Locate the apache-maven-<version> directory.
	mavenHome, err := findMavenHome(extractDir, version)
	if err != nil {
		return "", err
	}

	// Sanity check: bin directory must exist.
	binDir := filepath.Join(mavenHome, "bin")
	if info, err := os.Stat(binDir); err != nil || !info.IsDir() {
		return "", fmt.Errorf("Maven installation incomplete, bin directory not found at %s", binDir)
	}

	// Configure environment (shell rc / Windows registry).
	if warning, err := configureEnvironment(mavenHome, opts); err != nil {
		return "", fmt.Errorf("failed to set environment variables: %w", err)
	} else if warning != "" {
		fmt.Printf("Warning: %s\n", warning)
	}

	return mavenHome, nil
}

// findUsableMaven returns the path to an already-installed Maven whose version
// is >= the requested version, or "" if none is available.
func findUsableMaven(requiredVersion string) string {
	mvnPath, err := finder.FindMaven()
	if err != nil || mvnPath == "" {
		return ""
	}

	// Ask Maven for its version.
	output, err := command.Version(mvnPath)
	if err != nil {
		// Found an executable but can't determine its version; trust it's usable
		// rather than forcing a reinstall. This is the safer choice for users
		// who have a broken-but-present mvn they intend to keep.
		return deriveMavenHome(mvnPath)
	}

	parsed, err := command.ParseVersion(output)
	if err != nil {
		return deriveMavenHome(mvnPath)
	}

	major, minor, patch := parseVersionTriple(requiredVersion)
	if parsed.IsAtLeast(major, minor, patch) {
		if parsed.Home != "" {
			return parsed.Home
		}
		return deriveMavenHome(mvnPath)
	}

	// Installed version is older than required — proceed with reinstall.
	return ""
}

// deriveMavenHome returns the likely MAVEN_HOME for a given mvn executable path,
// assuming the standard <prefix>/bin/mvn layout.
func deriveMavenHome(mvnPath string) string {
	dir := filepath.Dir(mvnPath) // bin
	return filepath.Dir(dir)     // maven home
}

// parseVersionTriple splits a "X.Y.Z" (or "X.Y") version into three ints.
// Missing components default to 0.
func parseVersionTriple(v string) (int, int, int) {
	parts := strings.Split(v, ".")
	major, _ := strconv.Atoi(safeIndex(parts, 0))
	minor, _ := strconv.Atoi(safeIndex(parts, 1))
	patch, _ := strconv.Atoi(safeIndex(parts, 2))
	return major, minor, patch
}

// safeIndex returns parts[i] or "" if out of range.
func safeIndex(parts []string, i int) string {
	if i < len(parts) {
		return parts[i]
	}
	return ""
}
