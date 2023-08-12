package game

type GameMap struct {
	CurrentLevel Level
}

func NewGameMap() GameMap {
	level := NewLevel()
	gameMap := GameMap{CurrentLevel: level}
	return gameMap
}
