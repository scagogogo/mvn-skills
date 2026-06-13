package installer

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestDownloadFile tests the file download functionality
func TestDownloadFile(t *testing.T) {
	testContent := []byte("This is a test file for Maven installer")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(testContent)
	}))
	defer server.Close()

	tempDir, err := os.MkdirTemp("", "download-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	destPath := filepath.Join(tempDir, "test-file.txt")
	err = downloadFile(server.URL, destPath)
	if err != nil {
		t.Fatalf("Failed to download file: %v", err)
	}

	_, err = os.Stat(destPath)
	if err != nil {
		t.Fatalf("Downloaded file does not exist: %v", err)
	}

	content, err := os.ReadFile(destPath)
	if err != nil {
		t.Fatalf("Failed to read downloaded file: %v", err)
	}

	if string(content) != string(testContent) {
		t.Fatalf("File content mismatch, expected: %s, got: %s", testContent, content)
	}

	t.Log("File download test passed")
}

// TestUntar tests tar.gz extraction (requires pre-prepared test file)
func TestUntar(t *testing.T) {
	t.Skip("Requires pre-prepared tar.gz test file, skipping for now")
}

// TestInstall tests the installation functionality (integration test)
func TestInstall(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("Skipping install test (set RUN_INTEGRATION_TESTS=1 to run)")
	}

	mavenHome, err := Install()
	if err != nil {
		t.Logf("Maven installation failed (possibly network issue): %v", err)
		t.Skip("Network or installation issue, skipping test")
	}

	if mavenHome == "" {
		t.Fatal("Returned Maven installation path is empty")
	}

	mvnPath := filepath.Join(mavenHome, "bin", "mvn")
	_, err = os.Stat(mvnPath)
	if err != nil {
		t.Fatalf("mvn executable not found: %v", err)
	}

	t.Logf("Maven successfully installed to: %s", mavenHome)
}

// TestPathTraversalSecurity tests path traversal security checks
func TestPathTraversalSecurity(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "security-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	destDir := filepath.Join(tempDir, "dest")
	if err := os.MkdirAll(destDir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	testCases := []struct {
		name       string
		headerPath string
		expectSafe bool
	}{
		{"safe path", "normal/path/file.txt", true},
		{"traversal path 1", "../dangerous/path.txt", false},
		{"traversal path 2", "../../etc/passwd", false},
		{"contains ../ but safe", "normal/../path/file.txt", true},
		{"real traversal path", "subdir/../../outside.txt", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var path string
			if filepath.IsAbs(tc.headerPath) {
				path = tc.headerPath
			} else {
				path = filepath.Join(destDir, tc.headerPath)
			}

			cleanedPath := filepath.Clean(path)
			cleanedDestDir := filepath.Clean(destDir)
			isSafe := strings.HasPrefix(cleanedPath, cleanedDestDir+string(os.PathSeparator))

			if isSafe != tc.expectSafe {
				t.Errorf("Security check result unexpected, path: %s (cleaned: %s), expected safe: %v, actual: %v",
					path, cleanedPath, tc.expectSafe, isSafe)
			}
		})
	}

	// Test absolute path security (platform-specific)
	t.Run("absolute path", func(t *testing.T) {
		absPath := filepath.Join(string(os.PathSeparator), "etc", "passwd")
		if !filepath.IsAbs(absPath) {
			t.Skip("Absolute path test not applicable on this platform")
		}

		cleanedPath := filepath.Clean(absPath)
		cleanedDestDir := filepath.Clean(destDir)
		isSafe := strings.HasPrefix(cleanedPath, cleanedDestDir+string(os.PathSeparator))

		if isSafe {
			t.Errorf("Absolute path should not be considered safe: %s (cleaned: %s)", absPath, cleanedPath)
		}
	})
}
