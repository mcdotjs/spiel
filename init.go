package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	ScreenWidth  int
	ScreenHeight int
	gopher  *ebiten.Image
	pikachu *ebiten.Image
)

type Game struct {
	Objects    []*GameObject
	background *ebiten.Image
}

type GameObject struct {
	Position Position
	Mover    Mover
	Img      *ebiten.Image
}

func init() {
	ScreenHeight = 800
	ScreenWidth = 1200
	var err error
	gopher, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	pikachu, _, err = ebitenutil.NewImageFromFile("pika.png")
	if err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
