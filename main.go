package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	g.count++
	if g.started && g.hideGame == false {
		g.UpdateObjectMovement()
		g.UpdateCollisions()
	}

	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.ended = false
		g.hideGame = false
		g.started = true
	}

	return nil
}

func (g *Game) drawBackground(screen *ebiten.Image, background *ebiten.Image) {
	w := background.Bounds().Dx()

	tileXCount := w / tileSize

	var xCount = (ScreenWidth + 100) / tileSize
	for _, l := range g.layers {
		for i, t := range l {

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xCount)*tileSize), float64((i/xCount)*tileSize))

			sx := (t % tileXCount) * tileSize
			sy := (t / tileXCount) * tileSize
			img := tilesSourceImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)

			screen.DrawImage(img, op)

		}

	}
}

func (g *Game) DrawDog(screen *ebiten.Image, playerObject *GameObject) {
	tileSize := 64
	frames := 5
	animationSpeed := 5
	sf := float64(2)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate((-float64(tileSize) / 2), -float64(tileSize)/2)
	op.GeoM.Translate(playerObject.Position.xDelta/sf, playerObject.Position.yDelta/sf)
	op.GeoM.Scale(sf, sf)
	i := (g.count / animationSpeed) % frames
	sx, sy := i*int(tileSize), 0
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("X: ", playerObject.Position.xDelta, sx))
	screen.DrawImage(dogImage.SubImage(image.Rect(sx, sy, sx+int(tileSize), sy+int(tileSize))).(*ebiten.Image), op)
}

func (g *Game) drawEndScreen(screen *ebiten.Image) {
	screen.Fill(color.White)
	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(0, 0)
	// scaleX := float64(ScreenWidth) / float64(w)
	// scaleY := float64(ScreenHeight) / float64(h)
	// op.GeoM.Scale(scaleX, scaleY)
	// screen.DrawImage(background, op)
}

// NOTE: this is more draw withi tiles
func (ob *GameObject) drawObstacle(screen *ebiten.Image) {

	w := tilesSourceImage.Bounds().Dx()
	tileXCount := w / tileSize

	var xCount = ObstacleWidth / tileSize
	for _, l := range ob.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xCount)*tileSize), float64((i/xCount)*tileSize))

			op.GeoM.Translate(ob.Position.xDelta, ob.Position.yDelta)
			sx := (t % tileXCount) * tileSize
			sy := (t / tileXCount) * tileSize
			img := tilesSourceImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)
			screen.DrawImage(img, op)
		}
	}

}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.started {
		g.drawBackground(screen, tilesSourceImage)
		for _, o := range g.Objects {
			g.DrawDog(screen, o)
			if g.debug {
				o.DrawDebug(screen)
			}
		}
		for _, o := range g.Obstacles {
			o.drawObstacle(screen)
			if g.debug {
				o.DrawDebug(screen)
			}
		}
	}

	if !g.started {

		g.notStarted(screen)
	}

	if g.ended {
		//g.drawEndScreen(screen)
		screen.Fill(color.Black)
		g.gameEnded(screen)

		g.DrawOwl(screen)

	}
}

func (g *Game) resetObstaclesPosition() {
	g.Obstacles = []*GameObject{}
	for i := 0; i < len(defaultXY); i++ {
		if val, ok := defaultXY[i]; !ok {
			log.Printf("some shit with obstacles generating", ok)
		} else {
			obstacle := &GameObject{
				Position: Position{
					yDelta: float64(val["y"]),
					xDelta: float64(val["x"]),
				},
				notImage: true,
				layers:   obsLayearsFromsky,
				Mover: &JustHorizontalMover{
					Speed: 2,
				},
			}
			g.Obstacles = append(g.Obstacles, obstacle)
		}
	}
}

func main() {
	MyFloppy := GameObject{
		Position: Position{
			yDelta: 100,
			xDelta: 100,
		},
		Mover: &KyeBoardMover{
			Speed: 5.0,
		},
		Img: dogImage,
	}
	// Obstacles := []*GameObject{
	// 	{
	// 		Position: Position{
	// 			yDelta: -100,
	// 			xDelta: 1300,
	// 		},
	// 		notImage: true,
	// 		layers:   obsLayearsFromsky,
	// 		Mover: &JustHorizontalMover{
	// 			Speed: 2,
	// 		},
	// 	},
	// 	{
	// 		Position: Position{
	// 			yDelta: 500,
	// 			xDelta: 1300,
	// 		},
	// 		notImage: true,
	// 		layers:   obsLayears,
	// 		Mover: &JustHorizontalMover{
	// 			Speed: 2,
	// 		},
	// 	},
	// 	{
	// 		Position: Position{
	// 			yDelta: 900,
	// 			xDelta: 1500,
	// 		},
	// 		notImage: true,
	// 		layers:   obsLayears,
	// 		Mover: &JustHorizontalMover{
	// 			Speed: 2,
	// 		},
	// 	},
	// 	{
	// 		Position: Position{
	// 			yDelta: 830,
	// 			xDelta: 1800,
	// 		},
	// 		notImage: true,
	// 		layers:   obsLayears,
	// 		Mover: &JustHorizontalMover{
	// 			Speed: 2,
	// 		},
	// 	},
	// 	{
	// 		Position: Position{
	// 			yDelta: 240,
	// 			xDelta: 2000,
	// 		},
	// 		notImage: true,
	// 		layers:   obsLayears,
	// 		Mover: &JustHorizontalMover{
	// 			Speed: 2,
	// 		},
	// 	},
	//
	// 	{
	// 		Position: Position{
	// 			yDelta: 900,
	// 			xDelta: 2000,
	// 		},
	// 		notImage: true,
	// 		layers:   obsLayears,
	// 		Mover: &JustHorizontalMover{
	// 			Speed: 2,
	// 		},
	// 	},
	// }
	//
	MyGame := &Game{
		Objects: []*GameObject{
			&MyFloppy,
		},
		debug:   false,
		layers:  gameLayer,
		started: false,
		ended:   false,
	}

	MyGame.resetObstaclesPosition()
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Not Floppy :)")
	if err := ebiten.RunGame(MyGame); err != nil {
		log.Fatal(err)
	}

}
