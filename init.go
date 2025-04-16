package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	ScreenWidth   int
	ScreenHeight  int
	gopher        *ebiten.Image
	pikachu       *ebiten.Image
	ObstacleImage *ebiten.Image
	ObstacleWidth float64
)

type Game struct {
	Objects    []*GameObject
	Obstacles  []*Obstacle
	background *ebiten.Image
	debug      bool
}

type Obstacle struct {
	Objects  []*GameObject
	Window   float64
	Mover    Mover
	Position Position
}

type GameObject struct {
	Position Position
	Mover    Mover
	Img      *ebiten.Image
	Points   []Point
}

func init() {
	ScreenHeight = 800
	ScreenWidth = 1200
	ObstacleWidth = 69
	var err error
	gopher, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	ObstacleImage = ebiten.NewImage(int(ObstacleWidth), ScreenHeight)
	ObstacleImage.Fill(color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
