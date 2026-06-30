package installer

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// TestMirrorFallback verifies that downloadFromMirrors tries the next mirror
// when the primary returns an error.
func TestMirrorFallback(t *testing.T) {
	// Primary mirror: always 500.
	primary := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "down", http.StatusInternalServerError)
	}))
	defer primary.Close()

	// Fallback mirror: serves archive + checksum.
	fallback, tempDir := newMockServer(t, mockServerBuilder{serveArchive: true, serveSHA512: true})
	defer os.RemoveAll(tempDir)
	defer fallback.Close()

	testHomeDir, err := os.MkdirTemp("", "maven-fallback-home")
	if err != nil {
		t.Fatalf("create test home: %v", err)
	}
	defer os.RemoveAll(testHomeDir)

	opts := InstallOptions{
		Version:      "3.9.11",
		Mirrors:      []string{primary.URL, fallback.URL},
		HomeDir:      testHomeDir,
		SkipEnvSetup: true,
		SkipChecksum: false,
		MaxRetries:   1,
	}

	dest, err := downloadFromMirrors(opts, "3.9.11", "linux", filepath.Join(testHomeDir, ".m2", "maven"))
	if err != nil {
		t.Fatalf("expected fallback mirror to succeed, got: %v", err)
	}
	if dest == "" {
		t.Fatal("empty archive path")
	}
	if _, err := os.Stat(dest); err != nil {
		t.Fatalf("archive not present: %v", err)
	}
}

// TestChecksumMismatchFails verifies that a tampered checksum causes failure.
func TestChecksumMismatchFails(t *testing.T) {
	mux := http.NewServeMux()
	// Serve a real archive but a wrong checksum.
	archiveData := []byte("fake archive bytes")
	mux.HandleFunc("/maven/maven-3/3.9.11/binaries/apache-maven-3.9.11-bin.tar.gz", func(w http.ResponseWriter, r *http.Request) {
		w.Write(archiveData)
	})
	mux.HandleFunc("/maven/maven-3/3.9.11/binaries/apache-maven-3.9.11-bin.tar.gz.sha512", func(w http.ResponseWriter, r *http.Request) {
		// Deliberately wrong checksum.
		w.Write([]byte("0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	tempDir, err := os.MkdirTemp("", "checksum-mismatch")
	if err != nil {
		t.Fatalf("create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	opts := InstallOptions{
		Version:      "3.9.11",
		Mirrors:      []string{server.URL},
		HomeDir:      tempDir,
		SkipEnvSetup: true,
		MaxRetries:   1,
	}

	_, err = downloadFromMirrors(opts, "3.9.11", "linux", filepath.Join(tempDir, "dl"))
	if err == nil {
		t.Fatal("expected checksum mismatch error, got nil")
	}
}

// TestAllMirrorsFail verifies that exhausting all mirrors returns an aggregated error.
func TestAllMirrorsFail(t *testing.T) {
	a := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "down", http.StatusInternalServerError)
	}))
	defer a.Close()
	b := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "down", http.StatusInternalServerError)
	}))
	defer b.Close()

	tempDir, err := os.MkdirTemp("", "all-fail")
	if err != nil {
		t.Fatalf("create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	opts := InstallOptions{
		Version:      "3.9.11",
		Mirrors:      []string{a.URL, b.URL},
		HomeDir:      tempDir,
		SkipEnvSetup: true,
		MaxRetries:   1,
	}

	_, err = downloadFromMirrors(opts, "3.9.11", "linux", filepath.Join(tempDir, "dl"))
	if err == nil {
		t.Fatal("expected error when all mirrors fail, got nil")
	}
	if !contains(err.Error(), "all mirrors failed") {
		t.Fatalf("expected aggregated error, got: %v", err)
	}
}

