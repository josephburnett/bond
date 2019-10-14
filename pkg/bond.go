package bond

import (
	"math/rand"
)

const (
	maxSum      = 20
	choiceCount = 3
)

type operator string

const (
	plus  operator = "plus"
	minus operator = "minus"
)

type problem struct {
	a, b, c int
	op      operator
	cs      []int
}

func NewProblem() (p problem) {
	if rand.Intn(2) == 0 {
		p.op = plus
		p.a = rand.Intn(maxSum) + 1     // 1..20
		p.b = rand.Intn(maxSum-p.a) + 1 // 1..(20-a)
		p.c = p.a + p.b                 // 2..20
	} else {
		p.op = minus
		p.a = rand.Intn(maxSum-2) + 2 // 2..20
		p.b = rand.Intn(p.a-1) + 1    // 1..(a-1)
		p.c = p.a - p.b               // 1..19
	}
	cs := map[int]bool{p.c: true}
	for len(cs) < choiceCount {
		c := rand.Intn(maxSum) + 1 // 1..20
		if _, ok := cs[c]; !ok {
			cs[c] = true
		}
	}
	p.cs = make([]int, 0)
	for c := range cs {
		p.cs = append(p.cs, c)
	}
	rand.Shuffle(len(p.cs), func(i, j int) {
		p.cs[i], p.cs[j] = p.cs[j], p.cs[i]
	})
	return
}
