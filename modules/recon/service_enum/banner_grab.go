package service_enum

import (
	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type BannerGrab struct {
	Reporter *reporting.ReportManager
}

func NewBannerGrab(r *reporting.ReportManager) *BannerGrab {
	return &BannerGrab{Reporter: r}
}

func (m *BannerGrab) Name() string { return "recon-banner-grab" }

func (m *BannerGrab) Description() string {
	return "Simulates service banner grabbing"
}

func (m *BannerGrab) Run(ctx *core.AppContext) error {
	report := `
Banner Grab Simulation:
- SSH 22: OpenSSH 8.2p1
- HTTP 80: Apache/2.4.52
- HTTPS 443: nginx 1.21
`
	m.Reporter.AddReport("Service Banner Grab", m.Name(), report)
	ctx.Logger.Info("Banner grab simulated")
	return nil
}
