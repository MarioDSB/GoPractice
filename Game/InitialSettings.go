package Game

import (
	"GoPractice/Characters"
	"GoPractice/Characters/PlayerClasses"
	"GoPractice/Equipments/Armors"
	"GoPractice/Equipments/Weapons"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type InitialSettings struct {
	player Characters.Player
}

func (settings *InitialSettings) StartUp() *Characters.Player {
	fmt.Println("Hello traveller!")
	fmt.Println("What is your name?")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	fmt.Printf("\nWelcome, %s!\n", name)
	fmt.Println("\nIn this game, your goal is to survive.")
	fmt.Println("Enemies will appear in your way and try to kill you.")
	fmt.Println("Good luck!")

	settings.setupPlayer(name)
	settings.setupEquipment()

	return &settings.player
}

func (settings *InitialSettings) initPlayer(option int, name string) {
	switch option {
	case 1:
		settings.player = *Characters.MakePlayer(PlayerClasses.MakeAssassin(), name)
	case 2:
		settings.player = *Characters.MakePlayer(PlayerClasses.MakeFootman(), name)
	case 3:
		settings.player = *Characters.MakePlayer(PlayerClasses.MakeKnight(), name)
	case 4:
		settings.player = *Characters.MakePlayer(PlayerClasses.MakeRanger(), name)
	}
}

func (settings *InitialSettings) setupPlayer(name string) {
	for {
		fmt.Println("\nChoose your class:")
		fmt.Println("    1) Assassin (Critical attack and evasion-based class)")
		fmt.Println("    2) Footman  (Damage and evasion-based class)")
		fmt.Println("    3) Knight   (Damage and defense-based class)")
		fmt.Println("    4) Ranger   (Critical attack and damage-based class, with some evasion)")

		var option int
		var err error

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Enter a number from 1 to 4, corresponding to the class you want to play.")
			continue
		}

		if option >= 1 && option <= 4 {
			settings.initPlayer(option, name)
			return
		} else {
			fmt.Println("Enter a number from 1 to 4, corresponding to the class you want to play.")
		}
	}
}

func (settings *InitialSettings) initWeapon(option int) {
	switch option {
	case 1:
		settings.player.AddEquipment(Weapons.MakeBowAndArrow())
	case 2:
		settings.player.AddEquipment(Weapons.MakeDoubleKnives())
	case 3:
		settings.player.AddEquipment(Weapons.MakeSpear())
	case 4:
		settings.player.AddEquipment(Weapons.MakeSwordAndShield())
	case 5:
		settings.player.AddEquipment(Weapons.MakeTwoHandedSword())
	}
}

func (settings *InitialSettings) initArmor(option int) {
	switch option {
	case 1:
		settings.player.AddEquipment(Armors.MakeNoArmor())
	case 2:
		settings.player.AddEquipment(Armors.MakeLightArmor())
	case 3:
		settings.player.AddEquipment(Armors.MakeMediumArmor())
	case 4:
		settings.player.AddEquipment(Armors.MakeHeavyArmor())
	}
}

func (settings *InitialSettings) getWeapon() {
	for {
		fmt.Println("\nNow choose your weapon:")
		fmt.Println("    1) Bow and Arrow        (Critical attack and damage-based weapon)")
		fmt.Println("    2) Double Knives        (Critical attack and evasion-based weapon)")
		fmt.Println("    3) Spear                (Defense and evasion-based weapon)")
		fmt.Println("    4) Sword and Shield     (Damage and defense-based option)")
		fmt.Println("    5) Two Handed Sword     (Damage-based weapon)")

		var option int
		var err error
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Enter a number from 1 to 5, corresponding to the class you want to play.")
			continue
		}

		if option >= 1 && option <= 5 {
			settings.initWeapon(option)
			return
		} else {
			fmt.Println("Enter a number from 1 to 5, corresponding to the class you want to play.")
		}
	}
}

func (settings *InitialSettings) getArmor() {
	for {
		fmt.Println("\nAnd finally, your armor:")
		fmt.Println("    1) No Armor         (Increased evasion)")
		fmt.Println("    2) Light Armor      (Slightly increased defense and evasion)")
		fmt.Println("    3) Medium Armor     (Increased defense)")
		fmt.Println("    4) Heavy Armor      (Increased defense and HP, decreased evasion)")

		var option int
		var err error
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Enter a number from 1 to 4, corresponding to the class you want to play.")
			continue
		}

		if option >= 1 && option <= 4 {
			settings.initArmor(option)
			return
		} else {
			fmt.Println("Enter a number from 1 to 4, corresponding to the class you want to play.")
		}
	}
}

func (settings *InitialSettings) setupEquipment() {
	settings.getWeapon()
	settings.getArmor()
}
