package access_control

import (
	"fmt"
	"time"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

// RoleAuditResult represents a single audit finding
type RoleAuditResult struct {
	Role  string
	Issue string
	Level string
}

// RoleAuditModule audits roles for common access control issues
type RoleAuditModule struct {
	Reporter *reporting.ReportManager
}

func NewRoleAuditModule(r *reporting.ReportManager) *RoleAuditModule {
	return &RoleAuditModule{Reporter: r}
}

func (m *RoleAuditModule) Name() string {
	return "auth-role-audit"
}

func (m *RoleAuditModule) Description() string {
	return "Audits roles for over-privilege and misconfiguration (educational)"
}

func (m *RoleAuditModule) Run(ctx *core.AppContext) error {
	ctx.Logger.Info("Running role audit module")

	findings := []RoleAuditResult{
		{
			Role:  "Admin",
			Issue: "Admin role assigned to non-privileged users",
			Level: "High",
		},
		{
			Role:  "ServiceAccount",
			Issue: "Service account has interactive login permissions",
			Level: "Medium",
		},
		{
			Role:  "Root",
			Issue: "Root role lacks separation of duties",
			Level: "Critical",
		},
	}

	report := "Role Audit Findings:\n\n"
	for _, f := range findings {
		report += fmt.Sprintf(
			"- Role: %s | Issue: %s | Severity: %s\n",
			f.Role, f.Issue, f.Level,
		)
	}

	m.Reporter.AddReport(
		"Role Configuration Audit",
		m.Name(),
		report,
	)

	ctx.Logger.Warnf(
		"Role audit completed with %d findings (%s)",
		len(findings),
		time.Now().UTC().Format(time.RFC3339),
	)

	return nil
}
