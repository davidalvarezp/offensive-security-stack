package core

import "github.com/davidalvarezp/offensive-security-stack/internal/config"

// ModuleLifecycle handles the start/stop lifecycle of modules.
type ModuleLifecycle struct {
	App *App
}

// StartModule safely starts a module.
func (l *ModuleLifecycle) StartModule(name string) error {
	mod, ok := l.App.Modules[name]
	if !ok {
		return &ModuleNotFoundError{name}
	}

	// Check permissions before running
	if err := CheckPermissions(l.App.Context.Config, mod); err != nil {
		return err
	}

	// Enforce guardrails
	guard := NewGuardrails(l.App.Context.Config, l.App.Context.Env)
	if err := guard.EnforceGuardrails(); err != nil {
		return err
	}

	// Run module
	return mod.Run(l.App.Context)
}

// StopModule placeholder for future lifecycle handling
func (l *ModuleLifecycle) StopModule(name string) error {
	// For now, nothing to clean up
	return nil
}

// ModuleNotFoundError is returned when a module is missing
type ModuleNotFoundError struct {
	Name string
}

func (e *ModuleNotFoundError) Error() string {
	return "module not found: " + e.Name
}
