package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	log.Println("Update")
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	log.Println("Draw")
	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
	screen.DrawImage(img, nil)
	ebitenutil.DebugPrint(screen, "Hello, World jkljkjk!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 888
}

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	log.Println("main")
	ebiten.SetWindowSize(640, 888)
	ebiten.SetWindowTitle("Hello, MIREC@@@@!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
