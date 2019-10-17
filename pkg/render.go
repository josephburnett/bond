package bond

import (
	"fmt"
	"strconv"
	"syscall/js"
)

type Event string

const (
	CORRECT   Event = "correct"
	INCORRECT Event = "incorrect"
	HINT      Event = "hint"
)

type HtmlView struct {
	p         Problem
	score     Score
	doc       js.Value
	svg       js.Value
	event     func(Event)
	toRelease []js.Func
}

func NewHtmlView(event func(Event)) *HtmlView {
	doc := js.Global().Get("document")
	svg := doc.Call("getElementById", "bond")
	svg.Call("setAttribute", "viewBox", "0 0 50 25")
	svg.Set("innerHTML", "")
	return &HtmlView{
		doc:       doc,
		svg:       svg,
		event:     event,
		toRelease: make([]js.Func, 0),
	}
}

func (v *HtmlView) SetProblem(p Problem) {
	v.p = p
}

func (v *HtmlView) Render() {

	v.svg.Set("innerHTML", "")
	v.release()

	// First number
	v.circle("white", 7, 5, 3)
	v.text(5, 6, strconv.Itoa(v.p.a), 3)

	// Second number
	v.circle("white", 23, 5, 3)
	v.text(21, 6, strconv.Itoa(v.p.b), 3)

	// Operator
	v.line(13, 5, 17, 5)
	if v.p.op == plus {
		v.line(15, 7, 15, 3)
	}

	a1, a2, b1, b2 := v.p.Breakout()

	// First number, part one
	if a1 != 0 {
		v.circle("#afa", 4, 12, 2)
		v.line(4, 10, 6, 8)
		v.text(3, 13, strconv.Itoa(a1), 2)
		v.numberLine("#afa", 2, 16, a1, 0)
	}

	// First number, part two
	if a2 != 0 {
		advance := 0
		if v.p.op == minus {
			advance = a1
		}
		v.circle("#faa", 10, 12, 2)
		v.line(10, 10, 8, 8)
		v.text(9, 13, strconv.Itoa(a2), 2)
		v.numberLine("#faa", 8, 16, a2, advance)
	}

	// Second number, part one
	if b1 != 0 {
		advance := 0
		if v.p.op == plus {
			advance = a2
		} else {
			advance = a1
		}
		v.circle("#faa", 20, 12, 2)
		v.line(20, 10, 22, 8)
		v.text(19, 13, strconv.Itoa(b1), 2)
		v.numberLine("#faa", 18, 16, b1, advance)
	}

	// Second number, part two
	if b2 != 0 {
		v.circle("#aaf", 26, 12, 2)
		v.line(26, 10, 24, 8)
		v.text(25, 13, strconv.Itoa(b2), 2)
		v.numberLine("#aaf", 24, 16, b2, 0)
	}

	// Equals
	v.line(29, 4, 32, 4)
	v.line(29, 6, 32, 6)

	// Choices
	for i, c := range v.p.cs {
		a := c
		answer := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			v.answer(a)
			return nil
		})
		v.toRelease = append(v.toRelease, answer)
		r := v.circle("white", 37, 6*i+4, 2)
		r.Set("onclick", answer)
		t := v.text(36, 6*i+5, strconv.Itoa(c), 2)
		t.Call("setAttribute", "id", fmt.Sprintf("answer-%v", i))
		t.Set("onclick", answer)
	}
}

func (v *HtmlView) release() {
	for _, fn := range v.toRelease {
		fn.Release()
	}
	v.toRelease = make([]js.Func, 0)
}

func (v *HtmlView) answer(c int) {
	if c == v.p.c {
		v.event(CORRECT)
	} else {
		v.event(INCORRECT)
	}
}

func (v *HtmlView) circle(color string, x, y, r int) js.Value {
	c := v.doc.Call("createElementNS", "http://www.w3.org/2000/svg", "circle")
	c.Call("setAttribute", "cx", x)
	c.Call("setAttribute", "cy", y)
	c.Call("setAttribute", "r", r)
	c.Call("setAttribute", "style", fmt.Sprintf("fill: %v; stroke: black; stroke-width: 0.25", color))
	v.svg.Call("appendChild", c)
	return c
}

func (v *HtmlView) text(x, y int, txt string, size int) js.Value {
	t := v.doc.Call("createElementNS", "http://www.w3.org/2000/svg", "text")
	t.Call("setAttribute", "x", x)
	t.Call("setAttribute", "y", y)
	t.Call("setAttribute", "style", fmt.Sprintf("font-size: %vpx", size))
	t.Set("innerHTML", txt)
	v.svg.Call("appendChild", t)
	return t
}

func (v *HtmlView) line(x1, y1, x2, y2 int) js.Value {
	l := v.doc.Call("createElementNS", "http://www.w3.org/2000/svg", "line")
	l.Call("setAttribute", "x1", x1)
	l.Call("setAttribute", "y1", y1)
	l.Call("setAttribute", "x2", x2)
	l.Call("setAttribute", "y2", y2)
	l.Call("setAttribute", "style", "stroke: black; stroke-width: 0.25")
	v.svg.Call("appendChild", l)
	return l
}

func (v *HtmlView) numberLine(color string, x, y, count, skip int) {
	row, col := 0, 0
	advance := func() {
		col += 1
		if col == 5 {
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
