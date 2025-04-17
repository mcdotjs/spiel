package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type KyeBoardMover struct {
	Speed float64
}

type JustHorizontalMover struct {
	Speed float64
	Index float64
}

type Mover interface {
	Move(obj HasPosition) error
}

type Position struct {
	yDelta float64
	xDelta float64
}
//NOTE: there will be more movin types
type HasPosition interface {
	GetPosition() *Position
}

func (g *GameObject) GetPosition() *Position {
	return &g.Position
}

func (a *JustHorizontalMover) Move(obj HasPosition) error {
	pos := obj.GetPosition()
	pos.xDelta -= a.Speed

	if pos.xDelta < float64(-ObstacleWidth) {
		pos.xDelta = float64(ObstacleWidth) + float64(ScreenWidth)
	}
	return nil
}

func (a *KyeBoardMover) Move(ob HasPosition) error {
	pos := ob.GetPosition()
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		pos.xDelta += a.Speed
		if pos.xDelta > 699 {
			pos.xDelta = 699
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		pos.xDelta -= a.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		pos.yDelta += a.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		pos.yDelta -= a.Speed
	}
	return nil
}


func (g *Game) UpdateObjectMovement() {
	for _, o := range g.Objects {
		o.Mover.Move(o)
	}

	for _, o := range g.Obstacles {
		o.Mover.Move(o)
		//g.Objects[0].Colliding(o)
	}

}
