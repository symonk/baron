package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
}

func NewGameData() GameData {
	return GameData{
		ScreenWidth:  80,
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

func CreateTiles() []MapTile {
	gameData := NewGameData()
	tiles := make([]MapTile, 0)
	for x := 0; x < gameData.ScreenWidth; x++ {
		for y := 0; y < gameData.ScreenHeight; y++ {
			if x == 0 || x == gameData.ScreenWidth-1 || y == 0 || y == gameData.ScreenHeight-1 {
				wall, _, err := ebitenutil.NewImageFromFile("assets/images/wall.png")
				if err != nil {
					panic(err)
				}
				tile := MapTile{
					TopLeftPixelX: x * gameData.TileWidth,
					TopLeftPixelY: y * gameData.TileHeight,
					Blocks:        true,
					Image:         wall,
				}
				tiles = append(tiles, tile)
			} else {
				floor, _, err := ebitenutil.NewImageFromFile("assets/images/floor.png")
				if err != nil {
					panic(err)
				}
				tile := MapTile{
					TopLeftPixelX: x * gameData.TileWidth,
					TopLeftPixelY: y * gameData.TileWidth,
					Blocks:        true,
					Image:         floor,
				}
				tiles = append(tiles, tile)
			}
		}
	}
	return tiles
}
