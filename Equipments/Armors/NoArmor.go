package Armors

import (
	"GoPractice/Stats"
)

type NoArmor struct {
	Armor
}

func (armor *NoArmor) GetArmorType() ArmorType {
	return NONE
}

func MakeNoArmor() *NoArmor {
	armor := MakeArmor(
		map[Stats.StatType]float32{
			Stats.EVASION: 0.10,
		},
	)

	return &NoArmor{*armor}
}
