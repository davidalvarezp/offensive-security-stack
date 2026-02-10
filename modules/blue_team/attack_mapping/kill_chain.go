package attack_mapping

import (
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type KillChainStage struct {
	Stage       string
	Description string
	Defenses    []string
}

type KillChainModule struct {
	Reporter *reporting.ReportManager
}

func NewKillChainModule(r *reporting.ReportManager) *KillChainModule {
	return &KillChainModule{Reporter: r}
}

func (m *KillChainModule) Name() string { return "blue-kill-chain" }

func (m *KillChainModule) Description() string {
	return "Maps attacks to the cyber kill chain (educational)"
}

func (m *KillChainModule) Run(ctx *core.AppContext) error {
	stages := []KillChainStage{
		{"Reconnaissance", "Information gathering", []string{"Asset inventory", "Rate limiting"}},
		{"Weaponization", "Payload preparation", []string{"Email filtering"}},
		{"Delivery", "Payload delivery", []string{"Web proxy", "Spam detection"}},
		{"Exploitation", "Triggering vulnerability", []string{"EDR", "WAF"}},
		{"Installation", "Persistence setup", []string{"Application allowlisting"}},
		{"C2", "Command & control", []string{"DNS monitoring"}},
		{"Actions on Objectives", "Data access/exfiltration", []string{"DLP"}},
	}

	var b strings.Builder
	b.WriteString("Cyber Kill Chain Mapping:\n\n")

	for _, s := range stages {
		b.WriteString("- " + s.Stage + "\n")
		b.WriteString("  Description: " + s.Description + "\n")
		b.WriteString("  Defensive Controls:\n")
		for _, d := range s.Defenses {
			b.WriteString("   - " + d + "\n")
		}
		b.WriteString("\n")
	}

	m.Reporter.AddReport("Kill Chain Mapping", m.Name(), b.String())
	ctx.Logger.Info("Kill chain mapping completed")
	return nil
}
