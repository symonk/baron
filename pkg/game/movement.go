package game

import "github.com/hajimehoshi/ebiten/v2"

func MovePlayer(g *Game) {
	players := g.WorldTags["players"]
	x, y := 0, 0

	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		y = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		y = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		x = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
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