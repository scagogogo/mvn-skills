package settings

import (
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

// ParseFile 从文件路径解析 settings.xml
func ParseFile(path string) (*Settings, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ParseReader(file)
}

// ParseReader 从 io.Reader 解析 settings.xml
func ParseReader(r io.Reader) (*Settings, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return ParseBytes(data)
}

// ParseBytes 从字节数组解析 settings.xml
func ParseBytes(data []byte) (*Settings, error) {
	var settings Settings
	if err := xml.Unmarshal(data, &settings); err != nil {
		return nil, err
	}
	return &settings, nil
}

// ParseDefault 尝试解析默认位置的 settings.xml
// 依次检查 ~/.m2/settings.xml 和 ${M2_HOME}/conf/settings.xml
func ParseDefault() (*Settings, error) {
	// 检查用户目录下的 settings.xml
	homeDir, err := os.UserHomeDir()
	if err == nil {
		userSettings := filepath.Join(homeDir, ".m2", "settings.xml")
		if _, err := os.Stat(userSettings); err == nil {
			return ParseFile(userSettings)
		}
	}

	// 检查 M2_HOME 下的 settings.xml
	m2Home := getMavenHome()
	if m2Home != "" {
		globalSettings := filepath.Join(m2Home, "conf", "settings.xml")
		if _, err := os.Stat(globalSettings); err == nil {
			return ParseFile(globalSettings)
		}
	}

	return nil, os.ErrNotExist
}

// GetDefaultSettingsPath 返回默认 settings.xml 路径
// 返回用户级 settings.xml 的路径（无论文件是否存在）
func GetDefaultSettingsPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(homeDir, ".m2", "settings.xml")
}

// GetMirrors 获取所有镜像配置
func (s *Settings) GetMirrors() []Mirror {
	if s.Mirrors == nil {
		return []Mirror{}
	}
	return s.Mirrors.Mirror
}

// GetServers 获取所有服务器认证信息
func (s *Settings) GetServers() []Server {
	if s.Servers == nil {
		return []Server{}
	}
	return s.Servers.Server
}

// GetProxies 获取所有代理配置
func (s *Settings) GetProxies() []Proxy {
	if s.Proxies == nil {
		return []Proxy{}
	}
	return s.Proxies.Proxy
}

// GetProfiles 获取所有 Profile
func (s *Settings) GetProfiles() []SettingsProfile {
	if s.Profiles == nil {
		return []SettingsProfile{}
	}
	return s.Profiles.Profile
}

// GetActiveProfileIds 获取激活的 Profile ID 列表
func (s *Settings) GetActiveProfileIds() []string {
	if s.ActiveProfiles == nil {
		return []string{}
	}
	return s.ActiveProfiles.ActiveProfile
}

// FindServer 根据 ID 查找服务器认证信息
func (s *Settings) FindServer(id string) *Server {
	for i := range s.GetServers() {
		if s.Servers.Server[i].Id == id {
			return &s.Servers.Server[i]
		}
	}
	return nil
}

// FindMirror 根据 ID 查找镜像配置
func (s *Settings) FindMirror(id string) *Mirror {
	for i := range s.GetMirrors() {
		if s.Mirrors.Mirror[i].Id == id {
			return &s.Mirrors.Mirror[i]
		}
	}
	return nil
}

// FindMirrorOf 查找匹配指定仓库 ID 的镜像
func (s *Settings) FindMirrorOf(repositoryId string) *Mirror {
	for i := range s.GetMirrors() {
		mirror := &s.Mirrors.Mirror[i]
		if mirror.MirrorOf == repositoryId || mirror.MirrorOf == "*" {
			return mirror
		}
	}
	return nil
}

// FindActiveProxy 查找第一个激活的代理
func (s *Settings) FindActiveProxy() *Proxy {
	for i := range s.GetProxies() {
		if s.Proxies.Proxy[i].Active {
			return &s.Proxies.Proxy[i]
		}
	}
	return nil
}

// getMavenHome 获取 Maven Home 目录
func getMavenHome() string {
	// 优先检查 M2_HOME
	if m2Home := os.Getenv("M2_HOME"); m2Home != "" {
		return m2Home
	}
	// 然后检查 MAVEN_HOME
	if mavenHome := os.Getenv("MAVEN_HOME"); mavenHome != "" {
		return mavenHome
	}
	// 尝试常见安装路径
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
