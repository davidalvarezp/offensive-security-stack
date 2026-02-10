package grading

import "errors"

type ValidationResult struct {
	Passed  bool
	Message string
}

type Validator interface {
	Validate(inputs map[string]string) ValidationResult
}

type SimpleValidator struct {
	RequiredKeys []string
}

func (v *SimpleValidator) Validate(inputs map[string]string) ValidationResult {
	for _, key := range v.RequiredKeys {
		if _, ok := inputs[key]; !ok {
			return ValidationResult{
				Passed:  false,
				Message: "missing required input: " + key,
			}
		}
	}
	return ValidationResult{
		Passed:  true,
		Message: "validation successful",
	}
}

func EnsureAuthorizedLab(labMode bool) error {
	if !labMode {
		return errors.New("labs can only be run in lab mode")
	}
	return nil
}
