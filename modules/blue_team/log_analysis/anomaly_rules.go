package log_analysis

import (
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type AnomalyRulesModule struct {
	Reporter *reporting.ReportManager
}

func NewAnomalyRulesModule(r *reporting.ReportManager) *AnomalyRulesModule {
	return &AnomalyRulesModule{Reporter: r}
}

func (m *AnomalyRulesModule) Name() string { return "blue-anomaly-rules" }

func (m *AnomalyRulesModule) Description() string {
	return "Defines behavioral anomaly rules (educational)"
}

func (m *AnomalyRulesModule) Run(ctx *core.AppContext) error {
	report := `
Anomaly Detection Rules:
- Login frequency > baseline
- Access outside business hours
- New device fingerprint
`
	m.Reporter.AddReport("Anomaly Detection Rules", m.Name(), report)
	ctx.Logger.Info("Anomaly rules defined")
	return nil
}
