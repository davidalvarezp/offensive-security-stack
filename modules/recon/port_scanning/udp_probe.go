package port_scanning

import (
	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type UDPProbe struct {
	Reporter *reporting.ReportManager
}

func NewUDPProbe(r *reporting.ReportManager) *UDPProbe {
	return &UDPProbe{Reporter: r}
}

func (m *UDPProbe) Name() string { return "recon-udp-probe" }

func (m *UDPProbe) Description() string {
	return "Simulates UDP port probing"
}

func (m *UDPProbe) Run(ctx *core.AppContext) error {
	report := `
UDP Probe Simulation:
- 53: Open
- 123: Open
- 161: Open
`
	m.Reporter.AddReport("UDP Probe Simulation", m.Name(), report)
	ctx.Logger.Info("UDP probe simulated")
	return nil
}
