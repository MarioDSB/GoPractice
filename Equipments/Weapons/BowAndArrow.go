package Weapons

import (
	"GoPractice/Equipments"
	"GoPractice/Stats"
)

type BowAndArrow struct {
	attributes map[Stats.StatType]float32
	Equipments.EquipmentInterface
	WeaponInterface
}

func (weapon *BowAndArrow) GetWeaponType() WeaponType {
	return BOW_AND_ARROW
}

func (weapon *BowAndArrow) GetAttributes() map[Stats.StatType]float32 {
	return weapon.attributes
}

func (weapon *BowAndArrow) GetAttribute(statType Stats.StatType) float32 {
	return weapon.attributes[statType]
}

func MakeBowAndArrow() *BowAndArrow {
	attributes := map[Stats.StatType]float32{
		Stats.CRIT: 0.10,
		Stats.DMG:  5.00,
	}

	return &BowAndArrow{attributes: attributes}
}
