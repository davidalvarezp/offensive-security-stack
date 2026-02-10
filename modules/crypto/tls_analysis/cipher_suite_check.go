package tls_analysis

import (
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type CipherSuiteModule struct {
	Reporter *reporting.ReportManager
}

func NewCipherSuiteModule(r *reporting.ReportManager) *CipherSuiteModule {
	return &CipherSuiteModule{Reporter: r}
}

func (m *CipherSuiteModule) Name() string { return "crypto-cipher-suite-check" }

func (m *CipherSuiteModule) Description() string {
	return "Reviews TLS cipher suites (educational)"
}

func (m *CipherSuiteModule) Run(ctx *core.AppContext) error {
	report := `
Cipher Suite Review:
- TLS_RSA_WITH_3DES_EDE_CBC_SHA (deprecated)
- TLS_AES_128_GCM_SHA256 (recommended)

Recommendation:
Disable legacy cipher suites
`
	m.Reporter.AddReport("TLS Cipher Suite Review", m.Name(), report)
	ctx.Logger.Warn("Weak cipher suites identified (simulated)")
	return nil
}
