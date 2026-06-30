package installer

import (
	"os/exec"
	"path/filepath"
)

// InstallLinux installs Maven on Linux systems.
// It first tries the system package manager (apt/dnf/yum/apk/pacman/zypper),
// then falls back to a binary tar.gz installation.
func InstallLinux() (string, error) {
	return InstallWithOptions(DefaultInstallOptions())
}

// linuxPackageManagerSpec describes how to install Maven via a system package manager.
type linuxPackageManagerSpec struct {
	detect string   // command used to detect the manager's presence (via LookPath)
	install []string // full command to install Maven
}

// linuxPackageManagers lists supported package managers in priority order.
// Detection is via exec.LookPath on the manager binary, which is more reliable
// than probing /etc/*-release files (works in containers and minimal images).
var linuxPackageManagers = []linuxPackageManagerSpec{
	{detect: "apt-get", install: []string{"sudo", "apt-get", "install", "-y", "maven"}},
	{detect: "dnf", install: []string{"sudo", "dnf", "install", "-y", "maven"}},
	{detect: "yum", install: []string{"sudo", "yum", "install", "-y", "maven"}},
	{detect: "apk", install: []string{"sudo", "apk", "add", "--no-cache", "maven"}},
	{detect: "pacman", install: []string{"sudo", "pacman", "-S", "--noconfirm", "maven"}},
	{detect: "zypper", install: []string{"sudo", "zypper", "install", "-y", "maven"}},
}

// tryPackageManagerLinux attempts to install Maven via the first available
// system package manager. Returns (true, mavenHome) on success.
func tryPackageManagerLinux() (bool, string) {
	for _, pm := range linuxPackageManagers {
		if !commandExists(pm.detect) {
			continue
		}
		cmd := exec.Command(pm.install[0], pm.install[1:]...)
		if err := cmd.Run(); err != nil {
			continue
		}
		mvnPath, err := exec.LookPath("mvn")
		if err != nil || mvnPath == "" {
			continue
		}
		// mvn typically lives at <prefix>/bin/mvn; mavenHome is two levels up.
		mavenHome := filepath.Dir(filepath.Dir(mvnPath))
		return true, mavenHome
	}
	return false, ""
}
