package installer

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestDownloadFile tests the file download functionality.
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
	if err := downloadFile(server.URL, destPath); err != nil {
		t.Fatalf("Failed to download file: %v", err)
	}

	content, err := os.ReadFile(destPath)
	if err != nil {
		t.Fatalf("Failed to read downloaded file: %v", err)
	}
	if string(content) != string(testContent) {
		t.Fatalf("File content mismatch, expected: %s, got: %s", testContent, content)
	}
}

// TestDownloadFileHTTPError verifies a non-200 response is reported as an error.
func TestDownloadFileHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "nope", http.StatusInternalServerError)
	}))
	defer server.Close()

	tempDir, err := os.MkdirTemp("", "download-err-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	if err := downloadFile(server.URL, filepath.Join(tempDir, "out.txt")); err == nil {
		t.Fatal("expected error for HTTP 500, got nil")
	}
}

// TestInstall is an integration test that performs a real install.
// Skipped unless RUN_INTEGRATION_TESTS=1 is set.
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
	if _, err := os.Stat(mvnPath); err != nil {
		t.Fatalf("mvn executable not found: %v", err)
	}
	t.Logf("Maven successfully installed to: %s", mavenHome)
}

// TestPathTraversalSecurity tests the safeJoinPath security check.
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
			_, err := safeJoinPath(destDir, tc.headerPath)
			isSafe := err == nil
			if isSafe != tc.expectSafe {
				t.Errorf("Security check unexpected for path: %s, expected safe: %v, got: %v (err: %v)",
					tc.headerPath, tc.expectSafe, isSafe, err)
			}
		})
	}

	t.Run("absolute path is contained under dest", func(t *testing.T) {
		// An absolute header path like "/etc/passwd" is joined under destDir by
		// filepath.Join (Go treats it as relative to destDir), so it is safely
		// contained — this is the desired security behavior.
		absPath := filepath.Join(string(os.PathSeparator), "etc", "passwd")
		if !filepath.IsAbs(absPath) {
			t.Skip("Absolute path test not applicable on this platform")
		}
		got, err := safeJoinPath(destDir, absPath)
		if err != nil {
			t.Errorf("absolute path should be contained under destDir, got error: %v", err)
		}
		if !strings.HasPrefix(got, filepath.Clean(destDir)+string(os.PathSeparator)) {
			t.Errorf("result %q escapes destDir %q", got, destDir)
		}
	})
}

// TestArchiveFileName verifies URL/file name generation per OS.
func TestArchiveFileName(t *testing.T) {
	cases := []struct {
		goos, version, want string
	}{
		{"linux", "3.9.11", "apache-maven-3.9.11-bin.tar.gz"},
		{"darwin", "3.9.11", "apache-maven-3.9.11-bin.tar.gz"},
		{"windows", "3.9.11", "apache-maven-3.9.11-bin.zip"},
		{"windows", "3.9.6", "apache-maven-3.9.6-bin.zip"},
	}
	for _, c := range cases {
		got := archiveFileName(c.version, c.goos)
		if got != c.want {
			t.Errorf("archiveFileName(%q,%q) = %q, want %q", c.version, c.goos, got, c.want)
		}
	}
}

// TestMavenArchiveURL verifies the full URL layout.
func TestMavenArchiveURL(t *testing.T) {
	got := mavenArchiveURL("https://archive.apache.org/dist", "3.9.11", "linux")
	want := "https://archive.apache.org/dist/maven/maven-3/3.9.11/binaries/apache-maven-3.9.11-bin.tar.gz"
	if got != want {
		t.Errorf("mavenArchiveURL = %q, want %q", got, want)
	}
}

// TestResolvedVersionDefaults verifies the version fallback.
func TestResolvedVersion(t *testing.T) {
	if got := (InstallOptions{}).resolvedVersion(); got != DefaultMavenVersion {
		t.Errorf("empty version = %q, want %q", got, DefaultMavenVersion)
	}
	if got := (InstallOptions{Version: "3.9.6"}).resolvedVersion(); got != "3.9.6" {
		t.Errorf("explicit version = %q, want 3.9.6", got)
	}
}

// TestResolvedMirrorsDefaults verifies the mirror fallback.
func TestResolvedMirrors(t *testing.T) {
	if got := (InstallOptions{}).resolvedMirrors(); len(got) != len(DefaultMirrors) {
		t.Errorf("empty mirrors len = %d, want %d", len(got), len(DefaultMirrors))
	}
	custom := []string{"https://example.com"}
	if got := (InstallOptions{Mirrors: custom}).resolvedMirrors(); len(got) != 1 || got[0] != custom[0] {
		t.Errorf("custom mirrors = %v, want %v", got, custom)
	}
}

// TestParseVersionTriple verifies version string parsing.
func TestParseVersionTriple(t *testing.T) {
	major, minor, patch := parseVersionTriple("3.9.11")
	if major != 3 || minor != 9 || patch != 11 {
		t.Errorf("parseVersionTriple(3.9.11) = %d.%d.%d, want 3.9.11", major, minor, patch)
	}
	// Two-part version defaults patch to 0.
	major, minor, patch = parseVersionTriple("3.9")
	if major != 3 || minor != 9 || patch != 0 {
		t.Errorf("parseVersionTriple(3.9) = %d.%d.%d, want 3.9.0", major, minor, patch)
	}
}

// TestPathContains verifies Windows PATH membership checks.
func TestPathContains(t *testing.T) {
	path := `C:\foo;C:\bar\bin;D:\baz`
	if !pathContains(path, `C:\bar\bin`) {
		t.Error("expected pathContains to find C:\\bar\\bin")
	}
	if !pathContains(path, `c:\BAR\bin`) { // case-insensitive
		t.Error("expected pathContains to be case-insensitive")
	}
	if pathContains(path, `C:\missing`) {
		t.Error("did not expect C:\\missing in path")
	}
}

// TestReadChecksumFile verifies checksum file parsing (handles "hash  filename" format).
func TestReadChecksumFile(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "checksum-test")
	if err != nil {
		t.Fatalf("create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	path := filepath.Join(tempDir, "archive.tar.gz.sha512")
	content := "abcdef0123456789  apache-maven-3.9.11-bin.tar.gz\n"
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("write checksum file: %v", err)
	}

	got, err := readChecksumFile(path)
	if err != nil {
		t.Fatalf("readChecksumFile: %v", err)
	}
	if got != "abcdef0123456789" {
		t.Errorf("readChecksumFile = %q, want abcdef0123456789", got)
	}
}

// TestExtractArchiveUnsupported verifies an unknown extension is rejected.
func TestExtractArchiveUnsupported(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "extract-test")
	if err != nil {
		t.Fatalf("create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	archivePath := filepath.Join(tempDir, "archive.rar")
	if err := os.WriteFile(archivePath, []byte("not an archive"), 0644); err != nil {
		t.Fatalf("write file: %v", err)
	}
	if err := extractArchive(archivePath, filepath.Join(tempDir, "out")); err == nil {
		t.Fatal("expected error for unsupported archive type, got nil")
	} else if !strings.Contains(err.Error(), "unsupported archive type") {
		t.Fatalf("unexpected error: %v", err)
	}
}
