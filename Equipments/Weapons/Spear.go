package Weapons

import (
	"GoPractice/Equipments"
	"GoPractice/Stats"
)

type Spear struct {
	attributes map[Stats.StatType]float32
	Equipments.EquipmentInterface
	WeaponInterface
}

func (weapon *Spear) GetWeaponType() WeaponType {
	return SPEAR
}

func (weapon *Spear) GetAttributes() map[Stats.StatType]float32 {
	return weapon.attributes
}

func (weapon *Spear) GetAttribute(statType Stats.StatType) float32 {
	return weapon.attributes[statType]
}

func MakeSpear() *Spear {
	attributes := map[Stats.StatType]float32{
		Stats.DEFENSE: 5.00,
		Stats.EVASION: 0.10,
	}

	return &Spear{attributes: attributes}
}
