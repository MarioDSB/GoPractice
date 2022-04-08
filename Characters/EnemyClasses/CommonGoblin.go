package EnemyClasses

import (
	stats "GoPractice/Stats"
	"fmt"
)

type CommonGoblin struct {
	attributes map[stats.StatType]float32
	EnemyInterface
}

func (goblin *CommonGoblin) GetEnemyName() string {
	return "Common Goblin"
}

func (goblin *CommonGoblin) GetAttributes() map[stats.StatType]float32 {
	return goblin.attributes
}

func (goblin *CommonGoblin) GetAttribute(statType stats.StatType) float32 {
	return goblin.attributes[statType]
}

func (goblin *CommonGoblin) GetClassType() EnemyType {
	return COMMON_GOBLIN
}

func (goblin *CommonGoblin) Injured(dmg float32) {
	goblin.attributes[stats.HP] = goblin.attributes[stats.HP] - dmg
}

func (goblin *CommonGoblin) GetEnemyStatus() string {
	return fmt.Sprintf(" (%.2f HP)", goblin.attributes[stats.HP])
}

func MakeCommonGoblin() *CommonGoblin {
	attributes := map[stats.StatType]float32{
		stats.CRIT:    0.10,
		stats.DMG:     4.00,
		stats.DEFENSE: 2.00,
		stats.EVASION: 0.10,
		stats.HP:      30.00,
		stats.LUCK:    0.05,
	}

	return &CommonGoblin{attributes: attributes}
}
