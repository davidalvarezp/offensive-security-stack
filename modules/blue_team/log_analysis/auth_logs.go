package log_analysis

import (
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type AuthLogAnalysisModule struct {
	Reporter *reporting.ReportManager
}

func NewAuthLogAnalysisModule(r *reporting.ReportManager) *AuthLogAnalysisModule {
	return &AuthLogAnalysisModule{Reporter: r}
}

func (m *AuthLogAnalysisModule) Name() string { return "blue-auth-log-analysis" }

func (m *AuthLogAnalysisModule) Description() string {
	return "Analyzes authentication logs (educational)"
}

func (m *AuthLogAnalysisModule) Run(ctx *core.AppContext) error {
	report := `
Authentication Log Findings:
- Multiple failed logins followed by success
- Login from new country
- MFA challenge bypassed (simulated)
`
	m.Reporter.AddReport("Auth Log Analysis", m.Name(), report)
	ctx.Logger.Warnf("Authentication log analysis completed")
	return nil
}
