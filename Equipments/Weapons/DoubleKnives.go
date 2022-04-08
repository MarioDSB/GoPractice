package Weapons

import (
	"GoPractice/Equipments"
	"GoPractice/Stats"
)

type DoubleKnives struct {
	attributes map[Stats.StatType]float32
	Equipments.EquipmentInterface
	WeaponInterface
}

func (weapon *DoubleKnives) GetWeaponType() WeaponType {
	return DOUBLE_KNIVES
}

func (weapon *DoubleKnives) GetAttributes() map[Stats.StatType]float32 {
	return weapon.attributes
}

func (weapon *DoubleKnives) GetAttribute(statType Stats.StatType) float32 {
	return weapon.attributes[statType]
}

func MakeDoubleKnives() *DoubleKnives {
	attributes := map[Stats.StatType]float32{
		Stats.CRIT:    0.10,
		Stats.EVASION: 0.10,
	}

	return &DoubleKnives{attributes: attributes}
}
