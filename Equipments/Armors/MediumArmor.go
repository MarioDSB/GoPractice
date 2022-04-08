package Armors

import (
	"GoPractice/Equipments"
	"GoPractice/Stats"
)

type MediumArmor struct {
	attributes map[Stats.StatType]float32
	ArmorInterface
	Equipments.EquipmentInterface
}

func (armor *MediumArmor) GetArmorType() ArmorType {
	return MEDIUM
}

func (armor *MediumArmor) GetAttributes() map[Stats.StatType]float32 {
	return armor.attributes
}

func (armor *MediumArmor) GetAttribute(statType Stats.StatType) float32 {
	return armor.attributes[statType]
}

func MakeMediumArmor() *MediumArmor {
	attributes := map[Stats.StatType]float32{
		Stats.DEFENSE: 5.00,
	}

	return &MediumArmor{attributes: attributes}
}
