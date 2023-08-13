package game

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/symonk/baron/pkg/world"
)

var (
	position    *ecs.Component
	renderables *ecs.Component
)

type Game struct {
	Map    GameMap
	World  *world.World
	ticks  int
	Player *Player
}

func NewGame() *Game {
	game := &Game{Map: NewGameMap(), World: world.NewWorld(), Player: NewPlayer()}
	game.Init()
	return game
}

func (g *Game) Init() {
	RegisterEntities(g)
}

func (g *Game) Update() error {
	g.ticks++
	if !g.Player.IsDead {
		MovePlayer(g)
	} else {
		ResetPlayer(g)
	}
	return nil
}

func (g *Game) AddTag(key string, elements ...any) {
	tag := ecs.BuildTag(elements...)
	g.World.Tags[key] = tag

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
