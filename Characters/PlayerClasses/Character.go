package PlayerClasses

import stats "GoPractice/Stats"

type CharacterInterface interface {
	GetAttributes() map[stats.StatType]float32
	SetAttribute(statType stats.StatType, value float32)
	GetClassType() ClassType
}
