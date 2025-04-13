package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Update() error {
	g.floppy.useMoves()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 123, 99, 0xff})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.floppy.floppyPosition.xDelta, g.floppy.floppyPosition.yDelta)
	screen.DrawImage(gopher, op)
	ebitenutil.DebugPrint(screen, "something work")
}

func main() {
	ebiten.SetWindowSize(640, 888)
	ebiten.SetWindowTitle("Hello, MIREC@@@@!")
	myGame := &Game{
		floppy: MyFloppy,
	}
	if err := ebiten.RunGame(myGame); err != nil {
		log.Fatal(err)
	}
}
