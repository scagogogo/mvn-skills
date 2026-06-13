package command

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildExecutable(t *testing.T) {
	// On Windows the executable is mvn.cmd, on other platforms it's mvn
	execName := "mvn"
	if runtime.GOOS == "windows" {
		execName = "mvn.cmd"
	}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple path",
			input:    "/opt/maven",
			expected: filepath.Join("/opt/maven", "bin", execName),
		},
		{
			name:     "path with trailing slash",
			input:    "/opt/maven/",
			expected: filepath.Join("/opt/maven/", "bin", execName),
		},
		{
			name:     "empty string",
			input:    "",
			expected: filepath.Join("", "bin", execName),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildExecutable(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}

	// Windows-specific test
	if runtime.GOOS == "windows" {
		t.Run("windows path", func(t *testing.T) {
			result := BuildExecutable("C:\\Program Files\\Apache\\maven")
			assert.Equal(t, "C:\\Program Files\\Apache\\maven\\bin\\mvn.cmd", result)
		})
	}
}

func TestOptions_WorkingDirectory(t *testing.T) {
	opts := &Options{
		Executable:       "mvn",
		Args:             []string{"-v"},
		WorkingDirectory: "/tmp",
	}
	assert.Equal(t, "/tmp", opts.WorkingDirectory)
	assert.Equal(t, "mvn", opts.Executable)
	assert.Equal(t, []string{"-v"}, opts.Args)
}
