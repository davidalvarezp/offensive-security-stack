package grading

import "fmt"

type Hint struct {
	ID        string
	Message   string
	Penalty   int
	Unlocked  bool
}

type HintManager struct {
	hints map[string][]Hint
}

func NewHintManager() *HintManager {
	return &HintManager{
		hints: make(map[string][]Hint),
	}
}

func (h *HintManager) AddHint(labID string, hint Hint) {
	h.hints[labID] = append(h.hints[labID], hint)
}

func (h *HintManager) UnlockHint(labID string, index int) (*Hint, error) {
	if index < 0 || index >= len(h.hints[labID]) {
		return nil, fmt.Errorf("invalid hint index")
	}
	h.hints[labID][index].Unlocked = true
	return &h.hints[labID][index], nil
}

func (h *HintManager) ListHints(labID string) []Hint {
	return h.hints[labID]
}
