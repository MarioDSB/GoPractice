package Characters

import (
	"GoPractice/Characters/PlayerClasses"
	"GoPractice/Equipments"
	"GoPractice/Stats"
	"fmt"
)

type Player struct {
	name  string
	class PlayerClasses.CharacterInterface
	Stats *Stats.Stats
}

func (player *Player) GetName() string {
	return player.name
}

func (player *Player) ImproveDmg() {
	newDmg := player.Stats.GetAttribute(Stats.DMG) + 2
	player.Stats.SetAttribute(Stats.DMG, newDmg)
}

func (player *Player) ImproveDefense() {
	newDefense := player.Stats.GetAttribute(Stats.DEFENSE) + 2
	player.Stats.SetAttribute(Stats.DEFENSE, newDefense)
}

func (player *Player) Injured(dmg float32) {
	newHp := player.Stats.GetAttribute(Stats.HP) - dmg
	player.Stats.SetAttribute(Stats.HP, newHp)
}

func (player *Player) AddEquipment(object interface{}) {
	equipment := object.(Equipments.EquipmentInterface)
	for key, value := range player.Stats.GetAttributes() {
		equipmentStatValue := equipment.GetStats().GetAttribute(key)
		player.Stats.GetAttributes()[key] = value + equipmentStatValue
	}
}

func (player *Player) CheckDead() bool {
	return player.Stats.GetAttribute(Stats.HP) <= 0
}

func (player *Player) PrintStats() {
	for key, value := range player.Stats.GetAttributes() {
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
		class: character,
		Stats: character.GetStats(),
		name:  name,
	}
}
