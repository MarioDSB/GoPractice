package EnemyClasses

import (
	stats "GoPractice/Stats"
)

type EnemyInterface interface {
	GetAttributes() map[stats.StatType]float32
	GetAttribute(statType stats.StatType) float32
	GetClassType() EnemyType
	GetEnemyName() string
	Injured(dmg float32)
	GetEnemyStatus() string
}
