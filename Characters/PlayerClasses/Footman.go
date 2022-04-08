package PlayerClasses

import stats "GoPractice/Stats"

type Footman struct {
	attributes map[stats.StatType]float32
	CharacterInterface
}

func (footman *Footman) GetAttributes() map[stats.StatType]float32 {
	return footman.attributes
}

func (footman *Footman) GetClassType() ClassType {
	return FOOTMAN
}

func (footman *Footman) SetAttribute(statType stats.StatType, value float32) {
	footman.attributes[statType] = value
}

func MakeFootman() *Footman {
	attributes := map[stats.StatType]float32{
		stats.CRIT:    0.10,
		stats.DMG:     10.00,
		stats.DEFENSE: 5.00,
		stats.EVASION: 0.10,
		stats.HP:      100.00,
		stats.LUCK:    0.05,
	}
	return &Footman{attributes: attributes}
}
