package EnemyClasses

import (
	stats "GoPractice/Stats"
)

type CommonGoblin struct {
	Enemy
}

func (goblin *CommonGoblin) GetEnemyName() string {
	return "Common Goblin"
}

func (goblin *CommonGoblin) GetClassType() EnemyType {
	return COMMON_GOBLIN
}

func MakeCommonGoblin() *CommonGoblin {
	enemy := MakeEnemy(
		map[stats.StatType]float32{
			stats.CRIT:    0.10,
			stats.DMG:     4.00,
			stats.DEFENSE: 2.00,
			stats.EVASION: 0.10,
			stats.HP:      30.00,
			stats.LUCK:    0.05,
		},
	)

	return &CommonGoblin{*enemy}
}
