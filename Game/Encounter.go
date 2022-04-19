package Game

import (
	"GoPractice/Characters"
	"GoPractice/Characters/EnemyClasses"
	"GoPractice/Stats"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Encounter struct {
	enemies []EnemyClasses.EnemyInterface
}

func checkDmgBelowZero(dmg float32) float32 {
	if dmg < 0 {
		return 0
	}
	return dmg
}

func (encounter *Encounter) getNumberOfEnemies() int {
	rand.Seed(time.Now().UnixNano())
	seed := rand.Intn(100)

	/*
	 * 1 enemy chance: 40%
	 * 2 enemies chance: 35%
	 * 3 enemies chance: 25%
	 */
	if seed < 40 {
		return 1
	} else if seed < 75 {
		return 2
	} else {
		return 3
	}
}

func (encounter *Encounter) generateEnemy() interface{} {
	rand.Seed(time.Now().UnixNano())
	seed := rand.Intn(100)

	/*
	 * Wolf chance p/enemy: 20%
	 * Common Goblin chance p/enemy: 20%
	 * Goblin Archer chance p/enemy: 15%
	 * Goblin Skulker chance p/enemy: 15%
	 * Goblin Soldier chance p/enemy: 15%
	 * Goblin Warlord chance p/enemy: 10%
	 * Goblin King chance p/enemy: 5%
	 */
	if seed < 20 {
		return EnemyClasses.MakeWolf()
	} else if seed <= 40 {
		return EnemyClasses.MakeCommonGoblin()
	} else if seed <= 55 {
		return EnemyClasses.MakeGoblinArcher()
	} else if seed <= 70 {
		return EnemyClasses.MakeGoblinSkulker()
	} else if seed <= 85 {
		return EnemyClasses.MakeGoblinSoldier()
	} else if seed <= 95 {
		return EnemyClasses.MakeGoblinWarlord()
	} else {
		return EnemyClasses.MakeGoblinKing()
	}
}

func (encounter *Encounter) generateEnemyList() {
	for enemies := 0; enemies < encounter.getNumberOfEnemies(); enemies++ {
		encounter.enemies = append(encounter.enemies, encounter.generateEnemy().(EnemyClasses.EnemyInterface))
	}
}

func (encounter *Encounter) battle(player *Characters.Player) (bool, float32) {
	battleEnd := false
	checkDead := false

loop:
	for !battleEnd {
		if checkDead {
			return false, 0
		}

		var option int
		var err error
		for {
			fmt.Printf("\nYou have %.2f HP.\n", player.Stats.GetAttribute(Stats.HP))
			fmt.Println("Enemies:")
			for index, enemy := range encounter.enemies {
				fmt.Printf("    %d) %s%s", index+1, enemy.GetEnemyName(), enemy.GetEnemyStatus())
			}

			fmt.Println("\nWhat do you want to do?")
			fmt.Println("    1) Attack")
			fmt.Println("    2) Defend")
			fmt.Println("    3) Run away")

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			option, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Enter a number from 1 to 3, corresponding to the action you want to take.")
				continue
			}

			switch option {
			case 1:
				if len(encounter.enemies) == 1 {
					encounter.attack(player, 1)
					checkDead = player.CheckDead()
					battleEnd = encounter.checkDeadEnemies()
					if battleEnd {
						break loop
					}

				} else {
					attackOption := encounter.attackOption()
					if attackOption != 0 {
						encounter.attack(player, attackOption)
						checkDead = player.CheckDead()
						battleEnd = encounter.checkDeadEnemies()
						if battleEnd {
							break loop
						}
					}
				}
				break
			case 2:
				encounter.defend(player)
				checkDead = player.CheckDead()
				battleEnd = encounter.checkDeadEnemies()
				if battleEnd {
					break loop
				}
				break
			case 3:
				if encounter.runAway() {
					return false, player.Stats.GetAttribute(Stats.HP)
				}
				break
			default:
				fmt.Println("Enter a number from 1 to 3, corresponding to the action you want to take.")
				break
			}

			if checkDead {
				return false, player.Stats.GetAttribute(Stats.HP)
			}
		}
	}

	fmt.Println("You win!")
	return true, player.Stats.GetAttribute(Stats.HP)
}

