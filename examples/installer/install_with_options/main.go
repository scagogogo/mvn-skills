package main

import (
	"fmt"

	"github.com/scagogogo/mvn-skills/pkg/installer"
)

func main() {
	// Configurable install: pin a version, use custom mirrors, skip env setup.
	opts := installer.InstallOptions{
		Version:      "3.9.11",
		Mirrors:      installer.DefaultMirrors,
		SkipEnvSetup: false,
		SkipChecksum: false,
		Force:        false,
		MaxRetries:   3,
	}

	mavenHome, err := installer.InstallWithOptions(opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Maven %s 安装完成: %s\n", opts.Version, mavenHome)
}
