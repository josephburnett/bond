package main

import (
	"fmt"
	"math/rand"
	"time"

	bond "github.com/josephburnett/bond/pkg"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bond.NewScore()
	var v *bond.HtmlView
	eventHandler := func(e bond.Event) {
		switch e {
		case bond.CORRECT:
			fmt.Println("Correct.")
			s.Correct()
			v.SetProblem(bond.NewProblem())
			v.Render()
		case bond.INCORRECT:
			fmt.Println("Incorrect.")
			s.Incorrect()
			v.Render()
		default:
			fmt.Printf("Unhandled event: %v\n", e)
		}

	}
	v = bond.NewHtmlView(eventHandler)
	v.SetScore(s)
	v.SetProblem(bond.NewProblem())
	v.Render()
	select {}
}
