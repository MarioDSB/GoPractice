package Weapons

import (
	"GoPractice/Stats"
)

type TwoHandedSword struct {
	Weapon
}

func (weapon *TwoHandedSword) GetWeaponType() WeaponType {
	return TWO_HANDED_SWORD
}

func MakeTwoHandedSword() *TwoHandedSword {
	weapon := MakeWeapon(
		map[Stats.StatType]float32{
			Stats.DMG: 10.00,
		},
	)

	return &TwoHandedSword{*weapon}
}
