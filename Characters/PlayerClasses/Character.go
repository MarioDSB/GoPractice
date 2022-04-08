package PlayerClasses

import "GoPractice/Stats"

type CharacterInterface interface {
	GetClassType() ClassType
	GetStats() *Stats.Stats
}

type Character struct {
	CharacterInterface
	stats Stats.Stats
}

func (character *Character) GetStats() *Stats.Stats {
	return &character.stats
}

func MakeCharacter(attributes map[Stats.StatType]float32) *Character {
	stats := Stats.MakeStats(attributes)
	return &Character{
		stats: *stats,
	}
}
