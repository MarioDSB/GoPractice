package Weapons

import (
	"GoPractice/Stats"
)

type Spear struct {
	Weapon
}

func (weapon *Spear) GetWeaponType() WeaponType {
	return SPEAR
}

func MakeSpear() *Spear {
	weapon := MakeWeapon(
		map[Stats.StatType]float32{
			Stats.DEFENSE: 5.00,
			Stats.EVASION: 0.10,
		},
	)

	return &Spear{*weapon}
}
