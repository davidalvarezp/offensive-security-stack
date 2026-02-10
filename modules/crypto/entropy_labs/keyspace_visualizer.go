package entropy_labs

import (
	"fmt"
	"strings"

	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/internal/reporting"
)

type KeyspaceVisualizerModule struct {
	Reporter *reporting.ReportManager
}

func NewKeyspaceVisualizerModule(r *reporting.ReportManager) *KeyspaceVisualizerModule {
	return &KeyspaceVisualizerModule{Reporter: r}
}

func (m *KeyspaceVisualizerModule) Name() string { return "crypto-keyspace-visualizer" }

func (m *KeyspaceVisualizerModule) Description() string {
	return "Visualizes cryptographic keyspace sizes (educational)"
}

func (m *KeyspaceVisualizerModule) Run(ctx *core.AppContext) error {
	report := `
Keyspace Comparison:
- 4-digit PIN: 10^4
- 8-char lowercase password: 26^8
- 128-bit key: 2^128

Visualization:
PIN  | ####
PWD  | ####################
KEY  | ##################################################
`
	m.Reporter.AddReport("Keyspace Visualization", m.Name(), report)
	ctx.Logger.Info("Keyspace visualization generated")
	return nil
}
