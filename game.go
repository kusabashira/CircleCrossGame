package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type Scene interface {
	React(g *Game, e termbox.Event) error
	SetView()
}

type Game struct {
	noCircleWin int
	noCrossWin  int
	running     bool
	scene       Scene
	firstTurn   Cell
}

func NewGame() *Game {
	return &Game{
		scene:   NewMenu(),
		running: false,
	}
}

func (g *Game) Init() error {
	if err := termbox.Init(); err != nil {
		return err
	}
	if err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault); err != nil {
		return err
	}
	return termbox.Flush()
}

func (g *Game) Close() {
	g.running = false
	termbox.Close()
}

func (g *Game) SetFirstTurn(c Cell) {
	g.firstTurn = c
}

func (g *Game) GetFirstTurn() Cell {
	return g.firstTurn
}

func (g *Game) SetScene(s Scene) {
	g.scene = s
}

func (g *Game) CountUpCircleWin() {
	g.noCircleWin++
}

func (g *Game) CountUpCrossWin() {
	g.noCrossWin++
}

func (g *Game) setBaseView() {
	x, y := tbxCenterXY()
	tbxSetText(x-9, y-4, "Circle Cross Game",
		termbox.ColorGreen, termbox.ColorDefault)
	tbxSetText(x-14, y+4, "(Press q or Ctrl-c to exit)",
		termbox.ColorGreen, termbox.ColorDefault)
	if !(g.noCircleWin == 0 && g.noCrossWin == 0) {
		tbxSetText(x-12, y-1, fmt.Sprintf("Circle: %d", g.noCircleWin),
			termbox.ColorGreen, termbox.ColorDefault)
		tbxSetText(x-12, y+1, fmt.Sprintf(" Cross: %d", g.noCrossWin),
			termbox.ColorGreen, termbox.ColorDefault)
	}
	tbxSetFrame(x-2, y-2, x+2, y+2, termbox.ColorGreen)
}

func (g *Game) setView() error {
	if err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault); err != nil {
		return err
	}

	g.setBaseView()
	g.scene.SetView()
	return termbox.Flush()
}

func (g *Game) react(e termbox.Event) error {
	return g.scene.React(g, e)
}

func (g *Game) Main() error {
	g.running = true
	for g.running {
		if err := g.setView(); err != nil {
			return err
		}

		e := termbox.PollEvent()
		if err := g.react(e); err != nil {
			return err
		}
	}
	return nil
}
