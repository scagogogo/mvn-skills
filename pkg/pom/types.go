package pom

import "encoding/xml"

// Project 表示 Maven POM 文件的根结构
type Project struct {
	XMLName          xml.Name     `xml:"project"`
	ModelVersion     string       `xml:"modelVersion"`
	GroupId          string       `xml:"groupId"`
	ArtifactId       string       `xml:"artifactId"`
	Version          string       `xml:"version"`
	Packaging        string       `xml:"packaging"`
	Name             string       `xml:"name"`
	Description      string       `xml:"description"`
	URL              string       `xml:"url"`
	InceptionYear    string       `xml:"inceptionYear"`
	Parent           *Parent      `xml:"parent"`
	Properties       *Properties  `xml:"properties"`
	Dependencies     *Dependencies `xml:"dependencies"`
	DependencyManagement *DependencyManagement `xml:"dependencyManagement"`
	Modules          *Modules     `xml:"modules"`
	Build            *Build       `xml:"build"`
	Profiles         *Profiles    `xml:"profiles"`
	Repositories     *Repositories `xml:"repositories"`
	PluginRepositories *Repositories `xml:"pluginRepositories"`
	CiManagement     *CiManagement `xml:"ciManagement"`
	Scm              *Scm         `xml:"scm"`
	IssueManagement  *IssueManagement `xml:"issueManagement"`
	Organization     *Organization `xml:"organization"`
	Licenses         *Licenses    `xml:"licenses"`
	Developers       *Developers `xml:"developers"`
	Contributors     *Contributors `xml:"contributors"`
	MailingLists     *MailingLists `xml:"mailingLists"`
	Prerequisites    *Prerequisites `xml:"prerequisites"`
	DistributionManagement *DistributionManagement `xml:"distributionManagement"`
}

// Parent 表示 POM 的父项目信息
type Parent struct {
	GroupId      string `xml:"groupId"`
	ArtifactId   string `xml:"artifactId"`
	Version      string `xml:"version"`
	RelativePath string `xml:"relativePath"`
}

// Properties 表示 Maven 属性
type Properties struct {
	Entries []PropertyEntry
}

// PropertyEntry 表示单个 Maven 属性
type PropertyEntry struct {
	Key   string
	Value string
}

// Dependency 表示 Maven 依赖
type Dependency struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
	Type       string `xml:"type"`
	Scope      string `xml:"scope"`
	Classifier string `xml:"classifier"`
	Optional   bool   `xml:"optional"`
	Exclusions *Exclusions `xml:"exclusions"`
}

// Exclusions 表示依赖排除列表
type Exclusions struct {
	Exclusion []Exclusion `xml:"exclusion"`
}

// Exclusion 表示单个排除项
type Exclusion struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
}

// Dependencies 表示依赖列表
type Dependencies struct {
	Dependency []Dependency `xml:"dependency"`
}

// DependencyManagement 表示依赖管理
type DependencyManagement struct {
	Dependencies *Dependencies `xml:"dependencies"`
}

// Modules 表示子模块列表
type Modules struct {
	Module []string `xml:"module"`
}

// Build 表示构建配置
type Build struct {
	DefaultGoal    string       `xml:"defaultGoal"`
	SourceDirectory string      `xml:"sourceDirectory"`
	TestSourceDirectory string   `xml:"testSourceDirectory"`
	Resources      *Resources  `xml:"resources"`
	TestResources  *Resources  `xml:"testResources"`
	Plugins        *Plugins    `xml:"plugins"`
	PluginManagement *PluginManagement `xml:"pluginManagement"`
	Filters        *Filters     `xml:"filters"`
	FinalName      string      `xml:"finalName"`
	Directory      string      `xml:"directory"`
	OutputDirectory string     `xml:"outputDirectory"`
	TestOutputDirectory string  `xml:"testOutputDirectory"`
}

// Resource 表示资源文件配置
type Resource struct {
	Directory    string   `xml:"directory"`
	TargetPath   string   `xml:"targetPath"`
	Filtering    bool     `xml:"filtering"`
	Includes     *Includes `xml:"includes"`
	Excludes     *Excludes `xml:"excludes"`
}

// Resources 表示资源列表
type Resources struct {
	Resource []Resource `xml:"resource"`
}

// Includes 表示包含模式
type Includes struct {
	Include []string `xml:"include"`
}

// Excludes 表示排除模式
type Excludes struct {
	Exclude []string `xml:"exclude"`
}

// Plugin 表示 Maven 插件
type Plugin struct {
	GroupId      string     `xml:"groupId"`
	ArtifactId  string     `xml:"artifactId"`
	Version     string     `xml:"version"`
	Extensions  bool       `xml:"extensions"`
	Inherited   bool       `xml:"inherited"`
	Configuration *Configuration `xml:"configuration"`
	Executions  *Executions `xml:"executions"`
	Dependencies *Dependencies `xml:"dependencies"`
}

