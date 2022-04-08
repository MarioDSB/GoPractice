package PlayerClasses

import stats "GoPractice/Stats"

type Assassin struct {
	attributes map[stats.StatType]float32
	CharacterInterface
}

func (assassin *Assassin) GetAttributes() map[stats.StatType]float32 {
	return assassin.attributes
}

func (assassin *Assassin) SetAttribute(statType stats.StatType, value float32) {
	assassin.attributes[statType] = value
}

func (assassin *Assassin) GetClassType() ClassType {
	return ASSASSIN
}

func MakeAssassin() *Assassin {
	attributes := map[stats.StatType]float32{
		stats.CRIT:    0.33,
		stats.DMG:     7.00,
		stats.DEFENSE: 3.00,
		stats.EVASION: 0.20,
		stats.HP:      80.00,
		stats.LUCK:    0.05,
	}
	return &Assassin{attributes: attributes}
}
