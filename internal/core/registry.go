package core

import "fmt"

// ModuleRegistry holds all registered modules in OSS
type ModuleRegistry struct {
	modules map[string]Module
}

// NewModuleRegistry creates a new registry
func NewModuleRegistry() *ModuleRegistry {
	return &ModuleRegistry{
		modules: make(map[string]Module),
	}
}

// Register adds a module to the registry
func (r *ModuleRegistry) Register(mod Module) error {
	if _, exists := r.modules[mod.Name()]; exists {
		return fmt.Errorf("module already registered: %s", mod.Name())
	}
	r.modules[mod.Name()] = mod
	return nil
}

// Get retrieves a module by name
func (r *ModuleRegistry) Get(name string) (Module, error) {
	mod, ok := r.modules[name]
	if !ok {
		return nil, fmt.Errorf("module not found: %s", name)
	}
	return mod, nil
}

// List returns all registered modules
func (r *ModuleRegistry) List() []Module {
	mods := []Module{}
	for _, m := range r.modules {
		mods = append(mods, m)
	}
	return mods
}
