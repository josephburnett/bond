package view

import (
	"fmt"
	"strconv"
	"syscall/js"

	"github.com/josephburnett/bond/pkg/bond"
)

const (
	CLICKABLE = "#00b"
	Y_OFFSET  = 2
)

type Html struct {
	p         bond.Problem
	s         *bond.Score
	doc       js.Value
	svg       js.Value
	event     func(bond.Event)
	toRelease []js.Func
	hint      bool
}

func NewHtml(event func(bond.Event)) *Html {
	doc := js.Global().Get("document")
	svg := doc.Call("getElementById", "bond")
	svg.Call("setAttribute", "viewBox", "0 0 53 25")
	svg.Set("innerHTML", "")
	return &Html{
		doc:       doc,
		svg:       svg,
		event:     event,
		toRelease: make([]js.Func, 0),
	}
}

func (v *Html) SetProblem(p bond.Problem) {
	v.p = p
}

func (v *Html) SetScore(s *bond.Score) {
	v.s = s
}

func (v *Html) SetHint(h bool) {
	v.hint = h
}

func (v *Html) Render() {

	v.svg.Set("innerHTML", "")
	v.release()

	// First number
	v.circle("white", "black", 7, 5+Y_OFFSET, 3)
	v.text(5, 6+Y_OFFSET, strconv.Itoa(v.p.A), 3, "black")

	// Second number
	v.circle("white", "black", 23, 5+Y_OFFSET, 3)
	v.text(21, 6+Y_OFFSET, strconv.Itoa(v.p.B), 3, "black")

	// Operator
	switch v.p.Op {
	case bond.Minus:
		v.line(13, 5+Y_OFFSET, 17, 5+Y_OFFSET)
	case bond.Plus:
		v.line(13, 5+Y_OFFSET, 17, 5+Y_OFFSET)
		if v.p.Op == bond.Plus {
			v.line(15, 7+Y_OFFSET, 15, 3+Y_OFFSET)
		}
	case bond.Times:
		v.line(14, 6+Y_OFFSET, 16, 4+Y_OFFSET)
		v.line(16, 6+Y_OFFSET, 14, 4+Y_OFFSET)
	}

	switch {
	case !v.hint:
		h := v.text(14, 18+Y_OFFSET, "?", 5, CLICKABLE)
		h.Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			v.event(bond.HINT)
			return nil
		}))
	case v.p.Op == bond.Plus, v.p.Op == bond.Minus:
		a1, a2, b1, b2 := v.p.Breakout()

		// First number, part one
		if a1 != 0 {
			v.circle("#afa", "black", 4, 12+Y_OFFSET, 2)
			v.line(4, 10+Y_OFFSET, 6, 8+Y_OFFSET)
			v.text(3, 13+Y_OFFSET, strconv.Itoa(a1), 2, "black")
			v.numberLine("#afa", 2, 16+Y_OFFSET, a1, 0, 5)
		}

		// First number, part two
		if a2 != 0 {
			advance := 0
			if v.p.Op == bond.Minus {
				advance = a1
			}
			v.circle("#faa", "black", 10, 12+Y_OFFSET, 2)
			v.line(10, 10+Y_OFFSET, 8, 8+Y_OFFSET)
			v.text(9, 13+Y_OFFSET, strconv.Itoa(a2), 2, "black")
			v.numberLine("#faa", 8, 16+Y_OFFSET, a2, advance, 5)
		}

		// Second number, part one
		if b1 != 0 {
			advance := 0
			if v.p.Op == bond.Plus {
				advance = a2
			} else {
				advance = a1
			}
			v.circle("#faa", "black", 20, 12+Y_OFFSET, 2)
			v.line(20, 10+Y_OFFSET, 22, 8+Y_OFFSET)
			v.text(19, 13+Y_OFFSET, strconv.Itoa(b1), 2, "black")
			v.numberLine("#faa", 18, 16+Y_OFFSET, b1, advance, 5)
		}

		// Second number, part two
		if b2 != 0 {
			v.circle("#aaf", "black", 26, 12+Y_OFFSET, 2)
			v.line(26, 10+Y_OFFSET, 24, 8+Y_OFFSET)
			v.text(25, 13+Y_OFFSET, strconv.Itoa(b2), 2, "black")
			v.numberLine("#aaf", 24, 16+Y_OFFSET, b2, 0, 5)
		}
	case v.p.Op == bond.Times:
		colors := []string{
			"#afa",
			"#faa",
			"#aaf",
		}
		groups := v.p.Groups()
		for i, g := range groups {
			color := colors[i%3]
			for k := 0; k < g[0]; k++ {
				x := 4 + i*13
				y := 12 + Y_OFFSET + k*2
				v.numberLine(color, x, y, g[1], 0, 10)
			}

		}
	}

	// Equals
	v.line(29, 4+Y_OFFSET, 32, 4+Y_OFFSET)
	v.line(29, 6+Y_OFFSET, 32, 6+Y_OFFSET)

	// Choices
	for i, c := range v.p.Cs {
		a := c
		answer := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			v.answer(a)
			return nil
		})
		v.toRelease = append(v.toRelease, answer)
		r := v.circle("white", CLICKABLE, 37, 6*i+4+Y_OFFSET, 2)
		r.Set("onclick", answer)
		size := 2
		if c > 99 {
			size = 1
		}
		t := v.text(36, 6*i+5+Y_OFFSET, strconv.Itoa(c), size, CLICKABLE)
		t.Call("setAttribute", "id", fmt.Sprintf("answer-%v", i))
		t.Set("onclick", answer)
	}

	// Score
	r := v.doc.Call("createElementNS", "http://www.w3.org/2000/svg", "rect")
	r.Call("setAttribute", "x", 0)
	r.Call("setAttribute", "y", 0)
	r.Call("setAttribute", "height", 1)
	r.Call("setAttribute", "width", 1+int(float64(52)*(float64(v.s.Current)/(float64(v.s.Goal)))))
	r.Call("setAttribute", "fill", "green")
	v.svg.Call("appendChild", r)
	v.text(45, 3+Y_OFFSET, fmt.Sprintf("Score: %v", v.s.Wins), 1, "green")

	// Celebrate!
	if v.s.Celebrate {
		v.text(45, 8+Y_OFFSET, "+1", 5, "green")
	}
}

