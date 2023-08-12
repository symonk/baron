package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	MaxHealth     int
	CurrentHealth int
	Image         *ebiten.Image
}

func (p *Player) Move() {

}

func NewPlayer() *Player {
	baseImage, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")
	if err != nil {
		panic(err)
	}
	return &Player{
		MaxHealth:     100,
		CurrentHealth: 100,
		Image:         baseImage,
	}
}
