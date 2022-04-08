package Armors

import (
	"GoPractice/Equipments"
	"GoPractice/Stats"
)

type ArmorInterface interface {
	GetArmorType() ArmorType
}

type Armor struct {
	stats Stats.Stats
	ArmorInterface
	Equipments.EquipmentInterface
}

func (armor *Armor) GetStats() *Stats.Stats {
	return &armor.stats
}

func MakeArmor(attributes map[Stats.StatType]float32) *Armor {
	stats := Stats.MakeStats(attributes)
	return &Armor{stats: *stats}
}
