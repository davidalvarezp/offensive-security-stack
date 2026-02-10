package reporting

import (
	"encoding/json"
	"fmt"
	"os"
)

// ExportJSON exports all reports as a JSON file
func (rm *ReportManager) ExportJSON(path string) error {
	data, err := json.MarshalIndent(rm.Reports, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal reports to JSON: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write JSON report file: %w", err)
	}

	return nil
}
