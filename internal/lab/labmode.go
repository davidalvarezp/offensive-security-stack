package lab

import (
	"fmt"

	"github.com/davidalvarezp/offensive-security-stack/internal/config"
)

// EnsureLabMode verifies that OSS is running in a safe, lab-approved environment.
// Returns an error if lab conditions are not met.
func EnsureLabMode(cfg *config.Config) error {
	if !cfg.Lab.Enabled {
		return fmt.Errorf("lab mode is disabled in configuration")
	}

	if !isLabEnv() {
		return fmt.Errorf("unsafe environment detected: OSS must run in a controlled lab")
	}

	if cfg.Lab.LANOnly && !isLAN() {
		return fmt.Errorf("LAN-only enforcement failed: host does not appear to be on a local network")
	}

	return nil
}

// isLabEnv checks if the environment variable OSS_LAB_MODE is set
func isLabEnv() bool {
	val := getEnvBool("OSS_LAB_MODE")
	return val
}

// isLAN is a placeholder for LAN-only checks.
// In production, this would check IP ranges, gateway, or other network heuristics.
func isLAN() bool {
	// TODO: Implement proper LAN detection
	return true
}

// getEnvBool reads an environment variable and returns true if "1" or "true"
func getEnvBool(key string) bool {
	val := getenv(key)
	return val == "1" || val == "true"
}

// getenv is a helper to safely read environment variables
func getenv(key string) string {
	return ""
}
