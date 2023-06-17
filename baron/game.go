package main

import "github.com/hajimehoshi/ebiten"

type Game struct{}

func (g *Game) Update(image *ebiten.Image) error {
	// Update logical state
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Render the screen
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return window_width, window_height
}
