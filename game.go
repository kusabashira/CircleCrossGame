package main

import (
	"github.com/nsf/termbox-go"
)

type Scene interface {
	React(g *Game, e termbox.Event) error
	SetView(g *Game) error
}

type Game struct {
	NoCircleWin int
	NoCrossWin  int
	Running     bool
	Scene       Scene
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
	if err := tbxClear(); err != nil {
		return err
	}
	return termbox.Flush()
}

func (g *Game) Close() {
	g.Running = false
	termbox.Close()
}

func (g *Game) Main() error {
	g.Running = true
	for g.Running {
		if err := g.Scene.SetView(g); err != nil {
			return err
		}

		e := termbox.PollEvent()
		if err := g.Scene.React(g, e); err != nil {
			return err
		}
	}
	return nil
}
