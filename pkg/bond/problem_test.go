package bond

import (
	"fmt"
	"testing"
)

func TestNewProblemFuzz(t *testing.T) {
	params := &Parameters{
		Max:         30,
		ChoiceCount: 3,
		Operators: []Operator{
			Plus,
			Minus,
			Times,
		},
	}
	for i := 0; i < 1000; i++ {
		p := params.NewProblem()
		if p.A > params.Max {
			t.Errorf("Wanted max A %v. Got %v", params.Max, p.A)
		}
		if p.B > params.Max {
			t.Errorf("Wanted max B %v. Got %v", params.Max, p.B)
		}
		if len(p.Cs) != params.ChoiceCount {
			t.Errorf("Wanted %v choices. Got %v", params.ChoiceCount, len(p.Cs))
		}
		if p.A < 0 || p.B < 0 || p.C < 0 {
			t.Errorf("Wanted non-negative A, B and C. Got %v, %v and %v", p.A, p.B, p.C)
		}
		switch p.Op {
		case Plus:
			if p.A+p.B != p.C {
				t.Errorf("Wanted a+b=c. Got %v+%v=%v", p.A, p.B, p.C)
			}
			maxC := params.Max + params.Max
			if p.C > maxC {
				t.Errorf("Wanted max C %v. Got %v", maxC, p.C)
			}
		case Minus:
			if p.A-p.B != p.C {
				t.Errorf("Wanted a-b=c. Got %v-%v=%v", p.A, p.B, p.C)
			}
			if p.C > params.Max {
				t.Errorf("Wanted max C %v. Got %v", params.Max, p.C)
			}
		case Times:
			if p.A*p.B != p.C {
				t.Errorf("Wanted a*b=c. Got %v*%v=%v", p.A, p.B, p.C)
			}
			maxC := params.Max * params.Max
			if p.C > maxC {
				t.Errorf("Wanted max C %v. Got %v", maxC, p.C)
			}
		}

	}
}

func TestBreakout(t *testing.T) {
	cases := []struct {
		p              Problem
		a1, a2, b1, b2 int
	}{{
		p: Problem{A: 0, B: 0, Op: Plus},
	}, {
		p:  Problem{A: 3, B: 3, Op: Plus},
		a2: 3,
		b1: 3,
	}, {
		p:  Problem{A: 10, B: 10, Op: Plus},
		a1: 10,
		b1: 10,
	}, {
		p:  Problem{A: 11, B: 9, Op: Plus},
		a1: 10,
		a2: 1,
		b1: 9,
	}, {
		p:  Problem{A: 11, B: 10, Op: Plus},
		a1: 10,
		a2: 1,
		b1: 9,
		b2: 1,
	}, {
		p:  Problem{A: 0, B: 1, Op: Plus},
		b1: 1,
	}, {
		p:  Problem{A: 10, B: 10, Op: Minus},
		a2: 10,
		b1: 10,
	}, {
		p:  Problem{A: 5, B: 5, Op: Minus},
		a2: 5,
		b1: 5,
	}, {
		p:  Problem{A: 10, B: 5, Op: Minus},
		a1: 5,
		a2: 5,
		b1: 5,
	}, {
		p:  Problem{A: 12, B: 2, Op: Minus},
		a1: 10,
		a2: 2,
		b1: 2,
	}}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%+v", c.p), func(t *testing.T) {
			a1, a2, b1, b2 := c.p.Breakout()
			if a1 != c.a1 {
				t.Errorf("want first number left %v. got %v", c.a1, a1)
			}
			if a2 != c.a2 {
				t.Errorf("want first number right %v. got %v", c.a2, a2)
			}
			if b1 != c.b1 {
				t.Errorf("want second number left %v. got %v", c.b1, b1)
			}
			if b2 != c.b2 {
				t.Errorf("want second number right %v. got %v", c.b2, b2)
			}
		})
	}
}
