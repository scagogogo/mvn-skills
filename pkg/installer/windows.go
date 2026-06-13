package installer

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// InstallWindows installs Maven on Windows systems
func InstallWindows() (string, error) {
	mavenURL := "https://archive.apache.org/dist/maven/maven-3/3.9.11/binaries/apache-maven-3.9.11-bin.zip"

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	// Create installation directory
	mavenDir := filepath.Join(homeDir, ".m2", "maven")
	if err := os.MkdirAll(mavenDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create Maven installation directory: %w", err)
	}

	// Download Maven
	zipPath := filepath.Join(mavenDir, "maven.zip")
	if err := downloadFileWindows(mavenURL, zipPath); err != nil {
		return "", fmt.Errorf("failed to download Maven: %w", err)
	}

	// Extract Maven
	extractDir := filepath.Join(mavenDir, "maven-install")
	if err := unzipWindows(zipPath, extractDir); err != nil {
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
	if err := setEnvVarsWindows(mavenHome); err != nil {
		return "", fmt.Errorf("failed to set environment variables: %w", err)
	}

	// Remove temporary zip file
	os.Remove(zipPath)

	return mavenHome, nil
}

// downloadFileWindows downloads a file from the given URL
func downloadFileWindows(url, destPath string) error {
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

// unzipWindows extracts a zip archive
func unzipWindows(zipPath, destDir string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	for _, f := range r.File {
		fpath := filepath.Join(destDir, f.Name)

		// Check for path traversal vulnerability
		if !strings.HasPrefix(fpath, filepath.Clean(destDir)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}

	return nil
}

// setEnvVarsWindows configures persistent environment variables using SETX
func setEnvVarsWindows(mavenHome string) error {
	cmds := []struct {
		name string
		args []string
	}{
		{"setx", []string{"MAVEN_HOME", mavenHome}},
		{"setx", []string{"PATH", fmt.Sprintf("%%PATH%%;%s", filepath.Join(mavenHome, "bin"))}},
	}

	for _, cmd := range cmds {
		c := exec.Command(cmd.name, cmd.args...)
		if err := c.Run(); err != nil {
			return fmt.Errorf("failed to execute %s %v: %w", cmd.name, cmd.args, err)
		}
	}

	return nil
}
