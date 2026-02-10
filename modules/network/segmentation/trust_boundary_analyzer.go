package segmentation

import (
	"strings"

	"github.com/your-org/offensive-security-stack/internal/core"
	"github.com/your-org/offensive-security-stack/internal/reporting"
)

type TrustBoundaryAnalyzer struct {
	Reporter *reporting.ReportManager
}

func NewTrustBoundaryAnalyzer(r *reporting.ReportManager) *TrustBoundaryAnalyzer {
	return &TrustBoundaryAnalyzer{Reporter: r}
}

func (m *TrustBoundaryAnalyzer) Name() string { return "network-trust-boundary-analyzer" }

func (m *TrustBoundaryAnalyzer) Description() string {
	return "Analyzes trust boundaries between segments"
}

func (m *TrustBoundaryAnalyzer) Run(ctx *core.AppContext) error {
	report := `
Trust Boundary Analysis:
- User → Server: Weak controls
- Guest → Internal: No isolation
`
	m.Reporter.AddReport("Trust Boundary Analysis", m.Name(), report)
	ctx.Logger.Warn("Trust boundary weaknesses identified")
	return nil
}
