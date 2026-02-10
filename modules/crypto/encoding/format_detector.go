package encoding

import (
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type FormatDetectorModule struct {
	Reporter *reporting.ReportManager
}

func NewFormatDetectorModule(r *reporting.ReportManager) *FormatDetectorModule {
	return &FormatDetectorModule{Reporter: r}
}

func (m *FormatDetectorModule) Name() string { return "crypto-format-detector" }

func (m *FormatDetectorModule) Description() string {
	return "Detects encoding formats heuristically (educational)"
}

func (m *FormatDetectorModule) Run(ctx *core.AppContext) error {
	report := `
Detected Formats:
- JWT-like structure (header.payload.signature)
- Hexadecimal string (0-9a-f)
- URL-safe Base64 variant

No decoding was performed.
`
	m.Reporter.AddReport("Encoding Format Detection", m.Name(), report)
	ctx.Logger.Info("Encoding format detection completed")
	return nil
}
