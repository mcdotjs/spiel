package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	ScreenWidth       int
	ScreenHeight      int
	gopher            *ebiten.Image
	pikachu           *ebiten.Image
	ObstacleImage     *ebiten.Image
	dogImage          *ebiten.Image
	dogImageBack      *ebiten.Image
	ObstacleWidth     int
	TilesImage        *ebiten.Image
	tilesSourceImage  *ebiten.Image
	tileSize          int = 32
	tilesVieport      int = 1400
	padding           int = -100
	bg                image.Rectangle
	MouseButtonLeft   MouseButton
	MouseButtonRight  MouseButton
	MouseButtonMiddle MouseButton
)

type MouseButton int

type Game struct {
	Objects    []*GameObject
	Obstacles  []*GameObject
	started    bool
	ended      bool
	hideGame   bool
	background *ebiten.Image
	debug      bool
	layers     [][]int
	count      int
	viewport   viewport
}

type GameObject struct {
	Position  Position
	Mover     Mover
	Img       *ebiten.Image
	layers    [][]int
	notImage  bool
	Amplitude float64
}

type viewport struct {
	x16 int
	y16 int
}

func init() {
	ScreenHeight = 800
	ScreenWidth = 1200
	ObstacleWidth = 96
	var err error
	gopher, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	tilesSourceImage, _, err = ebitenutil.NewImageFromFile("dungeon2.png")
	if err != nil {
		log.Fatal(err)
	}

	dogImageBack, _, err = ebitenutil.NewImageFromFile("wolfback.png")
	if err != nil {
		log.Fatal(err)
	}

	dogImage, _, err = ebitenutil.NewImageFromFile("wolf.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
