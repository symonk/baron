package world

import (
	"github.com/bytearena/ecs"
)

type World struct {
	World     *ecs.Manager
	WorldTags map[string]ecs.Tag
}

func NewWorld() *World {
	tags := make(map[string]ecs.Tag)
	return &World{
		World:     ecs.NewManager(),
		WorldTags: tags,
	}
}

func (g *World) Init() {

}
