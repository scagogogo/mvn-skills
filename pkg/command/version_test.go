package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	// This test requires Maven to be installed on the system
	executable := "mvn"
	version, err := Version(executable)
	if err != nil {
		// Maven not installed in test environment
		t.Skip("Maven not installed, skipping Version test")
	}
	assert.NotEmpty(t, version)
	t.Log(version)
}

func TestParseVersion(t *testing.T) {
	output := `Apache Maven 3.9.11 (5786ffe7daa69c7e7e46b3a7e1e4e0db7b88d9c3)
Maven home: /usr/local/Cellar/maven/3.9.11/libexec
Java version: 17.0.8, vendor: Eclipse Adoptium, runtime: /usr/local/Cellar/openjdk/17.0.8/libexec/openjdk.jdk/Contents/Home/jdk-17.0.8/Contents/Home
Default locale: en_US, platform encoding: UTF-8
OS name: "mac os x", version: "14.0", arch: "aarch64", family: "mac"`

	v, err := ParseVersion(output)
	assert.Nil(t, err)
	assert.Equal(t, 3, v.Major)
	assert.Equal(t, 9, v.Minor)
	assert.Equal(t, 11, v.Patch)
	assert.Equal(t, "3.9.11", v.Raw)
	assert.Equal(t, "/usr/local/Cellar/maven/3.9.11/libexec", v.Home)
}

func TestParseVersionTwoPart(t *testing.T) {
	output := `Apache Maven 3.8 (some hash)
Maven home: /opt/maven
Java version: 11.0.12, vendor: Oracle Corporation`

	v, err := ParseVersion(output)
	assert.Nil(t, err)
	assert.Equal(t, 3, v.Major)
	assert.Equal(t, 8, v.Minor)
	assert.Equal(t, 0, v.Patch)
	assert.Equal(t, "3.8", v.Raw)
}

func TestParseVersionEmpty(t *testing.T) {
	_, err := ParseVersion("")
	assert.NotNil(t, err)
}

func TestParseVersionInvalid(t *testing.T) {
	_, err := ParseVersion("not a valid version output")
	assert.NotNil(t, err)
}

func TestMavenVersionIsAtLeast(t *testing.T) {
	v := &MavenVersion{Major: 3, Minor: 9, Patch: 11}

	assert.True(t, v.IsAtLeast(3, 9, 11))
	assert.True(t, v.IsAtLeast(3, 9, 10))
	assert.True(t, v.IsAtLeast(3, 8, 0))
	assert.True(t, v.IsAtLeast(2, 0, 0))
	assert.False(t, v.IsAtLeast(3, 9, 12))
	assert.False(t, v.IsAtLeast(3, 10, 0))
	assert.False(t, v.IsAtLeast(4, 0, 0))
}

func TestMavenVersionIsAtLeastMinor(t *testing.T) {
	v := &MavenVersion{Major: 3, Minor: 9, Patch: 11}

	assert.True(t, v.IsAtLeastMinor(3, 9))
	assert.True(t, v.IsAtLeastMinor(3, 8))
	assert.True(t, v.IsAtLeastMinor(2, 0))
	assert.False(t, v.IsAtLeastMinor(3, 10))
	assert.False(t, v.IsAtLeastMinor(4, 0))
}

func TestMavenVersionString(t *testing.T) {
	v := &MavenVersion{Major: 3, Minor: 9, Patch: 11, Raw: "3.9.11"}
	assert.Equal(t, "3.9.11", v.String())
}
