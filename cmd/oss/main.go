package main

import (
	"os"

	"github.com/spf13/cobra"

	// Internal packages
	"github.com/davidalvarezp/offensive-security-stack/internal/cli"
	"github.com/davidalvarezp/offensive-security-stack/internal/config"
	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/lab"
	"github.com/davidalvarezp/offensive-security-stack/internal/logging"
)

func main() {
	// ---------------------------------------------------------------------
	// 1. Initialize configuration
	// ---------------------------------------------------------------------
	cfg, err := config.Load()
	if err != nil {
		printFatal("failed to load configuration", err)
	}

	// ---------------------------------------------------------------------
	// 2. Initialize logging (audit-first)
	// ---------------------------------------------------------------------
	logger, err := logging.New(cfg.Logging)
	if err != nil {
		printFatal("failed to initialize logging", err)
	}
	defer logger.Sync()

	// ---------------------------------------------------------------------
	// 3. Enforce lab-only guardrails EARLY
	// ---------------------------------------------------------------------
	if err := lab.EnsureLabMode(cfg); err != nil {
		logger.Fatal("lab mode check failed", err)
	}

	// ---------------------------------------------------------------------
	// 4. Print banner (after safety checks)
	// ---------------------------------------------------------------------
	PrintBanner(Version)

	// ---------------------------------------------------------------------
	// 5. Initialize application core
	// ---------------------------------------------------------------------
	appCtx := core.NewContext(core.ContextOptions{
		Config: cfg,
		Logger: logger,
	})

	app := core.NewApp(appCtx)

	// ---------------------------------------------------------------------
	// 6. Initialize CLI (commands, flags, menu)
	// ---------------------------------------------------------------------
	rootCmd := cli.NewRootCommand(app)

	// Enable shell completion if requested
	rootCmd.CompletionOptions.DisableDefaultCmd = false

	// ---------------------------------------------------------------------
	// 7. Execute CLI
	// ---------------------------------------------------------------------
	if err := execute(rootCmd); err != nil {
		logger.Error("execution failed", err)
		os.Exit(1)
	}
}

// execute runs the root Cobra command.
func execute(cmd *cobra.Command) error {
	return cmd.Execute()
}

// printFatal prints a fatal startup error without relying on the logger.
// Used only before logging is initialized.
func printFatal(msg string, err error) {
	os.Stderr.WriteString("FATAL: " + msg + "\n")
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
	}
	os.Exit(1)
}
