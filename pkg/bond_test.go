package bond

import (
	"fmt"
	"testing"
)

func TestBreakout(t *testing.T) {
	cases := []struct {
		p              problem
		a1, a2, b1, b2 int
	}{{
		p: problem{a: 0, b: 0, op: plus},
	}, {
		p:  problem{a: 3, b: 3, op: plus},
		a2: 3,
		b1: 3,
	}, {
		p:  problem{a: 10, b: 10, op: plus},
		a1: 10,
		b1: 10,
	}, {
		p:  problem{a: 11, b: 9, op: plus},
		a1: 10,
		a2: 1,
		b1: 9,
	}, {
		p:  problem{a: 11, b: 10, op: plus},
		a1: 10,
		a2: 1,
		b1: 9,
		b2: 1,
	}, {
		p:  problem{a: 0, b: 1, op: plus},
		b1: 1,
	}, {
		p:  problem{a: 10, b: 10, op: minus},
		a2: 10,
		b1: 10,
	}, {
		p:  problem{a: 5, b: 5, op: minus},
		a2: 5,
		b1: 5,
	}, {
		p:  problem{a: 10, b: 5, op: minus},
		a1: 5,
		a2: 5,
		b1: 5,
	}, {
		p:  problem{a: 12, b: 2, op: minus},
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
