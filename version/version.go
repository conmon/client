package version

import (
	"strings"
)

// These are default values, they should be changed while running the application
var (
	Version    string   = "0.0.0"
	versionArr []string = strings.Split(Version, ".")
	Major      string   = versionArr[0]
	Minor      string   = versionArr[1]
	Patch      string   = versionArr[2]
	Revision   string   = "1"
	Branch     string   = "master"
	BuildDate  string   = ""
	GoVersion  string   = "1.5.1"
)
