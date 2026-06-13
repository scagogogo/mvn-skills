package main

import (
	"fmt"
	"log"

	"github.com/scagogogo/mvn-skills/pkg/settings"
)

func main() {
	s, err := settings.ParseDefault()
	if err != nil {
		log.Fatalf("Failed to parse settings.xml: %v", err)
	}

	fmt.Printf("Local Repository: %s\n", s.LocalRepository)
	fmt.Printf("Offline Mode: %v\n", s.Offline)

	fmt.Println("\nMirrors:")
	for _, mirror := range s.GetMirrors() {
		fmt.Printf("  %s: %s -> %s (mirrorOf: %s)\n",
			mirror.Id, mirror.MirrorOf, mirror.URL, mirror.MirrorOf)
	}

	fmt.Println("\nServers:")
	for _, server := range s.GetServers() {
		fmt.Printf("  %s: %s\n", server.Id, server.Username)
	}

	fmt.Println("\nProxies:")
	for _, proxy := range s.GetProxies() {
		fmt.Printf("  %s (%s): %s:%d (active: %v)\n",
			proxy.Id, proxy.Protocol, proxy.Host, proxy.Port, proxy.Active)
	}

	fmt.Println("\nActive Profiles:")
	for _, id := range s.GetActiveProfileIds() {
		fmt.Printf("  %s\n", id)
	}

	// Find a specific server
	server := s.FindServer("github")
	if server != nil {
		fmt.Printf("\nGitHub server username: %s\n", server.Username)
	}

	// Find mirror for central repository
	mirror := s.FindMirrorOf("central")
	if mirror != nil {
		fmt.Printf("Mirror for central: %s (%s)\n", mirror.Id, mirror.URL)
	}
}
