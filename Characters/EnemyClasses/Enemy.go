package EnemyClasses

import (
	"GoPractice/Stats"
	"fmt"
)

type EnemyInterface interface {
	GetClassType() EnemyType
	GetEnemyName() string
	Injured(dmg float32)
	GetEnemyStatus() string
	GetStats() *Stats.Stats
}

type Enemy struct {
	EnemyInterface
	stats Stats.Stats
}

func (enemy *Enemy) Injured(dmg float32) {
	newHp := enemy.stats.GetAttribute(Stats.HP) - dmg
	enemy.stats.SetAttribute(Stats.HP, newHp)
}

func (enemy *Enemy) GetEnemyStatus() string {
	return fmt.Sprintf(" (%.2f HP)", enemy.stats.GetAttribute(Stats.HP))
}

func (enemy *Enemy) GetStats() *Stats.Stats {
	return &enemy.stats
}

func MakeEnemy(attributes map[Stats.StatType]float32) *Enemy {
	stats := Stats.MakeStats(attributes)
	return &Enemy{
		stats: *stats,
	}
}
