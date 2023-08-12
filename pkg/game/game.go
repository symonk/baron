package game

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/symonk/baron/pkg/world"
)

type Game struct {
	Map       GameMap
	GameWorld *world.World
	ticks     int
}

func NewGame() *Game {
	game := &Game{Map: NewGameMap(), GameWorld: world.NewWorld()}
	game.registerComponents()
	return game
}

func (g *Game) registerComponents() {
	playerComponent := g.GameWorld.Manager.NewComponent()
	g.GameWorld.Manager.NewEntity().AddComponent(playerComponent, &Player{})
	playerTags := ecs.BuildTag(playerComponent)
	g.GameWorld.Tags[PlayerEntity] = playerTags
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
