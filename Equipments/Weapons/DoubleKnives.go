package Weapons

import (
	"GoPractice/Stats"
)

type DoubleKnives struct {
	Weapon
}

func (weapon *DoubleKnives) GetWeaponType() WeaponType {
	return DOUBLE_KNIVES
}

func MakeDoubleKnives() *DoubleKnives {
	weapon := MakeWeapon(
		map[Stats.StatType]float32{
			Stats.CRIT:    0.10,
			Stats.EVASION: 0.10,
		},
	)

	return &DoubleKnives{*weapon}
}
