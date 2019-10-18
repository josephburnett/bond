package bond

type Score struct {
	goal, current, wins int
	celebrate           bool
}

func NewScore() *Score {
	return &Score{
		goal: 20,
	}
}

func (s *Score) Correct() {
	s.celebrate = false
	s.current += 5
	if s.current >= s.goal {
		s.celebrate = true
		s.wins += 1
		s.current = 0
	}
}

func (s *Score) Incorrect() {
	s.celebrate = false
	s.current -= 1
	if s.current < 0 {
		s.current = 0
	}
}
