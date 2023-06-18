package main

import (
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	audioPlayer   *AudioPlayer
	audioPlayerCh chan *AudioPlayer
	errCh         chan *error
}

func (g *Game) Update() error {
	// Update logical state
	return nil
}

func NewGame() (*Game, error) {
	audioContext := audio.NewContext(sampleRate)
	g := &Game{
		audioPlayerCh:  make(chan **Player)
		errorCh: make(chan error),
	}
	player, err := AudioPlayer{}
	g.audioPlayer = playe
	return g, nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	screen.DrawImage(loadingBackgroundImg, options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}
