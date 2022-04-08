package PlayerClasses

import stats "GoPractice/Stats"

type Knight struct {
	attributes map[stats.StatType]float32
	CharacterInterface
}

func (knight *Knight) GetAttributes() map[stats.StatType]float32 {
	return knight.attributes
}

func (knight *Knight) GetClassType() ClassType {
	return KNIGHT
}

func (knight *Knight) SetAttribute(statType stats.StatType, value float32) {
	knight.attributes[statType] = value
}

func MakeKnight() *Knight {
	attributes := map[stats.StatType]float32{
		stats.CRIT:    0.33,
		stats.DMG:     7.00,
		stats.DEFENSE: 3.00,
		stats.EVASION: 0.20,
		stats.HP:      80.00,
		stats.LUCK:    0.05,
	}
	return &Knight{attributes: attributes}
}
