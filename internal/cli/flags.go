package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Flag names (centralized to avoid duplication and typos)
const (
	flagConfigFile = "config"
	flagProfile    = "profile"
	flagVerbose    = "verbose"
	flagQuiet      = "quiet"
	flagNoColor    = "no-color"
	flagDryRun     = "dry-run"
)

// BindGlobalFlags defines and binds global CLI flags.
//
// These flags apply to the entire application and all modules.
// Configuration precedence (highest → lowest):
//   1. CLI flags
//   2. Environment variables
//   3. Config file
//   4. Defaults
func BindGlobalFlags(cmd *cobra.Command) {
	flags := cmd.PersistentFlags()

	// ------------------------------------------------------------------
	// Configuration & Profiles
	// ------------------------------------------------------------------

	flags.StringP(
		flagConfigFile,
		"c",
		"",
		"Path to configuration file (YAML, JSON, TOML)",
	)

	flags.StringP(
		flagProfile,
		"p",
		"default",
		"Configuration profile to use",
	)

	// ------------------------------------------------------------------
	// Output & UX
	// ------------------------------------------------------------------

	flags.BoolP(
		flagVerbose,
		"v",
		false,
		"Enable verbose output (debug-level logging)",
	)

	flags.Bool(
		flagQuiet,
		false,
		"Suppress non-essential output (errors only)",
	)

	flags.Bool(
		flagNoColor,
		false,
		"Disable colored output",
	)

	// ------------------------------------------------------------------
	// Execution Safety
	// ------------------------------------------------------------------

	flags.Bool(
		flagDryRun,
		false,
		"Show what would be executed without performing any actions",
	)

	// ------------------------------------------------------------------
	// Bind flags to Viper
	// ------------------------------------------------------------------

	mustBind(flagConfigFile, flags.Lookup(flagConfigFile))
	mustBind(flagProfile, flags.Lookup(flagProfile))
	mustBind(flagVerbose, flags.Lookup(flagVerbose))
	mustBind(flagQuiet, flags.Lookup(flagQuiet))
	mustBind(flagNoColor, flags.Lookup(flagNoColor))
	mustBind(flagDryRun, flags.Lookup(flagDryRun))
}

// mustBind binds a Cobra flag to Viper and panics on failure.
//
// This is acceptable here because:
// - Flags are defined at startup
// - Failure indicates a programming error
func mustBind(key string, flag *cobra.Flag) {
	if err := viper.BindPFlag(key, flag); err != nil {
		panic("failed to bind flag: " + key)
	}
}
