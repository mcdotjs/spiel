package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	playerXGrow int = 20
	playerYGrow int = 20
)

func (o *GameObject) GetBoundsOfPlayer() image.Rectangle {
	pos := o.Position
	if o.notImage == false {
		//width, height := o.Img.Bounds().Dx(), o.Img.Bounds().Dy()
		img := image.Rect(
			int(pos.xDelta)-2*playerXGrow,
			int(pos.yDelta)-3*playerYGrow,
			int(pos.xDelta)+2*playerXGrow,
			int(pos.yDelta),
		)
		return img
	}

	width, height := (tx+1)*tileSize, (ty+1)*tileSize
	img := image.Rect(
		int(o.Position.xDelta),
		int(o.Position.yDelta),
		int(o.Position.xDelta)+width,
		int(o.Position.yDelta)+height,
	)
	return img
}

func getTileBounds(obstaclePos *Position) image.Rectangle {
	img := image.Rect(
		int(obstaclePos.xDelta)+tx*tileSize,
		int(obstaclePos.yDelta)+ty*tileSize,
		int(obstaclePos.xDelta)+(tx+1)*tileSize,
		int(obstaclePos.yDelta)+(ty+1)*tileSize,
	)
	return img
}

func (o *GameObject) DrawDebug(screen *ebiten.Image) {
	playerBounds := o.GetBoundsOfPlayer()
	vector.DrawFilledRect(screen, float32(playerBounds.Min.X), float32(playerBounds.Min.Y), float32(playerBounds.Dx()), float32(playerBounds.Dy()), color.RGBA{255, 0, 0, 28}, true)
}
