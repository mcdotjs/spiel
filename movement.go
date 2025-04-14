package main

import "github.com/hajimehoshi/ebiten/v2"

type KyeBoardMover struct {
	Speed float64
}

type JustHorizontalMover struct {
	Speed float64
}

type RandomMover struct {
	Speed        float64
	RandomFactor int
}

type Mover interface {
	Move(*Position) error
}

type Position struct {
	yDelta float64
	xDelta float64
}

func (a *JustHorizontalMover) Move(f *Position) error {
	f.xDelta -= a.Speed
	return nil
}

func (a *KyeBoardMover) Move(f *Position) error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		f.xDelta += a.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		f.xDelta -= a.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		f.yDelta += a.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		f.yDelta -= a.Speed
	}
	return nil
}

func (a *RandomMover) Move(f *Position) error {
	f.yDelta -= a.Speed
	return nil
}

func (g *Game) UpdateObjectMovement() {
	for _, o := range g.Objects {
		o.Mover.Move(&o.Position)
	}
}
