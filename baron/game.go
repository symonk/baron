package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	isMenu bool
}

func (g *Game) Update() error {
	// Update logical state
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	screen.DrawImage(loadingBackgroundImg, options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}
