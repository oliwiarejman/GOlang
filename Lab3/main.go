package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	empty = iota
	tree
	burning
)

type forest [][]int

func main() {
	rand.Seed(time.Now().UnixNano())

	width, height := 10, 10

	f := makeForest(width, height)

	plantTrees(f)

	fmt.Println("Początkowy stan lasu:")
	printForest(f)

	burnForest(f)

	burnedPercent := calculateBurnedPercent(f)
	fmt.Printf("Procent spalonych drzew: %.2f%%\n", burnedPercent)

	fmt.Println("Stan lasu po pożarze:")
	printForest(f)
}

func makeForest(width, height int) forest {
	f := make(forest, height)
	for i := range f {
		f[i] = make([]int, width)
	}
	return f
}

func plantTrees(f forest) {
	for y := range f {
		for x := range f[y] {
			if rand.Intn(2) == 0 {
				f[y][x] = tree
			} else {
				f[y][x] = empty 
			}
		}
	}
}


func burnForest(f forest) {

	firstBurningX := rand.Intn(len(f[0]))
	firstBurningY := rand.Intn(len(f))
	f[firstBurningY][firstBurningX] = burning


	igniteNeighbors(f, firstBurningX, firstBurningY)
	f[firstBurningY][firstBurningX] = empty

	for {
		changed := false
		for y := range f {
			for x := range f[y] {
				if f[y][x] == tree && isBurning(f, x, y) {
					f[y][x] = burning
					igniteNeighbors(f, x, y)
					f[y][x] = empty
					changed = true
				}
			}
		}
		if !changed {
			break
		}
	}
}



func isBurning(f forest, x, y int) bool {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}
			ny, nx := y+dy, x+dx
			if ny >= 0 && ny < len(f) && nx >= 0 && nx < len(f[ny]) && f[ny][nx] == burning {
				return true
			}
		}
	}
	return false
}


func igniteNeighbors(f forest, x, y int) {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			ny, nx := y+dy, x+dx
			if ny >= 0 && ny < len(f) && nx >= 0 && nx < len(f[ny]) && (dx != 0 || dy != 0) && f[ny][nx] == tree {
				f[ny][nx] = burning
			}
		}
	}
}


func calculateBurnedPercent(f forest) float64 {
	totalTrees := 0
	burnedTrees := 0
	for y := range f {
		for x := range f[y] {
			if f[y][x] == tree {
				totalTrees++
			} else if f[y][x] == burning {
				burnedTrees++
			}
		}
	}
	return float64(burnedTrees) / float64(totalTrees) * 100
}

func printForest(f forest) {
	for _, row := range f {
		for _, cell := range row {
			switch cell {
			case empty:
				fmt.Print(" ")
			case tree:
				fmt.Print("T")
			case burning:
				fmt.Print("X")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
