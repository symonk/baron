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
	return GameData{ScreenWidth: 80,
	ScreenHeight: 50,
	TileWidth: 16,
	TileHeight: 16}
}