package password_audit

import (
	"fmt"
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

// PasswordPolicy represents a conceptual password policy
type PasswordPolicy struct {
	MinLength      int
	RequireUpper   bool
	RequireLower   bool
	RequireNumber  bool
	RequireSymbol  bool
	DisallowReuse  bool
	Description   string
}

// PolicyFinding represents a policy weakness
type PolicyFinding struct {
	Rule  string
	Issue string
	Risk  string
}

// PolicyValidatorModule evaluates password policies (educational)
type PolicyValidatorModule struct {
	Reporter *reporting.ReportManager
}

func NewPolicyValidatorModule(r *reporting.ReportManager) *PolicyValidatorModule {
	return &PolicyValidatorModule{Reporter: r}
}

func (m *PolicyValidatorModule) Name() string {
	return "auth-password-policy-validator"
}

func (m *PolicyValidatorModule) Description() string {
	return "Validates password policies against best practices (educational)"
}

func (m *PolicyValidatorModule) Run(ctx *core.AppContext) error {
	ctx.Logger.Info("Running password policy validator")

	// Example policy (simulated)
	policy := PasswordPolicy{
		MinLength:     8,
		RequireUpper:  false,
		RequireLower:  true,
		RequireNumber: true,
		RequireSymbol: false,
		DisallowReuse: false,
		Description:  "Legacy internal policy",
	}

	var findings []PolicyFinding

	if policy.MinLength < 12 {
		findings = append(findings, PolicyFinding{
			Rule:  "Minimum Length",
			Issue: "Password length below recommended minimum (12)",
			Risk:  "Medium",
		})
	}

	if !policy.RequireUpper {
		findings = append(findings, PolicyFinding{
			Rule:  "Uppercase Requirement",
			Issue: "Uppercase characters not required",
			Risk:  "Low",
		})
	}

	if !policy.RequireSymbol {
		findings = append(findings, PolicyFinding{
			Rule:  "Symbol Requirement",
			Issue: "Special characters not required",
			Risk:  "Low",
		})
	}

	if !policy.DisallowReuse {
		findings = append(findings, PolicyFinding{
			Rule:  "Password Reuse",
			Issue: "Password reuse is permitted",
			Risk:  "High",
		})
	}

	var report strings.Builder
	report.WriteString("Password Policy Review\n\n")
	report.WriteString(fmt.Sprintf("Policy Description: %s\n\n", policy.Description))

	if len(findings) == 0 {
		report.WriteString("No policy issues identified.\n")
	} else {
		for _, f := range findings {
			report.WriteString(fmt.Sprintf(
				"- Rule: %s\n  Issue: %s\n  Risk: %s\n\n",
				f.Rule, f.Issue, f.Risk,
			))
		}
	}

	m.Reporter.AddReport(
		"Password Policy Validation",
		m.Name(),
		report.String(),
	)

	ctx.Logger.Warnf("Password policy validation completed with %d findings", len(findings))
	return nil
}
