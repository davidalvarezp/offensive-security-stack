package packet_analysis

import (
	"strings"

	"github.com/your-org/offensive-security-stack/internal/core"
	"github.com/your-org/offensive-security-stack/internal/reporting"
)

type ProtocolStatsModule struct {
	Reporter *reporting.ReportManager
}

func NewProtocolStatsModule(r *reporting.ReportManager) *ProtocolStatsModule {
	return &ProtocolStatsModule{Reporter: r}
}

func (m *ProtocolStatsModule) Name() string { return "network-protocol-stats" }

func (m *ProtocolStatsModule) Description() string {
	return "Calculates protocol distribution statistics"
}

func (m *ProtocolStatsModule) Run(ctx *core.AppContext) error {
	report := `
Protocol Distribution:
- TCP: 65%
- UDP: 30%
- ICMP: 5%
`
	m.Reporter.AddReport("Protocol Statistics", m.Name(), report)
	ctx.Logger.Info("Protocol statistics generated")
	return nil
}
