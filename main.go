package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Update() error {
	//g.Objects[0].Colliding(g.Objects[1])
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

	for _, o := range g.Obstacles {
		for _, s := range o.Objects {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(o.Position.xDelta, o.Position.yDelta+s.Position.yDelta)
			screen.DrawImage(s.Img, op)

			if g.debug {
				s.GetPoints()
				s.DrawBorders()
			}

		}

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

	Obstacle1 := GameObject{
		Position: Position{
			yDelta: -float64(ScreenHeight / 2),
			xDelta: 400,
		},
		Mover: &JustHorizontalMover{
			Speed: 3,
		},
		Img: ObstacleImage,
	}

	Obstacle2 := GameObject{
		Position: Position{
			yDelta: float64(ScreenHeight/2) + 300,
			xDelta: 400,
		},
		Mover: &JustHorizontalMover{
			Speed: 3,
		},
		Img: ObstacleImage,
	}

	ObstacleParent := Obstacle{
		Position: Position{
			yDelta: 0,
			xDelta: 400,
		},
		Objects: []*GameObject{
			&Obstacle1, &Obstacle2,
		},
		Mover: &JustHorizontalMover{
			Speed: 3,
		},
	}
	Obstacle10 := GameObject{
		Position: Position{
			yDelta: -float64(ScreenHeight / 2),
			xDelta: 800,
		},
		Img: ObstacleImage,
	}

	Obstacle20 := GameObject{
		Position: Position{
			yDelta: float64(ScreenHeight/2) + 300,
			xDelta: 800,
		},
		Img: ObstacleImage,
	}

	ObstacleParent0 := Obstacle{
		Position: Position{
			yDelta: 0,
			xDelta: 800,
		},
		Objects: []*GameObject{
			&Obstacle10, &Obstacle20,
		},
		Mover: &JustHorizontalMover{
			Speed: 3,
		},
	}

	MyGame := &Game{
		Objects: []*GameObject{
			&MyFloppy,
		},
		Obstacles: []*Obstacle{
			&ObstacleParent, &ObstacleParent0,
		},
		debug: true,
	}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	if err := ebiten.RunGame(MyGame); err != nil {
		log.Fatal(err)
	}
}
