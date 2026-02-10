package logging

import (
	"fmt"
	"os"
	"time"
)

// AuditLogger records a permanent log of actions/events in OSS
type AuditLogger struct {
	FilePath string
	file     *os.File
}

// NewAuditLogger initializes the audit logger and creates the file if needed
func NewAuditLogger(filePath string) (*AuditLogger, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		return nil, fmt.Errorf("failed to open audit log: %w", err)
	}

	return &AuditLogger{
		FilePath: filePath,
		file:     file,
	}, nil
}

// Record writes an audit entry with timestamp and level
func (a *AuditLogger) Record(level, message string) {
	entry := fmt.Sprintf("%s [%s] %s\n", time.Now().UTC().Format(time.RFC3339), level, message)
	if _, err := a.file.WriteString(entry); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write audit log: %v\n", err)
	}
}

// Close closes the audit file
func (a *AuditLogger) Close() error {
	if a.file != nil {
		return a.file.Close()
	}
	return nil
}
