package Weapons

import (
	"GoPractice/Equipments"
	"GoPractice/Stats"
)

type TwoHandedSword struct {
	attributes map[Stats.StatType]float32
	Equipments.EquipmentInterface
	WeaponInterface
}

func (weapon *TwoHandedSword) GetWeaponType() WeaponType {
	return TWO_HANDED_SWORD
}

func (weapon *TwoHandedSword) GetAttributes() map[Stats.StatType]float32 {
	return weapon.attributes
}

func (weapon *TwoHandedSword) GetAttribute(statType Stats.StatType) float32 {
	return weapon.attributes[statType]
}

func MakeTwoHandedSword() *TwoHandedSword {
	attributes := map[Stats.StatType]float32{
		Stats.DMG: 10.00,
	}

	return &TwoHandedSword{attributes: attributes}
}
