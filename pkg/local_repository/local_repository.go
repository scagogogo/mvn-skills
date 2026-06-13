package local_repository

import (
	"errors"
	"fmt"
	"github.com/scagogogo/mvn-skills/pkg/command"
	"os"
	"path/filepath"
	"strings"
)

// DefaultLocalRepositoryDirectory is the default repository location, ${user.home}/.m2/repository
// Windows 7: C:/Documents and Settings/<username>/.m2/repository
// Windows 10: C:/Users/<username>/.m2/repository
// Linux: /home/<username>/.m2/repository
// Mac: /Users/<username>/.m2/repository
var DefaultLocalRepositoryDirectory string

func init() {
	dir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	DefaultLocalRepositoryDirectory = filepath.Join(dir, ".m2", "repository")
}

// ParseLocalRepositoryDirectory parses the local repository location
func ParseLocalRepositoryDirectory(executable string) string {

	// Try to find the repository location from the installed Maven
	directory, err := command.GetLocalRepositoryDirectory(executable)
	if err == nil && directory != "" {
		return directory
	}

	// If not found, return the default repository location
	return DefaultLocalRepositoryDirectory
}

// BuildDirectory constructs the relative path for a GAV
func BuildDirectory(groupId, artifactId, version string) string {
	return filepath.Join(strings.ReplaceAll(groupId, ".", "/"), artifactId, version)
}

// FindDirectory locates the GAV in the local repository
func FindDirectory(localRepositoryDirectory string, groupId, artifactId, version string) (string, error) {
	gavDirectory := filepath.Join(localRepositoryDirectory, BuildDirectory(groupId, artifactId, version))
	stat, err := os.Stat(gavDirectory)
	if err != nil {
		return "", err
	}
	if stat.IsDir() {
		return gavDirectory, nil
	} else {
		return "", errors.New("not a directory")
	}
}

// FindJar locates the JAR file for the given GAV in the local repository
func FindJar(localRepositoryDirectory string, groupId, artifactId, version string) (string, error) {
	return FindJarWithClassifier(localRepositoryDirectory, groupId, artifactId, version, "")
}

// FindJarWithClassifier locates the JAR file for the given GAV and classifier in the local repository
// When classifier is an empty string, it is equivalent to FindJar, locating the main artifact
// When classifier is non-empty, it locates the artifact with the classifier, such as "sources" or "javadoc"
func FindJarWithClassifier(localRepositoryDirectory string, groupId, artifactId, version, classifier string) (string, error) {
	directory, err := FindDirectory(localRepositoryDirectory, groupId, artifactId, version)
	if err != nil {
		return "", err
	}

	var jarPath string
	if classifier == "" {
		jarPath = filepath.Join(directory, fmt.Sprintf("%s-%s.jar", artifactId, version))
	} else {
		jarPath = filepath.Join(directory, fmt.Sprintf("%s-%s-%s.jar", artifactId, version, classifier))
	}

	stat, err := os.Stat(jarPath)
	if err != nil {
		return "", err
	}
	if !stat.IsDir() {
		return jarPath, nil
	} else {
		return "", errors.New("is directory, need a file")
	}
}
