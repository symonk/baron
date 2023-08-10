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
	level := g.Map.Dungeons[0].Levels[0]
	level.DrawLevel(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	gd := NewGameData()
	return gd.ScreenWidth * gd.TileWidth, gd.ScreenHeight * gd.TileHeight
}
