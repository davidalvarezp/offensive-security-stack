package hash_labs

import (
	"fmt"
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

// HashSignature represents a known hash format pattern (educational)
type HashSignature struct {
	Name        string
	Length      int
	CharsetHint string
	Description string
}

// HashIdentifierModule demonstrates how hashes are classified
type HashIdentifierModule struct {
	Reporter *reporting.ReportManager
}

func NewHashIdentifierModule(r *reporting.ReportManager) *HashIdentifierModule {
	return &HashIdentifierModule{Reporter: r}
}

func (m *HashIdentifierModule) Name() string {
	return "auth-hash-identifier"
}

func (m *HashIdentifierModule) Description() string {
	return "Identifies hash types based on format (educational, non-operational)"
}

func (m *HashIdentifierModule) Run(ctx *core.AppContext) error {
	ctx.Logger.Info("Running hash identifier lab")

	signatures := []HashSignature{
		{
			Name:        "MD5",
			Length:      32,
			CharsetHint: "hex",
			Description: "Legacy fast hash, insecure",
		},
		{
			Name:        "SHA1",
			Length:      40,
			CharsetHint: "hex",
			Description: "Deprecated cryptographic hash",
		},
		{
			Name:        "SHA256",
			Length:      64,
			CharsetHint: "hex",
			Description: "Modern cryptographic hash",
		},
		{
			Name:        "bcrypt",
			Length:      60,
			CharsetHint: "base64-like",
			Description: "Adaptive password hash",
		},
	}

	exampleHash := "5f4dcc3b5aa765d61d8327deb882cf99"
	report := fmt.Sprintf("Analyzing hash: %s\n\n", exampleHash)

	for _, sig := range signatures {
		if len(exampleHash) == sig.Length {
			report += fmt.Sprintf(
				"Possible Match: %s\n- Charset: %s\n- Notes: %s\n\n",
				sig.Name, sig.CharsetHint, sig.Description,
			)
		}
	}

	if !strings.Contains(report, "Possible Match") {
		report += "No known hash signatures matched.\n"
	}

	m.Reporter.AddReport(
		"Hash Identification Demo",
		m.Name(),
		report,
	)

	ctx.Logger.Info("Hash identification lab completed")
	return nil
}
