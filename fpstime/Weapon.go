package main

import "fmt"

type Weapon struct {
	name string
	ammo int8
	dmg  int8
}

func (w Weapon) GetName() string {
	return w.name
}

func (w Weapon) GetAmmo() int8 {
	return w.ammo
}

func (w Weapon) GetDmg() int8 {
	return w.dmg
}

func (w Weapon) UseWeapon(player *Player) {
	player.doHurt(w.GetDmg())
}

func (w Weapon) String() string {
	return fmt.Sprintf("%s: Ammo:%d | Damage:%d |", w.name, w.ammo, w.dmg)
}
