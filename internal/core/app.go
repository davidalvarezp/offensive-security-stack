package core

// App is the main OSS application struct.
// It holds the context and orchestrates modules and commands.
type App struct {
	Context *AppContext
	Modules map[string]Module
}

// Module interface defines what a module must implement
type Module interface {
	Name() string
	Description() string
	Run() error
}

// NewApp initializes the main application
func NewApp(ctx *AppContext) *App {
	return &App{
		Context: ctx,
		Modules: make(map[string]Module),
	}
}

// RegisterModule adds a new module to the application
func (a *App) RegisterModule(m Module) {
	if _, exists := a.Modules[m.Name()]; exists {
		a.Context.Logger.Warnf("module %s already registered", m.Name())
		return
	}
	a.Modules[m.Name()] = m
	a.Context.Logger.Infof("module %s registered", m.Name())
}
