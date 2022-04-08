package Equipments

import "GoPractice/Stats"

type EquipmentInterface interface {
	GetStats() *Stats.Stats
}
