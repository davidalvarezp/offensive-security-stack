package dhcp_security

import (
	"strings"

	"github.com/your-org/offensive-security-stack/internal/core"
	"github.com/your-org/offensive-security-stack/internal/reporting"
)

type RogueDHCPDetector struct {
	Reporter *reporting.ReportManager
}

func NewRogueDHCPDetector(r *reporting.ReportManager) *RogueDHCPDetector {
	return &RogueDHCPDetector{Reporter: r}
}

func (m *RogueDHCPDetector) Name() string { return "network-rogue-dhcp-detector" }

func (m *RogueDHCPDetector) Description() string {
	return "Detects unauthorized DHCP servers (simulated)"
}

func (m *RogueDHCPDetector) Run(ctx *core.AppContext) error {
	report := `
Rogue DHCP Detection:
- DHCP offers from unknown MAC
- Conflicting gateway option

Mitigation:
- DHCP snooping
`
	m.Reporter.AddReport("Rogue DHCP Detection", m.Name(), report)
	ctx.Logger.Warn("Rogue DHCP server detected (simulated)")
	return nil
}
