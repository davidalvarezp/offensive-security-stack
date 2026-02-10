package logging

import (
	"fmt"
	"log"
	"os"
)

// Logger wraps the standard Go logger with levels and optional file output.
type Logger struct {
	Verbose    bool
	Quiet      bool
	NoColor    bool
	Audit      *AuditLogger
	stdLogger  *log.Logger
}

// NewLogger initializes a Logger with optional audit support.
func NewLogger(verbose, quiet, noColor bool, audit *AuditLogger) *Logger {
	l := &Logger{
		Verbose:   verbose,
		Quiet:     quiet,
		NoColor:   noColor,
		Audit:     audit,
		stdLogger: log.New(os.Stdout, "", log.LstdFlags),
	}
	return l
}

// Info prints informational messages
func (l *Logger) Info(msg string, args ...interface{}) {
	if l.Quiet {
		return
	}
	l.printf("INFO", msg, args...)
}

// Warn prints warning messages
func (l *Logger) Warnf(msg string, args ...interface{}) {
	if l.Quiet {
		return
	}
	l.printf("WARN", msg, args...)
}

// Error prints error messages
func (l *Logger) Error(msg string, args ...interface{}) {
	l.printf("ERROR", msg, args...)
}

// Debug prints debug messages only if verbose is enabled
func (l *Logger) Debug(msg string, args ...interface{}) {
	if l.Verbose {
		l.printf("DEBUG", msg, args...)
	}
}

// printf formats and outputs a log message
func (l *Logger) printf(level, msg string, args ...interface{}) {
	formatted := fmt.Sprintf(msg, args...)
	if !l.NoColor {
		switch level {
		case "INFO":
			formatted = "\033[34m" + formatted + "\033[0m"
		case "WARN":
			formatted = "\033[33m" + formatted + "\033[0m"
		case "ERROR":
			formatted = "\033[31m" + formatted + "\033[0m"
		case "DEBUG":
			formatted = "\033[36m" + formatted + "\033[0m"
		}
	}
	l.stdLogger.Println(formatted)

	// Audit log if enabled
	if l.Audit != nil {
		l.Audit.Record(level, formatted)
	}
}
