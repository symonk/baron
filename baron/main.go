package baron

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	ebiten.DebugPrint(screen, "Baron")
	return nil
}

func main() {
	if err := ebiten.Run(update, 640, 640, 2, "Baron"); err != nil {
		log.Fatal(err)
	}
}
