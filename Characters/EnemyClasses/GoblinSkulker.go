package EnemyClasses

import (
	stats "GoPractice/Stats"
)

type GoblinSkulker struct {
	Enemy
}

func (goblin *GoblinSkulker) GetEnemyName() string {
	return "Goblin Skulker"
}

func (goblin *GoblinSkulker) GetClassType() EnemyType {
	return GOBLIN_SKULKER
}

func MakeGoblinSkulker() *GoblinSkulker {
	enemy := MakeEnemy(
		map[stats.StatType]float32{
			stats.CRIT:    0.20,
			stats.DMG:     4.00,
			stats.DEFENSE: 0.00,
			stats.EVASION: 0.10,
			stats.HP:      30.00,
			stats.LUCK:    0.05,
		},
	)
	return &GoblinSkulker{*enemy}
}
