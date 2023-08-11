package main

type GameMap struct {
	Dungeons     []Dungeon
	CurrentLevel Level
}

func NewGameMap() GameMap {
	level := NewLevel()
	levels := []Level{level}
	dungeon := Dungeon{Name: "Default", Levels: levels}
	dungeons := []Dungeon{dungeon}
	gameMap := GameMap{Dungeons: dungeons, CurrentLevel: level}
	return gameMap
}
