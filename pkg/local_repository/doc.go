// Package local_repository provides utilities for working with the Maven
// local repository (~/.m2/repository).
//
// It can locate JAR files by GAV coordinates, find directories, and
// build artifact paths:
//
//	repoDir := local_repository.ParseLocalRepositoryDirectory("mvn")
//	jarPath, err := local_repository.FindJar(repoDir, "org.springframework", "spring-core", "5.3.21")
//	sourcesJar, err := local_repository.FindJarWithClassifier(repoDir, "org.springframework", "spring-core", "5.3.21", "sources")
package local_repository
