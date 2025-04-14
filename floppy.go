package main

import "github.com/hajimehoshi/ebiten/v2"

type Floppy struct {
	floppyPosition  Position
	movementHandler Movement
}

type Enemy struct {
	enemyPosition   Position
	movementHandler Movement
}

type Movement struct {
	speed float64
}

type Position struct {
	yDelta float64
	xDelta float64
}

type YMovementWithKeyPress interface {
	yAxis(*Position) error
}

type XMovementWithKeyPress interface {
	xAxis(*Position) error
}

type OnlyXMovement interface {
	goToLeft(*Position) error
}

var MyFloppy = Floppy{
	floppyPosition: Position{
		yDelta: 0,
		xDelta: 0,
	},
	movementHandler: Movement{
		speed: 9.0,
		//keys!!!!
	},
}

var MyEnemy = Enemy{
	enemyPosition: Position{
		yDelta: 100,
		xDelta: 300,
	},
	movementHandler: Movement{
		speed: 0.3,
	},
}

func (a *Movement) goToLeft(f *Position) error {
	f.xDelta -= a.speed
	return nil
}

func (a *Movement) xAxis(f *Position) error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		f.xDelta += a.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		f.xDelta -= a.speed
	}
	return nil
}

func (a *Movement) yAxis(f *Position) error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		f.yDelta += a.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		f.yDelta -= a.speed
	}
	return nil
}

func (f *Floppy) initMoves() {
	f.movementHandler.yAxis(&f.floppyPosition)
	f.movementHandler.xAxis(&f.floppyPosition)
}

func (e *Enemy) initEnemy() {
	e.movementHandler.goToLeft(&e.enemyPosition)
}

