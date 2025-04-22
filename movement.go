package main

import (
	"image"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mover interface {
	Move(*GameObject) error
}

type HasPosition interface {
	GetPosition() *Position
}

type KyeBoardMover struct {
	Speed          float64
}

type HorizontalMover struct {
	Speed     float64
	Amplitude float64
}

type Position struct {
	yDelta float64
	xDelta float64
}

func (a *HorizontalMover) Move(obj *GameObject) error {
	pos := &obj.Position
	ampl := a.Amplitude

	pos.xDelta -= 1
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		pos.xDelta -= 3
	}

	timeFactor := float64(time.Now().UnixNano()) / 1e9
	offset := ampl * math.Cos(timeFactor*2)

	pos.yDelta += offset
	if pos.xDelta < float64(-ObstacleWidth) {
		pos.xDelta = float64(ObstacleWidth) + float64(ScreenWidth)
	}
	return nil
}

func (a *KyeBoardMover) Move(ob *GameObject) error {
	pos := &ob.Position
	*ob.movingRightNow = false
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		ob.metres += 0.3
		*ob.movingRightNow = true
		*ob.goingForward = true
		pos.xDelta += a.Speed
		if pos.xDelta > float64(ScreenWidth/2) {
			pos.xDelta = float64(ScreenWidth / 2)
		}

	}

	if pos.yDelta > float64(ScreenHeight) {
		pos.yDelta = float64(screenHeight)
	}

	if pos.yDelta < 80 {
		pos.yDelta = 80
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {

		ob.metres -= 0.3
		*ob.goingForward = false
		*ob.movingRightNow = true
		pos.xDelta -= a.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		ob.metres -= 0.08
		pos.yDelta += a.Speed * 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		ob.metres -= 0.08
		pos.yDelta -= a.Speed * 2
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
