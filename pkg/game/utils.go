package game

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func GetRandomInt(num int) int {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
	return int(x.Int64())
}

func GetDiceRoll(num int) int {
	return GetRandomInt(num) + 1
}

// Loads an image from the assets folder.
// Todo: Stop using relative pathing!
func LoadImageFromAssets(name string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("assets/images/%s", name))
	if err != nil {
		panic(err)
	}
	return img
}
