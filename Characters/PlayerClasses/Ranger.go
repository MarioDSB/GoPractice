package PlayerClasses

import stats "GoPractice/Stats"

type Ranger struct {
	attributes map[stats.StatType]float32
	CharacterInterface
}

func (ranger *Ranger) GetAttributes() map[stats.StatType]float32 {
	return ranger.attributes
}

func (ranger *Ranger) GetClassType() ClassType {
	return RANGER
}

func (ranger *Ranger) SetAttribute(statType stats.StatType, value float32) {
	ranger.attributes[statType] = value
}

func MakeRanger() *Ranger {
	attributes := map[stats.StatType]float32{
		stats.CRIT:    0.20,
		stats.DMG:     10.00,
		stats.DEFENSE: 3.00,
		stats.EVASION: 0.33,
		stats.HP:      60.00,
		stats.LUCK:    0.05,
	}
	return &Ranger{attributes: attributes}
}
