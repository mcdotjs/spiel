package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand/v2"
)

type KyeBoardMover struct {
	Speed float64
}

type JustHorizontalMover struct {
	Speed float64
	Index float64
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
	random := []int{-300, -350, -288, -188, -200, -250, -100, 0, 30, 50, 100}
	r := rand.IntN(10)
	if f.xDelta < -ObstacleWidth {
		f.xDelta = ObstacleWidth + float64(ScreenWidth)
		fmt.Println("ran", r)
		f.yDelta = float64(random[r])
	}
	return nil
}

func (a *KyeBoardMover) Move(f *Position) error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		f.xDelta += a.Speed
		if f.xDelta > 699 {
			f.xDelta = 699
		}
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

type Collision interface {
	Colliding(*GameObject) error
}

func (o *GameObject) Colliding(other *Obstacle) error {
	//fmt.Println("colllllllideing", o.Position.xDelta, other.Position.xDelta)
	return nil
}

func (g *Game) UpdateObjectMovement() {
	for _, o := range g.Objects {
		o.Mover.Move(&o.Position)
	}

	for _, o := range g.Obstacles {
		o.Mover.Move(&o.Position)
	}
}

func (g *Game) UpdateCollisions() {
	//NOTE: for all ostacles na s floppyn volat Collision
}
