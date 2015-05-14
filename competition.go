package main

import "github.com/nsf/termbox-go"

type Competition struct {
	x     int
	y     int
	turn  Cell
	board *Board
}

func NewCompetition(first Cell) *Competition {
	return &Competition{
		turn:  first,
		board: NewBoard(),
	}
}

func (c *Competition) SupressCursorXY() {
	switch {
	case c.x < 0:
		c.x = 0
	case 2 < c.x:
		c.x = 2
	}
	switch {
	case c.y < 0:
		c.y = 0
	case 2 < c.y:
		c.y = 2
	}
}

func (c *Competition) React(g *Game, e termbox.Event) error {
	switch e.Type {
	case termbox.EventError:
		return e.Err
	case termbox.EventKey:
		switch e.Key {
		case termbox.KeyCtrlC, termbox.KeyEsc:
			g.Close()
		case termbox.KeyArrowLeft:
			c.x--
		case termbox.KeyArrowRight:
			c.x++
		case termbox.KeyArrowUp:
			c.y--
		case termbox.KeyArrowDown:
			c.y++
		case termbox.KeySpace, termbox.KeyEnter:
			c.SupressCursorXY()
			err := c.board.Put(c.x, c.y, c.turn)
			switch err {
			case IndexOutOfBoard:
				return err
			case AlreadyPlaced:
				return nil
			case nil:
				c.turn = c.turn.Reverse()
			}
		}
	}
	if c.board.Finished() {
		w := c.board.Winner()
		switch w {
		case Circle:
			g.NoCircleWin++
		case Cross:
			g.NoCrossWin++
		}
		g.Scene = NewResult(c.board, w)
	}
	c.SupressCursorXY()
	return nil
}

func (c *Competition) SetView() {
	cenX, cenY := tbxCenterXY()
	switch c.turn {
	case Circle:
		tbxSetText(cenX+4, cenY-1, "Circle",
			termbox.ColorGreen|termbox.AttrBold, termbox.ColorDefault)
		tbxSetText(cenX+4, cenY+1, "Cross",
			termbox.ColorGreen, termbox.ColorDefault)
	case Cross:
		tbxSetText(cenX+4, cenY-1, "Circle",
			termbox.ColorGreen, termbox.ColorDefault)
		tbxSetText(cenX+4, cenY+1, "Cross",
			termbox.ColorGreen|termbox.AttrBold, termbox.ColorDefault)
	}

	termbox.SetCursor(cenX+c.x-1, cenY+c.y-1)
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			termbox.SetCell(cenX+x-1, cenY+y-1, c.board[x][y].View(),
				termbox.ColorGreen, termbox.ColorDefault)
		}
	}
}
