package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/symonk/baron/pkg/world"
)

type Game struct {
	Map       GameMap
	GameWorld *world.World
	ticks     int
}

func NewGame() *Game {
	return &Game{Map: NewGameMap(), GameWorld: world.NewWorld()}
}

func (g *Game) registerComponents() {
	playerComponent := g.GameWorld.World.NewComponent()
	g.GameWorld.World.NewEntity().AddComponent(playerComponent, &Player{})
}

func (g *Game) Update() error {
	g.ticks++
	MovePlayer(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Map.CurrentLevel
	level.DrawLevel(screen)
	DrawEntity(g, level, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	gd := NewGameData()
	return gd.ScreenWidth * gd.TileWidth, gd.ScreenHeight * gd.TileHeight
}
