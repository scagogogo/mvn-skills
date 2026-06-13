package pom

import (
	"encoding/xml"
	"io"
	"os"
)

// ParseFile parses a POM from a file path
func ParseFile(path string) (*Project, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ParseReader(file)
}

// ParseReader parses a POM from an io.Reader
func ParseReader(r io.Reader) (*Project, error) {
	// Read all content first to support re-parsing
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return ParseBytes(data)
}

// ParseBytes parses a POM from a byte slice
func ParseBytes(data []byte) (*Project, error) {
	var project Project
	if err := xml.Unmarshal(data, &project); err != nil {
		return nil, err
	}
	return &project, nil
}

// GetGAV gets the project's GAV coordinates
// If the project itself does not have a groupId/version, it inherits from the parent
func (p *Project) GetGAV() (groupId, artifactId, version string) {
	artifactId = p.ArtifactId

	if p.GroupId != "" {
		groupId = p.GroupId
	} else if p.Parent != nil {
		groupId = p.Parent.GroupId
	}

	if p.Version != "" {
		version = p.Version
	} else if p.Parent != nil {
		version = p.Parent.Version
	}

	return groupId, artifactId, version
}

// GetDependencies gets the project's dependency list
// Returns an empty slice rather than nil when there are no dependencies
func (p *Project) GetDependencies() []Dependency {
	if p.Dependencies == nil {
		return []Dependency{}
	}
	return p.Dependencies.Dependency
}

// GetModules gets the project's submodule list
// Returns an empty slice when it is not a multi-module project
func (p *Project) GetModules() []string {
	if p.Modules == nil {
		return []string{}
	}
	return p.Modules.Module
}

// GetPlugins gets the build plugin list
func (p *Project) GetPlugins() []Plugin {
	if p.Build == nil || p.Build.Plugins == nil {
		return []Plugin{}
	}
	return p.Build.Plugins.Plugin
}

// GetProfiles gets the Profile list
func (p *Project) GetProfiles() []Profile {
	if p.Profiles == nil {
		return []Profile{}
	}
	return p.Profiles.Profile
}

// GetRepositories gets the repository list
func (p *Project) GetRepositories() []Repository {
	if p.Repositories == nil {
		return []Repository{}
	}
	return p.Repositories.Repository
}

// IsMultiModule checks whether this is a multi-module project
func (p *Project) IsMultiModule() bool {
	return p.Modules != nil && len(p.Modules.Module) > 0
}

// HasParent checks whether there is a parent POM
func (p *Project) HasParent() bool {
	return p.Parent != nil
}

// FindDependency finds the dependency with the specified groupId:artifactId
func (p *Project) FindDependency(groupId, artifactId string) *Dependency {
	for i := range p.GetDependencies() {
		dep := &p.Dependencies.Dependency[i]
		if dep.GroupId == groupId && dep.ArtifactId == artifactId {
			return dep
		}
	}
	return nil
}

// FindPlugin finds the plugin with the specified groupId:artifactId
func (p *Project) FindPlugin(groupId, artifactId string) *Plugin {
	for i := range p.GetPlugins() {
		plugin := &p.Build.Plugins.Plugin[i]
		if plugin.GroupId == groupId && plugin.ArtifactId == artifactId {
			return plugin
		}
	}
	return nil
}

// GetProperties returns project properties as a map
func (p *Project) GetProperties() map[string]string {
	result := make(map[string]string)
	if p.Properties != nil {
		for _, entry := range p.Properties.Entries {
			result[entry.Key] = entry.Value
		}
	}
	return result
}

// GetLicenses returns the project license list
func (p *Project) GetLicenses() []License {
	if p.Licenses == nil {
		return []License{}
	}
	return p.Licenses.License
}

// GetDevelopers returns the project developer list
func (p *Project) GetDevelopers() []Developer {
	if p.Developers == nil {
		return []Developer{}
	}
	return p.Developers.Developer
}

// GetScm returns the SCM information, or nil if not defined
func (p *Project) GetScm() *Scm {
	return p.Scm
}

// GetBuild returns the Build configuration, or nil if not defined
func (p *Project) GetBuild() *Build {
	return p.Build
}

// GetPackaging returns the packaging type (defaults to "jar")
func (p *Project) GetPackaging() string {
	if p.Packaging == "" {
		return "jar"
	}
	return p.Packaging
}
