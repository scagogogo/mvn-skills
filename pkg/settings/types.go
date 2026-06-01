package settings

import "encoding/xml"

// Settings 表示 Maven settings.xml 的根结构
type Settings struct {
	XMLName           xml.Name           `xml:"settings"`
	Xmlns             string             `xml:"xmlns,attr"`
	LocalRepository   string             `xml:"localRepository"`
	InteractiveMode   bool               `xml:"interactiveMode"`
	UsePluginRegistry bool               `xml:"usePluginRegistry"`
	Offline           bool               `xml:"offline"`
	PluginGroups      *PluginGroups      `xml:"pluginGroups"`
	Servers           *Servers           `xml:"servers"`
	Mirrors           *Mirrors           `xml:"mirrors"`
	Proxies           *Proxies           `xml:"proxies"`
	Profiles          *SettingsProfiles  `xml:"profiles"`
	ActiveProfiles    *ActiveProfiles    `xml:"activeProfiles"`
}

// PluginGroup 表示插件组前缀
type PluginGroup struct {
	GroupId string `xml:"groupId"`
}

// PluginGroups 表示插件组列表
type PluginGroups struct {
	PluginGroup []PluginGroup `xml:"pluginGroup"`
}

// Server 表示服务器认证信息
type Server struct {
	Id       string `xml:"id"`
	Username string `xml:"username"`
	Password string `xml:"password"`
	PrivateKey string `xml:"privateKey"`
	Passphrase string `xml:"passphrase"`
	FilePermissions string `xml:"filePermissions"`
	DirectoryPermissions string `xml:"directoryPermissions"`
	Configuration *Configuration `xml:"configuration"`
}

// Servers 表示服务器列表
type Servers struct {
	Server []Server `xml:"server"`
}

// Mirror 表示仓库镜像配置
type Mirror struct {
	Id       string `xml:"id"`
	Name     string `xml:"name"`
	URL      string `xml:"url"`
	MirrorOf string `xml:"mirrorOf"`
	Layout   string `xml:"layout"`
	MirrorOfLayouts string `xml:"mirrorOfLayouts"`
	Blocked  bool   `xml:"blocked"`
}

// Mirrors 表示镜像列表
type Mirrors struct {
	Mirror []Mirror `xml:"mirror"`
}

// Proxy 表示代理配置
type Proxy struct {
	Id       string `xml:"id"`
	Active   bool   `xml:"active"`
	Protocol string `xml:"protocol"`
	Username string `xml:"username"`
	Password string `xml:"password"`
	Port     int    `xml:"port"`
	Host     string `xml:"host"`
	NonProxyHosts string `xml:"nonProxyHosts"`
}

// Proxies 表示代理列表
type Proxies struct {
	Proxy []Proxy `xml:"proxy"`
}

// SettingsProfile 表示 settings.xml 中的 Profile
type SettingsProfile struct {
	Id                 string              `xml:"id"`
	Activation         *SettingsActivation `xml:"activation"`
	Properties         *SettingsProperties `xml:"properties"`
	Repositories       *SettingsRepositories `xml:"repositories"`
	PluginRepositories *SettingsRepositories `xml:"pluginRepositories"`
}

// SettingsActivation 表示 settings Profile 的激活条件
type SettingsActivation struct {
	ActiveByDefault bool                `xml:"activeByDefault"`
	Jdk            string              `xml:"jdk"`
	Os             *SettingsActivationOs `xml:"os"`
	Property       *SettingsActivationProperty `xml:"property"`
	File           *SettingsActivationFile `xml:"file"`
}

// SettingsActivationOs 表示操作系统激活条件
type SettingsActivationOs struct {
	Name    string `xml:"name"`
	Family  string `xml:"family"`
	Arch    string `xml:"arch"`
	Version string `xml:"version"`
}

// SettingsActivationProperty 表示属性激活条件
type SettingsActivationProperty struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

// SettingsActivationFile 表示文件激活条件
type SettingsActivationFile struct {
	Exists  string `xml:"exists"`
	Missing string `xml:"missing"`
}

// SettingsProperties 表示 settings Profile 中的属性
type SettingsProperties struct {
	Entries []SettingsPropertyEntry
}

// SettingsPropertyEntry 表示单个属性
type SettingsPropertyEntry struct {
	Key   string
	Value string
}

// SettingsRepository 表示 settings 中的仓库配置
type SettingsRepository struct {
	Id        string               `xml:"id"`
	Name      string               `xml:"name"`
	URL       string               `xml:"url"`
	Layout    string               `xml:"layout"`
	Releases  *SettingsRepoPolicy `xml:"releases"`
	Snapshots *SettingsRepoPolicy `xml:"snapshots"`
}

// SettingsRepoPolicy 表示仓库策略
type SettingsRepoPolicy struct {
	Enabled        bool   `xml:"enabled"`
	UpdatePolicy   string `xml:"updatePolicy"`
	ChecksumPolicy string `xml:"checksumPolicy"`
}

// SettingsRepositories 表示仓库列表
type SettingsRepositories struct {
	Repository []SettingsRepository `xml:"repository"`
}

// SettingsProfiles 表示 Profile 列表
type SettingsProfiles struct {
	Profile []SettingsProfile `xml:"profile"`
}

// ActiveProfiles 表示激活的 Profile 列表
type ActiveProfiles struct {
	ActiveProfile []string `xml:"activeProfile"`
}

// Configuration 表示通用配置（简化版）
type Configuration struct {
	Elements []ConfigElement
}

// ConfigElement 表示配置元素
type ConfigElement struct {
	XMLName  xml.Name
	Content  string          `xml:",chardata"`
	Children []ConfigElement `xml:",any"`
}
