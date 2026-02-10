package web_config

import (
	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type TLSConfig struct {
	Reporter *reporting.ReportManager
}

func NewTLSConfig(r *reporting.ReportManager) *TLSConfig {
	return &TLSConfig{Reporter: r}
}

func (m *TLSConfig) Name() string { return "web-tls-config" }

func (m *TLSConfig) Description() string {
	return "Simulates TLS configuration analysis"
}

func (m *TLSConfig) Run(ctx *core.AppContext) error {
	report := `
TLS Configuration Analysis:
- Supported Protocols: TLS 1.0, 1.2, 1.3
- Weak Cipher: TLS_RSA_WITH_3DES_EDE_CBC_SHA
- Recommended: Disable TLS 1.0 and weak ciphers
`
	m.Reporter.AddReport("TLS Config Analysis", m.Name(), report)
	ctx.Logger.Warn("TLS configuration analysis completed (simulated)")
	return nil
}
