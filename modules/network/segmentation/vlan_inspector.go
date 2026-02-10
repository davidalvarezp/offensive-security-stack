package segmentation

import (
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type VLANInspector struct {
	Reporter *reporting.ReportManager
}

func NewVLANInspector(r *reporting.ReportManager) *VLANInspector {
	return &VLANInspector{Reporter: r}
}

func (m *VLANInspector) Name() string { return "network-vlan-inspector" }

func (m *VLANInspector) Description() string {
	return "Inspects VLAN segmentation (educational)"
}

func (m *VLANInspector) Run(ctx *core.AppContext) error {
	report := `
VLAN Inspection:
- VLAN 10: Users
- VLAN 20: Servers
- VLAN 99: Management

Issue:
Management VLAN accessible from user VLAN
`
	m.Reporter.AddReport("VLAN Inspection", m.Name(), report)
	ctx.Logger.Warn("VLAN segmentation issue detected")
	return nil
}
