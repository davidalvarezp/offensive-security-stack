package scenarios

import (
	"github.com/davidalvarezp/offensive-security-stack/internal/core"
	"github.com/davidalvarezp/offensive-security-stack/modules/labs/grading"
)

type AuthLab struct {
	ID        string
	Validator grading.Validator
}

func NewAuthLab() *AuthLab {
	return &AuthLab{
		ID: "auth-lab-101",
		Validator: &grading.SimpleValidator{
			RequiredKeys: []string{"password_policy", "hash_type"},
		},
	}
}

func (l *AuthLab) Run(ctx *core.AppContext, inputs map[string]string) grading.ValidationResult {
	ctx.Logger.Info("Running authentication lab (educational)")
	return l.Validator.Validate(inputs)
}

func (l *AuthLab) Description() string {
	return "Analyze authentication weaknesses and policy design"
}
