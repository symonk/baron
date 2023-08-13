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
	Map       GameMap
	GameWorld *world.World
	ticks     int
	Player    *Player
}

func NewGame() *Game {
	game := &Game{Map: NewGameMap(), GameWorld: world.NewWorld(), Player: NewPlayer()}
	game.registerPlayerToWorld()
	return game
}

// Create an entity and register our player as a componeent.
func (g *Game) registerPlayerToWorld() {
	playerEntity := NewPlayer()
	position = g.GameWorld.Manager.NewComponent()
	renderables = g.GameWorld.Manager.NewComponent()
	player := g.GameWorld.Manager.NewComponent()
	g.GameWorld.Manager.NewEntity().AddComponent(player, playerEntity).AddComponent(renderables, &Renderable{Image: LoadImageFromAssets("player.png")}).AddComponent(position, &Position{X: 20, Y: 20})
	players := ecs.BuildTag(player, position)
	g.GameWorld.Tags[PlayersView] = players
	renderables := ecs.BuildTag(renderables, position)
	g.GameWorld.Tags[RenderablesView] = renderables
	g.Player = playerEntity
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

func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Map.CurrentLevel
	level.DrawLevel(screen)
	DrawEntity(g, level, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	gd := NewGameData()
	return gd.ScreenWidth * gd.TileWidth, gd.ScreenHeight * gd.TileHeight
}
