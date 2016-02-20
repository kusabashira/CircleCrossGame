package main

import (
	"errors"
)

type Cell int

const (
	Empty Cell = iota
	Circle
	Cross
)

func (c *Cell) Reversed() Cell {
	switch *c {
	case Circle:
		return Cross
	case Cross:
		return Circle
	default:
		return Empty
	}
}

func (c *Cell) Appearance() rune {
	switch *c {
	case Circle:
		return 'o'
	case Cross:
		return 'x'
	default:
		return ' '
	}
}

var (
	ErrAlreadyPlaced   = errors.New("Already placed")
	ErrIndexOutOfBoard = errors.New("Index out of board")
)

type Board [3][3]Cell

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) Put(x, y int, c Cell) error {
	if x < 0 || x >= 3 || y < 0 || y >= 3 {
		return ErrIndexOutOfBoard
	}
	if b[x][y] != Empty {
		return ErrAlreadyPlaced
	}
	b[x][y] = c
	return nil
}

func (b *Board) Winner() Cell {
	if b[1][1] != Empty &&
		(b[0][0] == b[1][1] && b[1][1] == b[2][2]) ||
		(b[0][2] == b[1][1] && b[1][1] == b[2][0]) {
		return b[1][1]
	}
	for x := 0; x < 3; x++ {
		if b[x][0] != Empty && b[x][0] == b[x][1] && b[x][1] == b[x][2] {
			return b[x][0]
		}
	}
	for y := 0; y < 3; y++ {
		if b[0][y] != Empty && b[0][y] == b[1][y] && b[1][y] == b[2][y] {
			return b[0][y]
		}
	}
	return Empty
}

func (b *Board) Finished() bool {
	if b.Winner() != Empty {
		return true
	}
	noEmpty := 0
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if b[x][y] == Empty {
				noEmpty++
			}
		}
	}
	return noEmpty == 0
}
