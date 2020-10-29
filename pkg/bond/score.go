package bond

type Score struct {
	Goal, Current, Wins int
	Celebrate           bool
}

func NewScore() *Score {
	return &Score{
		Goal: 20,
	}
}

func (s *Score) Correct() {
	s.Celebrate = false
	s.Current += 5
	if s.Current >= s.Goal {
		s.Celebrate = true
		s.Wins += 1
		s.Current = 0
	}
}

func (s *Score) Incorrect() {
	s.Celebrate = false
	s.Current -= 1
	if s.Current < 0 {
		s.Current = 0
	}
}
