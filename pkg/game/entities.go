package game

import "github.com/bytearena/ecs"

/*
What are Entities?
What are Components?
What are Tags?
*/

// Register all of the ECS entities.
func RegisterEntities(game *Game) {
	renderable := game.World.Manager.NewComponent()
	registerPlayer(game, renderable)
	RegisterBosses(game, renderable)

}

// Register the core player to the game world
func registerPlayer(g *Game, renderable *ecs.Component) *Player {
	playerEntity := NewPlayer()
	position = g.World.Manager.NewComponent()
	renderables = g.World.Manager.NewComponent()
	player := g.World.Manager.NewComponent()
	g.World.Manager.NewEntity().AddComponent(player, playerEntity).AddComponent(renderables, &Renderable{Image: LoadImageFromAssets("player.png")}).AddComponent(position, &Position{X: 20, Y: 20})
	players := ecs.BuildTag(player, position)
	g.World.Tags[PlayersView] = players
	renderables := ecs.BuildTag(renderables, position)
	g.World.Tags[RenderablesView] = renderables
	return playerEntity
}

func RegisterBosses(g *Game, renderable *ecs.Component) {
	bossObject := &Monster{image: LoadImageFromAssets("boss1.png")}
	bossComponent := g.World.Manager.NewComponent()
	bossRenderable := g.World.Manager.NewComponent()
	bossPosition := g.World.Manager.NewComponent()
	g.World.Manager.NewEntity().AddComponent(bossComponent, bossObject).AddComponent(bossRenderable, &Renderable{Image: LoadImageFromAssets("boss1.png")}).AddComponent(bossPosition, &Position{X: 20, Y: 5})
	bosses := ecs.BuildTag(bossComponent, position)
	g.World.Tags[BossesView] = bosses
}
