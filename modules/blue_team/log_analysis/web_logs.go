package log_analysis

import (
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type WebLogAnalysisModule struct {
	Reporter *reporting.ReportManager
}

func NewWebLogAnalysisModule(r *reporting.ReportManager) *WebLogAnalysisModule {
	return &WebLogAnalysisModule{Reporter: r}
}

func (m *WebLogAnalysisModule) Name() string { return "blue-web-log-analysis" }

func (m *WebLogAnalysisModule) Description() string {
	return "Analyzes web access logs (educational)"
}

func (m *WebLogAnalysisModule) Run(ctx *core.AppContext) error {
	report := `
Web Log Findings:
- Repeated 403 responses
- SQL injection-like patterns (simulated)
- Abnormal user-agent frequency
`
	m.Reporter.AddReport("Web Log Analysis", m.Name(), report)
	ctx.Logger.Info("Web log analysis completed")
	return nil
}
