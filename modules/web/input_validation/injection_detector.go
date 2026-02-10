package input_validation

import (
	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type InjectionDetector struct {
	Reporter *reporting.ReportManager
}

func NewInjectionDetector(r *reporting.ReportManager) *InjectionDetector {
	return &InjectionDetector{Reporter: r}
}

func (m *InjectionDetector) Name() string { return "web-injection-detector" }

func (m *InjectionDetector) Description() string {
	return "Simulates detection of injection patterns in inputs"
}

func (m *InjectionDetector) Run(ctx *core.AppContext) error {
	report := `
Injection Detection Simulation:
- SQL-like input patterns detected
- XSS-like patterns highlighted
- No real injection performed
`
	m.Reporter.AddReport("Injection Detection", m.Name(), report)
	ctx.Logger.Warn("Injection patterns reviewed (simulated)")
	return nil
}
