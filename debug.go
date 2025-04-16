package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// "image/color"
//
// "github.com/hajimehoshi/ebiten/v2"
// "github.com/hajimehoshi/ebiten/v2/vector"

type Point struct {
	x, y float64
}

type Line struct {
	Start Point
	End   Point
}

type Cell struct {
	Lines []Line
}

func (o *GameObject) GetPoints(img *ebiten.Image) ([]Point, error) {
	o.Points = make([]Point, 4)
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	o.Points[0] = Point{x: float64(0), y: float64(0)}
	o.Points[1] = Point{x: float64(w), y: float64(0)}
	o.Points[2] = Point{x: float64(w), y: float64(h)}
	o.Points[3] = Point{x: float64(0), y: float64(h)}
	fmt.Println("print points", o.Points)
	return o.Points, nil
}

func (o *GameObject) DrawBorders() error {
	for i, point := range o.Points {
		currentPoint := point
		anotherPoint := o.Points[0]
		if i < 3 {
			anotherPoint = o.Points[i+1]
		}
		vector.StrokeLine(o.Img, float32(currentPoint.x), float32(currentPoint.y), float32(anotherPoint.x), float32(anotherPoint.y), 1, color.RGBA{255, 0, 0, 255}, true)
	}
	return nil
}

type HasWalls interface {
	drawWalls() error
}
