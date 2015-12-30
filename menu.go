package main

import "github.com/nsf/termbox-go"

type Menu struct {
	firstTurn Cell
}

func NewMenu() *Menu {
	return &Menu{
		firstTurn: Circle,
	}
}

func (m *Menu) React(g *Game, e termbox.Event) error {
	switch e.Type {
	case termbox.EventError:
		return e.Err
	case termbox.EventKey:
		switch e.Ch {
		case 'q':
			g.Close()
		}
		switch e.Key {
		case termbox.KeyCtrlC:
			g.Close()
		case termbox.KeyArrowLeft, termbox.KeyArrowRight:
			m.firstTurn = m.firstTurn.Reversed()
		case termbox.KeySpace, termbox.KeyEnter:
			t := m.firstTurn
			g.SetFirstTurn(t)
			g.SetScene(NewCompetition(t))
		}
	}
	return nil
}

func (m *Menu) SetView() {
	termbox.HideCursor()

	x, y := tbxCenterXY()
	tbxSetText(x+4, y-1, "Which is first?",
		termbox.ColorGreen, termbox.ColorDefault)
	switch m.firstTurn {
	case Circle:
		tbxSetText(x+4, y+1, "Circle",
			termbox.ColorGreen|termbox.AttrBold, termbox.ColorDefault)
		tbxSetText(x+12, y+1, "Cross",
			termbox.ColorGreen, termbox.ColorDefault)
	case Cross:
		tbxSetText(x+4, y+1, "Circle",
			termbox.ColorGreen, termbox.ColorDefault)
		tbxSetText(x+12, y+1, "Cross",
			termbox.ColorGreen|termbox.AttrBold, termbox.ColorDefault)
	}
}
