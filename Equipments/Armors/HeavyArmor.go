package Armors

import (
	"GoPractice/Equipments"
	"GoPractice/Stats"
)

type HeavyArmor struct {
	attributes map[Stats.StatType]float32
	ArmorInterface
	Equipments.EquipmentInterface
}

func (armor *HeavyArmor) GetArmorType() ArmorType {
	return HEAVY
}

func (armor *HeavyArmor) GetAttributes() map[Stats.StatType]float32 {
	return armor.attributes
}

func (armor *HeavyArmor) GetAttribute(statType Stats.StatType) float32 {
	return armor.attributes[statType]
}

func MakeHeavyArmor() *HeavyArmor {
	attributes := map[Stats.StatType]float32{
		Stats.DEFENSE: 5.00,
		Stats.HP:      20.00,
		Stats.EVASION: -0.05,
	}

	return &HeavyArmor{attributes: attributes}
}
