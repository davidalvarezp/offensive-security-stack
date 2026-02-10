package scenarios

import (
	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/modules/labs/grading"
)

type WebLab struct {
	ID        string
	Validator grading.Validator
}

func NewWebLab() *WebLab {
	return &WebLab{
		ID: "web-lab-301",
		Validator: &grading.SimpleValidator{
			RequiredKeys: []string{"input_validation", "session_controls"},
		},
	}
}

func (l *WebLab) Run(ctx *core.AppContext, inputs map[string]string) grading.ValidationResult {
	ctx.Logger.Info("Running web security lab (educational)")
	return l.Validator.Validate(inputs)
}

func (l *WebLab) Description() string {
	return "Explore web application security fundamentals"
}
