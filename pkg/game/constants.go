package game

import "github.com/bytearena/ecs"

const (
	GameName     = "Baron"
	PlayerEntity = "player"
	Renderables  = "renderables"
)

var (
	globalPosition   *ecs.Component
	globalRenderable *ecs.Component
)
