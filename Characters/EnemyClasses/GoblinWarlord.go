package EnemyClasses

import (
	stats "GoPractice/Stats"
	"fmt"
)

type GoblinWarlord struct {
	attributes map[stats.StatType]float32
	EnemyInterface
}

func (goblin *GoblinWarlord) GetEnemyName() string {
	return "Goblin Warlord"
}

func (goblin *GoblinWarlord) GetAttributes() map[stats.StatType]float32 {
	return goblin.attributes
}

func (goblin *GoblinWarlord) GetAttribute(statType stats.StatType) float32 {
	return goblin.attributes[statType]
}

func (goblin *GoblinWarlord) GetClassType() EnemyType {
	return GOBLIN_WARLORD
}

func (goblin *GoblinWarlord) Injured(dmg float32) {
	goblin.attributes[stats.HP] = goblin.attributes[stats.HP] - dmg
}

func (goblin *GoblinWarlord) GetEnemyStatus() string {
	return fmt.Sprintf(" (%.2f HP)", goblin.attributes[stats.HP])
}

func MakeGoblinWarlord() *GoblinWarlord {
	attributes := map[stats.StatType]float32{
		stats.CRIT:    0.05,
		stats.DMG:     10.00,
		stats.DEFENSE: 7.00,
		stats.EVASION: 0.00,
		stats.HP:      72.00,
		stats.LUCK:    0.05,
	}
	return &GoblinWarlord{attributes: attributes}
}
