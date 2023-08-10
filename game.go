package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	gd    GameData
	Tiles []MapTile
}

func NewGame() *Game {
	return &Game{gd: NewGameData(), Tiles: CreateTiles()}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for x := 0; x < g.gd.ScreenWidth; x++ {
		for y := 0; y < g.gd.ScreenHeight; y++ {
			tile := g.Tiles[GetIndexFromXY(x, y)]
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(tile.TopLeftPixelX), float64(tile.TopLeftPixelY))
			screen.DrawImage(tile.Image, options)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1280, 800
}
