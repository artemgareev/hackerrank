package bomber_man

import (
	"fmt"
)

const ExplosionSign = "O"
const EmptyCellSign = "."

//https://www.hackerrank.com/challenges/bomber-man/problem
func BomberMan(n int, grid []string) {
	copyGrid := func(grid []string) []string {
		var newGrid []string
		newGrid = append(newGrid, grid...)
		return newGrid
	}
	printGrid := func(grid []string) {
		for _, line := range grid {
			fmt.Println(line)
		}
		fmt.Println()
	}

	if n == 0 || n == 1 {
		printGrid(grid)
		return
	}

	plantedBomb := plantBomb(grid, copyGrid(grid))
	explodeBomb1 := explodeBomb(grid, copyGrid(plantedBomb))
	explodeBomb2 := explodeBomb(explodeBomb1, copyGrid(plantedBomb))

	if (n-1)%4 == 0 {
		printGrid(explodeBomb2)
	} else if (n-1)%2 == 0 {
		printGrid(explodeBomb1)
	} else {
		printGrid(plantedBomb)
	}
}

func plantBomb(baseGrid []string, grid []string) []string {
	x := len(baseGrid[0])
	y := len(baseGrid)
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if baseGrid[i][j:j+1] == EmptyCellSign {
				grid[i] = changeAtIndex(grid[i], ExplosionSign, j)
			}
		}
	}
	return grid
}

func explodeBomb(baseGrid []string, grid []string) []string {
	x := len(baseGrid[0])
	y := len(baseGrid)

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if baseGrid[i][j:j+1] == ExplosionSign {
				grid[i] = changeAtIndex(grid[i], EmptyCellSign, j)
				if i < y-1 {
					grid[i+1] = changeAtIndex(grid[i+1], EmptyCellSign, j)
				}
				if i > 0 {
					grid[i-1] = changeAtIndex(grid[i-1], EmptyCellSign, j)
				}
				if j < x-1 {
					grid[i] = changeAtIndex(grid[i], EmptyCellSign, j+1)
				}
				if j > 0 {
					grid[i] = changeAtIndex(grid[i], EmptyCellSign, j-1)
				}
			}
		}
	}

	return grid
}

func changeAtIndex(str string, replacement string, index int) string {
	return str[:index] + replacement + str[index+1:]
}