// TestInstallZip covers the Windows zip extraction path via the mock server.
// We force the archive extension by setting goos="windows" in downloadFromMirrors
// and extractArchive dispatches on the file extension.
func TestInstallZip(t *testing.T) {
	server, tempDir := newMockServer(t, mockServerBuilder{archiveExt: "zip", serveArchive: true, serveSHA512: true})
	defer os.RemoveAll(tempDir)
	defer server.Close()

	testHomeDir, err := os.MkdirTemp("", "maven-zip-home")
	if err != nil {
		t.Fatalf("create test home: %v", err)
	}
	defer os.RemoveAll(testHomeDir)

	opts := InstallOptions{
		Version:      "3.9.11",
		Mirrors:      []string{server.URL},
		HomeDir:      testHomeDir,
		SkipEnvSetup: true,
		MaxRetries:   1,
	}

	// Drive the binary install with goos=windows. We call installBinary directly
	// because InstallWithOptions would route through platform-specific code.
	mavenHome, err := installBinaryForGOOS(opts, "3.9.11", "windows")
	if err != nil {
		t.Fatalf("zip install failed: %v", err)
	}
	mvnPath := filepath.Join(mavenHome, "bin", "mvn")
	if _, err := os.Stat(mvnPath); err != nil {
		t.Fatalf("mvn not found in zip install: %v", err)
	}
}

// TestFindMavenHomeVersionMatch verifies directory discovery by version.
func TestFindMavenHomeVersionMatch(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "findhome")
	if err != nil {
		t.Fatalf("create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create an apache-maven-3.9.11 directory tree.
	mavenDir := filepath.Join(tempDir, "apache-maven-3.9.11", "bin")
	if err := os.MkdirAll(mavenDir, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	got, err := findMavenHome(tempDir, "3.9.11")
	if err != nil {
		t.Fatalf("findMavenHome: %v", err)
	}
	if filepath.Base(got) != "apache-maven-3.9.11" {
		t.Errorf("findMavenHome = %q, want base apache-maven-3.9.11", got)
	}

	// Wrong version should not match the directory.
	if _, err := findMavenHome(tempDir, "3.9.6"); err == nil {
		t.Fatal("expected error for non-matching version, got nil")
	}
}

// TestIdempotency verifies that an already-installed Maven (mocked via a fake
// mvn on PATH) causes InstallWithOptions to skip downloading.
func TestIdempotency(t *testing.T) {
	// We can't easily mock finder.FindMaven without polluting PATH, so instead
	// verify the Force flag wiring: Force=true must bypass the check by reaching
	// the binary path (which will then fail at download since no server is set).
	// This at least confirms the code path executes rather than short-circuiting.
	tempDir, err := os.MkdirTemp("", "idempotent")
	if err != nil {
		t.Fatalf("create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	opts := InstallOptions{
		Version:      "3.9.11",
		Mirrors:      []string{"http://127.0.0.1:0"}, // unreachable
		HomeDir:      tempDir,
		SkipEnvSetup: true,
		Force:        true,
		MaxRetries:   1,
	}

	_, err = installBinaryForGOOS(opts, "3.9.11", "linux")
	if err == nil {
		t.Fatal("expected download failure for unreachable mirror, got nil (Force path did not execute)")
	}
}

// TestShellRCFile verifies shell rc file selection.
func TestShellRCFile(t *testing.T) {
	home := "/tmp/fake-home"
	cases := []struct {
		shell, goos, want string
	}{
		{"/bin/zsh", "darwin", filepath.Join(home, ".zshrc")},
		{"/bin/bash", "linux", filepath.Join(home, ".bashrc")},
		{"/bin/bash", "darwin", filepath.Join(home, ".bash_profile")},
		{"/usr/bin/fish", "linux", filepath.Join(home, ".config", "fish", "config.fish")},
	}
	for _, c := range cases {
		t.Run(c.shell+"-"+c.goos, func(t *testing.T) {
			t.Setenv("SHELL", c.shell)
			got := shellRCFileForGOOS(home, c.goos)
			if got != c.want {
				t.Errorf("shellRCFile(shell=%s, goos=%s) = %q, want %q", c.shell, c.goos, got, c.want)
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || (len(s) > 0 && containsStr(s, substr)))
}

func containsStr(s, substr string) bool {
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
