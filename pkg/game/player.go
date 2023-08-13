package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerStats struct {
	MaxHealth     int
	CurrentHealth int
	MaxMana       int
	CurrentMana   int
}

type Player struct {
	Image *ebiten.Image
	Stats PlayerStats
	CurrX int
	CurrY int
}

func (p *Player) Move() {

}

func NewPlayer() *Player {
	baseImage := LoadImageFromAssets("player.png")
	return &Player{
		Stats: PlayerStats{MaxHealth: 100, CurrentHealth: 100, MaxMana: 100, CurrentMana: 100},
		Image: baseImage,
	}
}
