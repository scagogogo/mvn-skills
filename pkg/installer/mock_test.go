package installer

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// mockMaven describes the small Maven tree placed inside test archives.
type mockMaven struct {
	Version   string
	BinFiles  []string
	LibFiles  []string
	ConfFiles []string
}

func createMockMaven() mockMaven {
	return mockMaven{
		Version:   "3.9.11",
		BinFiles:  []string{"mvn", "mvnDebug"},
		LibFiles:  []string{"maven-core.jar", "maven-model.jar"},
		ConfFiles: []string{"settings.xml", "logging/simplelogger.properties"},
	}
}

// createMockMavenTarGz writes a fake apache-maven-<version> tar.gz to outputPath.
func createMockMavenTarGz(t *testing.T, outputPath string) error {
	t.Helper()
	mock := createMockMaven()
	return writeMockTarGz(outputPath, mock)
}

// writeMockTarGz builds a tar.gz archive matching the Maven layout for the given mock.
func writeMockTarGz(outputPath string, mock mockMaven) error {
	var buf bytes.Buffer
	gzWriter := gzip.NewWriter(&buf)
	tarWriter := tar.NewWriter(gzWriter)

	baseName := "apache-maven-" + mock.Version

	// Helper to add a directory entry.
	addDir := func(name string) error {
		return tarWriter.WriteHeader(&tar.Header{
			Name:     name,
			Mode:     0755,
			Typeflag: tar.TypeDir,
		})
	}
	// Helper to add a regular file.
	addFile := func(name string, content []byte, mode int64) error {
		if err := tarWriter.WriteHeader(&tar.Header{
			Name:     name,
			Mode:     mode,
			Size:     int64(len(content)),
			Typeflag: tar.TypeReg,
		}); err != nil {
			return err
		}
		_, err := tarWriter.Write(content)
		return err
	}

	if err := addDir(baseName + "/"); err != nil {
		return err
	}
	if err := addDir(baseName + "/bin/"); err != nil {
		return err
	}
	for _, f := range mock.BinFiles {
		if err := addFile(baseName+"/bin/"+f, []byte("#!/bin/sh\necho mock maven"), 0755); err != nil {
			return err
		}
	}
	if err := addDir(baseName + "/lib/"); err != nil {
		return err
	}
	for _, f := range mock.LibFiles {
		if err := addFile(baseName+"/lib/"+f, []byte("mock jar"), 0644); err != nil {
			return err
		}
	}
	if err := addDir(baseName + "/conf/"); err != nil {
		return err
	}
	for _, f := range mock.ConfFiles {
		dir := filepath.Dir(f)
		if dir != "." {
			if err := addDir(baseName + "/conf/" + dir + "/"); err != nil {
				return err
			}
		}
		if err := addFile(baseName+"/conf/"+f, []byte("# mock config\n"), 0644); err != nil {
			return err
		}
	}

	if err := tarWriter.Close(); err != nil {
		return err
	}
	if err := gzWriter.Close(); err != nil {
		return err
	}
	return os.WriteFile(outputPath, buf.Bytes(), 0644)
}

// writeMockZip builds a zip archive matching the Maven layout.
func writeMockZip(outputPath string, mock mockMaven) error {
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)
	baseName := "apache-maven-" + mock.Version

	addDir := func(name string) error {
		w, err := zw.Create(name)
		if err != nil {
			return err
		}
		_, err = w.Write(nil)
		return err
	}
	addFile := func(name, content string) error {
		w, err := zw.Create(name)
		if err != nil {
			return err
		}
		_, err = w.Write([]byte(content))
		return err
	}

	if err := addDir(baseName + "/"); err != nil {
		return err
	}
	if err := addDir(baseName + "/bin/"); err != nil {
		return err
	}
	for _, f := range mock.BinFiles {
		if err := addFile(baseName+"/bin/"+f, "#!/bin/sh\necho mock maven"); err != nil {
			return err
		}
	}
	if err := addDir(baseName + "/lib/"); err != nil {
		return err
	}
	for _, f := range mock.LibFiles {
		if err := addFile(baseName+"/lib/"+f, "mock jar"); err != nil {
			return err
		}
	}
	if err := addDir(baseName + "/conf/"); err != nil {
		return err
	}
	for _, f := range mock.ConfFiles {
		if err := addFile(baseName+"/conf/"+f, "# mock config\n"); err != nil {
			return err
		}
	}
	if err := zw.Close(); err != nil {
		return err
	}
	return os.WriteFile(outputPath, buf.Bytes(), 0644)
}

