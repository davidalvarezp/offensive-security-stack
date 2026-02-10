package hash_labs

import (
	"fmt"
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

// ParsedHash represents a parsed hash structure (conceptual)
type ParsedHash struct {
	Raw        string
	HasSalt   bool
	HasPrefix bool
	Format    string
}

// HashParserModule demonstrates parsing hash formats
type HashParserModule struct {
	Reporter *reporting.ReportManager
}

func NewHashParserModule(r *reporting.ReportManager) *HashParserModule {
	return &HashParserModule{Reporter: r}
}

func (m *HashParserModule) Name() string {
	return "auth-hash-parser"
}

func (m *HashParserModule) Description() string {
	return "Parses hash structures to explain components (educational)"
}

func (m *HashParserModule) Run(ctx *core.AppContext) error {
	ctx.Logger.Info("Running hash parser lab")

	sample := "$2b$12$abcdefghijklmnopqrstuvABCDEFGHIJKLMNO1234567890"

	parsed := ParsedHash{
		Raw:        sample,
		HasSalt:   strings.Contains(sample, "$"),
		HasPrefix: strings.HasPrefix(sample, "$2b$"),
		Format:    "bcrypt-style",
	}

	report := fmt.Sprintf(
		"Parsed Hash Example:\n\nRaw: %s\n\nDetected Properties:\n- Format: %s\n- Has Prefix: %t\n- Contains Salt: %t\n\n",
		parsed.Raw,
		parsed.Format,
		parsed.HasPrefix,
		parsed.HasSalt,
	)

	report += "⚠️ No cryptographic operations were performed.\n"
	report += "This lab demonstrates structural analysis only.\n"

	m.Reporter.AddReport(
		"Hash Structure Parsing",
		m.Name(),
		report,
	)

	ctx.Logger.Info("Hash parser lab completed")
	return nil
}
