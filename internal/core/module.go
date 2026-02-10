package core

import "fmt"

// Module interface defines a standard module in OSS.
type Module interface {
	Name() string
	Description() string
	Run(ctx *AppContext) error
}

// BaseModule provides common fields for modules.
type BaseModule struct {
	moduleName string
	desc       string
}

// Name returns the module name.
func (m *BaseModule) Name() string {
	return m.moduleName
}

// Description returns the module description.
func (m *BaseModule) Description() string {
	return m.desc
}

// Run is a placeholder; override in real modules.
func (m *BaseModule) Run(ctx *AppContext) error {
	return fmt.Errorf("module %s does not implement Run()", m.moduleName)
}
