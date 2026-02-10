package password_audit

import (
	"fmt"
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

// ReuseScenario represents a simulated password reuse case
type ReuseScenario struct {
	User        string
	Occurrences int
	Risk        string
}

// ReuseDetectorModule demonstrates password reuse detection (simulation)
type ReuseDetectorModule struct {
	Reporter *reporting.ReportManager
}

func NewReuseDetectorModule(r *reporting.ReportManager) *ReuseDetectorModule {
	return &ReuseDetectorModule{Reporter: r}
}

func (m *ReuseDetectorModule) Name() string {
	return "auth-password-reuse-detector"
}

func (m *ReuseDetectorModule) Description() string {
	return "Detects simulated password reuse patterns (educational)"
}

func (m *ReuseDetectorModule) Run(ctx *core.AppContext) error {
	ctx.Logger.Info("Running password reuse detector")

	// Simulated reuse data (no real passwords)
	scenarios := []ReuseScenario{
		{
			User:        "alice",
			Occurrences: 3,
			Risk:        "High",
		},
		{
			User:        "bob",
			Occurrences: 2,
			Risk:        "Medium",
		},
		{
			User:        "charlie",
			Occurrences: 1,
			Risk:        "Low",
		},
	}

	var report strings.Builder
	report.WriteString("Password Reuse Analysis (Simulated)\n\n")

	for _, s := range scenarios {
		report.WriteString(fmt.Sprintf(
			"- User: %s\n  Reuse Count: %d\n  Risk Level: %s\n\n",
			s.User, s.Occurrences, s.Risk,
		))
	}

	report.WriteString("⚠️ No real passwords were analyzed.\n")
	report.WriteString("This demo illustrates why reuse increases compromise impact.\n")

	m.Reporter.AddReport(
		"Password Reuse Detection",
		m.Name(),
		report.String(),
	)

	ctx.Logger.Warnf("Password reuse detector identified %d reuse scenarios", len(scenarios))
	return nil
}
