package EnemyClasses

import (
	stats "GoPractice/Stats"
)

type Wolf struct {
	Enemy
}

func (wolf *Wolf) GetEnemyName() string {
	return "Wolf"
}

func (wolf *Wolf) GetClassType() EnemyType {
	return WOLF
}

func MakeWolf() *Wolf {
	enemy := MakeEnemy(
		map[stats.StatType]float32{
			stats.CRIT:    0.10,
			stats.DMG:     4.00,
			stats.DEFENSE: 0.00,
			stats.EVASION: 0.10,
			stats.HP:      20.00,
			stats.LUCK:    0.05,
		},
	)

	return &Wolf{*enemy}
}
