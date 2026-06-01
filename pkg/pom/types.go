package pom

import "encoding/xml"

// Project represents the root structure of a Maven POM file
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
	Developers       *Developers  `xml:"developers"`
	Contributors     *Contributors `xml:"contributors"`
	MailingLists     *MailingLists `xml:"mailingLists"`
	Prerequisites    *Prerequisites `xml:"prerequisites"`
	DistributionManagement *DistributionManagement `xml:"distributionManagement"`
}

// Parent represents the parent project information in a POM
type Parent struct {
	GroupId      string `xml:"groupId"`
	ArtifactId   string `xml:"artifactId"`
	Version      string `xml:"version"`
	RelativePath string `xml:"relativePath"`
}

// Properties represents Maven properties
type Properties struct {
	Entries []PropertyEntry
}

// PropertyEntry represents a single Maven property
type PropertyEntry struct {
	Key   string
	Value string
}

// Dependency represents a Maven dependency
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

// Exclusions represents the dependency exclusion list
type Exclusions struct {
	Exclusion []Exclusion `xml:"exclusion"`
}

// Exclusion represents a single exclusion entry
type Exclusion struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
}

// Dependencies represents the dependency list
type Dependencies struct {
	Dependency []Dependency `xml:"dependency"`
}

// DependencyManagement represents dependency management
type DependencyManagement struct {
	Dependencies *Dependencies `xml:"dependencies"`
}

// Modules represents the submodule list
type Modules struct {
	Module []string `xml:"module"`
}

// Build represents the build configuration
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

// Resource represents the resource file configuration
type Resource struct {
	Directory    string   `xml:"directory"`
	TargetPath   string   `xml:"targetPath"`
	Filtering    bool     `xml:"filtering"`
	Includes     *Includes `xml:"includes"`
	Excludes     *Excludes `xml:"excludes"`
}

// Resources represents the resource list
type Resources struct {
	Resource []Resource `xml:"resource"`
}

// Includes represents the include patterns
type Includes struct {
	Include []string `xml:"include"`
}

// Excludes represents the exclude patterns
type Excludes struct {
	Exclude []string `xml:"exclude"`
}

// Plugin represents a Maven plugin
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

// Configuration represents the plugin configuration (generic XML structure)
type Configuration struct {
	Elements []ConfigElement
}

// ConfigElement represents a configuration element
type ConfigElement struct {
	XMLName  xml.Name
	Content  string    `xml:",chardata"`
	Children []ConfigElement `xml:",any"`
}

// Execution represents a plugin execution
type Execution struct {
	Id            string       `xml:"id"`
	Phase         string       `xml:"phase"`
	Goals         *Goals       `xml:"goals"`
	Inherited     bool         `xml:"inherited"`
	Configuration *Configuration `xml:"configuration"`
}

// Goals represents the goal list
type Goals struct {
	Goal []string `xml:"goal"`
}

// Executions represents the execution list
type Executions struct {
	Execution []Execution `xml:"execution"`
}

// PluginManagement represents plugin management
type PluginManagement struct {
	Plugins *Plugins `xml:"plugins"`
}

// Plugins represents the plugin list
type Plugins struct {
	Plugin []Plugin `xml:"plugin"`
}

// Filters represents the build filters
type Filters struct {
	Filter []string `xml:"filter"`
}

// Profile represents a Maven Profile
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

// ProfileBuild represents the build configuration within a Profile
type ProfileBuild struct {
	DefaultGoal    string       `xml:"defaultGoal"`
	Resources      *Resources   `xml:"resources"`
	TestResources  *Resources   `xml:"testResources"`
	Plugins        *Plugins     `xml:"plugins"`
	PluginManagement *PluginManagement `xml:"pluginManagement"`
	Filters        *Filters     `xml:"filters"`
}

// Activation represents the Profile activation condition
type Activation struct {
	ActiveByDefault bool           `xml:"activeByDefault"`
	Jdk            string         `xml:"jdk"`
	Os             *ActivationOs  `xml:"os"`
	Property       *ActivationProperty `xml:"property"`
	File           *ActivationFile `xml:"file"`
}

// ActivationOs represents the operating system activation condition
type ActivationOs struct {
	Name    string `xml:"name"`
	Family  string `xml:"family"`
	Arch    string `xml:"arch"`
	Version string `xml:"version"`
}

// ActivationProperty represents the property activation condition
type ActivationProperty struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

// ActivationFile represents the file activation condition
type ActivationFile struct {
	Exists    string `xml:"exists"`
	Missing   string `xml:"missing"`
}

// Profiles represents the Profile list
type Profiles struct {
	Profile []Profile `xml:"profile"`
}

// Repository represents a Maven repository
type Repository struct {
	Id        string       `xml:"id"`
	Name      string       `xml:"name"`
	URL       string       `xml:"url"`
	Layout    string       `xml:"layout"`
	Releases  *RepoPolicy `xml:"releases"`
	Snapshots *RepoPolicy `xml:"snapshots"`
}

// RepoPolicy represents the repository policy
type RepoPolicy struct {
	Enabled        bool   `xml:"enabled"`
	UpdatePolicy   string `xml:"updatePolicy"`
	ChecksumPolicy string `xml:"checksumPolicy"`
}

