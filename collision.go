package main

import (
	"fmt"
)

var (
	tx int
	ty int
)

var c int = 1

func (g *Game) UpdateCollisions() {
	playerObj := g.Objects[0] //floppy
	playerBounds := playerObj.GetBoundsOfPlayer()

	tilesPerRow := ObstacleWidth / tileSize
	for _, obstacle := range g.Obstacles {
		for _, layer := range obstacle.layers {
			for tileIdx, tileValue := range layer {
				if tileValue == 0 {
					continue
				}
				tx = tileIdx % tilesPerRow
				ty = tileIdx / tilesPerRow

				tileBounds := getTileBounds(&obstacle.Position)

				if playerBounds.Overlaps(tileBounds) {
					g.ended = true
					g.hideGame = true
					g.resetObstaclesPosition()
					g.Objects[0].Position.xDelta = 0
					g.Objects[0].Position.yDelta = 100
					handleCollision(playerObj, obstacle)
					return
				}
			}

		}
	}
}

func isCollidableTile(tileType int) bool {
	//TODO: implement ... nice :)
	// Define which tile types should cause collisions
	// For example, if tile type 22 is a solid obstacle:

	return tileType == 22
}

func handleCollision(player *GameObject, obstacle *GameObject) {
	c++
	// For a Flappy Bird style game, this might end the game
	fmt.Println("Collision detected! Game over.", c)
	// Implement game over logic
	// Or push the player back, reduce health, etc.
}
