package game

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	MaxHealth     int
	CurrentHealth int
}

func NewPlayer() *Player {
	return &Player{
		MaxHealth:     100,
		CurrentHealth: 100,
	}
}

type Position struct {
	X, Y int
}

type Renderable struct {
	Image *ebiten.Image
}

type Movable struct{}
