package main

import (
	"fmt"

	"github.com/scagogogo/mvn-skills/pkg/installer"
)

func main() {
	// Simple install: uses sensible defaults (latest pinned version, official
	// mirrors with regional fallbacks, SHA512 verification).
	mavenHome, err := installer.Install()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Maven已成功安装到: %s\n", mavenHome)
	fmt.Printf("可执行文件位于: %s/bin/mvn\n", mavenHome)
}
