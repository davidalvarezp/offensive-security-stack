package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// NewCompletionCommand creates the `completion` command.
//
// This command generates shell completion scripts for OSS.
// Supported shells:
//   - bash
//   - zsh
//   - fish
//   - powershell
//
// Example usage:
//   oss completion bash
//   oss completion zsh
func NewCompletionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "completion [shell]",
		Short:                 "Generate shell completion scripts",
		Long:                  completionLongDescription(),
		DisableFlagsInUseLine: true,
		Args:                  cobra.ExactValidArgs(1),
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return generateCompletion(cmd, args[0])
		},
	}

	return cmd
}

// generateCompletion writes the completion script for the given shell to stdout.
func generateCompletion(cmd *cobra.Command, shell string) error {
	switch shell {
	case "bash":
		return cmd.Root().GenBashCompletion(os.Stdout)

	case "zsh":
		return cmd.Root().GenZshCompletion(os.Stdout)

	case "fish":
		return cmd.Root().GenFishCompletion(os.Stdout, true)

	case "powershell":
		return cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)

	default:
		return fmt.Errorf("unsupported shell: %s", shell)
	}
}

// completionLongDescription returns the long help text for the completion command.
func completionLongDescription() string {
	return `Generate shell completion scripts for Offensive Security Stack (OSS).

Shell completions improve the command-line experience by enabling:
  - Tab completion for commands and subcommands
  - Automatic flag suggestions
  - Inline command descriptions

Supported shells:
  - bash
  - zsh
  - fish
  - powershell

Installation examples:

Bash (current session):
  source <(oss completion bash)

Bash (persistent):
  oss completion bash > /etc/bash_completion.d/oss

Zsh:
  oss completion zsh > "${fpath[1]}/_oss"

Fish:
  oss completion fish > ~/.config/fish/completions/oss.fish

PowerShell:
  oss completion powershell > oss.ps1
  Then source the script from your PowerShell profile.

Note:
  You may need elevated privileges to install completions system-wide.
`
}
