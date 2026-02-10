package main

import (
	"fmt"
	"runtime"
)

// Versioning variables.
// These are meant to be set at build time using -ldflags.
//
// Example:
// go build -ldflags "\
//   -X main.Version=1.0.0 \
//   -X main.GitCommit=$(git rev-parse --short HEAD) \
//   -X main.BuildDate=$(date -u +%Y-%m-%dT%H:%M:%SZ)"
var (
	// Version is the semantic version of the application.
	// Fallback is "dev" when not set via ldflags.
	Version = "dev"

	// GitCommit is the short commit hash.
	GitCommit = "unknown"

	// BuildDate is the UTC build timestamp (RFC3339).
	BuildDate = "unknown"
)

// VersionInfo returns a human-readable version string.
func VersionInfo() string {
	return fmt.Sprintf(
		"%s (commit: %s, built: %s, go: %s)",
		Version,
		GitCommit,
		BuildDate,
		runtime.Version(),
	)
}
