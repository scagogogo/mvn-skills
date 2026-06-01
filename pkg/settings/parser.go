package settings

import (
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

// ParseFile parses a settings.xml from a file path
func ParseFile(path string) (*Settings, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ParseReader(file)
}

// ParseReader parses a settings.xml from an io.Reader
func ParseReader(r io.Reader) (*Settings, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return ParseBytes(data)
}

// ParseBytes parses a settings.xml from a byte slice
func ParseBytes(data []byte) (*Settings, error) {
	var settings Settings
	if err := xml.Unmarshal(data, &settings); err != nil {
		return nil, err
	}
	return &settings, nil
}

// ParseDefault attempts to parse the settings.xml at the default location
// Checks ~/.m2/settings.xml first, then ${M2_HOME}/conf/settings.xml
func ParseDefault() (*Settings, error) {
	// Check the settings.xml in the user's home directory
	homeDir, err := os.UserHomeDir()
	if err == nil {
		userSettings := filepath.Join(homeDir, ".m2", "settings.xml")
		if _, err := os.Stat(userSettings); err == nil {
			return ParseFile(userSettings)
		}
	}

	// Check the settings.xml under M2_HOME
	m2Home := getMavenHome()
	if m2Home != "" {
		globalSettings := filepath.Join(m2Home, "conf", "settings.xml")
		if _, err := os.Stat(globalSettings); err == nil {
			return ParseFile(globalSettings)
		}
	}

	return nil, os.ErrNotExist
}

// GetDefaultSettingsPath returns the default settings.xml path
// Returns the user-level settings.xml path (regardless of whether the file exists)
func GetDefaultSettingsPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(homeDir, ".m2", "settings.xml")
}

// GetMirrors gets all mirror configurations
func (s *Settings) GetMirrors() []Mirror {
	if s.Mirrors == nil {
		return []Mirror{}
	}
	return s.Mirrors.Mirror
}

// GetServers gets all server authentication information
func (s *Settings) GetServers() []Server {
	if s.Servers == nil {
		return []Server{}
	}
	return s.Servers.Server
}

// GetProxies gets all proxy configurations
func (s *Settings) GetProxies() []Proxy {
	if s.Proxies == nil {
		return []Proxy{}
	}
	return s.Proxies.Proxy
}

// GetProfiles gets all Profiles
func (s *Settings) GetProfiles() []SettingsProfile {
	if s.Profiles == nil {
		return []SettingsProfile{}
	}
	return s.Profiles.Profile
}

// GetActiveProfileIds gets the active Profile ID list
func (s *Settings) GetActiveProfileIds() []string {
	if s.ActiveProfiles == nil {
		return []string{}
	}
	return s.ActiveProfiles.ActiveProfile
}

// FindServer finds server authentication information by ID
func (s *Settings) FindServer(id string) *Server {
	for i := range s.GetServers() {
		if s.Servers.Server[i].Id == id {
			return &s.Servers.Server[i]
		}
	}
	return nil
}

// FindMirror finds a mirror configuration by ID
func (s *Settings) FindMirror(id string) *Mirror {
	for i := range s.GetMirrors() {
		if s.Mirrors.Mirror[i].Id == id {
			return &s.Mirrors.Mirror[i]
		}
	}
	return nil
}

// FindMirrorOf finds a mirror matching the specified repository ID
func (s *Settings) FindMirrorOf(repositoryId string) *Mirror {
	for i := range s.GetMirrors() {
		mirror := &s.Mirrors.Mirror[i]
		if mirror.MirrorOf == repositoryId || mirror.MirrorOf == "*" {
			return mirror
		}
	}
	return nil
}

// FindActiveProxy finds the first active proxy
func (s *Settings) FindActiveProxy() *Proxy {
	for i := range s.GetProxies() {
		if s.Proxies.Proxy[i].Active {
			return &s.Proxies.Proxy[i]
		}
	}
	return nil
}

// GetPluginGroups returns the plugin group prefixes
func (s *Settings) GetPluginGroups() []string {
	if s.PluginGroups == nil {
		return []string{}
	}
	return s.PluginGroups.PluginGroup
}

// GetLocalRepository returns the local repository path, or empty string if not set
func (s *Settings) GetLocalRepository() string {
	return s.LocalRepository
}

// IsOffline returns whether offline mode is enabled
func (s *Settings) IsOffline() bool {
	return s.Offline
}

// FindProfile finds a profile by ID
func (s *Settings) FindProfile(id string) *SettingsProfile {
	for i := range s.GetProfiles() {
		if s.Profiles.Profile[i].Id == id {
			return &s.Profiles.Profile[i]
		}
	}
	return nil
}

// getMavenHome gets the Maven home directory
func getMavenHome() string {
	// Check M2_HOME first
	if m2Home := os.Getenv("M2_HOME"); m2Home != "" {
		return m2Home
	}
	// Then check MAVEN_HOME
	if mavenHome := os.Getenv("MAVEN_HOME"); mavenHome != "" {
		return mavenHome
	}
	// Try common installation paths
	switch runtime.GOOS {
	case "linux":
		candidates := []string{"/usr/share/maven", "/opt/maven", "/usr/local/maven"}
		for _, dir := range candidates {
			if _, err := os.Stat(filepath.Join(dir, "bin", "mvn")); err == nil {
				return dir
			}
		}
	case "darwin":
		candidates := []string{"/usr/local/Cellar/maven", "/opt/homebrew/Cellar/maven"}
		for _, dir := range candidates {
			if entries, err := os.ReadDir(dir); err == nil && len(entries) > 0 {
				return filepath.Join(dir, entries[0].Name(), "libexec")
			}
		}
	}
	return ""
}