package game

type Player struct {
	MaxHealth     int
	CurrentHealth int
}

func (p *Player) Move() {

}

func NewPlayer() *Player {
	return &Player{
		MaxHealth:     100,
		CurrentHealth: 100,
	}
}
