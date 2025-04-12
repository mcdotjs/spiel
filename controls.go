package main

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) controls() {

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.floppy.yDelta += 3.9
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.floppy.yDelta -= 3.9
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.floppy.xDelta += 3.9
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.floppy.xDelta -= 3.9
	}

}
