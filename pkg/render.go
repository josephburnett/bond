package bond

import (
	"fmt"
	"strconv"
	"syscall/js"
)

type HtmlView struct {
	p   problem
	doc js.Value
	svg js.Value
}

func NewHtmlView(p problem) *HtmlView {
	doc := js.Global().Get("document")
	svg := doc.Call("getElementById", "bond")
	svg.Set("innerHTML", "")
	return &HtmlView{
		p:   p,
		doc: doc,
		svg: svg,
	}
}

func (v *HtmlView) Render() {

	// First number
	v.circle("white", 7, 10, 3)
	v.text(5, 11, strconv.Itoa(v.p.a), 3)

	// Second number
	v.circle("white", 23, 10, 3)
	v.text(21, 11, strconv.Itoa(v.p.b), 3)

	// Operator
	v.line(13, 10, 17, 10)
	if v.p.op == plus {
		v.line(15, 12, 15, 8)
	}

	a1, a2, b1, b2 := v.p.Breakout()

	// First number, part one
	if a1 != 0 {
		v.circle("#afa", 4, 17, 2)
		v.line(4, 15, 6, 13)
		v.text(3, 18, strconv.Itoa(a1), 2)
	}

	// First number, part two
	if a2 != 0 {
		v.circle("#faa", 10, 17, 2)
		v.line(10, 15, 8, 13)
		v.text(9, 18, strconv.Itoa(a2), 2)
	}

	// Second number, part one
	if b1 != 0 {
		v.circle("#faa", 20, 17, 2)
		v.line(20, 15, 22, 13)
		v.text(19, 18, strconv.Itoa(b1), 2)
	}

	// Second number, part two
	if b2 != 0 {
		v.circle("#aaf", 26, 17, 2)
		v.line(26, 15, 24, 13)
		v.text(25, 18, strconv.Itoa(b2), 2)
	}
}

func (v *HtmlView) circle(color string, x, y, r int) {
	c := v.doc.Call("createElementNS", "http://www.w3.org/2000/svg", "circle")
	c.Call("setAttribute", "cx", x)
	c.Call("setAttribute", "cy", y)
	c.Call("setAttribute", "r", r)
	c.Call("setAttribute", "style", fmt.Sprintf("fill: %v; stroke: black; stroke-width: 0.25", color))
	v.svg.Call("appendChild", c)
}

func (v *HtmlView) text(x, y int, txt string, size int) {
	t := v.doc.Call("createElementNS", "http://www.w3.org/2000/svg", "text")
	t.Call("setAttribute", "x", x)
	t.Call("setAttribute", "y", y)
	t.Call("setAttribute", "style", fmt.Sprintf("font-size: %vpx", size))
	t.Set("innerHTML", txt)
	v.svg.Call("appendChild", t)
}

func (v *HtmlView) line(x1, y1, x2, y2 int) {
	l := v.doc.Call("createElementNS", "http://www.w3.org/2000/svg", "line")
	l.Call("setAttribute", "x1", x1)
	l.Call("setAttribute", "y1", y1)
	l.Call("setAttribute", "x2", x2)
	l.Call("setAttribute", "y2", y2)
	l.Call("setAttribute", "style", "stroke: black; stroke-width: 0.25")
	v.svg.Call("appendChild", l)
}
