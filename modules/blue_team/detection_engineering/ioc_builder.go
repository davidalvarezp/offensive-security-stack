package detection_engineering

import (
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type IOC struct {
	Type  string
	Value string
}

type IOCBuilderModule struct {
	Reporter *reporting.ReportManager
}

func NewIOCBuilderModule(r *reporting.ReportManager) *IOCBuilderModule {
	return &IOCBuilderModule{Reporter: r}
}

func (m *IOCBuilderModule) Name() string { return "blue-ioc-builder" }

func (m *IOCBuilderModule) Description() string {
	return "Builds example Indicators of Compromise (educational)"
}

func (m *IOCBuilderModule) Run(ctx *core.AppContext) error {
	iocs := []IOC{
		{"IP", "192.168.10.45"},
		{"Domain", "malicious.example"},
		{"Hash", "deadbeefdeadbeefdeadbeefdeadbeef"},
	}

	var b strings.Builder
	b.WriteString("IOC Examples:\n\n")

	for _, i := range iocs {
		b.WriteString("- " + i.Type + ": " + i.Value + "\n")
	}

	m.Reporter.AddReport("IOC Builder", m.Name(), b.String())
	ctx.Logger.Info("IOC builder completed")
	return nil
}
