package Game

import (
	characters "GoPractice/Characters"
	"GoPractice/Equipments/Other"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Store struct {
}

func (store *Store) ShowStoreInterface(player *characters.Player, purchase bool) bool {
	fmt.Println("\nWelcome to the store traveller!")
	if purchase {
		fmt.Println("\nI'm sorry, but I'm out of stock for today...")
		fmt.Println("Maybe you can come back tomorrow...?")
		return false
	}

	option := 0
	var err error
	for option == 0 {
		fmt.Println("\nWhat have you come here for? (You can only buy one thing per turn)")
		fmt.Println("    1) Buy and use an health potion")
		fmt.Println("    2) Improve weapon damage")
		fmt.Println("    3) Improve armor defense")
		fmt.Println("    4) Leave the store")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Enter a number from 1 to 4, corresponding to the action you want to take.")
			continue
		}

		switch option {
		case 1:
			option = store.potionDescription()
			break
		case 2:
			option = store.weaponDescription()
			break
		case 3:
			option = store.armorDescription()
			break
		case 4:
			store.leave()
			return false
		default:
			fmt.Println("Enter a number from 1 to 4, corresponding to the action you want to take.")
			option = 0
		}
	}

	store.makePurchase(player, option)
	return true
}

func (store *Store) makePurchase(player *characters.Player, option int) {
	switch option {
	case 1:
		healthPotion := Other.MakeHealthPotion()
		player.AddEquipment(healthPotion)
		break
	case 2:
		player.ImproveDmg()
		break
	case 3:
		player.ImproveDefense()
		break
	default:
		break
	}
}

func (store *Store) potionDescription() int {
	var option int
	var err error
	for {
		fmt.Println("\nBuying a health potion increases your health.")
		fmt.Println("Do you want to buy a health potion?")
		fmt.Println("1) Yes              2) No")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Enter 1, if you want to buy the item, 2 if you don't.")
			continue
		}

		switch option {
		case 1:
			return 1
		case 2:
			return 0
		default:
			fmt.Println("Enter 1, if you want to buy the item, 2 if you don't.")
			break
		}
	}
}

func (store *Store) weaponDescription() int {
	var option int
	var err error
	for {
		fmt.Println("\nAre you sure you want to upgrade your weapon?")
		fmt.Println("1) Yes              2) No")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Enter 1, if you want to buy the item, 2 if you don't.")
			continue
		}

		switch option {
		case 1:
			return 2
		case 2:
			return 0
		default:
			fmt.Println("Enter 1, if you want to buy the item, 2 if you don't.")
			break
		}
	}
}

func (store *Store) armorDescription() int {
	var option int
	var err error
	for {
		fmt.Println("\nAre you sure you want to upgrade your armor?")
		fmt.Println("1) Yes              2) No")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Enter 1, if you want to buy the item, 2 if you don't.")
			continue
		}

		switch option {
		case 1:
			return 3
		case 2:
			return 0
		default:
			fmt.Println("Enter 1, if you want to buy the item, 2 if you don't.")
			break
		}
	}
}

func (store *Store) leave() {
	fmt.Println("\nThank you! And come back anytime!")
}
