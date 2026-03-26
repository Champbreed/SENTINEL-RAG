package audit

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// AuditResult represents the groundedness data from the RAI Toolbox
type AuditResult struct {
	Target    string  `json:"target"`
	Score     float64 `json:"score"`
	Compliant bool    `json:"compliant"`
}

// FetchRAIData bridges the Python RAI backend with our Go TUI
func FetchRAIData() ([]AuditResult, error) {
	// Using the venv path directly ensures we use the patched environment
	cmd := exec.Command("./venv/bin/python3", "scripts/rai_audit.py")
	
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("RAI Backend Error: %w", err)
	}

	var results []AuditResult
	if err := json.Unmarshal(stdout, &results); err != nil {
		return nil, fmt.Errorf("JSON Parsing Error: %w", err)
	}

	return results, nil
}
