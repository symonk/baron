package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadBackgroundImage() *ebiten.Image {
	name := "./assets/background.png"
	background, _, err := ebitenutil.NewImageFromFile(name)
	if err != nil {
		panic(err)
	}
	return background
}

func LoadLevelBackground(level int) *ebiten.Image {
	path := fmt.Sprintf("./assets/%d/.png", level)
	bg, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(err)
	}
	return bg
}
