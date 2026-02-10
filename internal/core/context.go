package core

import (
	"log"

	"github.com/davidalvarezp/offensive-security-stack/internal/config"
)

// ContextOptions are used to initialize the AppContext
type ContextOptions struct {
	Config *config.Config
	Logger *Logger
}

// AppContext holds all global application state and shared resources
type AppContext struct {
	Config *config.Config
	Logger *Logger
	Env    *Environment
}

// NewContext creates a new application context
func NewContext(opts ContextOptions) *AppContext {
	if opts.Config == nil {
		log.Panic("config cannot be nil")
	}
	if opts.Logger == nil {
		log.Panic("logger cannot be nil")
	}

	env := NewEnvironment(opts.Config)

	return &AppContext{
		Config: opts.Config,
		Logger: opts.Logger,
		Env:    env,
	}
}
