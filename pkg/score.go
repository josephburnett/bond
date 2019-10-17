package bond

type Score struct {
	goal, current, wins int
}

func NewScore() *Score {
	return &Score{
		goal: 20,
	}
}

func (s *Score) Correct() {
	s.current += 5
	if s.current >= s.goal {
		s.wins += 1
		s.current = 0
	}
}

func (s *Score) Incorrect() {
	s.current -= 1
	if s.current < 0 {
		s.current = 0
	}
}
