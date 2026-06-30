package installer

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// DefaultMavenVersion is the Maven version installed by default when no version is specified.
// Centralized here so upgrading only requires changing this single constant.
const DefaultMavenVersion = "3.9.11"

// archiveType returns the archive file extension used on the current OS.
// Windows uses .zip, Linux/macOS use .tar.gz.
func archiveType(goos string) string {
	if goos == "windows" {
		return "zip"
	}
	return "tar.gz"
}

// archiveFileName builds the binary archive file name for a given version and OS.
// Example: "apache-maven-3.9.11-bin.tar.gz"
func archiveFileName(version, goos string) string {
	return fmt.Sprintf("apache-maven-%s-bin.%s", version, archiveType(goos))
}

// mavenDirName returns the top-level directory name inside the archive.
// Example: "apache-maven-3.9.11"
func mavenDirName(version string) string {
	return fmt.Sprintf("apache-maven-%s", version)
}

// mavenArchiveURL builds the full download URL for a Maven binary archive.
// mirror example: "https://archive.apache.org/dist"
// Result example: "https://archive.apache.org/dist/maven/maven-3/3.9.11/binaries/apache-maven-3.9.11-bin.tar.gz"
func mavenArchiveURL(mirror, version, goos string) string {
	return fmt.Sprintf(
		"%s/maven/maven-3/%s/binaries/%s",
		mirror,
		version,
		archiveFileName(version, goos),
	)
}

// mavenChecksumURL builds the URL for the SHA512 checksum file accompanying the archive.
func mavenChecksumURL(mirror, version, goos string) string {
	return mavenArchiveURL(mirror, version, goos) + ".sha512"
}

// DefaultMirrors is the ordered list of download mirrors tried by Install.
// Official sources first, then regional mirrors that accelerate downloads in China.
var DefaultMirrors = []string{
	"https://archive.apache.org/dist",
	"https://dlcdn.apache.org",
	"https://mirrors.aliyun.com/apache",
	"https://mirrors.tuna.tsinghua.edu.cn/apache",
}

// InstallOptions holds configurable options for Maven installation.
type InstallOptions struct {
	// Version is the Maven version to install (e.g. "3.9.11").
	// Empty means DefaultMavenVersion.
	Version string

	// Mirrors is the ordered list of download mirrors to try.
	// Empty means DefaultMirrors.
	Mirrors []string

	// HomeDir is the user home directory used for installation.
	// Empty means the real user home directory.
	HomeDir string

	// SkipEnvSetup skips shell/registry environment variable configuration.
	// Useful for testing or when the caller manages PATH manually.
	SkipEnvSetup bool

	// SkipChecksum skips SHA512 verification of the downloaded archive.
	// Useful when mirrors don't host the .sha512 file.
	SkipChecksum bool

	// Force reinstalls even if a usable Maven is already present.
	Force bool

	// MaxRetries is the number of retry attempts per mirror on download failure.
	// Zero or negative means 1 attempt (no retries).
	MaxRetries int
}

// DefaultInstallOptions returns the default installation options.
func DefaultInstallOptions() InstallOptions {
	return InstallOptions{
		Version:    DefaultMavenVersion,
		Mirrors:    append([]string(nil), DefaultMirrors...),
		MaxRetries: 3,
	}
}

// resolvedVersion returns the effective version, falling back to the default.
func (o InstallOptions) resolvedVersion() string {
	if o.Version == "" {
		return DefaultMavenVersion
	}
	return o.Version
}

// resolvedMirrors returns the effective mirror list, falling back to defaults.
func (o InstallOptions) resolvedMirrors() []string {
	if len(o.Mirrors) == 0 {
		return DefaultMirrors
	}
	return o.Mirrors
}

// installBaseDir returns the directory under which Maven is extracted.
// Convention: <homeDir>/.m2/maven
func (o InstallOptions) installBaseDir() (string, error) {
	homeDir := o.HomeDir
	if homeDir == "" {
		var err error
		homeDir, err = os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get user home directory: %w", err)
		}
	}
	return filepath.Join(homeDir, ".m2", "maven"), nil
}

// currentGOOS returns the target OS. Extracted for testability.
func currentGOOS() string {
	return runtime.GOOS
}
