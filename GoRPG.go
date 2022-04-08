package main

import (
	"GoPractice/Characters"
	"GoPractice/Game"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Core struct {
	player *Characters.Player
	fights int
}

func (core *Core) showMainInterface() {
	purchase := false
	option := 0
	var err error
	for {
		fmt.Printf("\n\nPlayer's name:    %s\n", core.player.GetName())
		fmt.Printf("Fights won:    %d\n", core.fights)
		fmt.Println("\n1) Play an encounter")
		fmt.Println("2) Open the store")
		fmt.Println("3) Show stats")
		fmt.Println("4) End the game")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Enter a number from 1 to 4, corresponding to the action you want to take.")
			continue
		}

		switch option {

		case 1:
			encounter := Game.Encounter{}
			result, hp := encounter.RandomEncounter(core.player)
			if !result {
				if hp > 0 {
					// Ran away
					fmt.Println("You ran, and so, you won't be able to purchase anything.")
					break
				} else {
					// Died
					fmt.Println("You died...")
					return
				}
			}

			core.fights++
			purchase = false
			break
		case 2:
			store := Game.Store{}
			purchase = store.ShowStoreInterface(core.player, purchase)
			break
		case 3:
			core.player.PrintStats()
			break
		case 4:
			if core.LeaveGame() {
				return
			}
			break
		default:
			fmt.Println("Enter a number from 1 to 4, corresponding to the action you want to take.")
		}
	}
}

func (core *Core) LeaveGame() bool {
	fmt.Println("\nAre you sure you want to leave the game?")
	fmt.Println("(Your results aren't saved)")
	fmt.Println("1) Yes              2) No")

	var option int
	var err error
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Enter 1, if you want to leave the game, 2 if you wish to continue.")
			continue
		}

		switch option {
		case 1:
			fmt.Println("\nHope you come back!")
			return true
		case 2:
			return false
		default:
			fmt.Println("Enter 1, if you want to leave the game, 2 if you wish to continue.")
			break
		}
	}
}

func main() {
	initialSettings := Game.InitialSettings{}
	player := initialSettings.StartUp()
	core := Core{
		player: player,
		fights: 0,
	}

	core.showMainInterface()
}