// Repositories represents the repository list
type Repositories struct {
	Repository []Repository `xml:"repository"`
}

// Scm represents the SCM (Source Code Management) information
type Scm struct {
	Connection          string `xml:"connection"`
	DeveloperConnection string `xml:"developerConnection"`
	Tag                 string `xml:"tag"`
	URL                 string `xml:"url"`
}

// CiManagement represents the CI management information
type CiManagement struct {
	System    string       `xml:"system"`
	URL       string       `xml:"url"`
	Notifiers *Notifiers   `xml:"notifiers"`
}

// Notifiers represents the notifier list
type Notifiers struct {
	Notifier []Notifier `xml:"notifier"`
}

// Notifier represents a notifier
type Notifier struct {
	Type          string `xml:"type"`
	SendOnError   bool   `xml:"sendOnError"`
	SendOnFailure bool   `xml:"sendOnFailure"`
	SendOnSuccess bool   `xml:"sendOnSuccess"`
	SendOnWarning bool   `xml:"sendOnWarning"`
	Address       string `xml:"address"`
	Configuration *Configuration `xml:"configuration"`
}

// IssueManagement represents the issue tracking system information
type IssueManagement struct {
	System string `xml:"system"`
	URL    string `xml:"url"`
}

// Organization represents the organization information
type Organization struct {
	Name string `xml:"name"`
	URL  string `xml:"url"`
}

// License represents the license information
type License struct {
	Name string `xml:"name"`
	URL  string `xml:"url"`
	Distribution string `xml:"distribution"`
	Comments string `xml:"comments"`
}

// Licenses represents the license list
type Licenses struct {
	License []License `xml:"license"`
}

// Developer represents the developer information
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

// Roles represents the role list
type Roles struct {
	Role []string `xml:"role"`
}

// Developers represents the developer list
type Developers struct {
	Developer []Developer `xml:"developer"`
}

// Contributor represents the contributor information
type Contributor struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
	URL   string `xml:"url"`
	Organization string `xml:"organization"`
	OrganizationURL string `xml:"organizationUrl"`
	Roles *Roles `xml:"roles"`
	Timezone string `xml:"timezone"`
}

// Contributors represents the contributor list
type Contributors struct {
	Contributor []Contributor `xml:"contributor"`
}

// MailingList represents a mailing list
type MailingList struct {
	Name      string `xml:"name"`
	Subscribe string `xml:"subscribe"`
	Unsubscribe string `xml:"unsubscribe"`
	Post      string `xml:"post"`
	Archive   string `xml:"archive"`
	OtherArchives *OtherArchives `xml:"otherArchives"`
}

// OtherArchives represents other archives
type OtherArchives struct {
	OtherArchive []string `xml:"otherArchive"`
}

// MailingLists represents the mailing list collection
type MailingLists struct {
	MailingList []MailingList `xml:"mailingList"`
}

// Prerequisites represents the prerequisites
type Prerequisites struct {
	Maven string `xml:"maven"`
}

// DistributionManagement represents the distribution management
type DistributionManagement struct {
	Repository       *DeploymentRepository `xml:"repository"`
	SnapshotRepository *DeploymentRepository `xml:"snapshotRepository"`
	Site             *Site              `xml:"site"`
	DownloadURL      string             `xml:"downloadUrl"`
	Relocation       *Relocation        `xml:"relocation"`
	Status           string             `xml:"status"`
}

// DeploymentRepository represents a deployment repository
type DeploymentRepository struct {
	Id        string `xml:"id"`
	Name      string `xml:"name"`
	URL       string `xml:"url"`
	Layout    string `xml:"layout"`
	UniqueVersion bool `xml:"uniqueVersion"`
}

// Site represents the site deployment information
type Site struct {
	Id   string `xml:"id"`
	Name string `xml:"name"`
	URL  string `xml:"url"`
}

// Relocation represents the relocation information
type Relocation struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
	Message    string `xml:"message"`
}

// Reports represents the report configuration
type Reports struct {}

// Reporting represents the report generation configuration
type Reporting struct {
	ExcludeDefaults bool    `xml:"excludeDefaults"`
	OutputDirectory string  `xml:"outputDirectory"`
	Plugins         *ReportPlugins `xml:"plugins"`
}

// ReportPlugin represents a report plugin
type ReportPlugin struct {
	GroupId      string       `xml:"groupId"`
	ArtifactId  string       `xml:"artifactId"`
	Version     string       `xml:"version"`
	Inherited   bool         `xml:"inherited"`
	Configuration *Configuration `xml:"configuration"`
	ReportSets  *ReportSets  `xml:"reportSets"`
}

// ReportSet represents a report set
type ReportSet struct {
	Id            string       `xml:"id"`
	Reports       *Reports     `xml:"reports"`
	Inherited     bool         `xml:"inherited"`
	Configuration *Configuration `xml:"configuration"`
}

// ReportSets represents the report set list
type ReportSets struct {
	ReportSet []ReportSet `xml:"reportSet"`
}

// ReportPlugins represents the report plugin list
type ReportPlugins struct {
	ReportPlugin []ReportPlugin `xml:"reportPlugin"`
}