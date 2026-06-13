package command

// DependencyGetOption holds the options for the dependency:get goal
type DependencyGetOption struct {
	GroupId            string   // -DgroupId
	ArtifactId         string   // -DartifactId
	Version            string   // -Dversion
	Classifier         string   // -Dclassifier (optional)
	Type               string   // -Dtype (optional, e.g. "jar", "war", "pom")
	RemoteRepositories []string // -DremoteRepositories (optional)
	Packaging          string   // alias for Type, for clarity
}

// ToArgs converts DependencyGetOption to Maven command arguments
func (o *DependencyGetOption) ToArgs() []string {
	args := []string{"dependency:get"}
	args = append(args, "-DgroupId="+o.GroupId)
	args = append(args, "-DartifactId="+o.ArtifactId)
	args = append(args, "-Dversion="+o.Version)
	if o.Classifier != "" {
		args = append(args, "-Dclassifier="+o.Classifier)
	}
	packaging := o.Packaging
	if packaging == "" {
		packaging = o.Type
	}
	if packaging != "" {
		args = append(args, "-Dpackaging="+packaging)
	}
	for _, repo := range o.RemoteRepositories {
		args = append(args, "-DremoteRepositories="+repo)
	}
	return args
}

// DeployDeployFileOption holds the options for the deploy:deploy-file goal
type DeployDeployFileOption struct {
	File         string // -Dfile
	GroupId      string // -DgroupId
	ArtifactId   string // -DartifactId
	Version      string // -Dversion
	Packaging    string // -Dpackaging (e.g. "jar", "war", "pom")
	RepositoryId string // -DrepositoryId
	URL          string // -Durl
	PomFile      string // -DpomFile (optional, replaces -DgroupId/ArtifactId/Version/Packaging)
	Classifier   string // -Dclassifier (optional)
	Sources      string // -Dsources (optional source JAR)
	Javadoc      string // -Djavadoc (optional javadoc JAR)
}

// ToArgs converts DeployDeployFileOption to Maven command arguments
func (o *DeployDeployFileOption) ToArgs() []string {
	args := []string{"deploy:deploy-file"}
	if o.File != "" {
		args = append(args, "-Dfile="+o.File)
	}
	if o.PomFile != "" {
		args = append(args, "-DpomFile="+o.PomFile)
	} else {
		if o.GroupId != "" {
			args = append(args, "-DgroupId="+o.GroupId)
		}
		if o.ArtifactId != "" {
			args = append(args, "-DartifactId="+o.ArtifactId)
		}
		if o.Version != "" {
			args = append(args, "-Dversion="+o.Version)
		}
		if o.Packaging != "" {
			args = append(args, "-Dpackaging="+o.Packaging)
		}
	}
	if o.RepositoryId != "" {
		args = append(args, "-DrepositoryId="+o.RepositoryId)
	}
	if o.URL != "" {
		args = append(args, "-Durl="+o.URL)
	}
	if o.Classifier != "" {
		args = append(args, "-Dclassifier="+o.Classifier)
	}
	if o.Sources != "" {
		args = append(args, "-Dsources="+o.Sources)
	}
	if o.Javadoc != "" {
		args = append(args, "-Djavadoc="+o.Javadoc)
	}
	return args
}

// InstallFileOption holds the options for the install:install-file goal
type InstallFileOption struct {
	File       string // -Dfile
	GroupId    string // -DgroupId
	ArtifactId string // -DartifactId
	Version    string // -Dversion
	Packaging  string // -Dpackaging (e.g. "jar", "war", "pom", "aar", "ear")
	PomFile    string // -DpomFile (optional, replaces -DgroupId/ArtifactId/Version/Packaging)
	Classifier string // -Dclassifier (optional)
	Sources    string // -Dsources (optional)
	Javadoc    string // -Djavadoc (optional)
}

// ToArgs converts InstallFileOption to Maven command arguments
func (o *InstallFileOption) ToArgs() []string {
	args := []string{"install:install-file"}
	if o.File != "" {
		args = append(args, "-Dfile="+o.File)
	}
	if o.PomFile != "" {
		args = append(args, "-DpomFile="+o.PomFile)
	} else {
		if o.GroupId != "" {
			args = append(args, "-DgroupId="+o.GroupId)
		}
		if o.ArtifactId != "" {
			args = append(args, "-DartifactId="+o.ArtifactId)
		}
		if o.Version != "" {
			args = append(args, "-Dversion="+o.Version)
		}
		packaging := o.Packaging
		if packaging == "" {
			packaging = "jar" // default packaging
		}
		args = append(args, "-Dpackaging="+packaging)
	}
	if o.Classifier != "" {
		args = append(args, "-Dclassifier="+o.Classifier)
	}
	if o.Sources != "" {
		args = append(args, "-Dsources="+o.Sources)
	}
	if o.Javadoc != "" {
		args = append(args, "-Djavadoc="+o.Javadoc)
	}
	return args
}

// ArchetypeGenerateOption holds the options for the archetype:generate goal
type ArchetypeGenerateOption struct {
	ArchetypeGroupId    string // -DarchetypeGroupId
	ArchetypeArtifactId string // -DarchetypeArtifactId
	ArchetypeVersion    string // -DarchetypeVersion
	GroupId             string // -DgroupId (new project)
	ArtifactId          string // -DartifactId (new project)
	Version             string // -Dversion (new project)
	Package             string // -Dpackage (optional, defaults to groupId)
	OutputDirectory     string // -DoutputDirectory (optional)
	InteractiveMode     bool   // -DinteractiveMode=false for batch mode
}

// ToArgs converts ArchetypeGenerateOption to Maven command arguments
func (o *ArchetypeGenerateOption) ToArgs() []string {
	args := []string{"archetype:generate"}
	if o.ArchetypeGroupId != "" {
		args = append(args, "-DarchetypeGroupId="+o.ArchetypeGroupId)
	}
	if o.ArchetypeArtifactId != "" {
		args = append(args, "-DarchetypeArtifactId="+o.ArchetypeArtifactId)
	}
	if o.ArchetypeVersion != "" {
		args = append(args, "-DarchetypeVersion="+o.ArchetypeVersion)
	}
	if o.GroupId != "" {
		args = append(args, "-DgroupId="+o.GroupId)
	}
	if o.ArtifactId != "" {
		args = append(args, "-DartifactId="+o.ArtifactId)
	}
	if o.Version != "" {
		args = append(args, "-Dversion="+o.Version)
	}
	if o.Package != "" {
		args = append(args, "-Dpackage="+o.Package)
	}
	if o.OutputDirectory != "" {
		args = append(args, "-DoutputDirectory="+o.OutputDirectory)
	}
	if !o.InteractiveMode {
		args = append(args, "-DinteractiveMode=false")
	}
	return args
}
