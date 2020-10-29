package bond

import (
	"math/rand"
)

type Operator string

const (
	Plus  Operator = "plus"
	Minus          = "minus"
	Times          = "times"
)

type Parameters struct {
	Max         int
	ChoiceCount int
	Operators   []Operator
}

type Problem struct {
	A, B, C int
	Op      Operator
	Cs      []int
}

func (param *Parameters) NewProblem() (p Problem) {
	op := param.Operators[rand.Intn(len(param.Operators))]
	switch op {
	case Plus:
		p.Op = Plus
		// a: 0..20
		p.A = rand.Intn(param.Max + 1)
		maxB := param.Max + 1 - p.A
		// b: 0..(20-a)
		if maxB == 0 {
			p.B = 0
		} else {
			p.B = rand.Intn(maxB)
		}
		// c: 0..20
		p.C = p.A + p.B
	case Minus:
		p.Op = Minus
		// a: 0..20
		p.A = rand.Intn(param.Max + 1)
		// b: 0..A
		if p.A == 0 {
			p.B = 0
		} else {
			p.B = rand.Intn(p.A + 1)
		}
		p.C = p.A - p.B // 0..20
	}
	cs := map[int]bool{p.C: true}
	for len(cs) < param.ChoiceCount {
		c := rand.Intn(param.Max) + 1 // 1..20
		if _, ok := cs[c]; !ok {
			cs[c] = true
		}
	}
	p.Cs = make([]int, 0)
	for c := range cs {
		p.Cs = append(p.Cs, c)
	}
	rand.Shuffle(len(p.Cs), func(i, j int) {
		p.Cs[i], p.Cs[j] = p.Cs[j], p.Cs[i]
	})
	return
}

func (p Problem) Breakout() (a1, a2, b1, b2 int) {
	switch p.Op {
	case Plus:
		switch {
		case p.A < 10:
			a2 = p.A
			a1 = 0
		default:
			a1 = 10
			a2 = p.A - 10
		}
		switch {
		case a2+p.B <= 10:
			b1 = p.B
			b2 = 0
		default:
			b1 = 10 - a2
			b2 = p.B - b1
		}
	case Minus:
		a1 = p.A - p.B
		a2 = p.B
		b1 = p.A - a1
		b2 = 0
	}
	// No breakout for times.
	return
}
