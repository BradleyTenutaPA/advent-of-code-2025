package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// The forklifts can only access a roll of paper if there are fewer than four rolls of paper in the eight adjacent positions.
// Positions to check
// CCC
// C@C
// CCC

type PaperRoll struct {
	accessible bool
	removed    bool
	value      byte
}

func main() {
	currentDir, _ := os.Getwd()
	path := filepath.Join(currentDir, "input.txt")
	dat, _ := os.ReadFile(path)
	rows := strings.Split(string(dat), "\n")

	// Create a grid of PaperRolls.
	grid := make([][]PaperRoll, len(rows))
	for i, row := range rows {
		grid[i] = make([]PaperRoll, len(row))
		for j := 0; j < len(row); j++ {
			grid[i][j] = PaperRoll{
				accessible: false,
				removed:    false,
				value:      row[j],
			}
		}
	}

	numberOfRollsRemoved := 0
	rollsWereRemoved := true

	for rollsWereRemoved {
		rollsWereRemoved = false

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				// Check each position in the grid.
				if grid[i][j].value == 64 && !grid[i][j].removed && isAccessible(grid, i, j) {
					grid[i][j].accessible = true
					grid[i][j].removed = true

					numberOfRollsRemoved++
					rollsWereRemoved = true
				}
			}
		}
	}

	//fmt.Println("Updated Grid:", grid)
	fmt.Println("Number of rolls removed:", numberOfRollsRemoved)
}

func isAccessible(grid [][]PaperRoll, row int, col int) bool {
	indexesToCheck := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	numberOfAdjacentRolls := 0
	for _, indexToCheck := range indexesToCheck {
		if isValidIndex(grid, indexToCheck[0]+row, indexToCheck[1]+col) {
			paperRoll := grid[indexToCheck[0]+row][indexToCheck[1]+col]
			//fmt.Println("Position to check", paperRoll)
			if paperRoll.value == 64 && !paperRoll.removed {
				numberOfAdjacentRolls++
			}
		}
	}
	return numberOfAdjacentRolls < 4
}

func isValidIndex(grid [][]PaperRoll, row int, col int) bool {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {
		return false
	}
	return true
}
