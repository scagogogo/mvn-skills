// Package settings provides a parser for Maven settings.xml configuration files.
//
// Basic usage:
//
//	settings, err := settings.ParseDefault()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, mirror := range settings.GetMirrors() {
//	    fmt.Printf("Mirror %s: %s -> %s\n", mirror.Id, mirror.MirrorOf, mirror.URL)
//	}
//
// ParseDefault() looks for settings.xml in ~/.m2/settings.xml first,
// then falls back to ${M2_HOME}/conf/settings.xml. All collection accessors
// return empty slices (not nil) when the corresponding elements are absent.
package settings
