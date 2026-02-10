package attack_mapping

import (
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type DetectionGap struct {
	Area  string
	Issue string
	Risk  string
}

type GapAnalysisModule struct {
	Reporter *reporting.ReportManager
}

func NewGapAnalysisModule(r *reporting.ReportManager) *GapAnalysisModule {
	return &GapAnalysisModule{Reporter: r}
}

func (m *GapAnalysisModule) Name() string { return "blue-gap-analysis" }

func (m *GapAnalysisModule) Description() string {
	return "Identifies detection and visibility gaps (educational)"
}

func (m *GapAnalysisModule) Run(ctx *core.AppContext) error {
	gaps := []DetectionGap{
		{"Authentication Logs", "No MFA event correlation", "High"},
		{"DNS Monitoring", "No anomaly detection", "Medium"},
		{"Endpoint Telemetry", "Limited script visibility", "High"},
	}

	var b strings.Builder
	b.WriteString("Detection Gap Analysis:\n\n")

	for _, g := range gaps {
		b.WriteString("- Area: " + g.Area + "\n")
		b.WriteString("  Issue: " + g.Issue + "\n")
		b.WriteString("  Risk: " + g.Risk + "\n\n")
	}

	m.Reporter.AddReport("Detection Gap Analysis", m.Name(), b.String())
	ctx.Logger.Warnf("Gap analysis completed with %d gaps", len(gaps))
	return nil
}
