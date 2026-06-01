package settings

import "encoding/xml"

// Settings represents the root structure of a Maven settings.xml file
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

// PluginGroup represents a plugin group prefix
type PluginGroup struct {
	GroupId string `xml:"groupId"`
}

// PluginGroups represents the plugin group list
type PluginGroups struct {
	PluginGroup []string `xml:"pluginGroup"`
}

// Server represents server authentication information
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

// Servers represents the server list
type Servers struct {
	Server []Server `xml:"server"`
}

// Mirror represents a repository mirror configuration
type Mirror struct {
	Id       string `xml:"id"`
	Name     string `xml:"name"`
	URL      string `xml:"url"`
	MirrorOf string `xml:"mirrorOf"`
	Layout   string `xml:"layout"`
	MirrorOfLayouts string `xml:"mirrorOfLayouts"`
	Blocked  bool   `xml:"blocked"`
}

// Mirrors represents the mirror list
type Mirrors struct {
	Mirror []Mirror `xml:"mirror"`
}

// Proxy represents a proxy configuration
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

// Proxies represents the proxy list
type Proxies struct {
	Proxy []Proxy `xml:"proxy"`
}

// SettingsProfile represents a Profile in settings.xml
type SettingsProfile struct {
	Id                 string              `xml:"id"`
	Activation         *SettingsActivation `xml:"activation"`
	Properties         *SettingsProperties `xml:"properties"`
	Repositories       *SettingsRepositories `xml:"repositories"`
	PluginRepositories *SettingsRepositories `xml:"pluginRepositories"`
}

// SettingsActivation represents the activation condition for a settings Profile
type SettingsActivation struct {
	ActiveByDefault bool                `xml:"activeByDefault"`
	Jdk            string              `xml:"jdk"`
	Os             *SettingsActivationOs `xml:"os"`
	Property       *SettingsActivationProperty `xml:"property"`
	File           *SettingsActivationFile `xml:"file"`
}

// SettingsActivationOs represents the operating system activation condition
type SettingsActivationOs struct {
	Name    string `xml:"name"`
	Family  string `xml:"family"`
	Arch    string `xml:"arch"`
	Version string `xml:"version"`
}

// SettingsActivationProperty represents the property activation condition
type SettingsActivationProperty struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

// SettingsActivationFile represents the file activation condition
type SettingsActivationFile struct {
	Exists  string `xml:"exists"`
	Missing string `xml:"missing"`
}

// SettingsProperties represents the properties in a settings Profile
type SettingsProperties struct {
	Entries []SettingsPropertyEntry
}

// SettingsPropertyEntry represents a single property
type SettingsPropertyEntry struct {
	Key   string
	Value string
}

// SettingsRepository represents a repository configuration in settings
type SettingsRepository struct {
	Id        string               `xml:"id"`
	Name      string               `xml:"name"`
	URL       string               `xml:"url"`
	Layout    string               `xml:"layout"`
	Releases  *SettingsRepoPolicy `xml:"releases"`
	Snapshots *SettingsRepoPolicy `xml:"snapshots"`
}

// SettingsRepoPolicy represents the repository policy
type SettingsRepoPolicy struct {
	Enabled        bool   `xml:"enabled"`
	UpdatePolicy   string `xml:"updatePolicy"`
	ChecksumPolicy string `xml:"checksumPolicy"`
}

// SettingsRepositories represents the repository list
type SettingsRepositories struct {
	Repository []SettingsRepository `xml:"repository"`
}

// SettingsProfiles represents the Profile list
type SettingsProfiles struct {
	Profile []SettingsProfile `xml:"profile"`
}

// ActiveProfiles represents the active Profile list
type ActiveProfiles struct {
	ActiveProfile []string `xml:"activeProfile"`
}

// Configuration represents generic configuration (simplified version)
type Configuration struct {
	Elements []ConfigElement
}

// ConfigElement represents a configuration element
type ConfigElement struct {
	XMLName  xml.Name
	Content  string          `xml:",chardata"`
	Children []ConfigElement `xml:",any"`
}