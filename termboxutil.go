package main

import "github.com/nsf/termbox-go"

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
