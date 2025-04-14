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

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 123, 99, 0xff})
	for _, o := range g.Objects {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(o.Position.xDelta, o.Position.yDelta)
		screen.DrawImage(o.Img, op)
	}

	ebitenutil.DebugPrint(screen, "something work")
}

func main() {
	MyFloppy := GameObject{
		Position: Position{
			yDelta: 0,
			xDelta: 0,
		},
		Mover: &KyeBoardMover{
			Speed: 9.0,
		},
		Img: gopher,
	}

	MyEnemy := GameObject{
		Position: Position{
			yDelta: 100,
			xDelta: 300,
		},
		Mover: &JustHorizontalMover{
			Speed: 0.3,
		},
		Img: pikachu,
	}

	MyGame := &Game{
		Objects: []*GameObject{
			&MyFloppy,
			&MyEnemy,
		},
	}
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Hello, MIREC@@@@!")

	if err := ebiten.RunGame(MyGame); err != nil {
		log.Fatal(err)
	}
}