// Configuration 表示插件配置（通用 XML 结构）
type Configuration struct {
	Elements []ConfigElement
}

// ConfigElement 表示配置元素
type ConfigElement struct {
	XMLName  xml.Name
	Content  string    `xml:",chardata"`
	Children []ConfigElement `xml:",any"`
}

// Execution 表示插件执行配置
type Execution struct {
	Id            string       `xml:"id"`
	Phase         string       `xml:"phase"`
	Goals         *Goals       `xml:"goals"`
	Inherited     bool         `xml:"inherited"`
	Configuration *Configuration `xml:"configuration"`
}

// Goals 表示目标列表
type Goals struct {
	Goal []string `xml:"goal"`
}

// Executions 表示执行列表
type Executions struct {
	Execution []Execution `xml:"execution"`
}

// PluginManagement 表示插件管理
type PluginManagement struct {
	Plugins *Plugins `xml:"plugins"`
}

// Plugins 表示插件列表
type Plugins struct {
	Plugin []Plugin `xml:"plugin"`
}

// Filters 表示构建过滤器
type Filters struct {
	Filter []string `xml:"filter"`
}

// Profile 表示 Maven Profile
type Profile struct {
	Id             string       `xml:"id"`
	Activation     *Activation  `xml:"activation"`
	Properties     *Properties  `xml:"properties"`
	Dependencies   *Dependencies `xml:"dependencies"`
	DependencyManagement *DependencyManagement `xml:"dependencyManagement"`
	Modules        *Modules     `xml:"modules"`
	Build          *ProfileBuild `xml:"build"`
	Repositories   *Repositories `xml:"repositories"`
	PluginRepositories *Repositories `xml:"pluginRepositories"`
	Reports        *Reports     `xml:"reports"`
	Reporting      *Reporting   `xml:"reporting"`
}

// ProfileBuild 表示 Profile 中的构建配置
type ProfileBuild struct {
	DefaultGoal    string       `xml:"defaultGoal"`
	Resources      *Resources   `xml:"resources"`
	TestResources  *Resources   `xml:"testResources"`
	Plugins        *Plugins     `xml:"plugins"`
	PluginManagement *PluginManagement `xml:"pluginManagement"`
	Filters        *Filters     `xml:"filters"`
}

// Activation 表示 Profile 激活条件
type Activation struct {
	ActiveByDefault bool           `xml:"activeByDefault"`
	Jdk            string         `xml:"jdk"`
	Os             *ActivationOs  `xml:"os"`
	Property       *ActivationProperty `xml:"property"`
	File           *ActivationFile `xml:"file"`
}

// ActivationOs 表示操作系统激活条件
type ActivationOs struct {
	Name    string `xml:"name"`
	Family  string `xml:"family"`
	Arch    string `xml:"arch"`
	Version string `xml:"version"`
}

// ActivationProperty 表示属性激活条件
type ActivationProperty struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

// ActivationFile 表示文件激活条件
type ActivationFile struct {
	Exists    string `xml:"exists"`
	Missing   string `xml:"missing"`
}

// Profiles 表示 Profile 列表
type Profiles struct {
	Profile []Profile `xml:"profile"`
}

// Repository 表示 Maven 仓库
type Repository struct {
	Id        string       `xml:"id"`
	Name      string       `xml:"name"`
	URL       string       `xml:"url"`
	Layout    string       `xml:"layout"`
	Releases  *RepoPolicy `xml:"releases"`
	Snapshots *RepoPolicy `xml:"snapshots"`
}

// RepoPolicy 表示仓库策略
type RepoPolicy struct {
	Enabled        bool   `xml:"enabled"`
	UpdatePolicy   string `xml:"updatePolicy"`
	ChecksumPolicy string `xml:"checksumPolicy"`
}

// Repositories 表示仓库列表
type Repositories struct {
	Repository []Repository `xml:"repository"`
}

// Scm 表示源代码管理信息
type Scm struct {
	Connection          string `xml:"connection"`
	DeveloperConnection string `xml:"developerConnection"`
	Tag                 string `xml:"tag"`
	URL                 string `xml:"url"`
}

// CiManagement 表示 CI 管理信息
type CiManagement struct {
	System    string       `xml:"system"`
	URL       string       `xml:"url"`
	Notifiers *Notifiers   `xml:"notifiers"`
}

// Notifiers 表示通知者列表
type Notifiers struct {
	Notifier []Notifier `xml:"notifier"`
}

