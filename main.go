package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Update() error {
	g.floppy.initMoves()
	g.enemy.initEnemy()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 123, 99, 0xff})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.floppy.floppyPosition.xDelta, g.floppy.floppyPosition.yDelta)
	screen.DrawImage(gopher, op)

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(g.enemy.enemyPosition.xDelta, g.enemy.enemyPosition.yDelta)
	screen.DrawImage(gopher, op2)
	ebitenutil.DebugPrint(screen, "something work")
}

func main() {
	ebiten.SetWindowSize(640, 888)
	ebiten.SetWindowTitle("Hello, MIREC@@@@!")
	myGame := &Game{
		floppy: MyFloppy,
		enemy:  MyEnemy,
	}
	if err := ebiten.RunGame(myGame); err != nil {
		log.Fatal(err)
	}
}
