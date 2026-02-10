package tls_analysis

import (
	"strings"

	"github.com/your-org/offensive-security-stack/internal/core"
	"github.com/your-org/offensive-security-stack/internal/reporting"
)

type ProtocolDowngradeModule struct {
	Reporter *reporting.ReportManager
}

func NewProtocolDowngradeModule(r *reporting.ReportManager) *ProtocolDowngradeModule {
	return &ProtocolDowngradeModule{Reporter: r}
}

func (m *ProtocolDowngradeModule) Name() string { return "crypto-protocol-downgrade" }

func (m *ProtocolDowngradeModule) Description() string {
	return "Demonstrates TLS downgrade risks (educational)"
}

func (m *ProtocolDowngradeModule) Run(ctx *core.AppContext) error {
	report := `
Protocol Downgrade Scenario:
- Client supports TLS 1.0–1.3
- Server allows TLS 1.0
- Attacker forces downgrade

Mitigation:
Disable legacy protocol versions
`
	m.Reporter.AddReport("TLS Protocol Downgrade", m.Name(), report)
	ctx.Logger.Warn("Protocol downgrade risk demonstrated")
	return nil
}
