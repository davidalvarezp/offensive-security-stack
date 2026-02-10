package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoadProfile loads a named configuration profile.
// Profiles allow different sets of configurations for labs, exercises, or students.
func LoadProfile(name string) error {
	if name == "" || name == "default" {
		return nil // nothing to do
	}

	// Profiles can be stored under a subkey "profiles.<name>"
	profileKey := fmt.Sprintf("profiles.%s", name)
	if !viper.IsSet(profileKey) {
		return fmt.Errorf("profile %s not found in configuration", name)
	}

	sub := viper.Sub(profileKey)
	if sub == nil {
		return fmt.Errorf("failed to load profile %s", name)
	}

	// Merge profile values into global viper config
	for _, key := range sub.AllKeys() {
		viper.Set(key, sub.Get(key))
	}

	return nil
}
