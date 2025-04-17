package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var (
	ScreenWidth      int
	ScreenHeight     int
	gopher           *ebiten.Image
	pikachu          *ebiten.Image
	ObstacleImage    *ebiten.Image
	ObstacleWidth    int
	TilesImage       *ebiten.Image
	tilesSourceImage *ebiten.Image
	tileSize         int = 32
)

type Game struct {
	Objects    []*GameObject
	Obstacles  []*GameObject
	started    bool
	ended      bool
	hideGame   bool
	background *ebiten.Image
	debug      bool
	layers     [][]int
}

type GameObject struct {
	Position Position
	Mover    Mover
	Img      *ebiten.Image
	layers   [][]int
	notImage bool
}

func init() {
	ScreenHeight = 800
	ScreenWidth = 1200
	ObstacleWidth = 128
	var err error
	gopher, _, err = ebitenutil.NewImageFromFile("gopher.png")
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
