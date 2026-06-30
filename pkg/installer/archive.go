package installer

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// extractArchive extracts a Maven binary archive to destDir.
// It auto-detects .tar.gz (Linux/macOS) and .zip (Windows) by file extension.
// All entries are validated against path traversal before being written.
func extractArchive(archivePath, destDir string) error {
	switch {
	case strings.HasSuffix(archivePath, ".tar.gz") || strings.HasSuffix(archivePath, ".tgz"):
		return untarTo(archivePath, destDir)
	case strings.HasSuffix(archivePath, ".zip"):
		return unzipTo(archivePath, destDir)
	default:
		return fmt.Errorf("unsupported archive type: %s", archivePath)
	}
}

// safeJoinPath joins name under destDir and rejects path traversal.
// Returns the cleaned absolute path within destDir, or an error if name escapes destDir.
func safeJoinPath(destDir, name string) (string, error) {
	path := filepath.Join(destDir, name)
	cleaned := filepath.Clean(path)
	cleanedDest := filepath.Clean(destDir)
	if cleaned == cleanedDest {
		// name resolved to destDir itself; treat as illegal (no file name)
		return "", fmt.Errorf("illegal file path: %s", name)
	}
	prefix := cleanedDest + string(os.PathSeparator)
	if !strings.HasPrefix(cleaned, prefix) {
		return "", fmt.Errorf("illegal file path: %s", name)
	}
	return cleaned, nil
}

// untarTo extracts a tar.gz archive to destDir.
func untarTo(tarPath, destDir string) error {
	file, err := os.Open(tarPath)
	if err != nil {
		return fmt.Errorf("failed to open tar file: %w", err)
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzr.Close()

	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	tr := tar.NewReader(gzr)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read tar entry: %w", err)
		}

		path, err := safeJoinPath(destDir, header.Name)
		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
		case tar.TypeReg, tar.TypeRegA:
			if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
				return fmt.Errorf("failed to create parent directory: %w", err)
			}
			outFile, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				return fmt.Errorf("failed to create file: %w", err)
			}
			if _, err := io.Copy(outFile, tr); err != nil {
				outFile.Close()
				return fmt.Errorf("failed to write file: %w", err)
			}
			outFile.Close()
		case tar.TypeSymlink:
			// Skip symlinks for security; they could escape destDir.
			continue
		}
	}
	return nil
}

// unzipTo extracts a zip archive to destDir.
func unzipTo(zipPath, destDir string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %w", err)
	}
	defer r.Close()

	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	for _, f := range r.File {
		path, err := safeJoinPath(destDir, f.Name)
		if err != nil {
			return err
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(path, 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return fmt.Errorf("failed to create parent directory: %w", err)
		}

		outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return fmt.Errorf("failed to open zip entry: %w", err)
		}

		if _, err := io.Copy(outFile, rc); err != nil {
			outFile.Close()
			rc.Close()
			return fmt.Errorf("failed to write file: %w", err)
		}
		outFile.Close()
		rc.Close()
	}
	return nil
}

// findMavenHome walks extractDir and returns the path to the apache-maven-* directory.
// version may be empty to match any "apache-maven-*" directory.
func findMavenHome(extractDir, version string) (string, error) {
	target := "apache-maven-"
	var mavenHome string
	err := filepath.Walk(extractDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		base := info.Name()
		if !strings.HasPrefix(base, target) {
			return nil
		}
		if version != "" && base != mavenDirName(version) {
			return nil
		}
		mavenHome = path
		return filepath.SkipDir
	})
	if err != nil {
		return "", fmt.Errorf("failed to find Maven directory: %w", err)
	}
	if mavenHome == "" {
		return "", fmt.Errorf("Maven installation directory not found")
	}
	return mavenHome, nil
}
