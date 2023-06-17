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

type Player struct {
	x      int64
	y      int64
	weapon Weapon
	armor  Armor
}
