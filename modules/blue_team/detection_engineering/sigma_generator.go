package detection_engineering

import (
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type SigmaGeneratorModule struct {
	Reporter *reporting.ReportManager
}

func NewSigmaGeneratorModule(r *reporting.ReportManager) *SigmaGeneratorModule {
	return &SigmaGeneratorModule{Reporter: r}
}

func (m *SigmaGeneratorModule) Name() string { return "blue-sigma-generator" }

func (m *SigmaGeneratorModule) Description() string {
	return "Generates example Sigma rule structures (educational)"
}

func (m *SigmaGeneratorModule) Run(ctx *core.AppContext) error {
	rule := `
title: Suspicious PowerShell Execution
status: experimental
logsource:
  product: windows
  service: powershell
detection:
  selection:
    CommandLine|contains: "Invoke-Expression"
  condition: selection
`

	m.Reporter.AddReport("Sigma Rule Example", m.Name(), rule)
	ctx.Logger.Info("Sigma rule example generated")
	return nil
}
