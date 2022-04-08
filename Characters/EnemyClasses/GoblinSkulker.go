package EnemyClasses

import (
	stats "GoPractice/Stats"
	"fmt"
)

type GoblinSkulker struct {
	attributes map[stats.StatType]float32
	EnemyInterface
}

func (goblin *GoblinSkulker) GetEnemyName() string {
	return "Goblin Skulker"
}

func (goblin *GoblinSkulker) GetAttributes() map[stats.StatType]float32 {
	return goblin.attributes
}

func (goblin *GoblinSkulker) GetAttribute(statType stats.StatType) float32 {
	return goblin.attributes[statType]
}

func (goblin *GoblinSkulker) GetClassType() EnemyType {
	return GOBLIN_SKULKER
}

func (goblin *GoblinSkulker) Injured(dmg float32) {
	goblin.attributes[stats.HP] = goblin.attributes[stats.HP] - dmg
}

func (goblin *GoblinSkulker) GetEnemyStatus() string {
	return fmt.Sprintf(" (%.2f HP)", goblin.attributes[stats.HP])
}

func MakeGoblinSkulker() *GoblinSkulker {
	attributes := map[stats.StatType]float32{
		stats.CRIT:    0.20,
		stats.DMG:     4.00,
		stats.DEFENSE: 0.00,
		stats.EVASION: 0.10,
		stats.HP:      30.00,
		stats.LUCK:    0.05,
	}
	return &GoblinSkulker{attributes: attributes}
}
