package main

import (
	"fmt"
	"math/rand/v2"

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
type HasPosition interface {
	GetPosition() *Position
}

func (g *GameObject) GetPosition() *Position {
	return &g.Position
}

func (o *Obstacle) GetPosition() *Position {
	return &o.Position
}

func (a *JustHorizontalMover) Move(obj HasPosition) error {

	pos := obj.GetPosition()
	pos.xDelta -= a.Speed
	fmt.Println("Moving:", pos.xDelta)
	if obstacle, ok := obj.(*Obstacle); ok {
		for _, obObj := range obstacle.Objects {
			obObj.Position.xDelta = pos.xDelta
		}
	}
	random := []int{-300, -350, -288, -188, -200, -250, -100, 0, 30, 50, 100}
	r := rand.IntN(10)
	if pos.xDelta < float64(-ObstacleWidth) {
		pos.xDelta = float64(ObstacleWidth) + float64(ScreenWidth)
		fmt.Println("ran", r)
		pos.yDelta = float64(random[r])
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

type Collision interface {
	Colliding(*Obstacle) error
}

func (o *GameObject) Colliding(other *Obstacle) error {
	for _, obstacle := range other.Objects {
		fmt.Println("colllllllideingXXXX", obstacle.Position.xDelta, o.Position.xDelta)
	}
	return nil
}

func (g *Game) UpdateObjectMovement() {
	for _, o := range g.Objects {
		o.Mover.Move(o)
	}

	for _, o := range g.Obstacles {
		o.Mover.Move(o)
		fmt.Println("ll", o.GetPosition().xDelta)
		//g.Objects[0].Colliding(o)
	}

}

func (g *Game) UpdateCollisions() {
	//NOTE: for all ostacles na s floppyn volat Collision
}
