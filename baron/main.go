package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("./assets/background.png")
	if err != nil {
		log.Fatal(err)
	}
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
