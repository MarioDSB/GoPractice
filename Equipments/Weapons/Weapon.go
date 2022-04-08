package Weapons

import (
	"GoPractice/Equipments"
	"GoPractice/Stats"
)

type WeaponInterface interface {
	GetWeaponType() WeaponType
}

type Weapon struct {
	stats Stats.Stats
	WeaponInterface
	Equipments.EquipmentInterface
}

func (weapon *Weapon) GetStats() *Stats.Stats {
	return &weapon.stats
}

func MakeWeapon(attributes map[Stats.StatType]float32) *Weapon {
	stats := Stats.MakeStats(attributes)
	return &Weapon{stats: *stats}
}
