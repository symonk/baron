package main

import (
	"github.com/bytearena/ecs"
)

func InitializeWorld() (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()
	return manager, tags
}
