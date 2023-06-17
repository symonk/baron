package main

import (
	"github.com/hajimehoshi/ebiten"
)

func main() {
	ebiten.SetWindowTitle(window_title)
	ebiten.SetWindowSize(window_width, window_height)
	baron := &Game{}
	if err := ebiten.RunGame(baron); err != nil {
		panic(err)
	}
}
