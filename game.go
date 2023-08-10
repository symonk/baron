package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	Map GameMap
}

func NewGame() *Game {
	return &Game{Map: NewGameMap()}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	gd := NewGameData()
	level := g.Map.Dungeons[0].Levels[0]
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := level.Tiles[level.GetIndexFromXY(x, y)]
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(tile.TopLeftPixelX), float64(tile.TopLeftPixelY))
			screen.DrawImage(tile.Image, options)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	gd := NewGameData()
	return gd.ScreenWidth * gd.TileWidth, gd.ScreenHeight * gd.TileHeight
}
