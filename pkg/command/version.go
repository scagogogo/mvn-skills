package command

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Version gets the Maven version
func Version(executable string) (string, error) {
	return ExecForStdout(executable, "-v")
}

// ParseVersion parses the Maven version string and returns a MavenVersion struct.
// The version string typically looks like:
//
//	Apache Maven 3.9.11 (...)
//	Maven home: /usr/local/Cellar/maven/3.9.11/libexec
//	Java version: 17.0.8, vendor: Eclipse Adoptium
func ParseVersion(versionOutput string) (*MavenVersion, error) {
	lines := strings.Split(versionOutput, "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("empty version output")
	}

	// Extract version from first line, e.g. "Apache Maven 3.9.11 (abc123)"
	re := regexp.MustCompile(`Apache Maven\s+(\d+(?:\.\d+)*)`)
	matches := re.FindStringSubmatch(lines[0])
	if len(matches) < 2 {
		return nil, fmt.Errorf("could not parse Maven version from: %s", lines[0])
	}

	mv := &MavenVersion{Raw: matches[1]}
	parts := strings.Split(matches[1], ".")
	if len(parts) >= 1 {
		mv.Major, _ = strconv.Atoi(parts[0])
	}
	if len(parts) >= 2 {
		mv.Minor, _ = strconv.Atoi(parts[1])
	}
	if len(parts) >= 3 {
		mv.Patch, _ = strconv.Atoi(parts[2])
	}

	// Extract Maven home from second line if available
	if len(lines) >= 2 {
		homeRe := regexp.MustCompile(`Maven home:\s+(.+)`)
		homeMatches := homeRe.FindStringSubmatch(lines[1])
		if len(homeMatches) >= 2 {
			mv.Home = strings.TrimSpace(homeMatches[1])
		}
	}

	// Extract Java version from third line if available
	if len(lines) >= 3 {
		javaRe := regexp.MustCompile(`Java version:\s+(\S+)`)
		javaMatches := javaRe.FindStringSubmatch(lines[2])
		if len(javaMatches) >= 2 {
			mv.JavaVersion = javaMatches[1]
		}
	}

	return mv, nil
}

// MavenVersion represents a parsed Maven version with semantic comparison support
type MavenVersion struct {
	Major       int    // Major version number
	Minor       int    // Minor version number
	Patch       int    // Patch version number
	Raw         string // Raw version string (e.g. "3.9.11")
	Home        string // Maven home directory
	JavaVersion string // Java version string
}

// String returns the raw version string
func (v *MavenVersion) String() string {
	return v.Raw
}

// IsAtLeast returns true if this version is >= the specified version
func (v *MavenVersion) IsAtLeast(major, minor, patch int) bool {
	if v.Major != major {
		return v.Major > major
	}
	if v.Minor != minor {
		return v.Minor > minor
	}
	return v.Patch >= patch
}

// IsAtLeastMinor returns true if this version is >= the specified major.minor
func (v *MavenVersion) IsAtLeastMinor(major, minor int) bool {
	if v.Major != major {
		return v.Major > major
	}
	return v.Minor >= minor
}
