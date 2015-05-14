package main

import "github.com/nsf/termbox-go"

type Result struct {
	replay bool
	winner Cell
	board  *Board
}

func NewResult(b *Board, winner Cell) *Result {
	return &Result{
		replay: true,
		winner: winner,
		board:  b,
	}
}

func (r *Result) React(g *Game, e termbox.Event) error {
	switch e.Type {
	case termbox.EventError:
		return e.Err
	case termbox.EventKey:
		switch e.Key {
		case termbox.KeyCtrlC, termbox.KeyEsc:
			g.Close()
		case termbox.KeyArrowLeft, termbox.KeyArrowRight:
			r.replay = !r.replay
		case termbox.KeySpace, termbox.KeyEnter:
			if r.replay {
				g.Scene = NewCompetition(r.winner.Reverse())
			} else {
				g.Close()
			}
		}
	}
	return nil
}

func (r *Result) SetView() {
	termbox.HideCursor()

	x, y := tbxCenterXY()
	switch r.winner {
	case Circle:
		tbxSetText(x+4, y-1, "Circle win!",
			termbox.ColorGreen, termbox.ColorDefault)
	case Cross:
		tbxSetText(x+4, y-1, "Cross win!",
			termbox.ColorGreen, termbox.ColorDefault)
	default:
		tbxSetText(x+4, y-1, "Draw...",
			termbox.ColorGreen, termbox.ColorDefault)
	}

	if r.replay {
		tbxSetText(x+4, y+1, "Replay",
			termbox.ColorGreen|termbox.AttrBold, termbox.ColorDefault)
		tbxSetText(x+12, y+1, "Quit",
			termbox.ColorGreen, termbox.ColorDefault)
	} else {
		tbxSetText(x+4, y+1, "Replay",
			termbox.ColorGreen, termbox.ColorDefault)
		tbxSetText(x+12, y+1, "Quit",
			termbox.ColorGreen|termbox.AttrBold, termbox.ColorDefault)
	}

	x1, y1 := tbxCenterXY()
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			termbox.SetCell(x1+x-1, y1+y-1, r.board[x][y].View(),
				termbox.ColorGreen, termbox.ColorDefault)
		}
	}
}
