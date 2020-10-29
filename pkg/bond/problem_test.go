package bond

import (
	"fmt"
	"testing"
)

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
