package Armors

import (
	"GoPractice/Equipments"
	"GoPractice/Stats"
)

type NoArmor struct {
	attributes map[Stats.StatType]float32
	ArmorInterface
	Equipments.EquipmentInterface
}

func (armor *NoArmor) GetArmorType() ArmorType {
	return NONE
}

func (armor *NoArmor) GetAttributes() map[Stats.StatType]float32 {
	return armor.attributes
}

func (armor *NoArmor) GetAttribute(statType Stats.StatType) float32 {
	return armor.attributes[statType]
}

func MakeNoArmor() *NoArmor {
	attributes := map[Stats.StatType]float32{
		Stats.EVASION: 0.10,
	}

	return &NoArmor{attributes: attributes}
}
