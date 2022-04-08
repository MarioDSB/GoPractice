package Weapons

import (
	"GoPractice/Stats"
)

type BowAndArrow struct {
	Weapon
}

func (weapon *BowAndArrow) GetWeaponType() WeaponType {
	return BOW_AND_ARROW
}

func MakeBowAndArrow() *BowAndArrow {
	weapon := MakeWeapon(
		map[Stats.StatType]float32{
			Stats.CRIT: 0.10,
			Stats.DMG:  5.00,
		},
	)

	return &BowAndArrow{*weapon}
}
