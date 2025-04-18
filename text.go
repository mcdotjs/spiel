package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	mplusFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.ArcadeN_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

func DrawMultiLineText(screen *ebiten.Image, lines []struct {
	Text  string
	Color color.Color
}, x, y float64, fontSize float64, lineSpacing float64) {

	currentY := y
	for _, line := range lines {
		op := &text.DrawOptions{}
		op.GeoM.Translate(x, currentY)
		op.ColorScale.ScaleWithColor(line.Color)

		text.Draw(
			screen,
			line.Text,
			&text.GoTextFace{
				Source: mplusFaceSource,
				Size:   fontSize,
			},
			op,
		)

		currentY += lineSpacing
	}
}

func (g *Game) notStarted(screen *ebiten.Image) {
	textLines := []struct {
		Text  string
		Color color.Color
	}{
		{"WELCOME!!!", color.RGBA{255, 222, 33, 255}},
		{"Press ENTER to start game", color.RGBA{255, 255, 255, 255}},
		{"Use ARROW KEYS to move", color.RGBA{200, 200, 100, 255}},
		{"Avoid obstacles!", color.RGBA{255, 100, 100, 255}},
	}
	DrawMultiLineText(screen, textLines, 100, 60, 28.0, 37.0)
	nvimLines := []struct {
		Text  string
		Color color.Color
	}{
		{"no mouse needed...", color.RGBA{255, 255, 255, 255}},
		{"floppy is nvim user btw", color.RGBA{255, 100, 180, 255}},
	}
	DrawMultiLineText(screen, nvimLines, 900, 700, 12.0, 18.0)
}

func (g *Game) gameEnded(screen *ebiten.Image) {
	textLines := []struct {
		Text  string
		Color color.Color
	}{
		{"SO SORRY!!", color.RGBA{255, 222, 33, 255}},
		{"Press ENTER", color.RGBA{255, 255, 255, 255}},
		{"to restart game", color.RGBA{255, 255, 255, 255}},
	}
	DrawMultiLineText(screen, textLines, 100, 60, 28.0, 44.0)
}
