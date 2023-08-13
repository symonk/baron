package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type EntityQueryManager struct {
}

// Draw images on top of the tiles.
// Each Renderable has a current position.
func DrawRenderables(g *Game, level Level, screen *ebiten.Image) {
	queryResult := g.World.Manager.Query(g.World.Tags[RenderablesView])
	for _, result := range queryResult {
		pos := result.Components[position].(*Position)
		img := result.Components[renderables].(*Renderable).Image

		index := level.GetIndexFromXY(pos.X, pos.Y)
		tile := level.Tiles[index]
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
		screen.DrawImage(img, options)
	}
}
