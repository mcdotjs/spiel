package main

import (
	"fmt"
)

var (
	tx int
	ty int
)


func (g *Game) UpdateCollisions() {
	playerObj := g.Objects[0]
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

func handleCollision(player *GameObject, obstacle *GameObject) {
	fmt.Println("Collision detected! Game over." )
}
