package Armors

import (
	"GoPractice/Stats"
)

type LightArmor struct {
	Armor
}

func (armor *LightArmor) GetArmorType() ArmorType {
	return LIGHT
}

func MakeLightArmor() *LightArmor {
	armor := MakeArmor(
		map[Stats.StatType]float32{
			Stats.DEFENSE: 3.00,
			Stats.EVASION: 0.05,
		},
	)

	return &LightArmor{*armor}
}
