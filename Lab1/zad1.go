package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const numBoxes = 3
	const winBox = 1

	var stayWins, switchWins int
	totalPlays := 10000

	for i := 0; i < totalPlays; i++ {
		boxes := make([]int, numBoxes)
		boxes[winBox] = 1

		playerChoice := rand.Intn(numBoxes)

		var revealedBox int
		for revealedBox == winBox || revealedBox == playerChoice {
			revealedBox = rand.Intn(numBoxes)
		}

		var newChoice int
		for newChoice == playerChoice || newChoice == revealedBox {
			newChoice = rand.Intn(numBoxes)
		}

		if playerChoice == winBox {
			stayWins++
		} else if newChoice == winBox {
			switchWins++
		}
	}

	fmt.Printf("Po %d rozgrywkach:\n", totalPlays)
	fmt.Printf("Wygrane po pozostaniu przy wyborze: %d\n", stayWins)
	fmt.Printf("Wygrane po zmianie wyboru: %d\n", switchWins)
}
