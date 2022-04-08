package Other

import (
	"GoPractice/Equipments"
	"GoPractice/Stats"
)

type HealthPotion struct {
	stats Stats.Stats
	Equipments.EquipmentInterface
}

func (potion *HealthPotion) GetStats() *Stats.Stats {
	return &potion.stats
}

func MakeHealthPotion() *HealthPotion {
	potion := Stats.MakeStats(
		map[Stats.StatType]float32{
			Stats.HP: 20.0,
		},
	)

	return &HealthPotion{stats: *potion}
}
