package core

import (
	"errors"
	"fmt"

	"github.com/davidalvarezp/offensive-security-stack/internal/config"
)

// Guardrails enforces safety restrictions for OSS
type Guardrails struct {
	Config *config.Config
	Env    *Environment
}

// NewGuardrails initializes guardrails based on config and environment
func NewGuardrails(cfg *config.Config, env *Environment) *Guardrails {
	return &Guardrails{
		Config: cfg,
		Env:    env,
	}
}

// EnforceLabMode checks if OSS is running in a safe lab environment
func (g *Guardrails) EnforceLabMode() error {
	if !g.Config.Lab.Enabled {
		return errors.New("lab mode is not enabled in configuration")
	}
	if !g.Env.IsLab {
		return errors.New("unsafe environment detected; OSS must run in a lab")
	}
	if g.Config.Lab.LANOnly && g.Env.HostIP == "unknown" {
		return errors.New("cannot detect host IP; LAN-only enforcement failed")
	}
	return nil
}

// EnforceGuardrails runs all safety checks before executing modules
func (g *Guardrails) EnforceGuardrails() error {
	if err := g.EnforceLabMode(); err != nil {
		return fmt.Errorf("guardrail violation: %w", err)
	}
	// Future guardrails can be added here
	return nil
}
