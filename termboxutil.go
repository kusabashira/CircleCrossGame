package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func tbxClear() error {
	return termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func tbxCenterXY() (x, y int) {
	winW, winH := termbox.Size()
	return winW / 2, winH / 2
}

func tbxSetFrame(x1, y1, x2, y2 int, bg termbox.Attribute) {
	for x := x1; x <= x2; x++ {
		termbox.SetCell(x, y1, ' ', bg, bg)
		termbox.SetCell(x, y2, ' ', bg, bg)
	}
	for y := y1; y <= y2; y++ {
		termbox.SetCell(x1, y, ' ', bg, bg)
		termbox.SetCell(x2, y, ' ', bg, bg)
	}
}

func tbxSetText(x, y int, s string, fg, bg termbox.Attribute) {
	for i, ch := range s {
		termbox.SetCell(x+i, y, ch, fg, bg)
	}
}

func tbxSetBaseView(g *Game) {
	x, y := tbxCenterXY()

	tbxSetText(x-9, y-4, "Circle Cross Game",
		termbox.ColorGreen, termbox.ColorDefault)
	tbxSetText(x-15, y+4, "(Press Ctrl+c or ESC to exit)",
		termbox.ColorGreen, termbox.ColorDefault)

	if !(g.NoCircleWin == 0 && g.NoCrossWin == 0) {
		tbxSetText(x-12, y-1, fmt.Sprintf("Circle: %d", g.NoCircleWin),
			termbox.ColorGreen, termbox.ColorDefault)
		tbxSetText(x-12, y+1, fmt.Sprintf(" Cross: %d", g.NoCrossWin),
			termbox.ColorGreen, termbox.ColorDefault)
	}

	tbxSetFrame(x-2, y-2, x+2, y+2, termbox.ColorGreen)
}
