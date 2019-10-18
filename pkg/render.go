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

	CLICKABLE = "#00b"

	Y_OFFSET = 2
)

type HtmlView struct {
	p         Problem
	s         *Score
	doc       js.Value
	svg       js.Value
	event     func(Event)
	toRelease []js.Func
	hint      bool
}

func NewHtmlView(event func(Event)) *HtmlView {
	doc := js.Global().Get("document")
	svg := doc.Call("getElementById", "bond")
	svg.Call("setAttribute", "viewBox", "0 0 53 25")
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

func (v *HtmlView) SetScore(s *Score) {
	v.s = s
}

func (v *HtmlView) SetHint(h bool) {
	v.hint = h
}

func (v *HtmlView) Render() {

	v.svg.Set("innerHTML", "")
	v.release()

	// First number
	v.circle("white", "black", 7, 5+Y_OFFSET, 3)
	v.text(5, 6+Y_OFFSET, strconv.Itoa(v.p.a), 3, "black")

	// Second number
	v.circle("white", "black", 23, 5+Y_OFFSET, 3)
	v.text(21, 6+Y_OFFSET, strconv.Itoa(v.p.b), 3, "black")

	// Operator
	v.line(13, 5+Y_OFFSET, 17, 5+Y_OFFSET)
	if v.p.op == plus {
		v.line(15, 7+Y_OFFSET, 15, 3+Y_OFFSET)
	}

	if !v.hint {
		h := v.text(14, 18+Y_OFFSET, "?", 5, CLICKABLE)
		h.Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			v.event(HINT)
			return nil
		}))

	} else {
		a1, a2, b1, b2 := v.p.Breakout()

		// First number, part one
		if a1 != 0 {
			v.circle("#afa", "black", 4, 12+Y_OFFSET, 2)
			v.line(4, 10+Y_OFFSET, 6, 8+Y_OFFSET)
			v.text(3, 13+Y_OFFSET, strconv.Itoa(a1), 2, "black")
			v.numberLine("#afa", 2, 16+Y_OFFSET, a1, 0)
		}

		// First number, part two
		if a2 != 0 {
			advance := 0
			if v.p.op == minus {
				advance = a1
			}
			v.circle("#faa", "black", 10, 12+Y_OFFSET, 2)
			v.line(10, 10+Y_OFFSET, 8, 8+Y_OFFSET)
			v.text(9, 13+Y_OFFSET, strconv.Itoa(a2), 2, "black")
			v.numberLine("#faa", 8, 16+Y_OFFSET, a2, advance)
		}

		// Second number, part one
		if b1 != 0 {
			advance := 0
			if v.p.op == plus {
				advance = a2
			} else {
				advance = a1
			}
			v.circle("#faa", "black", 20, 12+Y_OFFSET, 2)
			v.line(20, 10+Y_OFFSET, 22, 8+Y_OFFSET)
			v.text(19, 13+Y_OFFSET, strconv.Itoa(b1), 2, "black")
			v.numberLine("#faa", 18, 16+Y_OFFSET, b1, advance)
		}

		// Second number, part two
		if b2 != 0 {
			v.circle("#aaf", "black", 26, 12+Y_OFFSET, 2)
			v.line(26, 10+Y_OFFSET, 24, 8+Y_OFFSET)
			v.text(25, 13+Y_OFFSET, strconv.Itoa(b2), 2, "black")
			v.numberLine("#aaf", 24, 16+Y_OFFSET, b2, 0)
		}
	}

	// Equals
	v.line(29, 4+Y_OFFSET, 32, 4+Y_OFFSET)
	v.line(29, 6+Y_OFFSET, 32, 6+Y_OFFSET)

	// Choices
	for i, c := range v.p.cs {
		a := c
		answer := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			v.answer(a)
			return nil
		})
		v.toRelease = append(v.toRelease, answer)
		r := v.circle("white", CLICKABLE, 37, 6*i+4+Y_OFFSET, 2)
		r.Set("onclick", answer)
		t := v.text(36, 6*i+5+Y_OFFSET, strconv.Itoa(c), 2, CLICKABLE)
		t.Call("setAttribute", "id", fmt.Sprintf("answer-%v", i))
		t.Set("onclick", answer)
	}

	// Score
	r := v.doc.Call("createElementNS", "http://www.w3.org/2000/svg", "rect")
	r.Call("setAttribute", "x", 0)
	r.Call("setAttribute", "y", 0)
	r.Call("setAttribute", "height", 1)
	r.Call("setAttribute", "width", 1+int(float64(52)*(float64(v.s.current)/(float64(v.s.goal)))))
	r.Call("setAttribute", "fill", "green")
	v.svg.Call("appendChild", r)
	v.text(45, 3+Y_OFFSET, fmt.Sprintf("Score: %v", v.s.wins), 1, "green")
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

func (v *HtmlView) circle(fill, stroke string, x, y, r int) js.Value {
	c := v.doc.Call("createElementNS", "http://www.w3.org/2000/svg", "circle")
	c.Call("setAttribute", "cx", x)
	c.Call("setAttribute", "cy", y)
	c.Call("setAttribute", "r", r)
	c.Call("setAttribute", "style", fmt.Sprintf("fill: %v; stroke: %v; stroke-width: 0.25", fill, stroke))
	v.svg.Call("appendChild", c)
	return c
}

func (v *HtmlView) text(x, y int, txt string, size int, color string) js.Value {
	t := v.doc.Call("createElementNS", "http://www.w3.org/2000/svg", "text")
	t.Call("setAttribute", "x", x)
	t.Call("setAttribute", "y", y)
	t.Call("setAttribute", "style", fmt.Sprintf("font-size: %vpx; fill: %v", size, color))
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
