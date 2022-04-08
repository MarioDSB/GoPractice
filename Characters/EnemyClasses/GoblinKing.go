package EnemyClasses

import (
	stats "GoPractice/Stats"
)

type GoblinKing struct {
	Enemy
}

func (goblin *GoblinKing) GetEnemyName() string {
	return "Goblin King"
}

func (goblin *GoblinKing) GetClassType() EnemyType {
	return GOBLIN_KING
}

func MakeGoblinKing() *GoblinKing {
	enemy := MakeEnemy(
		map[stats.StatType]float32{
			stats.CRIT:    0.10,
			stats.DMG:     10.00,
			stats.DEFENSE: 10.00,
			stats.EVASION: 0.10,
			stats.HP:      120.00,
			stats.LUCK:    0.05,
		},
	)

	return &GoblinKing{*enemy}
}
