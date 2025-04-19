package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	back := playerObject.Mover.xIsGrowing()
	sf := float64(2)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate((-float64(tileSize) / 2), -float64(tileSize)/2)
	op.GeoM.Translate(playerObject.Position.xDelta/sf, playerObject.Position.yDelta/sf)
	op.GeoM.Scale(sf, sf)
	i := 0
	eee := playerObject.Mover.isMoving()
	if eee == false {
		i = 0
	} else {
		i = (g.count / animationSpeed) % frames
	}
	sx, sy := i*int(tileSize), 0
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Is moving right now?: ", eee, back))
	if back == false {
		screen.DrawImage(dogImageBack.SubImage(image.Rect(sx, sy, sx+int(tileSize), sy+int(tileSize))).(*ebiten.Image), op)
	} else {
		screen.DrawImage(dogImage.SubImage(image.Rect(sx, sy, sx+int(tileSize), sy+int(tileSize))).(*ebiten.Image), op)
	}
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
				Amplitude: 3,
				notImage:  true,
				layers:    obsLayearsFromsky,
				Mover: &JustHorizontalMover{
					Speed: 2,
				},
			}
			g.Obstacles = append(g.Obstacles, obstacle)
		}
	}
}

func main() {
	//NOTE: learn Mirko
	j := true
	f := false
	MyFloppy := GameObject{
		Position: Position{
			yDelta: 100,
			xDelta: 100,
		},
		Mover: &KyeBoardMover{
			Speed:          5.0,
			movingRightNow: &j,
			goingForward:   &f,
		},
		Img: dogImage,
	}

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
	ebiten.SetWindowTitle("Running Dago :)")
	if err := ebiten.RunGame(MyGame); err != nil {
		log.Fatal(err)
	}

}
