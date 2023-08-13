package game

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type MapTile struct {
	PixelX   int
	PixelY   int
	Passable bool
	Image    *ebiten.Image
}

func NewTile(x int, y int, passable bool, sprite *ebiten.Image) *MapTile {
	return &MapTile{
		PixelX:   x,
		PixelY:   y,
		Passable: passable,
		Image:    sprite,
	}
}

type Level struct {
	Tiles []*MapTile
	gd    GameData
}

func NewLevel() Level {
	level := Level{gd: NewGameData()}
	tiles := level.CreateTiles()
	level.Tiles = tiles
	return level
}

func (l *Level) GetIndexFromXY(x, y int) int {
	return (y * l.gd.ScreenWidth) + x
}

func (l *Level) CreateTiles() []*MapTile {
	tiles := make([]*MapTile, l.gd.ScreenHeight*l.gd.ScreenWidth)
	for x := 0; x < l.gd.ScreenWidth; x++ {
		for y := 0; y < l.gd.ScreenHeight; y++ {
			index := l.GetIndexFromXY(x, y)
			if x == 0 || x == l.gd.ScreenWidth-1 || y == 0 || y == l.gd.ScreenHeight-1 {
				wall := LoadImageFromAssets("wall.png")
				tile := NewTile(x*l.gd.TileWidth, y*l.gd.TileHeight, true, wall)
				tiles[index] = tile

			} else {
				floor := LoadImageFromAssets("floor.png")
				tile := NewTile(x*l.gd.TileWidth, y*l.gd.TileHeight, false, floor)
				tiles[index] = tile
			}
		}
	}
	return tiles
}

func (l *Level) DrawLevel(screen *ebiten.Image) {
	gd := NewGameData()
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := l.Tiles[l.GetIndexFromXY(x, y)]
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, options)
		}
	}
}
