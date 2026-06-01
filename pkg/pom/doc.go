// Package pom provides a parser for Maven POM (Project Object Model) XML files.
//
// Basic usage:
//
//	project, err := pom.ParseFile("pom.xml")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	groupId, artifactId, version := project.GetGAV()
//	fmt.Printf("Project: %s:%s:%s\n", groupId, artifactId, version)
//
// The parser handles XML namespaces correctly and supports GAV inheritance
// from parent POMs. All collection accessors return empty slices (not nil)
// when the corresponding elements are absent.
package pom
