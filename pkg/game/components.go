package game

import "github.com/hajimehoshi/ebiten/v2"

type Position struct {
	X, Y int
}

type Renderable struct {
	Image *ebiten.Image
}

type Movable struct{}

func (m *Movable) Move() {

}
