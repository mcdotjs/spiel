package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Objects    []*GameObject
	background *ebiten.Image
}

type GameObject struct {
	Position Position
	Mover    Mover
}

var (
	gopher *ebiten.Image
)

func init() {
	var err error
	gopher, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 888
}
