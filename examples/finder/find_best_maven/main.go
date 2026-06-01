package main

import (
	"fmt"
	"log"

	"github.com/scagogogo/mvn-sdk/pkg/command"
	"github.com/scagogogo/mvn-sdk/pkg/finder"
)

func main() {
	projectDir := "."

	// Check if project has Maven Wrapper
	if finder.HasMavenWrapper(projectDir) {
		fmt.Println("Project has Maven Wrapper")
	} else {
		fmt.Println("Project does not have Maven Wrapper")
	}

	// Find the best Maven executable (prefers Wrapper, falls back to system Maven)
	maven, err := finder.FindBestMaven(projectDir)
	if err != nil {
		log.Fatalf("No Maven found: %v", err)
	}
	fmt.Printf("Using Maven: %s\n", maven)

	// Run a command with the found Maven
	version, err := command.Version(maven)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Maven version: %s\n", version)
}
