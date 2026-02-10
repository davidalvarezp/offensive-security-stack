package attack_mapping

import (
	"fmt"
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type MitreTechnique struct {
	ID   string
	Name string
	Tactic string
}

type MitreMapperModule struct {
	Reporter *reporting.ReportManager
}

func NewMitreMapperModule(r *reporting.ReportManager) *MitreMapperModule {
	return &MitreMapperModule{Reporter: r}
}

func (m *MitreMapperModule) Name() string { return "blue-mitre-mapper" }

func (m *MitreMapperModule) Description() string {
	return "Maps events to MITRE ATT&CK techniques (educational)"
}

func (m *MitreMapperModule) Run(ctx *core.AppContext) error {
	techniques := []MitreTechnique{
		{"T1078", "Valid Accounts", "Persistence"},
		{"T1059", "Command and Scripting Interpreter", "Execution"},
		{"T1046", "Network Service Discovery", "Discovery"},
	}

	var b strings.Builder
	b.WriteString("MITRE ATT&CK Technique Mapping:\n\n")

	for _, t := range techniques {
		b.WriteString(fmt.Sprintf("- %s (%s)\n  Tactic: %s\n\n", t.ID, t.Name, t.Tactic))
	}

	m.Reporter.AddReport("MITRE ATT&CK Mapping", m.Name(), b.String())
	ctx.Logger.Info("MITRE mapping completed")
	return nil
}
