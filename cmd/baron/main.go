package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/symonk/baron/pkg/game"
)

func main() {
	gameInstance := game.NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle(game.GameName)
	ebiten.SetTPS(30)
	ebiten.SetWindowSize(1280, 1024)
	if err := ebiten.RunGame(gameInstance); err != nil {
		panic(err)
	}
}
