package Armors

import (
	"GoPractice/Stats"
)

type HeavyArmor struct {
	Armor
}

func (armor *HeavyArmor) GetArmorType() ArmorType {
	return HEAVY
}

func MakeHeavyArmor() *HeavyArmor {
	armor := MakeArmor(
		map[Stats.StatType]float32{
			Stats.DEFENSE: 5.00,
			Stats.HP:      20.00,
			Stats.EVASION: -0.05,
		},
	)

	return &HeavyArmor{*armor}
}
