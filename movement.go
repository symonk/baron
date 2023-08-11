package main

import "github.com/hajimehoshi/ebiten/v2"

func MovePlayer(g *Game) {
	players := g.WorldTags["players"]
	x, y := 0, 0

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		y = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		y = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		x = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x = 1
	}

	level := g.Map.CurrentLevel
	for _, result := range g.World.Query(players) {
		pos := result.Components[globalPosition].(*Position)
		index := level.GetIndexFromXY(pos.X+x, pos.Y+y)
		tile := level.Tiles[index]
		if !tile.Blocks {
			pos.X += x
			pos.Y += y
		}
	}
}
