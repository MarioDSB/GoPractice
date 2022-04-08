package PlayerClasses

import stats "GoPractice/Stats"

type Footman struct {
	Character
}

func (footman *Footman) GetClassType() ClassType {
	return FOOTMAN
}

func MakeFootman() *Footman {
	character := MakeCharacter(
		map[stats.StatType]float32{
			stats.CRIT:    0.10,
			stats.DMG:     10.00,
			stats.DEFENSE: 5.00,
			stats.EVASION: 0.10,
			stats.HP:      100.00,
			stats.LUCK:    0.05,
		},
	)
	return &Footman{*character}
}
