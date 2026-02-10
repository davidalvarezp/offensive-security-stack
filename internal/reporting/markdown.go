package reporting

import (
	"fmt"
	"os"
)

// ExportMarkdown exports all reports to a Markdown file
func (rm *ReportManager) ExportMarkdown(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create Markdown report: %w", err)
	}
	defer f.Close()

	fmt.Fprintln(f, "# OSS Reports")
	fmt.Fprintln(f, "Generated:", rm.Reports[0].Timestamp.Format("2006-01-02 15:04:05 UTC"))

	for _, r := range rm.Reports {
		fmt.Fprintf(f, "\n## %s - %s\n", r.ModuleName, r.Title)
		fmt.Fprintf(f, "Timestamp: %s\n\n", r.Timestamp.Format(time.RFC3339))
		fmt.Fprintf(f, "```\n%s\n```\n", r.Content)
	}

	return nil
}
