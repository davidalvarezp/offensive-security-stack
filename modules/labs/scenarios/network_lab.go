package scenarios

import (
	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/modules/labs/grading"
)

type NetworkLab struct {
	ID        string
	Validator grading.Validator
}

func NewNetworkLab() *NetworkLab {
	return &NetworkLab{
		ID: "network-lab-201",
		Validator: &grading.SimpleValidator{
			RequiredKeys: []string{"segmentation", "firewall_rules"},
		},
	}
}

func (l *NetworkLab) Run(ctx *core.AppContext, inputs map[string]string) grading.ValidationResult {
	ctx.Logger.Info("Running network security lab (simulation)")
	return l.Validator.Validate(inputs)
}

func (l *NetworkLab) Description() string {
	return "Understand network segmentation and lateral movement prevention"
}
