package EnemyClasses

import (
	stats "GoPractice/Stats"
)

type GoblinSoldier struct {
	Enemy
}

func (goblin *GoblinSoldier) GetEnemyName() string {
	return "Goblin Soldier"
}

func (goblin *GoblinSoldier) GetClassType() EnemyType {
	return GOBLIN_SOLDIER
}

func MakeGoblinSoldier() *GoblinSoldier {
	enemy := MakeEnemy(
		map[stats.StatType]float32{
			stats.CRIT:    0.05,
			stats.DMG:     5.00,
			stats.DEFENSE: 5.00,
			stats.EVASION: 0.05,
			stats.HP:      40.00,
			stats.LUCK:    0.05,
		},
	)

	return &GoblinSoldier{*enemy}
}
