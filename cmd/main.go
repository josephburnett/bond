package main

import (
	"math/rand"
	"time"

	bond "github.com/josephburnett/bond/pkg"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var next func()
	var p bond.Problem
	var v *bond.HtmlView
	next = func() {
		p = bond.NewProblem()
		v = bond.NewHtmlView(p, next)
		v.Render()
	}
	next()
	select {}
}
