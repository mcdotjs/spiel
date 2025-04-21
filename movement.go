package main

import (
	"image"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

//NOTE:  you are welcome to tell me how to implement logic here

type KyeBoardMover struct {
	Speed          float64
	movingRightNow *bool
	goingForward   *bool
}

type JustHorizontalMover struct {
	Speed          float64
	movingRightNow *bool
	goingForward   *bool
}

type Mover interface {
	Move(obj HasPosition) error
	isMoving() bool
	xIsGrowing() bool
}

type Position struct {
	yDelta float64
	xDelta float64
}

// NOTE: there will be more movin types
type HasPosition interface {
	GetPosition() *Position
	GetAmplitude() *float64
}

func (o *KyeBoardMover) xIsGrowing() bool {
	return *o.goingForward
}

func (o *JustHorizontalMover) xIsGrowing() bool {
	return false
}

func (g *KyeBoardMover) isMoving() bool {
	return *g.movingRightNow
}

func (g *JustHorizontalMover) isMoving() bool {
	return *g.movingRightNow
}

func (g *GameObject) GetPosition() *Position {
	return &g.Position
}

func (g *GameObject) GetAmplitude() *float64 {
	return &g.Amplitude
}

func (a *JustHorizontalMover) Move(obj HasPosition) error {
	pos := obj.GetPosition()
	ampl := obj.GetAmplitude()

	pos.xDelta -= 1
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		pos.xDelta -= 3
	}

	timeFactor := float64(time.Now().UnixNano()) / 1e9
	offset := *ampl * math.Cos(timeFactor*2)

	pos.yDelta += offset
	if pos.xDelta < float64(-ObstacleWidth) {
		pos.xDelta = float64(ObstacleWidth) + float64(ScreenWidth)
	}
	return nil
}

func (a *KyeBoardMover) Move(ob HasPosition) error {
	pos := ob.GetPosition()
	*a.movingRightNow = false
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		*a.movingRightNow = true
		*a.goingForward = true
		pos.xDelta += a.Speed
		if pos.xDelta > float64(ScreenWidth/2) {
			pos.xDelta = float64(ScreenWidth / 2)
		}

	}

	if pos.yDelta > float64(ScreenHeight) {
		pos.yDelta = float64(screenHeight)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		*a.goingForward = false
		*a.movingRightNow = true
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

func (p *viewport) MoveX(fastness int, direction int) {
	s := image.Point{X: tileSize, Y: tileSize}
	maxX16 := s.X * 16

	p.x16 += (s.X/16 + fastness) * direction
	p.x16 %= maxX16
}

func (p *viewport) Position() (int, int) {
	return p.x16, p.y16
}

func (g *Game) UpdateObjectMovement() {
	g.Objects[0].Mover.Move(g.Objects[0])
	for _, o := range g.Obstacles {
		o.Mover.Move(o)
	}

}
