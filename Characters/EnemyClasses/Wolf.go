package EnemyClasses

import (
	stats "GoPractice/Stats"
	"fmt"
)

type Wolf struct {
	attributes map[stats.StatType]float32
	EnemyInterface
}

func (wolf *Wolf) GetEnemyName() string {
	return "Wolf"
}

func (wolf *Wolf) GetAttributes() map[stats.StatType]float32 {
	return wolf.attributes
}

func (wolf *Wolf) GetAttribute(statType stats.StatType) float32 {
	return wolf.attributes[statType]
}

func (wolf *Wolf) GetClassType() EnemyType {
	return WOLF
}

func (wolf *Wolf) Injured(dmg float32) {
	wolf.attributes[stats.HP] = wolf.attributes[stats.HP] - dmg
}

func (wolf *Wolf) GetEnemyStatus() string {
	return fmt.Sprintf(" (%.2f HP)", wolf.attributes[stats.HP])
}

func MakeWolf() *Wolf {
	attributes := map[stats.StatType]float32{
		stats.CRIT:    0.10,
		stats.DMG:     4.00,
		stats.DEFENSE: 0.00,
		stats.EVASION: 0.10,
		stats.HP:      20.00,
		stats.LUCK:    0.05,
	}

	return &Wolf{attributes: attributes}
}
