package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	ScreenWidth      int
	ScreenHeight     int
	gopher           *ebiten.Image
	pikachu          *ebiten.Image
	ObstacleImage    *ebiten.Image
	ObstacleWidth    int
	Pika             *ebiten.Image
	TilesImage       *ebiten.Image
	tilesSourceImage *ebiten.Image
	tileSize         int = 32
)

type Game struct {
	Objects    []*GameObject
	Obstacles  []*Obstacle
	background *ebiten.Image
	debug      bool
	layers     [][]int
}

type Obstacle struct {
	Objects  []*GameObject
	Window   float64
	Mover    Mover
	Position Position
	layers   [][]int
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
	tileSize = 32
	var err error
	gopher, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	// ObstacleImage = ebiten.NewImage(int(ObstacleWidth), ScreenHeight)
	// ObstacleImage.Fill(color.White)

	Pika, _, err = ebitenutil.NewImageFromFile("nature.png")
	if err != nil {
		log.Fatal(err)
	}

	tilesSourceImage, _, err = ebitenutil.NewImageFromFile("dungeon.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
