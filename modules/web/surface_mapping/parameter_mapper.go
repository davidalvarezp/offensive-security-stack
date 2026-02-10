package surface_mapping

import (
	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type ParameterMapper struct {
	Reporter *reporting.ReportManager
}

func NewParameterMapper(r *reporting.ReportManager) *ParameterMapper {
	return &ParameterMapper{Reporter: r}
}

func (m *ParameterMapper) Name() string { return "web-parameter-mapper" }

func (m *ParameterMapper) Description() string {
	return "Simulates mapping parameters for web endpoints"
}

func (m *ParameterMapper) Run(ctx *core.AppContext) error {
	report := `
Parameter Mapping:
- /login: username, password
- /api/v1/users: id, filter
- /search: q
`
	m.Reporter.AddReport("Parameter Mapping", m.Name(), report)
	ctx.Logger.Info("Parameters mapped (simulated)")
	return nil
}
