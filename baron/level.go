package main

import (
	"log",
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameData struct {
	ScreenWidth int
	ScreenHeight int
	TileWidth int
	TileHeight int
}

func NewGameData() GameData {
	return GameData{ScreenWidth: 1280,
	ScreenHeight: 1024,
	TileWidth: 16,
	TileHeight: 16}
}