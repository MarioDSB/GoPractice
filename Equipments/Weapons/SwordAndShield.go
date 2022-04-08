package Weapons

import (
	"GoPractice/Stats"
)

type SwordAndShield struct {
	Weapon
}

func (weapon *SwordAndShield) GetWeaponType() WeaponType {
	return SWORD_SHIELD
}

func MakeSwordAndShield() *SwordAndShield {
	weapon := MakeWeapon(
		map[Stats.StatType]float32{
			Stats.DEFENSE: 5.00,
			Stats.DMG:     5.00,
		},
	)

	return &SwordAndShield{*weapon}
}
