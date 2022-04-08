package EnemyClasses

import (
	stats "GoPractice/Stats"
	"fmt"
)

type GoblinArcher struct {
	attributes map[stats.StatType]float32
	EnemyInterface
}

func (goblin *GoblinArcher) GetEnemyName() string {
	return "Goblin Archer"
}

func (goblin *GoblinArcher) GetAttributes() map[stats.StatType]float32 {
	return goblin.attributes
}

func (goblin *GoblinArcher) GetAttribute(statType stats.StatType) float32 {
	return goblin.attributes[statType]
}

func (goblin *GoblinArcher) GetClassType() EnemyType {
	return GOBLIN_ARCHER
}

func (goblin *GoblinArcher) Injured(dmg float32) {
	goblin.attributes[stats.HP] = goblin.attributes[stats.HP] - dmg
}

func (goblin *GoblinArcher) GetEnemyStatus() string {
	return fmt.Sprintf(" (%.2f HP)", goblin.attributes[stats.HP])
}

func MakeGoblinArcher() *GoblinArcher {
	attributes := map[stats.StatType]float32{
		stats.CRIT:    0.15,
		stats.DMG:     4.00,
		stats.DEFENSE: 0.00,
		stats.EVASION: 0.15,
		stats.HP:      30.00,
		stats.LUCK:    0.05,
	}

	return &GoblinArcher{attributes: attributes}
}
