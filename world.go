package main

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func InitializeWorld() (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()
	player := manager.NewComponent()
	position := manager.NewComponent()
	renderable := manager.NewComponent()
	movable := manager.NewComponent()

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")
	if err != nil {
		panic(err)
	}

	manager.NewEntity().AddComponent(player, Player{}).AddComponent(renderable, &Renderable{Image: playerImg}).AddComponent(movable, Movable{}).AddComponent(position, &Position{X: 40, Y: 25})
	players := ecs.BuildTag(player, position)
	tags["players"] = players
	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables
	return manager, tags
}