func (encounter *Encounter) checkDeadEnemies() bool {
	// Array of indices of objects to remove from the encounter.enemies slice
	toRemove := make([]int, 0)

	for index, enemy := range encounter.enemies {
		if enemy.GetStats().GetAttribute(Stats.HP) <= 0 {
			fmt.Printf("\n%s has died.\n", enemy.GetEnemyName())
			toRemove = append(toRemove, index)
		}
	}

	// Remove objects with re-slicing (order matters)
	for _, index := range toRemove {
		encounter.enemies = append(encounter.enemies[:index], encounter.enemies[index+1:]...)
	}

	return len(encounter.enemies) == 0
}

func (encounter *Encounter) showDmg(dmg float32) string {
	return fmt.Sprintf(" (%.2f dmg)", dmg)
}

func (encounter *Encounter) statCheck(stat float32) bool {
	rand.Seed(time.Now().UnixNano())
	return stat*100 > float32(rand.Intn(100))
}

func (encounter *Encounter) attack(player *Characters.Player, attackOption int) {
	var dmg float32

	if encounter.statCheck(player.Stats.GetAttribute(Stats.CRIT)) {
		// Critical attack
		dmg = player.Stats.GetAttribute(Stats.DMG)*2 - encounter.enemies[attackOption-1].GetStats().GetAttribute(Stats.DEFENSE)
		dmg = checkDmgBelowZero(dmg)
		fmt.Printf("\nYou did critical damage!%s\n", encounter.showDmg(dmg))
	} else {
		// Normal attack
		dmg = player.Stats.GetAttribute(Stats.DMG) - encounter.enemies[attackOption-1].GetStats().GetAttribute(Stats.DEFENSE)
		dmg = checkDmgBelowZero(dmg)

		fmt.Printf("\nYou attacked.%s\n", encounter.showDmg(dmg))
	}

	encounter.enemies[attackOption-1].Injured(dmg)

	if encounter.checkDeadEnemies() {
		return
	}

	for _, enemy := range encounter.enemies {
		if encounter.statCheck(enemy.GetStats().GetAttribute(Stats.CRIT)) {
			// Critical dmg received
			if encounter.statCheck(player.Stats.GetAttribute(Stats.LUCK)) {
				// Lucked out
				dmg = enemy.GetStats().GetAttribute(Stats.DMG) - player.Stats.GetAttribute(Stats.DEFENSE)
				dmg = checkDmgBelowZero(dmg)

				fmt.Printf("%s critically attacks you, but luckily the dmg missed any vital spots.%s\n", enemy.GetEnemyName(), encounter.showDmg(dmg))
			} else {
				// Regular critical dmg received
				dmg = enemy.GetStats().GetAttribute(Stats.DMG)*2 - player.Stats.GetAttribute(Stats.DEFENSE)
				dmg = checkDmgBelowZero(dmg)

				fmt.Printf("%s critically attacks you.%s\n", enemy.GetEnemyName(), encounter.showDmg(dmg))
			}
		} else {
			// Normal dmg received
			if encounter.statCheck(player.Stats.GetAttribute(Stats.EVASION)) {
				// Evaded
				dmg = 0
				fmt.Printf("You are attacked by a %s, but you evade the dmg.%s\n", enemy.GetEnemyName(), encounter.showDmg(0))
			} else {
				// Regular normal dmg received
				dmg = enemy.GetStats().GetAttribute(Stats.DMG) - player.Stats.GetAttribute(Stats.DEFENSE)
				dmg = checkDmgBelowZero(dmg)

				fmt.Printf("You are attacked by a %s.%s\n", enemy.GetEnemyName(), encounter.showDmg(dmg))
			}
		}

		player.Injured(dmg)
	}
}

