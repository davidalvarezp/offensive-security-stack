package access_control

import (
	"fmt"
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

// PrivilegeMap represents role-to-privilege relationships
type PrivilegeMap map[string][]string

// PrivilegeMapperModule teaches privilege modeling and review
type PrivilegeMapperModule struct {
	Reporter *reporting.ReportManager
}

func NewPrivilegeMapperModule(r *reporting.ReportManager) *PrivilegeMapperModule {
	return &PrivilegeMapperModule{Reporter: r}
}

func (m *PrivilegeMapperModule) Name() string {
	return "auth-privilege-mapper"
}

func (m *PrivilegeMapperModule) Description() string {
	return "Maps roles to privileges for access control analysis (educational)"
}

func (m *PrivilegeMapperModule) Run(ctx *core.AppContext) error {
	ctx.Logger.Info("Running privilege mapper module")

	privileges := PrivilegeMap{
		"Viewer": {"read_reports"},
		"Editor": {"read_reports", "edit_content"},
		"Admin":  {"read_reports", "edit_content", "manage_users"},
		"Root":   {"*"},
	}

	var b strings.Builder
	b.WriteString("Role → Privilege Mapping:\n\n")

	for role, perms := range privileges {
		b.WriteString(fmt.Sprintf("- %s:\n", role))
		for _, p := range perms {
			b.WriteString(fmt.Sprintf("  - %s\n", p))
		}
	}

	m.Reporter.AddReport(
		"Privilege Mapping Overview",
		m.Name(),
		b.String(),
	)

	ctx.Logger.Info("Privilege mapper module completed")
	return nil
}
