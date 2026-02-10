package grading

type Score struct {
	MaxPoints   int
	Earned      int
	Penalties   int
	Completed   bool
}

type ScoreEngine struct {
	scores map[string]*Score
}

func NewScoreEngine() *ScoreEngine {
	return &ScoreEngine{
		scores: make(map[string]*Score),
	}
}

func (s *ScoreEngine) InitLab(labID string, max int) {
	s.scores[labID] = &Score{
		MaxPoints: max,
		Earned:    max,
	}
}

func (s *ScoreEngine) ApplyPenalty(labID string, points int) {
	if score, ok := s.scores[labID]; ok {
		score.Penalties += points
		score.Earned -= points
		if score.Earned < 0 {
			score.Earned = 0
		}
	}
}

func (s *ScoreEngine) CompleteLab(labID string) {
	if score, ok := s.scores[labID]; ok {
		score.Completed = true
	}
}

func (s *ScoreEngine) GetScore(labID string) *Score {
	return s.scores[labID]
}
