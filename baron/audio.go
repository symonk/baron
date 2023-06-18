package main

import (
	"io"
	"time"

	"github.com/hajimehoshi/ebiten/audio"
)

type audioStream interface {
	io.ReadSeeker
	Length() int64
}

type AudioPlayer struct {
	game         *Game
	audioContext *audio.Context
	audioPlayer  *audio.Player
	current      time.Duration
	total        time.Duration
	seBytes      []byte
	volume128    int
}

func NewAudioPlayer(game *Game, audioContext *audio.Context) (*AudioPlayer, error) {
	var stream audioStream
	stream, err = wav.
	player, err := audioContext.NewPlayer(s)
}

func (p *AudioPlayer) Close() error {
	return p.audioPlayer.Close()
}
