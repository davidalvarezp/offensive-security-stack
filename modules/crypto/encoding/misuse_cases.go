package encoding

import (
	"strings"

	"github.com/your-org/offensive-security-stack/internal/core"
	"github.com/your-org/offensive-security-stack/internal/reporting"
)

type EncodingMisuseModule struct {
	Reporter *reporting.ReportManager
}

func NewEncodingMisuseModule(r *reporting.ReportManager) *EncodingMisuseModule {
	return &EncodingMisuseModule{Reporter: r}
}

func (m *EncodingMisuseModule) Name() string { return "crypto-encoding-misuse" }

func (m *EncodingMisuseModule) Description() string {
	return "Demonstrates encoding misuse cases (educational)"
}

func (m *EncodingMisuseModule) Run(ctx *core.AppContext) error {
	report := `
Common Encoding Misuses:
- Base64 used as "encryption"
- URL encoding for obfuscation
- Hex encoding for secrets in config files

Impact:
- Secrets easily recoverable
`
	m.Reporter.AddReport("Encoding Misuse Cases", m.Name(), report)
	ctx.Logger.Warn("Encoding misuse cases reviewed")
	return nil
}
