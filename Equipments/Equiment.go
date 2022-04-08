package Equipments

import stats "GoPractice/Stats"

type EquipmentInterface interface {
	GetAttributes() map[stats.StatType]float32
	GetAttribute(statType stats.StatType) float32
}
