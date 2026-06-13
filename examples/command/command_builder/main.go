package main

import (
	"fmt"
	"log"

	"github.com/scagogogo/mvn-skills/pkg/command"
)

func main() {
	// CI/CD build using CommandBuilder
	output, err := command.NewCommandBuilder().
		WithExecutable("mvn").
		WithWorkingDirectory(".").
		WithBatchMode().
		WithNoTransferProgress().
		WithProfiles("ci").
		WithSkipTests().
		WithUpdateSnapshots().
		Clean()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)

	// Multi-module build with specific modules
	output, err = command.NewCommandBuilder().
		WithExecutable("mvn").
		WithProjects("module-a", "module-b").
		WithAlsoMake().
		WithBatchMode().
		Install()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
