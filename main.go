package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle(GameName)
	ebiten.SetWindowSize(1280, 800)
	ebiten.SetTPS(144)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
