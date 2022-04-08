package PlayerClasses

import stats "GoPractice/Stats"

type Ranger struct {
	Character
}

func (ranger *Ranger) GetClassType() ClassType {
	return RANGER
}

func MakeRanger() *Ranger {
	character := MakeCharacter(
		map[stats.StatType]float32{
			stats.CRIT:    0.20,
			stats.DMG:     10.00,
			stats.DEFENSE: 3.00,
			stats.EVASION: 0.33,
			stats.HP:      60.00,
			stats.LUCK:    0.05,
		},
	)
	return &Ranger{*character}
}
