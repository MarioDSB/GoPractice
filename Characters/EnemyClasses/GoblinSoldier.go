package EnemyClasses

import (
	stats "GoPractice/Stats"
	"fmt"
)

type GoblinSoldier struct {
	attributes map[stats.StatType]float32
	EnemyInterface
}

func (goblin *GoblinSoldier) GetEnemyName() string {
	return "Goblin Soldier"
}

func (goblin *GoblinSoldier) GetAttributes() map[stats.StatType]float32 {
	return goblin.attributes
}

func (goblin *GoblinSoldier) GetAttribute(statType stats.StatType) float32 {
	return goblin.attributes[statType]
}

func (goblin *GoblinSoldier) GetClassType() EnemyType {
	return GOBLIN_SOLDIER
}

func (goblin *GoblinSoldier) Injured(dmg float32) {
	goblin.attributes[stats.HP] = goblin.attributes[stats.HP] - dmg
}

func (goblin *GoblinSoldier) GetEnemyStatus() string {
	return fmt.Sprintf(" (%.2f HP)", goblin.attributes[stats.HP])
}

func MakeGoblinSoldier() *GoblinSoldier {
	attributes := map[stats.StatType]float32{
		stats.CRIT:    0.05,
		stats.DMG:     5.00,
		stats.DEFENSE: 5.00,
		stats.EVASION: 0.05,
		stats.HP:      40.00,
		stats.LUCK:    0.05,
	}

	return &GoblinSoldier{attributes: attributes}
}
