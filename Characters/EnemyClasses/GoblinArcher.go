package EnemyClasses

import (
	stats "GoPractice/Stats"
)

type GoblinArcher struct {
	Enemy
}

func (goblin *GoblinArcher) GetEnemyName() string {
	return "Goblin Archer"
}

func (goblin *GoblinArcher) GetClassType() EnemyType {
	return GOBLIN_ARCHER
}

func MakeGoblinArcher() *GoblinArcher {
	enemy := MakeEnemy(
		map[stats.StatType]float32{
			stats.CRIT:    0.15,
			stats.DMG:     4.00,
			stats.DEFENSE: 0.00,
			stats.EVASION: 0.15,
			stats.HP:      30.00,
			stats.LUCK:    0.05,
		},
	)

	return &GoblinArcher{*enemy}
}
