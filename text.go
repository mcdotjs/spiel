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

