package hash_labs

import (
	"fmt"
	"time"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

// CrackingScenario represents a simulated cracking scenario
type CrackingScenario struct {
	HashType   string
	Protection string
	Estimated  string
	Outcome    string
}

// CrackingSimulatorModule simulates cracking outcomes WITHOUT cracking
type CrackingSimulatorModule struct {
	Reporter *reporting.ReportManager
}

func NewCrackingSimulatorModule(r *reporting.ReportManager) *CrackingSimulatorModule {
	return &CrackingSimulatorModule{Reporter: r}
}

func (m *CrackingSimulatorModule) Name() string {
	return "auth-cracking-simulator"
}

func (m *CrackingSimulatorModule) Description() string {
	return "Simulates password cracking outcomes (NO real cracking)"
}

func (m *CrackingSimulatorModule) Run(ctx *core.AppContext) error {
	ctx.Logger.Warnf("Running cracking simulator (educational only)")

	scenarios := []CrackingScenario{
		{
			HashType:   "MD5",
			Protection: "None",
			Estimated:  "Milliseconds",
			Outcome:    "Trivially compromised",
		},
		{
			HashType:   "SHA1",
			Protection: "None",
			Estimated:  "Seconds",
			Outcome:    "Easily compromised",
		},
		{
			HashType:   "bcrypt",
			Protection: "Salt + Work Factor",
			Estimated:  "Years",
			Outcome:    "Impractical",
		},
	}

	report := "Password Cracking Simulation (NO real cracking):\n\n"

	for _, s := range scenarios {
		report += fmt.Sprintf(
			"- Hash: %s\n  Protection: %s\n  Estimated Effort: %s\n  Outcome: %s\n\n",
			s.HashType, s.Protection, s.Estimated, s.Outcome,
		)
	}

	report += "⚠️ This simulation uses theoretical estimates only.\n"
	report += "No hashes were attacked or reversed.\n"

	m.Reporter.AddReport(
		"Password Cracking Simulation",
		m.Name(),
		report,
	)

	ctx.Logger.Info(
		"Cracking simulator completed at %s",
		time.Now().UTC().Format(time.RFC3339),
	)

	return nil
}
