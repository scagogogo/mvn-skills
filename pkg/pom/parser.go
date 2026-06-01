package pom

import (
	"encoding/xml"
	"io"
	"os"
)

// ParseFile 从文件路径解析 POM
func ParseFile(path string) (*Project, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ParseReader(file)
}

// ParseReader 从 io.Reader 解析 POM
func ParseReader(r io.Reader) (*Project, error) {
	// 先读取全部内容以支持二次解析
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return ParseBytes(data)
}

// ParseBytes 从字节数组解析 POM
func ParseBytes(data []byte) (*Project, error) {
	var project Project
	if err := xml.Unmarshal(data, &project); err != nil {
		return nil, err
	}
	return &project, nil
}

// GetGAV 获取项目的 GAV 坐标
// 如果项目本身没有 groupId/version，则从 parent 继承
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

// GetDependencies 获取项目的依赖列表
// 如果没有依赖，返回空切片而非 nil
func (p *Project) GetDependencies() []Dependency {
	if p.Dependencies == nil {
		return []Dependency{}
	}
	return p.Dependencies.Dependency
}

// GetModules 获取项目的子模块列表
// 如果不是多模块项目，返回空切片
func (p *Project) GetModules() []string {
	if p.Modules == nil {
		return []string{}
	}
	return p.Modules.Module
}

// GetPlugins 获取构建插件列表
func (p *Project) GetPlugins() []Plugin {
	if p.Build == nil || p.Build.Plugins == nil {
		return []Plugin{}
	}
	return p.Build.Plugins.Plugin
}

// GetProfiles 获取 Profile 列表
func (p *Project) GetProfiles() []Profile {
	if p.Profiles == nil {
		return []Profile{}
	}
	return p.Profiles.Profile
}

// GetRepositories 获取仓库列表
func (p *Project) GetRepositories() []Repository {
	if p.Repositories == nil {
		return []Repository{}
	}
	return p.Repositories.Repository
}

// IsMultiModule 判断是否是多模块项目
func (p *Project) IsMultiModule() bool {
	return p.Modules != nil && len(p.Modules.Module) > 0
}

// HasParent 判断是否有父 POM
func (p *Project) HasParent() bool {
	return p.Parent != nil
}

// FindDependency 查找指定 groupId:artifactId 的依赖
func (p *Project) FindDependency(groupId, artifactId string) *Dependency {
	for i := range p.GetDependencies() {
		dep := &p.Dependencies.Dependency[i]
		if dep.GroupId == groupId && dep.ArtifactId == artifactId {
			return dep
		}
	}
	return nil
}

// FindPlugin 查找指定 groupId:artifactId 的插件
func (p *Project) FindPlugin(groupId, artifactId string) *Plugin {
	for i := range p.GetPlugins() {
		plugin := &p.Build.Plugins.Plugin[i]
		if plugin.GroupId == groupId && plugin.ArtifactId == artifactId {
			return plugin
		}
	}
	return nil
}
