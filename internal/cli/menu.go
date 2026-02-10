package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// ShowInteractiveMenu displays the OSS master menu.
// It lists all top-level commands (modules) and allows users to select one.
func ShowInteractiveMenu(rootCmd *cobra.Command) {
	for {
		modules := getModuleNames(rootCmd)
		modules = append(modules, "Exit") // always add exit option

		prompt := promptui.Select{
			Label: "Offensive Security Stack Modules",
			Items: modules,
			Size:  15,
		}

		index, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
			return
		}

		if result == "Exit" {
			fmt.Println("Exiting OSS...")
			return
		}

		selectedCmd := rootCmd.Commands()[index]
		if selectedCmd != nil {
			fmt.Printf("\n=== Module: %s ===\n\n", selectedCmd.Name())
			selectedCmd.Help() // show help for the selected module
		}

		fmt.Println("\nPress Enter to return to the main menu...")
		fmt.Scanln()
	}
}

// getModuleNames returns a slice of visible top-level command names
func getModuleNames(rootCmd *cobra.Command) []string {
	var modules []string
	for _, c := range rootCmd.Commands() {
		if !c.Hidden {
			name := strings.Title(c.Name())
			modules = append(modules, name)
		}
	}
	return modules
}