// Notifier 表示通知者
type Notifier struct {
	Type          string `xml:"type"`
	SendOnError   bool   `xml:"sendOnError"`
	SendOnFailure bool   `xml:"sendOnFailure"`
	SendOnSuccess bool   `xml:"sendOnSuccess"`
	SendOnWarning bool   `xml:"sendOnWarning"`
	Address       string `xml:"address"`
	Configuration *Configuration `xml:"configuration"`
}

// IssueManagement 表示问题追踪系统信息
type IssueManagement struct {
	System string `xml:"system"`
	URL    string `xml:"url"`
}

// Organization 表示组织信息
type Organization struct {
	Name string `xml:"name"`
	URL  string `xml:"url"`
}

// License 表示许可证信息
type License struct {
	Name string `xml:"name"`
	URL  string `xml:"url"`
	Distribution string `xml:"distribution"`
	Comments string `xml:"comments"`
}

// Licenses 表示许可证列表
type Licenses struct {
	License []License `xml:"license"`
}

// Developer 表示开发者信息
type Developer struct {
	Id    string `xml:"id"`
	Name  string `xml:"name"`
	Email string `xml:"email"`
	URL   string `xml:"url"`
	Organization string `xml:"organization"`
	OrganizationURL string `xml:"organizationUrl"`
	Roles *Roles `xml:"roles"`
	Timezone string `xml:"timezone"`
}

// Roles 表示角色列表
type Roles struct {
	Role []string `xml:"role"`
}

// Developers 表示开发者列表
type Developers struct {
	Developer []Developer `xml:"developer"`
}

// Contributor 表示贡献者信息
type Contributor struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
	URL   string `xml:"url"`
	Organization string `xml:"organization"`
	OrganizationURL string `xml:"organizationUrl"`
	Roles *Roles `xml:"roles"`
	Timezone string `xml:"timezone"`
}

// Contributors 表示贡献者列表
type Contributors struct {
	Contributor []Contributor `xml:"contributor"`
}

// MailingList 表示邮件列表
type MailingList struct {
	Name      string `xml:"name"`
	Subscribe string `xml:"subscribe"`
	Unsubscribe string `xml:"unsubscribe"`
	Post      string `xml:"post"`
	Archive   string `xml:"archive"`
	OtherArchives *OtherArchives `xml:"otherArchives"`
}

// OtherArchives 表示其他归档
type OtherArchives struct {
	OtherArchive []string `xml:"otherArchive"`
}

// MailingLists 表示邮件列表
type MailingLists struct {
	MailingList []MailingList `xml:"mailingList"`
}

// Prerequisites 表示前置条件
type Prerequisites struct {
	Maven string `xml:"maven"`
}

// DistributionManagement 表示分发管理
type DistributionManagement struct {
	Repository       *DeploymentRepository `xml:"repository"`
	SnapshotRepository *DeploymentRepository `xml:"snapshotRepository"`
	Site             *Site              `xml:"site"`
	DownloadURL      string             `xml:"downloadUrl"`
	Relocation       *Relocation        `xml:"relocation"`
	Status           string             `xml:"status"`
}

// DeploymentRepository 表示部署仓库
type DeploymentRepository struct {
	Id        string `xml:"id"`
	Name      string `xml:"name"`
	URL       string `xml:"url"`
	Layout    string `xml:"layout"`
	UniqueVersion bool `xml:"uniqueVersion"`
}

// Site 表示站点部署信息
type Site struct {
	Id   string `xml:"id"`
	Name string `xml:"name"`
	URL  string `xml:"url"`
}

// Relocation 表示重定位信息
type Relocation struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
	Message    string `xml:"message"`
}

// Reports 表示报告配置
type Reports struct {}

// Reporting 表示报告生成配置
type Reporting struct {
	ExcludeDefaults bool    `xml:"excludeDefaults"`
	OutputDirectory string  `xml:"outputDirectory"`
	Plugins         *ReportPlugins `xml:"plugins"`
}

// ReportPlugin 表示报告插件
type ReportPlugin struct {
	GroupId      string       `xml:"groupId"`
	ArtifactId  string       `xml:"artifactId"`
	Version     string       `xml:"version"`
	Inherited   bool         `xml:"inherited"`
	Configuration *Configuration `xml:"configuration"`
	ReportSets  *ReportSets  `xml:"reportSets"`
}

// ReportSet 表示报告集
type ReportSet struct {
	Id            string       `xml:"id"`
	Reports       *Reports     `xml:"reports"`
	Inherited     bool         `xml:"inherited"`
	Configuration *Configuration `xml:"configuration"`
}

// ReportSets 表示报告集列表
type ReportSets struct {
	ReportSet []ReportSet `xml:"reportSet"`
}

// ReportPlugins 表示报告插件列表
type ReportPlugins struct {
	ReportPlugin []ReportPlugin `xml:"reportPlugin"`
}
