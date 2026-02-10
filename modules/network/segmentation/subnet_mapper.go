package segmentation

import (
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type SubnetMapper struct {
	Reporter *reporting.ReportManager
}

func NewSubnetMapper(r *reporting.ReportManager) *SubnetMapper {
	return &SubnetMapper{Reporter: r}
}

func (m *SubnetMapper) Name() string { return "network-subnet-mapper" }

func (m *SubnetMapper) Description() string {
	return "Maps network subnets (educational)"
}

func (m *SubnetMapper) Run(ctx *core.AppContext) error {
	report := `
Subnet Mapping:
- 192.168.1.0/24 → Users
- 192.168.2.0/24 → Servers
- 192.168.3.0/24 → Guests
`
	m.Reporter.AddReport("Subnet Mapping", m.Name(), report)
	ctx.Logger.Info("Subnet mapping completed")
	return nil
}
