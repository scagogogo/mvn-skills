package finder

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaven(t *testing.T) {
	maven, err := FindMaven()
	if err != nil {
		// Maven not installed in test environment — should return NotFoundError
		var nfe *NotFoundError
		assert.True(t, errors.As(err, &nfe))
	} else {
		assert.NotEmpty(t, maven)
	}
}

func TestCheck(t *testing.T) {
	// Non-existent directory
	assert.False(t, Check("/nonexistent/maven/home"))

	// A directory that exists but doesn't have mvn in bin/
	tmpDir := t.TempDir()
	assert.False(t, Check(tmpDir))
}
