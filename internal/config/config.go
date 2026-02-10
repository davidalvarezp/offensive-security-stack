package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config represents the full OSS configuration object
type Config struct {
	Logging LoggingConfig
	Lab     LabConfig
	App     AppConfig
}

// LoggingConfig holds logging-related settings
type LoggingConfig struct {
	Level       string
	File        string
	EnableAudit bool
}

// LabConfig holds lab-mode enforcement settings
type LabConfig struct {
	Enabled bool
	LANOnly bool
}

// AppConfig holds miscellaneous application settings
type AppConfig struct {
	DefaultProfile string
	ColorOutput    bool
	DryRun         bool
	Verbose        bool
	Quiet          bool
}

// Load loads the OSS configuration from:
// 1. CLI flags (via viper bindings)
// 2. Environment variables
// 3. Configuration file
// 4. Defaults (from defaults.go)
func Load() (*Config, error) {
	// Set defaults
	setDefaults()

	// Automatically read environment variables with prefix OSS_
	viper.SetEnvPrefix("OSS")
	viper.AutomaticEnv()

	// Read config file if provided
	configFile := viper.GetString("config")
	if configFile != "" {
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("failed to read config file %s: %w", configFile, err)
		}
	}

	cfg := &Config{
		Logging: LoggingConfig{
			Level:       viper.GetString("logging.level"),
			File:        viper.GetString("logging.file"),
			EnableAudit: viper.GetBool("logging.enable_audit"),
		},
		Lab: LabConfig{
			Enabled: viper.GetBool("lab.enabled"),
			LANOnly: viper.GetBool("lab.lan_only"),
		},
		App: AppConfig{
			DefaultProfile: viper.GetString("app.default_profile"),
			ColorOutput:    viper.GetBool("app.color_output"),
			DryRun:         viper.GetBool("app.dry_run"),
			Verbose:        viper.GetBool("app.verbose"),
			Quiet:          viper.GetBool("app.quiet"),
		},
	}

	// Safety check: prevent running outside of lab if lab enforcement enabled
	if cfg.Lab.Enabled && !isLabEnvironment() {
		return nil, fmt.Errorf("lab mode is enabled but environment is not safe")
	}

	return cfg, nil
}

// isLabEnvironment checks environment variables or other conditions
// to ensure OSS is running in a safe, isolated lab.
func isLabEnvironment() bool {
	val := os.Getenv("OSS_LAB_MODE")
	return val == "true" || val == "1"
}
