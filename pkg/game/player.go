package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	MaxHealth     int
	CurrentHealth int
	Image         *ebiten.Image
	CurrX         int
	CurrY         int
}

func (p *Player) Move() {

}

func NewPlayer() *Player {
	baseImage := LoadImageFromAssets("player.png")
	return &Player{
		MaxHealth:     100,
		CurrentHealth: 100,
		Image:         baseImage,
	}
}
