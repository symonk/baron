package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func MovePlayer(g *Game) {
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
	for _, result := range g.World.Manager.Query(g.World.Tags[RenderablesView]) {
		pos := result.Components[position].(*Position)
		index := level.GetIndexFromXY(pos.X+x, pos.Y+y)
		tile := level.Tiles[index]
		if !tile.Passable {
			pos.X += x
			pos.Y += y
		} else {
			g.Player.TransformHealth(-10)
		}
	}
}

func ResetPlayer(g *Game) {
	for _, result := range g.World.Manager.Query(g.World.Tags[RenderablesView]) {
		pos := result.Components[position].(*Position)
		pos.X = 20
		pos.Y = 20
		g.Player.Revive()
	}
}
