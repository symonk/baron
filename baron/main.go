package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle(GameName)
	ebiten.SetWindowSize(1024, 768)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
