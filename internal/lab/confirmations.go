package lab

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

// ConfirmAction asks the user for explicit confirmation before performing an action.
// Returns true if confirmed, false otherwise.
func ConfirmAction(promptMessage string) bool {
	prompt := promptui.Prompt{
		Label:     promptMessage + " (yes/no)",
		IsConfirm: true,
		Default:   "no",
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Confirmation canceled or failed:", err)
		return false
	}

	confirmed := result == "y" || result == "yes"
	if !confirmed {
		fmt.Println("Action canceled by user.")
	}
	return confirmed
}
