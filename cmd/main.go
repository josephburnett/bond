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

}
