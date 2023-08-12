package world

import (
	"github.com/bytearena/ecs"
)

type World struct {
	Manager *ecs.Manager
	Tags    map[string]ecs.Tag
}

func NewWorld() *World {
	tags := make(map[string]ecs.Tag)
	return &World{
		Manager: ecs.NewManager(),
		Tags:    tags,
	}
}

func (g *World) Init() {

}