// mockServerBuilder configures a test HTTP server hosting a Maven archive.
type mockServerBuilder struct {
	version      string
	archiveExt   string // "tar.gz" or "zip"
	serveArchive bool   // if false, the archive endpoint returns 500
	serveSHA512  bool   // if false, the .sha512 endpoint returns 500
	mock         mockMaven
}

// newMockServer builds an httptest.Server that serves the archive and/or its
// SHA512 checksum under the standard Apache URL layout.
func newMockServer(t *testing.T, b mockServerBuilder) (*httptest.Server, string) {
	t.Helper()
	if b.version == "" {
		b.version = "3.9.11"
	}
	if b.archiveExt == "" {
		b.archiveExt = "tar.gz"
	}
	if b.mock.Version == "" {
		b.mock = createMockMaven()
		b.mock.Version = b.version
	}

	tempDir, err := os.MkdirTemp("", "maven-mock")
	if err != nil {
		t.Fatalf("create temp dir: %v", err)
	}

	archiveName := fmt.Sprintf("apache-maven-%s-bin.%s", b.version, b.archiveExt)
	archivePath := filepath.Join(tempDir, archiveName)
	switch b.archiveExt {
	case "tar.gz":
		if err := writeMockTarGz(archivePath, b.mock); err != nil {
			t.Fatalf("write mock tar.gz: %v", err)
		}
	case "zip":
		if err := writeMockZip(archivePath, b.mock); err != nil {
			t.Fatalf("write mock zip: %v", err)
		}
	}

	archiveData, err := os.ReadFile(archivePath)
	if err != nil {
		t.Fatalf("read archive: %v", err)
	}
	sum := sha512.Sum512(archiveData)
	checksum := hex.EncodeToString(sum[:])

	mux := http.NewServeMux()
	archivePathHandler := fmt.Sprintf("/maven/maven-3/%s/binaries/%s", b.version, archiveName)
	checksumPath := archivePathHandler + ".sha512"

	mux.HandleFunc(archivePathHandler, func(w http.ResponseWriter, r *http.Request) {
		if !b.serveArchive {
			http.Error(w, "archive disabled", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(archiveData)
	})
	mux.HandleFunc(checksumPath, func(w http.ResponseWriter, r *http.Request) {
		if !b.serveSHA512 {
			http.Error(w, "checksum disabled", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, checksum)
	})

	server := httptest.NewServer(mux)
	return server, tempDir
}

// TestInstallMacOSWithMock verifies the binary-install path end-to-end using a
// mock server: download → SHA512 verify → extract → locate mvn.
func TestInstallMacOSWithMock(t *testing.T) {
	server, tempDir := newMockServer(t, mockServerBuilder{serveArchive: true, serveSHA512: true})
	defer os.RemoveAll(tempDir)
	defer server.Close()

	testHomeDir, err := os.MkdirTemp("", "maven-test-home")
	if err != nil {
		t.Fatalf("create test home: %v", err)
	}
	defer os.RemoveAll(testHomeDir)

	options := InstallOptions{
		Version:      "3.9.11",
		Mirrors:      []string{server.URL},
		HomeDir:      testHomeDir,
		SkipEnvSetup: true,
		MaxRetries:   1,
	}

	mavenHome, err := InstallMacOSWithOptions(options)
	if err != nil {
		t.Fatalf("install failed: %v", err)
	}
	if mavenHome == "" {
		t.Fatal("empty maven home")
	}

	mvnPath := filepath.Join(mavenHome, "bin", "mvn")
	if _, err := os.Stat(mvnPath); err != nil {
		t.Fatalf("mvn not found: %v", err)
	}
	t.Logf("Maven installed to: %s", mavenHome)
}
