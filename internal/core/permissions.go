package core

import (
	"errors"
	"fmt"

	"github.com/davidalvarezp/offensive-security-stack/internal/config"
)

// CheckPermissions ensures the module is allowed to run under current config
func CheckPermissions(cfg *config.Config, mod Module) error {
	// Lab enforcement example
	if cfg.Lab.Enabled {
		// Example: block certain modules in dry-run mode
		if cfg.App.DryRun && mod.Name() == "malware" {
			return errors.New("malware simulation is blocked in dry-run mode")
		}
	}

	// Additional permission checks can be added here
	fmt.Printf("Permission check passed for module: %s\n", mod.Name())
	return nil
}
