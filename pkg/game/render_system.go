package game

import "github.com/hajimehoshi/ebiten/v2"

func ProcessRenderables(g *Game, level Level, screen *ebiten.Image) {
	for _, result := range g.World.Query(g.WorldTags["renderables"]) {
		pos := result.Components[globalPosition].(*Position)
		img := result.Components[globalRenderable].(*Renderable).Image

		index := level.GetIndexFromXY(pos.X, pos.Y)
		tile := level.Tiles[index]
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(tile.TopLeftPixelX), float64(tile.TopLeftPixelY))
		screen.DrawImage(img, options)
	}
}