package main

import "github.com/hajimehoshi/ebiten/v2"

type Floppy struct {
	floppyPosition  Position
	movementHandler AxisMovements
}

type Movement struct {
	speed float64
}

type Position struct {
	yDelta float64
	xDelta float64
}

type AxisMovements interface {
	xAxis(*Floppy) error
	yAxis(*Floppy) error
}

var MyFloppy = Floppy{
	floppyPosition: Position{
		yDelta: 0,
		xDelta: 0,
	},
	movementHandler: &Movement{
		speed: 9.0,
	},
}

func (a *Movement) xAxis(f *Floppy) error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		f.floppyPosition.xDelta += a.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		f.floppyPosition.xDelta -= a.speed
	}
	return nil
}

func (a *Movement) yAxis(f *Floppy) error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		f.floppyPosition.yDelta += a.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		f.floppyPosition.yDelta -= a.speed
	}
	return nil
}

func (f *Floppy) useMoves() {
	f.movementHandler.yAxis(f)
	f.movementHandler.xAxis(f)
}
