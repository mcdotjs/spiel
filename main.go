package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Update() error {
	g.UpdateObjectMovement()
	return nil
}

var MyFloppy = GameObject{
	Position: Position{
		yDelta: 0,
		xDelta: 0,
	},
	Mover: &KyeBoardMover{
		Speed: 9.0,
	},
}

var MyEnemy = GameObject{
	Position: Position{
		yDelta: 100,
		xDelta: 300,
	},
	Mover: &JustHorizontalMover{
		Speed: 0.3,
	},
}

var Third = GameObject{
	Position: Position{
		yDelta: 600,
		xDelta: 300,
	},
	Mover: &JustHorizontalMover{
		Speed: 1,
	},
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 123, 99, 0xff})
	for _, o := range g.Objects {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(o.Position.xDelta, o.Position.yDelta)
		screen.DrawImage(gopher, op)
	}

	ebitenutil.DebugPrint(screen, "something work")
}

func main() {
	ebiten.SetWindowSize(640, 888)
	ebiten.SetWindowTitle("Hello, MIREC@@@@!")

	myGame := &Game{
		Objects: []*GameObject{&MyFloppy, &MyEnemy, &Third},
	}
	if err := ebiten.RunGame(myGame); err != nil {
		log.Fatal(err)
	}
}
