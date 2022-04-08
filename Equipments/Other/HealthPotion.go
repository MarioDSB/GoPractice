package Other

import (
	"GoPractice/Equipments"
	"GoPractice/Stats"
)

type HealthPotion struct {
	attributes map[Stats.StatType]float32
	Equipments.EquipmentInterface
}

func (potion *HealthPotion) GetAttributes() map[Stats.StatType]float32 {
	return potion.attributes
}

func (potion *HealthPotion) GetAttribute(statType Stats.StatType) float32 {
	return potion.attributes[statType]
}

func MakeHealthPotion() *HealthPotion {
	attributes := map[Stats.StatType]float32{
		Stats.HP: 20.0,
	}

	return &HealthPotion{attributes: attributes}
}
