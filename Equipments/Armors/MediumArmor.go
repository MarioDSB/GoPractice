package Armors

import (
	"GoPractice/Stats"
)

type MediumArmor struct {
	Armor
}

func (armor *MediumArmor) GetArmorType() ArmorType {
	return MEDIUM
}

func MakeMediumArmor() *MediumArmor {
	armor := MakeArmor(
		map[Stats.StatType]float32{
			Stats.DEFENSE: 5.00,
		},
	)

	return &MediumArmor{*armor}
}
