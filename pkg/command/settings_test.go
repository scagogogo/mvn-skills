package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalRepository(t *testing.T) {
	executable := "mvn"
	repoDirectory, err := GetLocalRepositoryDirectory(executable)
	if err != nil {
		t.Skip("Maven not installed, skipping LocalRepository test")
	}
	assert.NotEmpty(t, repoDirectory)
	t.Log(repoDirectory)
}
