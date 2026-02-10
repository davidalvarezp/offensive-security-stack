package web_config

import (
	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type MisconfigScanner struct {
	Reporter *reporting.ReportManager
}

func NewMisconfigScanner(r *reporting.ReportManager) *MisconfigScanner {
	return &MisconfigScanner{Reporter: r}
}

func (m *MisconfigScanner) Name() string { return "web-misconfig-scanner" }

func (m *MisconfigScanner) Description() string {
	return "Simulates scanning for common web misconfigurations"
}

func (m *MisconfigScanner) Run(ctx *core.AppContext) error {
	report := `
Web Configuration Misconfigurations:
- Directory listing enabled
- Weak TLS cipher suites allowed
- Default credentials detected (simulated)
`
	m.Reporter.AddReport("Web Misconfig Scan", m.Name(), report)
	ctx.Logger.Warn("Web configuration misconfigurations reviewed")
	return nil
}
