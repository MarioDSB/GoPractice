package Weapons

import (
	"GoPractice/Equipments"
	"GoPractice/Stats"
)

type SwordAndShield struct {
	attributes map[Stats.StatType]float32
	Equipments.EquipmentInterface
	WeaponInterface
}

func (weapon *SwordAndShield) GetWeaponType() WeaponType {
	return SWORD_SHIELD
}

func (weapon *SwordAndShield) GetAttributes() map[Stats.StatType]float32 {
	return weapon.attributes
}

func (weapon *SwordAndShield) GetAttribute(statType Stats.StatType) float32 {
	return weapon.attributes[statType]
}

func MakeSwordAndShield() *SwordAndShield {
	attributes := map[Stats.StatType]float32{
		Stats.DEFENSE: 5.00,
		Stats.DMG:     5.00,
	}

	return &SwordAndShield{attributes: attributes}
}
