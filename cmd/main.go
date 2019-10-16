package main

import (
	"fmt"
	"math/rand"
	"time"

	bond "github.com/josephburnett/bond/pkg"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	p := bond.NewProblem()
	fmt.Printf("%+v\n", p)
	a1, a2, b1, b2 := p.Breakout()
	fmt.Printf("%v %v %v %v\n", a1, a2, b1, b2)
}
