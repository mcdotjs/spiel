package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	log.Println("Draw")
	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-5, 5)
	op.GeoM.Scale(1.5, 3)
	screen.DrawImage(img, op)
	ebitenutil.DebugPrint(screen, "Hello, World jkljkjk!")
}


func main() {
	ebiten.SetWindowSize(640, 888)
	ebiten.SetWindowTitle("Hello, MIREC@@@@!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