func (v *Html) release() {
	for _, fn := range v.toRelease {
		fn.Release()
	}
	v.toRelease = make([]js.Func, 0)
}

func (v *Html) answer(c int) {
	if c == v.p.C {
		v.event(bond.CORRECT)
	} else {
		v.event(bond.INCORRECT)
	}
}

func (v *Html) circle(fill, stroke string, x, y, r int) js.Value {
	c := v.doc.Call("createElementNS", "http://www.w3.org/2000/svg", "circle")
	c.Call("setAttribute", "cx", x)
	c.Call("setAttribute", "cy", y)
	c.Call("setAttribute", "r", r)
	c.Call("setAttribute", "style", fmt.Sprintf("fill: %v; stroke: %v; stroke-width: 0.25", fill, stroke))
	v.svg.Call("appendChild", c)
	return c
}

func (v *Html) text(x, y int, txt string, size int, color string) js.Value {
	t := v.doc.Call("createElementNS", "http://www.w3.org/2000/svg", "text")
	t.Call("setAttribute", "x", x)
	t.Call("setAttribute", "y", y)
	t.Call("setAttribute", "style", fmt.Sprintf("font-size: %vpx; fill: %v", size, color))
	t.Set("innerHTML", txt)
	v.svg.Call("appendChild", t)
	return t
}

func (v *Html) line(x1, y1, x2, y2 int) js.Value {
	l := v.doc.Call("createElementNS", "http://www.w3.org/2000/svg", "line")
	l.Call("setAttribute", "x1", x1)
	l.Call("setAttribute", "y1", y1)
	l.Call("setAttribute", "x2", x2)
	l.Call("setAttribute", "y2", y2)
	l.Call("setAttribute", "style", "stroke: black; stroke-width: 0.25")
	v.svg.Call("appendChild", l)
	return l
}

func (v *Html) numberLine(color string, x, y, count, skip int, maxCol int) {
	row, col := 0, 0
	advance := func() {
		col += 1
		if col == maxCol {
			col = 0
			row += 1
		}
	}
	for i := 0; i < skip; i++ {
		advance()
	}
	for count != 0 {
		r := v.doc.Call("createElementNS", "http://www.w3.org/2000/svg", "rect")
		r.Call("setAttribute", "x", x+col)
		r.Call("setAttribute", "y", y+row)
		r.Call("setAttribute", "height", 1)
		r.Call("setAttribute", "width", 1)
		r.Call("setAttribute", "fill", color)
		r.Call("setAttribute", "style", "stroke: black; stroke-width: 0.25")
		v.svg.Call("appendChild", r)
		count -= 1
		advance()
	}
}
