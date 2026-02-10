package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// NewRootCommand initializes the root command for OSS.
func NewRootCommand(appCtx interface{}) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "oss",
		Short: "Offensive Security Stack (Educational, Lab-Only)",
		Long: `Offensive Security Stack (OSS) is an open-source educational framework 
for teaching offensive security concepts in controlled, isolated labs. 

Use subcommands for specific modules (recon, network, web, auth, crypto, malware, blue-team, labs, offensive-tools).`,
		Version: Version,
		// Run the interactive menu if no subcommand is given
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Launching OSS Interactive Menu...\n")
			ShowInteractiveMenu(cmd)
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	// ----------------------------
	// Global Flags
	// ----------------------------
	BindGlobalFlags(rootCmd)

	// ----------------------------
	// Help & Completion
	// ----------------------------
	rootCmd.SetHelpCommand(NewHelpCommand())
	rootCmd.AddCommand(NewCompletionCommand())

	// ----------------------------
	// Version
	// ----------------------------
	rootCmd.SetVersionTemplate("OSS Version: {{.Version}}\nCommit: " + GitCommit + "\nBuilt: " + BuildDate + "\n")

	// ----------------------------
	// Register Modules / Subcommands
	// ----------------------------
	registerModules(rootCmd)

	return rootCmd
}

// registerModules registers all top-level OSS modules.
// Each module is a subcommand corresponding to its functionality.
func registerModules(rootCmd *cobra.Command) {
	// Example module registration.
	// Each of these modules should implement a Cobra command with Run/RunE
	// In production, each module should live in modules/<module_name> and expose a GetCommand()
	modules := []struct {
		Name    string
		Command *cobra.Command
	}{
		{"recon", &cobra.Command{Use: "recon", Short: "Host & service reconnaissance"}},
		{"network", &cobra.Command{Use: "network", Short: "Network analysis & monitoring"}},
		{"web", &cobra.Command{Use: "web", Short: "Web application security analysis"}},
		{"auth", &cobra.Command{Use: "auth", Short: "Authentication & authorization analysis"}},
		{"crypto", &cobra.Command{Use: "crypto", Short: "Cryptography & entropy labs"}},
		{"malware", &cobra.Command{Use: "malware", Short: "Malware lifecycle simulation"}},
		{"post-exploitation", &cobra.Command{Use: "post-exploitation", Short: "Post-exploitation analysis"}},
		{"blue-team", &cobra.Command{Use: "blue-team", Short: "Detection & mitigation"}},
		{"labs", &cobra.Command{Use: "labs", Short: "Lab exercises & scenarios"}},
		{"offensive-tools", &cobra.Command{Use: "offensive-tools", Short: "Educational demo exploits"}},
	}

	for _, m := range modules {
		rootCmd.AddCommand(m.Command)
	}
}
