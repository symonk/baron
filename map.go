package main

type GameMap struct {
	Dungeons []Dungeon
}

func NewGameMap() GameMap {
	level := NewLevel()
	levels := []Level{level}
	dungeon := Dungeon{Name: "Default", Levels: levels}
	dungeons := []Dungeon{dungeon}
	gameMap := GameMap{Dungeons: dungeons}
	return gameMap
}
