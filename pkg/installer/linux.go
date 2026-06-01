package installer

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// InstallLinux installs Maven on Linux systems
func InstallLinux() (string, error) {
	// Try package manager first
	if installed, path := tryPackageManagerLinux(); installed {
		return path, nil
	}

	// Fall back to binary installation
	return installFromBinaryLinux()
}

// tryPackageManagerLinux attempts to install Maven via system package manager
func tryPackageManagerLinux() (bool, string) {
	if _, err := os.Stat("/etc/debian_version"); err == nil {
		// Debian/Ubuntu
		cmd := exec.Command("sudo", "apt-get", "update")
		cmd.Run()

		cmd = exec.Command("sudo", "apt-get", "install", "-y", "maven")
		if err := cmd.Run(); err == nil {
			cmd = exec.Command("which", "mvn")
			output, err := cmd.Output()
			if err == nil && len(output) > 0 {
				mvnPath := strings.TrimSpace(string(output))
				mavenHome := filepath.Dir(filepath.Dir(mvnPath))
				return true, mavenHome
			}
		}
	} else if _, err := os.Stat("/etc/redhat-release"); err == nil {
		// RedHat/CentOS/Fedora
		cmd := exec.Command("sudo", "yum", "install", "-y", "maven")
		if err := cmd.Run(); err == nil {
			cmd = exec.Command("which", "mvn")
			output, err := cmd.Output()
			if err == nil && len(output) > 0 {
				mvnPath := strings.TrimSpace(string(output))
				mavenHome := filepath.Dir(filepath.Dir(mvnPath))
				return true, mavenHome
			}
		}
	}

	return false, ""
}

// installFromBinaryLinux installs Maven from a binary tar.gz archive
func installFromBinaryLinux() (string, error) {
	mavenURL := "https://archive.apache.org/dist/maven/maven-3/3.9.11/binaries/apache-maven-3.9.11-bin.tar.gz"

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	mavenDir := filepath.Join(homeDir, ".m2", "maven")
	if err := os.MkdirAll(mavenDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create Maven installation directory: %w", err)
	}

	// Download Maven
	tarPath := filepath.Join(mavenDir, "maven.tar.gz")
	if err := downloadFileLinux(mavenURL, tarPath); err != nil {
		return "", fmt.Errorf("failed to download Maven: %w", err)
	}

	// Extract Maven
	extractDir := filepath.Join(mavenDir, "maven-install")
	if err := untarLinux(tarPath, extractDir); err != nil {
		return "", fmt.Errorf("failed to extract Maven: %w", err)
	}

	// Find the extracted directory
	mavenHome := ""
	err = filepath.Walk(extractDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && strings.Contains(info.Name(), "apache-maven") {
			mavenHome = path
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to find Maven directory: %w", err)
	}

	if mavenHome == "" {
		return "", errors.New("Maven installation directory not found")
	}

	// Ensure bin directory exists
	binDir := filepath.Join(mavenHome, "bin")
	if _, err := os.Stat(binDir); os.IsNotExist(err) {
		return "", errors.New("Maven installation incomplete, bin directory not found")
	}

	// Set environment variables
	if err := setEnvironmentVarsLinux(mavenHome); err != nil {
		return "", fmt.Errorf("failed to set environment variables: %w", err)
	}

	// Remove temporary file
	os.Remove(tarPath)

	return mavenHome, nil
}

// downloadFileLinux downloads a file from the given URL
func downloadFileLinux(url, destPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed, HTTP status code: %d", resp.StatusCode)
	}

	file, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

// untarLinux extracts a tar.gz archive
func untarLinux(tarPath, destDir string) error {
	file, err := os.Open(tarPath)
	if err != nil {
		return err
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		path := filepath.Join(destDir, header.Name)

		// Check for path traversal vulnerability
		if !strings.HasPrefix(path, filepath.Clean(destDir)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			dir := filepath.Dir(path)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}

			outFile, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			if _, err := io.Copy(outFile, tr); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()
		}
	}

	return nil
}

// setEnvironmentVarsLinux configures environment variables in shell rc file
func setEnvironmentVarsLinux(mavenHome string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Detect user's shell
	var rcFile string
	shell := os.Getenv("SHELL")
	if strings.Contains(shell, "zsh") {
		rcFile = filepath.Join(homeDir, ".zshrc")
	} else {
		rcFile = filepath.Join(homeDir, ".bashrc")
	}

	// Read existing file content
	content, err := os.ReadFile(rcFile)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// Build environment variable entries
	envVars := fmt.Sprintf("\n# Maven environment variables\nexport MAVEN_HOME=%s\nexport PATH=$PATH:$MAVEN_HOME/bin\n", mavenHome)

	// Append only if not already configured
	if !strings.Contains(string(content), "MAVEN_HOME=") {
		f, err := os.OpenFile(rcFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := f.WriteString(envVars); err != nil {
			return err
		}
	}

	fmt.Println("Maven environment variables configured. Run 'source " + rcFile + "' or restart your terminal to apply.")

	return nil
}
