package arp_security

import (
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type ARPTableMonitor struct {
	Reporter *reporting.ReportManager
}

func NewARPTableMonitor(r *reporting.ReportManager) *ARPTableMonitor {
	return &ARPTableMonitor{Reporter: r}
}

func (m *ARPTableMonitor) Name() string { return "network-arp-table-monitor" }

func (m *ARPTableMonitor) Description() string {
	return "Monitors ARP table changes (simulated)"
}

func (m *ARPTableMonitor) Run(ctx *core.AppContext) error {
	report := `
ARP Table Monitoring:
- IP: 192.168.1.1 → MAC changed
- IP: 192.168.1.50 → New entry detected

Risk:
ARP poisoning may indicate MITM attempts
`
	m.Reporter.AddReport("ARP Table Monitoring", m.Name(), report)
	ctx.Logger.Warn("ARP table anomalies detected (simulated)")
	return nil
}
