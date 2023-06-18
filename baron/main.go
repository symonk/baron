package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var loadingBackgroundImg *ebiten.Image

func init() {
	loadingBackgroundImg = LoadBackgroundImage()
}

func main() {
	ebiten.SetWindowTitle(windowTitle)
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	baron := &Game{}
	if err := ebiten.RunGame(baron); err != nil {
		panic(err)
	}
}
