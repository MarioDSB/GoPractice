package EnemyClasses

import (
	stats "GoPractice/Stats"
	"fmt"
)

type GoblinKing struct {
	attributes map[stats.StatType]float32
	EnemyInterface
}

func (goblin *GoblinKing) GetEnemyName() string {
	return "Goblin King"
}

func (goblin *GoblinKing) GetAttributes() map[stats.StatType]float32 {
	return goblin.attributes
}

func (goblin *GoblinKing) GetAttribute(statType stats.StatType) float32 {
	return goblin.attributes[statType]
}

func (goblin *GoblinKing) GetClassType() EnemyType {
	return GOBLIN_KING
}

func (goblin *GoblinKing) Injured(dmg float32) {
	goblin.attributes[stats.HP] = goblin.attributes[stats.HP] - dmg
}

func (goblin *GoblinKing) GetEnemyStatus() string {
	return fmt.Sprintf(" (%.2f HP)", goblin.attributes[stats.HP])
}

func MakeGoblinKing() *GoblinKing {
	attributes := map[stats.StatType]float32{
		stats.CRIT:    0.10,
		stats.DMG:     10.00,
		stats.DEFENSE: 10.00,
		stats.EVASION: 0.10,
		stats.HP:      120.00,
		stats.LUCK:    0.05,
	}

	return &GoblinKing{attributes: attributes}
}
