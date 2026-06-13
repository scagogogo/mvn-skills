package command

// ArchetypeCreate creates a project
// executable: the path to the mvn executable
// directory: the location where the project will be generated
// groupId / artifactId /
func ArchetypeCreate(executable string, directory string, groupId, artifactId, version string) (string, error) {
	// Maven 3.0.5+ deprecated create; use generate to create a project
	// mvn archetype:create -DgroupId=org.sonatype.mavenbook.ch03 -DartifactId=simple -DpackageName=org.sonatype.mavenbook
	return ExecForStdout(executable, "archetype:generate", "-DoutputDirectory="+directory, "-DgroupId="+groupId, "-DartifactId="+artifactId, "-DarchetypeVersion="+version, "-DinteractiveMode=false")
}
