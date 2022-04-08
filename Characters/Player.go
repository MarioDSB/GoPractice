package Characters

import (
	"GoPractice/Characters/PlayerClasses"
	"GoPractice/Equipments"
	"GoPractice/Stats"
	"fmt"
)

type Player struct {
	baseAttributes map[Stats.StatType]float32
	name           string
	class          PlayerClasses.CharacterInterface
}

func (player *Player) GetName() string {
	return player.name
}

func (player *Player) GetAttributes() map[Stats.StatType]float32 {
	return player.class.GetAttributes()
}

func (player *Player) GetAttribute(statType Stats.StatType) float32 {
	return player.GetAttributes()[statType]
}

func (player *Player) SetAttribute(statType Stats.StatType, value float32) {
	player.class.SetAttribute(statType, value)
}

func (player *Player) ImproveDmg() {
	currentDmg := player.GetAttribute(Stats.DMG)
	player.class.SetAttribute(Stats.DMG, currentDmg+2)
}

func (player *Player) ImproveDefense() {
	currentDefense := player.GetAttribute(Stats.DEFENSE)
	player.SetAttribute(Stats.DEFENSE, currentDefense+2)
}

func (player *Player) Injured(dmg float32) {
	currentHp := player.GetAttribute(Stats.HP)
	player.SetAttribute(Stats.HP, currentHp-dmg)
}

func (player *Player) AddEquipment(object interface{}) {
	equipment := object.(Equipments.EquipmentInterface)
	for key, value := range player.baseAttributes {
		equipmentStatValue := equipment.GetAttribute(key)
		player.baseAttributes[key] = value + equipmentStatValue
	}
}

func (player *Player) CheckDead() bool {
	return player.GetAttribute(Stats.HP) <= 0
}

func (player *Player) PrintStats() {
	for key, value := range player.GetAttributes() {
		switch key {
		case Stats.HP:
			fmt.Printf("HP: %.2f\n", value)
			break
		case Stats.DMG:
			fmt.Printf("Damage: %.2f\n", value)
			break
		case Stats.CRIT:
			fmt.Printf("Critical Chance: %.2f\n", value)
			break
		case Stats.DEFENSE:
			fmt.Printf("Defense: %.2f\n", value)
			break
		case Stats.EVASION:
			fmt.Printf("Evasion: %.2f\n", value)
			break
		case Stats.LUCK:
			fmt.Printf("Luck: %.2f\n", value)
			break
		default:
			break
		}
	}
}

func MakePlayer(character PlayerClasses.CharacterInterface, name string) *Player {
	return &Player{
		class:          character,
		baseAttributes: character.GetAttributes(),
		name:           name,
	}
}
