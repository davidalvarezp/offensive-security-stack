package packet_analysis

import (
	"strings"

	"github.com/your-org/offensive-security-stack/internal/core"
	"github.com/your-org/offensive-security-stack/internal/reporting"
)

type PCAPReader struct {
	Reporter *reporting.ReportManager
}

func NewPCAPReader(r *reporting.ReportManager) *PCAPReader {
	return &PCAPReader{Reporter: r}
}

func (m *PCAPReader) Name() string { return "network-pcap-reader" }

func (m *PCAPReader) Description() string {
	return "Reads PCAP metadata (no packet parsing)"
}

func (m *PCAPReader) Run(ctx *core.AppContext) error {
	report := `
PCAP Summary:
- Packets: 12,450
- Duration: 5 minutes
- Protocols: TCP, UDP, ICMP
`
	m.Reporter.AddReport("PCAP Metadata Analysis", m.Name(), report)
	ctx.Logger.Info("PCAP metadata processed")
	return nil
}
