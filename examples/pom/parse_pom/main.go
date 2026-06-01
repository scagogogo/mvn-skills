package main

import (
	"fmt"
	"log"

	"github.com/scagogogo/mvn-sdk/pkg/pom"
)

func main() {
	project, err := pom.ParseFile("pom.xml")
	if err != nil {
		log.Fatal(err)
	}

	groupId, artifactId, version := project.GetGAV()
	fmt.Printf("Project: %s:%s:%s\n", groupId, artifactId, version)
	fmt.Printf("Packaging: %s\n", project.Packaging)
	fmt.Printf("Multi-module: %v\n", project.IsMultiModule())

	if project.HasParent() {
		fmt.Printf("Parent: %s:%s:%s\n",
			project.Parent.GroupId,
			project.Parent.ArtifactId,
			project.Parent.Version)
	}

	fmt.Println("\nDependencies:")
	for _, dep := range project.GetDependencies() {
		scope := dep.Scope
		if scope == "" {
			scope = "compile"
		}
		fmt.Printf("  %s:%s:%s (%s)\n", dep.GroupId, dep.ArtifactId, dep.Version, scope)
	}

	fmt.Println("\nModules:")
	for _, mod := range project.GetModules() {
		fmt.Printf("  %s\n", mod)
	}

	fmt.Println("\nPlugins:")
	for _, plugin := range project.GetPlugins() {
		fmt.Printf("  %s:%s:%s\n", plugin.GroupId, plugin.ArtifactId, plugin.Version)
	}
}
