package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
}

func NewGameData() GameData {
	return GameData{ScreenWidth: 80,
		ScreenHeight: 50,
		TileWidth:    16,
		TileHeight:   16}
}

type MapTile struct {
	TopLeftPixelX int
	TopLeftPixelY int
	Blocks        bool
	Image         *ebiten.Image
}

func GetIndexFromXY(x, y int) int {
	gameData := NewGameData()
	return (y * gameData.ScreenWidth) + x
}
