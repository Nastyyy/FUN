package main

import "fmt"

type Player struct {
	Name   string
	Id     int64
	health int8
	Weapon Weapon
}

func (p *Player) doHurt(dmg int8) int8 {
	p.health -= dmg
	return p.health
}

func (p Player) String() string {
	return fmt.Sprintf("%s: %d | %d | %s", p.Name, p.Id, p.health, p.Weapon)
}
