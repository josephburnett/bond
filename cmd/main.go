package main

import (
	"fmt"
	"math/rand"
	"time"

	bond "github.com/josephburnett/bond/pkg"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var v *bond.HtmlView
	eventHandler := func(e bond.Event) {
		switch e {
		case bond.CORRECT:
			fmt.Println("Correct.")
			p := bond.NewProblem()
			v.SetProblem(p)
			v.Render()
		case bond.INCORRECT:
			fmt.Println("Incorrect.")
		default:
			fmt.Printf("Unhandled event: %v\n", e)
		}

	}
	v = bond.NewHtmlView(eventHandler)
	p := bond.NewProblem()
	v.SetProblem(p)
	v.Render()
	select {}
}