func (encounter *Encounter) defend(player *Characters.Player) {
	var dmg float32

	fmt.Println()
	for _, enemy := range encounter.enemies {
		if encounter.statCheck(player.Stats.GetAttribute(Stats.EVASION)) {
			// Attack received evaded
			if encounter.statCheck(player.Stats.GetAttribute(Stats.LUCK) + 0.1) {
				// Lucked out
				if encounter.statCheck(player.Stats.GetAttribute(Stats.CRIT)) {
					// Critical counterattack
					dmg = player.Stats.GetAttribute(Stats.DMG)*2 - enemy.GetStats().GetAttribute(Stats.DEFENSE)
					fmt.Printf("You managed to evade an attack from %s, and countered, dealing critical damage!!%s\n", enemy.GetEnemyName(), encounter.showDmg(dmg))
				} else {
					// Normal counterattack
					dmg = player.Stats.GetAttribute(Stats.DMG) - enemy.GetStats().GetAttribute(Stats.DEFENSE)
					fmt.Printf("You managed to evade an attack from %s, and countered!%s\n", enemy.GetEnemyName(), encounter.showDmg(dmg))
				}

				enemy.Injured(dmg)
			} else {
				// Regular evade
				fmt.Printf("You managed to evade an attack from %s.%s\n", enemy.GetEnemyName(), encounter.showDmg(0))
			}
		} else {
			// Attack received not evaded
			if encounter.statCheck(enemy.GetStats().GetAttribute(Stats.CRIT)) {
				// Critical attack received
				if encounter.statCheck(player.Stats.GetAttribute(Stats.LUCK)) {
					// Lucked out
					dmg = enemy.GetStats().GetAttribute(Stats.DMG) - player.Stats.GetAttribute(Stats.DEFENSE)
					dmg = checkDmgBelowZero(dmg)

					fmt.Printf("%s managed to breach you defenses, but luckily the attack missed any vitar spots.%s\n", enemy.GetEnemyName(), encounter.showDmg(dmg))
				} else {
					// Regular critical attack received
					dmg = enemy.GetStats().GetAttribute(Stats.DMG)*2 - player.Stats.GetAttribute(Stats.DEFENSE)
					dmg = checkDmgBelowZero(dmg)

					fmt.Printf("%s manages to breach your defenses and deals critical damage.%s\n", enemy.GetEnemyName(), encounter.showDmg(dmg))
				}

				player.Injured(dmg)
			} else {
				// Regular attack received
				if encounter.statCheck(player.Stats.GetAttribute(Stats.LUCK)) {
					// Lucked out
					fmt.Printf("%s attacked you, but luckily you dodged it.%s\n", enemy.GetEnemyName(), encounter.showDmg(0))
				} else {
					// Regular attack received not lucked out
					dmg = enemy.GetStats().GetAttribute(Stats.DMG) - player.Stats.GetAttribute(Stats.DEFENSE)
					dmg = checkDmgBelowZero(dmg)

					fmt.Printf("%s attacks you.%s\n", enemy.GetEnemyName(), encounter.showDmg(dmg))
					player.Injured(dmg)
				}
			}
		}
	}
}

func (encounter *Encounter) attackOption() int {
	var option int
	var err error
	for {
		fmt.Println("\nWho do you want to attack")

		for index, enemy := range encounter.enemies {
			fmt.Printf("    %d) %s\n", index+1, enemy.GetEnemyName())
		}
		fmt.Printf("    %d) Go back\n", len(encounter.enemies)+1)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Enter a number from 1 to %d, corresponding to the action you want to take.", len(encounter.enemies)+1)
			continue
		}

		if option == len(encounter.enemies)+1 {
			return 0
		} else if option >= 1 && option <= len(encounter.enemies) {
			return option
		} else {
			fmt.Printf("Enter a number from 1 to %d, corresponding to the action you want to take.", len(encounter.enemies)+1)
		}
	}
}

func (encounter *Encounter) runAway() bool {
	var option int
	var err error
	for {
		fmt.Println("\nAre you sure you want to run away?")
		fmt.Println("(You won't recover your health)")
		fmt.Println("1) Yes              2) No")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Enter 1, if you want to run away, 2 if you don't.")
			continue
		}

		switch option {
		case 1:
			return true
		case 2:
			return false
		default:
			fmt.Println("Enter 1, if you want to run away, 2 if you don't.")
			break
		}
	}
}

func (encounter *Encounter) RandomEncounter(player *Characters.Player) (bool, float32) {
	encounter.generateEnemyList()
	return encounter.battle(player)
}
