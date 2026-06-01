// Package finder locates Maven executables on the system.
//
// It supports three search strategies:
//   - System PATH and environment variables (M2_HOME, MAVEN_HOME)
//   - Maven Wrapper (mvnw/mvnw.cmd) in project directories
//   - Best-effort: prefers project Wrapper, falls back to system Maven
//
// Basic usage:
//
//	maven, err := finder.FindMaven()
//	maven, err := finder.FindBestMaven("/path/to/project")
package finder
