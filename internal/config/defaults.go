package config

import (
	"github.com/spf13/viper"
)

// setDefaults sets all default configuration values
func setDefaults() {
	// Logging defaults
	viper.SetDefault("logging.level", "info")
	viper.SetDefault("logging.file", "")
	viper.SetDefault("logging.enable_audit", true)

	// Lab defaults
	viper.SetDefault("lab.enabled", true)
	viper.SetDefault("lab.lan_only", true)

	// Application defaults
	viper.SetDefault("app.default_profile", "default")
	viper.SetDefault("app.color_output", true)
	viper.SetDefault("app.dry_run", false)
	viper.SetDefault("app.verbose", false)
	viper.SetDefault("app.quiet", false)
}
