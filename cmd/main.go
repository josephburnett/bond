package main

import (
	"fmt"
	"math/rand"
	"time"

	bond "github.com/josephburnett/bond/pkg"
)

func main() {
	fmt.Println("It's alive!")
	rand.Seed(time.Now().UTC().UnixNano())
	p := bond.NewProblem()
	v := bond.NewHtmlView(p)
	v.Render()
}
