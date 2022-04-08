package PlayerClasses

import stats "GoPractice/Stats"

type Assassin struct {
	Character
}

func (assassin *Assassin) GetClassType() ClassType {
	return ASSASSIN
}

func MakeAssassin() *Assassin {
	character := MakeCharacter(
		map[stats.StatType]float32{
			stats.CRIT:    0.33,
			stats.DMG:     7.00,
			stats.DEFENSE: 3.00,
			stats.EVASION: 0.20,
			stats.HP:      80.00,
			stats.LUCK:    0.05,
		})
	return &Assassin{*character}
}
