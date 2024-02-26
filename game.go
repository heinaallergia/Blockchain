package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome to the Forest Adventure!")

	// Player stats
	var health int = 10
	var attack int = 5

	// Game loop
	for {
		// Display current status
		fmt.Println("\nYour health:", health)

		// Present options
		fmt.Println("1. Explore deeper")
		fmt.Println("2. Check stats")
		fmt.Println("3. Exit")

		// Get user input
		var choice int
		fmt.Scanf("%d", &choice)

		// Handle choice
		switch choice {
		case 1:
			encounterMonster := rand.Intn(2) // 50% chance of encountering a monster
			if encounterMonster == 1 {
				fmt.Println("You encounter a monster!")
				attackMonster(health, attack) // Function to handle the fight
			} else {
				fmt.Println("You explore deeper into the forest. Nothing happens this time.")
			}
		case 2:
			fmt.Println("Health:", health)
			fmt.Println("Attack:", attack)
		case 3:
			fmt.Println("Exiting the game...")
			break
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		// Check if game continues
		if health <= 0 {
			fmt.Println("You have been defeated. Game Over!")
			break
		}
	}
}

func attackMonster(health int, attack int) {
	monsterHealth := 8   // Monster stats
	monsterAttack := 3  // Monster stats

	fmt.Println("Battle begins!")

	for {
		// Player turn
		fmt.Println("You attack the monster, dealing", attack, "damage.")
		monsterHealth -= attack

		if monsterHealth <= 0 {
			fmt.Println("You defeated the monster!")
			break
		}

		// Monster turn
		fmt.Println("The monster attacks you, dealing", monsterAttack, "damage.")
		health -= monsterAttack

		if health <= 0 {
			fmt.Println("The monster defeated you.")
			break
		}

		// Display updated health
		fmt.Println("Your health:", health)
		fmt.Println("Monster health:", monsterHealth)
	}
}
