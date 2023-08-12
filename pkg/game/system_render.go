package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type EntityQueryManager struct {
}

// Draw images on top of the tiles.
// Each Renderable has a current position.
func DrawEntity(g *Game, level Level, screen *ebiten.Image) {
	frame0X, frame0Y := 0, 32
	frameWidth, frameheight := 32, 32
	frameCount := 3
	for _, result := range g.GameWorld.Manager.Query(g.GameWorld.Tags[RenderablesView]) {
		pos := result.Components[position].(*Position)
		img := result.Components[renderables].(*Renderable).Image

		index := level.GetIndexFromXY(pos.X, pos.Y)
		tile := level.Tiles[index]
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(tile.TopLeftPixelX), float64(tile.TopLeftPixelY))
		i := (g.ticks / 5) % frameCount
		sx, sy := frame0X+i*frameWidth, frame0Y
		sub := img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameheight)).(*ebiten.Image)
		screen.DrawImage(sub, options)
	}
}
