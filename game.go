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
	NoCircleWin int
	NoCrossWin  int
	Running     bool
	Scene       Scene
	FirstTurn   Cell
}

func NewGame() *Game {
	return &Game{
		Scene:   NewMenu(),
		Running: false,
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
	g.Running = false
	termbox.Close()
}

func (g *Game) setBaseView() {
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

func (g *Game) setView() error {
	if err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault); err != nil {
		return err
	}

	g.setBaseView()
	g.Scene.SetView()
	return termbox.Flush()
}

func (g *Game) react(e termbox.Event) error {
	return g.Scene.React(g, e)
}

func (g *Game) Main() error {
	g.Running = true
	for g.Running {
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
