package game

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	globalPosition   *ecs.Component
	globalRenderable *ecs.Component
)

func InitializeWorld() (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()
	player := manager.NewComponent()
	globalPosition = manager.NewComponent()
	globalRenderable = manager.NewComponent()
	movable := manager.NewComponent()

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")
	if err != nil {
		panic(err)
	}

	manager.NewEntity().AddComponent(player, Player{}).AddComponent(globalRenderable, &Renderable{Image: playerImg}).AddComponent(movable, Movable{}).AddComponent(globalPosition, &Position{X: 40, Y: 25})
	players := ecs.BuildTag(player, globalPosition)
	tags["players"] = players
	renderables := ecs.BuildTag(globalRenderable, globalPosition)
	tags["renderables"] = renderables
	return manager, tags
}
