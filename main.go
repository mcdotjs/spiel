package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

		if g.debug {
			o.GetPoints()
			o.DrawBorders()
		}
	}
	vector.DrawFilledCircle(screen, 0, 0, 78, color.White, true)
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

	// MyEnemy := GameObject{
	// 	Position: Position{
	// 		yDelta: 2,
	// 		xDelta: 3,
	// 	},
	// 	Mover: &JustHorizontalMover{
	// 		Speed: 1.3,
	// 	},
	// 	Img: pikachu,
	// }

	MyGame := &Game{
		Objects: []*GameObject{
			&MyFloppy,
		},
		debug: true,
	}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Hello, MIREC@@@@!")

	if err := ebiten.RunGame(MyGame); err != nil {

		log.Fatal(err)
	}
}
