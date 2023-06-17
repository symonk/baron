package main

import (
	"github.com/hajimehoshi/ebiten"
)

func main() {
	ebiten.SetWindowTitle(windowTitle)
	ebiten.SetWindowSize(windowWidth, windowHeight)
	baron := &Game{}
	if err := ebiten.RunGame(baron); err != nil {
		panic(err)
	}
}
