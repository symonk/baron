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
	Image  *ebiten.Image
	Stats  *PlayerStats
	CurrX  int
	CurrY  int
	IsDead bool
	Deaths int
}

func (p *Player) Move() {

}

func (p *Player) Revive() {
	p.Stats.CurrentHealth = 100
	p.Stats.MaxHealth = 100
	p.Stats.CurrentMana = 100
	p.Stats.MaxMana = 100
	p.IsDead = false
}

// Apply a damaging effect to the player
func (p *Player) TransformHealth(diff int) {
	p.Stats.CurrentHealth += diff
	if p.Stats.CurrentHealth < 1 {
		p.Deaths++
		p.IsDead = true
	}
}

func NewPlayer() *Player {
	baseImage := LoadImageFromAssets("player.png")
	return &Player{
		Stats: &PlayerStats{MaxHealth: 100, CurrentHealth: 100, MaxMana: 100, CurrentMana: 100},
		Image: baseImage,
	}
}
