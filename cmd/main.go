package main

import (
	"fmt"
	"math/rand"
	"time"

	bond "github.com/josephburnett/bond/pkg/bond"
	view "github.com/josephburnett/bond/pkg/view"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	params := bond.Parameters{
		Max:         30,
		MaxTimes:    12,
		ChoiceCount: 5,
		Operators: []bond.Operator{
			bond.Plus,
			bond.Minus,
			bond.Times,
		},
	}
	s := bond.NewScore()
	var v *view.Html
	eventHandler := func(e bond.Event) {
		switch e {
		case bond.CORRECT:
			s.Correct()
			v.SetProblem(params.NewProblem())
			v.SetHint(false)
			v.Render()
		case bond.INCORRECT:
			s.Incorrect()
			v.Render()
		case bond.HINT:
			v.SetHint(true)
			v.Render()
		default:
			fmt.Printf("Unhandled event: %v\n", e)
		}

	}
	v = view.NewHtml(eventHandler)
	v.SetScore(s)
	v.SetProblem(params.NewProblem())
	v.Render()
	select {}
}
