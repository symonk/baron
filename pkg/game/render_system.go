package game

import "github.com/hajimehoshi/ebiten/v2"

type EntityQueryManager struct {
}

func DrawEntity(g *Game, level Level, screen *ebiten.Image) {
	for _, result := range g.GameWorld.Manager.Query(g.GameWorld.Tags[Renderables]) {
		pos := result.Components[globalPosition].(*Position)
		img := result.Components[globalRenderable].(*Renderable).Image

		index := level.GetIndexFromXY(pos.X, pos.Y)
		tile := level.Tiles[index]
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(tile.TopLeftPixelX), float64(tile.TopLeftPixelY))
		screen.DrawImage(img, options)
	}
}
