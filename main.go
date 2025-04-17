package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	fmt.Println("program started")
	g.count++
	if g.started && g.hideGame == false {
		g.UpdateObjectMovement()
		g.UpdateCollisions()
	}

	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
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
	g.drawBackground(screen, tilesSourceImage)
	if g.started {
		//NOTE: zatial len gopher
		for _, o := range g.Objects {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(o.Position.xDelta, o.Position.yDelta)
			screen.DrawImage(o.Img, op)
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
		g.gameEnded(screen)
	}
	g.DrawOwl(screen)
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

	Obstacles := []*GameObject{
		{
			Position: Position{
				yDelta: -100,
				xDelta: 1300,
			},
			notImage: true,
			layers:   obsLayearsFromsky,
			Mover: &JustHorizontalMover{
				Speed: 2,
			},
		},
		{
			Position: Position{
				yDelta: 500,
				xDelta: 1300,
			},
			notImage: true,
			layers:   obsLayears,
			Mover: &JustHorizontalMover{
				Speed: 2,
			},
		},
		{
			Position: Position{
				yDelta: 500,
				xDelta: 1500,
			},
			notImage: true,
			layers:   obsLayears,
			Mover: &JustHorizontalMover{
				Speed: 2,
			},
		},
	}

	MyGame := &Game{
		Objects: []*GameObject{
			&MyFloppy,
		},
		Obstacles: Obstacles,
		debug:     true,
		layers:    gameLayer,
		started:   false,
		ended:     false,
	}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Not Floppy :)")
	if err := ebiten.RunGame(MyGame); err != nil {
		log.Fatal(err)
	}

}
