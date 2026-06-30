package installer

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// downloadResult holds the outcome of attempting to download from a single mirror.
type downloadResult struct {
	mirror string
	path   string // local archive path on success
	err    error
}

// downloadFromMirrors downloads the Maven archive for the given version from the
// first mirror that succeeds, verifying the SHA512 checksum unless SkipChecksum is set.
// Tries each mirror with up to MaxRetries attempts (with exponential backoff).
// On total failure, returns an aggregated error listing every failed mirror.
func downloadFromMirrors(opts InstallOptions, version, goos string, destDir string) (string, error) {
	mirrors := opts.resolvedMirrors()
	if len(mirrors) == 0 {
		return "", fmt.Errorf("no download mirrors configured")
	}

	maxRetries := opts.MaxRetries
	if maxRetries <= 0 {
		maxRetries = 1
	}

	if err := os.MkdirAll(destDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create download directory: %w", err)
	}

	archiveName := archiveFileName(version, goos)
	var failures []string

	for _, mirror := range mirrors {
		url := mavenArchiveURL(mirror, version, goos)
		destPath := filepath.Join(destDir, archiveName)

		var lastErr error
		for attempt := 1; attempt <= maxRetries; attempt++ {
			err := downloadFile(url, destPath)
			if err == nil {
				// Download succeeded; verify checksum unless skipped.
				if !opts.SkipChecksum {
					if vErr := verifyChecksum(mirror, version, goos, destPath); vErr != nil {
						os.Remove(destPath)
						lastErr = fmt.Errorf("checksum verification failed: %w", vErr)
						// Try next mirror on checksum failure rather than retrying same.
						break
					}
				}
				return destPath, nil
			}
			lastErr = err
			if attempt < maxRetries {
				backoff := time.Duration(math.Pow(2, float64(attempt-1))) * time.Second
				time.Sleep(backoff)
			}
		}
		failures = append(failures, fmt.Sprintf("%s: %v", mirror, lastErr))
		os.Remove(destPath)
	}

	return "", fmt.Errorf("all mirrors failed:\n  - %s", strings.Join(failures, "\n  - "))
}

// downloadFile downloads a single URL to destPath. No retry here; caller handles retries.
// Based on the original installer.downloadFile with added progress reporting to stderr.
func downloadFile(url, destPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP status code: %d", resp.StatusCode)
	}

	file, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write downloaded content: %w", err)
	}
	return nil
}

// verifyChecksum downloads the .sha512 file from mirror and verifies the local archive.
func verifyChecksum(mirror, version, goos, archivePath string) error {
	checksumURL := mavenChecksumURL(mirror, version, goos)
	checksumPath := archivePath + ".sha512"

	if err := downloadFile(checksumURL, checksumPath); err != nil {
		return fmt.Errorf("failed to download checksum: %w", err)
	}
	defer os.Remove(checksumPath)

	expected, err := readChecksumFile(checksumPath)
	if err != nil {
		return err
	}

	actual, err := sha512OfFile(archivePath)
	if err != nil {
		return err
	}

	if !strings.EqualFold(expected, actual) {
		return fmt.Errorf("checksum mismatch: expected %s, got %s", expected, actual)
	}
	return nil
}

// readChecksumFile reads a .sha512 file and returns the hex digest.
// These files may contain "  filename" suffix; we take the first whitespace-delimited token.
func readChecksumFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read checksum file: %w", err)
	}
	fields := strings.Fields(strings.TrimSpace(string(data)))
	if len(fields) == 0 {
		return "", fmt.Errorf("checksum file is empty")
	}
	return fields[0], nil
}

// sha512OfFile computes the SHA512 hex digest of a file.
func sha512OfFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	h := sha512.New()
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
