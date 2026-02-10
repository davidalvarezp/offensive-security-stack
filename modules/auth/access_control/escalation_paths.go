package access_control

import (
	"fmt"
	"time"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

// EscalationPath represents a theoretical privilege escalation path
// This is a DEMO STRUCTURE for teaching purposes only.
type EscalationPath struct {
	FromRole string
	ToRole   string
	Vector   string
	Risk     string
}

// EscalationPathsModule demonstrates common access-control escalation patterns
type EscalationPathsModule struct {
	Reporter *reporting.ReportManager
}

func NewEscalationPathsModule(r *reporting.ReportManager) *EscalationPathsModule {
	return &EscalationPathsModule{Reporter: r}
}

func (m *EscalationPathsModule) Name() string {
	return "auth-escalation-paths"
}

func (m *EscalationPathsModule) Description() string {
	return "Demonstrates theoretical access control escalation paths (educational)"
}

func (m *EscalationPathsModule) Run(ctx *core.AppContext) error {
	ctx.Logger.Info("Running access control escalation paths demo")

	paths := []EscalationPath{
		{
			FromRole: "User",
			ToRole:   "Admin",
			Vector:   "Insecure role assignment endpoint",
			Risk:     "High",
		},
		{
			FromRole: "Viewer",
			ToRole:   "Editor",
			Vector:   "Missing authorization check",
			Risk:     "Medium",
		},
		{
			FromRole: "ServiceAccount",
			ToRole:   "Root",
			Vector:   "Over-privileged API token",
			Risk:     "Critical",
		},
	}

	report := "Identified Theoretical Escalation Paths:\n\n"
	for _, p := range paths {
		report += fmt.Sprintf(
			"- %s → %s | Vector: %s | Risk: %s\n",
			p.FromRole, p.ToRole, p.Vector, p.Risk,
		)
	}

	m.Reporter.AddReport(
		"Access Control Escalation Paths",
		m.Name(),
		report,
	)

	ctx.Logger.Info("Escalation paths demo completed at %s", time.Now().UTC().Format(time.RFC3339))
	return nil
}
