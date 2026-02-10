package service_enum

import (
	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type ProtocolProbe struct {
	Reporter *reporting.ReportManager
}

func NewProtocolProbe(r *reporting.ReportManager) *ProtocolProbe {
	return &ProtocolProbe{Reporter: r}
}

func (m *ProtocolProbe) Name() string { return "recon-protocol-probe" }

func (m *ProtocolProbe) Description() string {
	return "Simulates protocol probing to identify services"
}

func (m *ProtocolProbe) Run(ctx *core.AppContext) error {
	report := `
Protocol Probe Simulation:
- TCP/22 → SSH
- TCP/80 → HTTP
- TCP/443 → HTTPS
`
	m.Reporter.AddReport("Protocol Probe", m.Name(), report)
	ctx.Logger.Info("Protocol probe simulated")
	return nil
}
