package Stats

type StatInterface interface {
	GetAttributes() map[StatType]float32
	GetAttribute(statType StatType) float32
	SetAttribute(statType StatType, value float32)
}

type Stats struct {
	attributes map[StatType]float32
	StatInterface
}

func (stats *Stats) GetAttributes() map[StatType]float32 {
	return stats.attributes
}

func (stats *Stats) GetAttribute(statType StatType) float32 {
	return stats.attributes[statType]
}

func (stats *Stats) SetAttribute(statType StatType, value float32) {
	stats.attributes[statType] = value
}

func MakeStats(attributes map[StatType]float32) *Stats {
	return &Stats{
		attributes: attributes,
	}
}
