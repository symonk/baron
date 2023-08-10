package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Level struct {
	Tiles []MapTile
}

func NewLevel() Level {
	level := Level{}
	tiles := level.CreateTiles()
	level.Tiles = tiles
	return level
}

type MapTile struct {
	TopLeftPixelX int
	TopLeftPixelY int
	Blocks        bool
	Image         *ebiten.Image
}

func (l *Level) GetIndexFromXY(x, y int) int {
	gameData := NewGameData()
	return (y * gameData.ScreenWidth) + x
}

func (l *Level) CreateTiles() []MapTile {
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
