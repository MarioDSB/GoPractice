package Armors

import (
	"GoPractice/Equipments"
	"GoPractice/Stats"
)

type LightArmor struct {
	attributes map[Stats.StatType]float32
	ArmorInterface
	Equipments.EquipmentInterface
}

func (armor *LightArmor) GetArmorType() ArmorType {
	return LIGHT
}

func (armor *LightArmor) GetAttributes() map[Stats.StatType]float32 {
	return armor.attributes
}

func (armor *LightArmor) GetAttribute(statType Stats.StatType) float32 {
	return armor.attributes[statType]
}

func MakeLightArmor() *LightArmor {
	attributes := map[Stats.StatType]float32{
		Stats.DEFENSE: 3.00,
		Stats.EVASION: 0.05,
	}

	return &LightArmor{attributes: attributes}
}
