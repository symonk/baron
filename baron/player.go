package main

const (
	Mage    = "Mage"
	Warrior = "Warrior"
	Ranger  = "Ranger"
)

type Weapon struct {
}

type Armor struct {
}

type PowerUp struct {
}

type Player struct {
	x       int
	y       int
	class   string
	weapon  Weapon
	armor   Armor
	speed   int
	powerUp *PowerUp
}
