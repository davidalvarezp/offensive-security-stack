package reporting

import (
	"fmt"
	"time"
)

// Report represents a single report entry
type Report struct {
	Title      string
	ModuleName string
	Timestamp  time.Time
	Content    string
}

// ReportManager manages multiple reports
type ReportManager struct {
	Reports []*Report
}

// NewReportManager creates a new empty report manager
func NewReportManager() *ReportManager {
	return &ReportManager{
		Reports: []*Report{},
	}
}

// AddReport adds a new report entry
func (rm *ReportManager) AddReport(title, moduleName, content string) {
	report := &Report{
		Title:      title,
		ModuleName: moduleName,
		Timestamp:  time.Now().UTC(),
		Content:    content,
	}
	rm.Reports = append(rm.Reports, report)
}

// ListReports lists all report titles
func (rm *ReportManager) ListReports() {
	fmt.Println("=== OSS Reports ===")
	for i, r := range rm.Reports {
		fmt.Printf("%d) [%s] %s - %s\n", i+1, r.ModuleName, r.Title, r.Timestamp.Format(time.RFC3339))
	}
	fmt.Println("===================")
}
