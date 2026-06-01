package settings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPluginGroups(t *testing.T) {
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<settings>
    <pluginGroups>
        <pluginGroup>org.mortbay.jetty</pluginGroup>
        <pluginGroup>org.codehaus.mojo</pluginGroup>
    </pluginGroups>
</settings>`
	s, err := ParseBytes([]byte(xml))
	assert.Nil(t, err)

	groups := s.GetPluginGroups()
	assert.Len(t, groups, 2)
	assert.Contains(t, groups, "org.mortbay.jetty")
	assert.Contains(t, groups, "org.codehaus.mojo")
}

func TestGetPluginGroupsEmpty(t *testing.T) {
	xml := `<?xml version="1.0" encoding="UTF-8"?><settings></settings>`
	s, err := ParseBytes([]byte(xml))
	assert.Nil(t, err)

	groups := s.GetPluginGroups()
	assert.NotNil(t, groups)
	assert.Empty(t, groups)
}

func TestGetLocalRepository(t *testing.T) {
	s, err := ParseBytes([]byte(testSettingsXML))
	assert.Nil(t, err)
	assert.Equal(t, "/home/user/.m2/repository", s.GetLocalRepository())
}

func TestIsOffline(t *testing.T) {
	s, err := ParseBytes([]byte(testSettingsXML))
	assert.Nil(t, err)
	assert.False(t, s.IsOffline())

	xml := `<?xml version="1.0" encoding="UTF-8"?><settings><offline>true</offline></settings>`
	s2, err := ParseBytes([]byte(xml))
	assert.Nil(t, err)
	assert.True(t, s2.IsOffline())
}

func TestFindProfile(t *testing.T) {
	s, err := ParseBytes([]byte(testSettingsXML))
	assert.Nil(t, err)

	profile := s.FindProfile("dev")
	assert.NotNil(t, profile)
	assert.Equal(t, "dev", profile.Id)

	profile = s.FindProfile("nonexistent")
	assert.Nil(t, profile)
}
