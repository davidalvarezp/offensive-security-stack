package surface_mapping

import (
	"github.com/your-org/offensive-security-stack/internal/core"
	"github.com/your-org/offensive-security-stack/internal/reporting"
)

type EndpointDiscovery struct {
	Reporter *reporting.ReportManager
}

func NewEndpointDiscovery(r *reporting.ReportManager) *EndpointDiscovery {
	return &EndpointDiscovery{Reporter: r}
}

func (m *EndpointDiscovery) Name() string { return "web-endpoint-discovery" }

func (m *EndpointDiscovery) Description() string {
	return "Simulates discovery of web endpoints"
}

func (m *EndpointDiscovery) Run(ctx *core.AppContext) error {
	report := `
Endpoint Discovery:
- /login
- /api/v1/users
- /admin
`
	m.Reporter.AddReport("Endpoint Discovery", m.Name(), report)
	ctx.Logger.Info("Web endpoints discovered (simulated)")
	return nil
}
