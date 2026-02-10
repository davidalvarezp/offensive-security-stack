package logging

import (
	"fmt"
	"time"
)

// Telemetry is a lightweight in-memory telemetry tracker for OSS modules.
// It collects metrics such as module execution time, usage counts, etc.
type Telemetry struct {
	ModuleRuns map[string]int
	RunTimes   map[string]time.Duration
}

// NewTelemetry initializes an empty telemetry tracker
func NewTelemetry() *Telemetry {
	return &Telemetry{
		ModuleRuns: make(map[string]int),
		RunTimes:   make(map[string]time.Duration),
	}
}

// RecordModuleRun logs a module execution
func (t *Telemetry) RecordModuleRun(moduleName string, duration time.Duration) {
	t.ModuleRuns[moduleName]++
	t.RunTimes[moduleName] += duration
}

// PrintSummary outputs a summary of telemetry
func (t *Telemetry) PrintSummary() {
	fmt.Println("=== OSS Telemetry Summary ===")
	for mod, runs := range t.ModuleRuns {
		totalTime := t.RunTimes[mod]
		avg := totalTime / time.Duration(runs)
		fmt.Printf("Module: %-20s | Runs: %-3d | Total: %-10s | Avg: %s\n",
			mod, runs, totalTime, avg)
	}
	fmt.Println("==============================")
}
