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
	ebiten.SetTPS(144)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	baron := &NewGame{}
	if err := ebiten.RunGame(baron); err != nil {
		panic(err)
	}
}
