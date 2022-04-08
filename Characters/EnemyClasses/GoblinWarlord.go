package EnemyClasses

import (
	stats "GoPractice/Stats"
)

type GoblinWarlord struct {
	Enemy
}

func (goblin *GoblinWarlord) GetEnemyName() string {
	return "Goblin Warlord"
}

func (goblin *GoblinWarlord) GetClassType() EnemyType {
	return GOBLIN_WARLORD
}

func MakeGoblinWarlord() *GoblinWarlord {
	enemy := MakeEnemy(
		map[stats.StatType]float32{
			stats.CRIT:    0.05,
			stats.DMG:     10.00,
			stats.DEFENSE: 7.00,
			stats.EVASION: 0.00,
			stats.HP:      72.00,
			stats.LUCK:    0.05,
		},
	)
	return &GoblinWarlord{*enemy}
}
