package command

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildExecutable(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple path",
			input:    "/opt/maven",
			expected: filepath.Join("/opt/maven", "bin", "mvn"),
		},
		{
			name:     "path with trailing slash",
			input:    "/opt/maven/",
			expected: filepath.Join("/opt/maven/", "bin", "mvn"),
		},
		{
			name:     "windows path",
			input:    "C:\\Program Files\\Apache\\maven",
			expected: filepath.Join("C:\\Program Files\\Apache\\maven", "bin", "mvn"),
		},
		{
			name:     "empty string",
			input:    "",
			expected: filepath.Join("", "bin", "mvn"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildExecutable(tt.input)
			assert.Equal(t, tt.expected, result)
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
