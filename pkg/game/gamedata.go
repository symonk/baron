package game

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
}

func NewGameData() GameData {
	return GameData{
		ScreenWidth:  40,
		ScreenHeight: 25,
		TileWidth:    32,
		TileHeight:   32}
}
