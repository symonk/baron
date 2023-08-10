package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	Tiles []MapTile
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1920, 1280
}