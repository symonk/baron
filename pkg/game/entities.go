package game

import "github.com/bytearena/ecs"

/*
What are Entities?
What are Components?
What are Tags?
*/

// Register all of the ECS entities.
func RegisterEntities(game *Game) {
	registerPlayer(game)

}

// Register the core player to the game world
func registerPlayer(g *Game) *Player {
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
