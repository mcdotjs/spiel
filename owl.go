package main

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1200
	screenHeight = 800
	frameOX      = 0
	frameOY      = 0
	frameWidth   = 15
	frameHeight  = 20
	frameCount   = 12
)

var (
	owl *ebiten.Image
)

func init() {
	var err error
	owl, _, err = ebitenutil.NewImageFromFile("owl_vander.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) DrawOwl(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(120, 13)
	op.GeoM.Scale(4, 4)
	i := (g.count / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(owl.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}
