package installer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// InstallWindows installs Maven on Windows systems.
// Kept for backward compatibility; delegates to installBinaryWithOptions.
func InstallWindows() (string, error) {
	return InstallWithOptions(DefaultInstallOptions())
}

// configureWindowsEnv sets user-level MAVEN_HOME and appends Maven's bin to the
// user PATH safely. Unlike the previous setx-based approach, it reads the current
// user PATH and writes it back with Maven prepended, avoiding the PATH truncation
// bug where setx "%%PATH%%;..." drops existing entries.
//
// Returns a warning (not error) when PATH cannot be set automatically — e.g. when
// the resulting value would exceed setx's 1024-char limit — so callers can surface
// manual instructions without failing the whole install.
func configureWindowsEnv(mavenHome string, opts InstallOptions) (string, error) {
	homeDir := opts.HomeDir
	if homeDir == "" {
		var err error
		homeDir, err = os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get user home directory: %w", err)
		}
	}
	_ = homeDir

	// 1. MAVEN_HOME — safe to set via setx (single value, well under 1024 chars).
	if err := setx("MAVEN_HOME", mavenHome); err != nil {
		return "", fmt.Errorf("failed to set MAVEN_HOME: %w", err)
	}

	// 2. PATH — read current user PATH and append Maven's bin if not already present.
	binDir := filepath.Join(mavenHome, "bin")
	currentPath, err := readUserPath()
	if err != nil {
		// Could not read PATH reliably; warn instead of failing.
		warning := fmt.Sprintf(
			"Could not read current user PATH (%v). Manually add to PATH: %s",
			err, binDir,
		)
		return warning, nil
	}

	if pathContains(currentPath, binDir) {
		// Already on PATH; nothing to do.
		return "", nil
	}

	newPath := strings.TrimRight(currentPath, ";") + ";" + binDir
	if len(newPath) > 1024 {
		// setx truncates at 1024 chars; refuse to write to avoid corrupting PATH.
		warning := fmt.Sprintf(
			"User PATH is too long to update safely (%d chars). Manually add to PATH: %s",
			len(newPath), binDir,
		)
		return warning, nil
	}

	if err := setx("PATH", newPath); err != nil {
		warning := fmt.Sprintf(
			"Failed to set PATH (%v). Manually add to PATH: %s",
			err, binDir,
		)
		return warning, nil
	}

	fmt.Println("MAVEN_HOME and PATH configured. Restart your terminal for changes to take effect.")
	return "", nil
}

// setx runs the Windows setx command to persistently set a user environment variable.
func setx(name, value string) error {
	cmd := exec.Command("setx", name, value)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("setx %s failed: %w (output: %s)", name, err, strings.TrimSpace(string(out)))
	}
	return nil
}

// readUserPath reads the current user-level PATH from the registry
// (HKCU\Environment\PATH), falling back to os.Getenv("PATH") if registry access fails.
func readUserPath() (string, error) {
	cmd := exec.Command("reg", "query", "HKCU\\Environment", "/v", "PATH")
	out, err := cmd.Output()
	if err == nil {
		// Parse: lines contain "    PATH    REG_EXPAND_SZ    <value>" or "    PATH    REG_SZ    <value>"
		for _, line := range strings.Split(string(out), "\n") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "PATH") {
				// Split on whitespace; value is everything after the type token.
				fields := strings.Fields(line)
				// fields[0]=PATH fields[1]=REG_... fields[2:]=value parts
				if len(fields) >= 3 {
					// Rejoin in case the path itself had spaces.
					idx := strings.Index(line, fields[2])
					if idx >= 0 {
						return strings.TrimSpace(line[idx:]), nil
					}
				}
			}
		}
	}
	// Fallback to process PATH (may include system + user paths).
	return os.Getenv("PATH"), nil
}

// pathContains reports whether dir is present in a semicolon-separated PATH string.
// Comparison is case-insensitive on Windows.
func pathContains(pathValue, dir string) bool {
	dir = strings.ToLower(filepath.Clean(dir))
	for _, entry := range strings.Split(pathValue, ";") {
		entry = strings.ToLower(strings.TrimSpace(entry))
		if entry == dir {
			return true
		}
	}
	return false
}
