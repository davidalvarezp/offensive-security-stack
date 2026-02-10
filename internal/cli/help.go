package cli

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// NewHelpCommand returns a customized help command.
// This replaces the default Cobra help command with:
// - Cleaner formatting
// - Module-aware sectioning
// - Optional extended help for demos/labs
func NewHelpCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "help [command]",
		Short:                 "Show help for OSS commands",
		Long:                  "Display usage, flags, and examples for Offensive Security Stack (OSS) commands.",
		DisableFlagsInUseLine: true,
		Args:                  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				showSubcommandHelp(cmd, args[0])
			} else {
				showRootHelp(cmd)
			}
		},
	}
	return cmd
}

// showRootHelp prints the root command help with module grouping
func showRootHelp(root *cobra.Command) {
	var b bytes.Buffer
	fmt.Fprintf(&b, "\n%s - %s\n\n", root.Name(), root.Short)
	fmt.Fprintf(&b, "Usage:\n  %s [command] [flags]\n\n", root.Use)
	fmt.Fprintf(&b, "Available Modules / Commands:\n")

	for _, c := range root.Commands() {
		// Skip hidden commands
		if c.Hidden {
			continue
		}
		name := c.Name()
		desc := strings.TrimSpace(c.Short)
		fmt.Fprintf(&b, "  %-20s %s\n", name, desc)
	}

	fmt.Fprintf(&b, "\nUse \"%s help [command]\" for more information on a specific command.\n\n", root.Name())
	os.Stdout.Write(b.Bytes())
}

// showSubcommandHelp prints detailed help for a given subcommand
func showSubcommandHelp(root *cobra.Command, name string) {
	sub := root.Commands()
	for _, c := range sub {
		if c.Name() == name {
			var b bytes.Buffer
			fmt.Fprintf(&b, "\nCommand: %s\n", c.Name())
			fmt.Fprintf(&b, "%s\n\n", c.Long)
			fmt.Fprintf(&b, "Usage:\n  %s\n\n", c.UseLine())

			if c.HasAvailableFlags() {
				fmt.Fprintf(&b, "Flags:\n")
				c.Flags().VisitAll(func(f *cobra.Flag) {
					def := f.DefValue
					if def == "" {
						def = "none"
					}
					fmt.Fprintf(&b, "  --%-15s %s (default: %s)\n", f.Name, f.Usage, def)
				})
			}

			if c.HasAvailableSubCommands() {
				fmt.Fprintf(&b, "\nSubcommands:\n")
				for _, sc := range c.Commands() {
					if sc.Hidden {
						continue
					}
					fmt.Fprintf(&b, "  %-20s %s\n", sc.Name(), sc.Short)
				}
			}

			fmt.Fprintf(&b, "\n")
			os.Stdout.Write(b.Bytes())
			return
		}
	}

	// Unknown command
	fmt.Fprintf(os.Stderr, "Unknown command: %s\n", name)
	fmt.Fprintf(os.Stderr, "Use \"%s help\" for a list of commands.\n", root.Name())
}
