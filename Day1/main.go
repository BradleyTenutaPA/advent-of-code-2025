package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 2")
	position := 50

	currentDir, _ := os.Getwd()
	path := filepath.Join(currentDir, "test-input.txt")
	dat, _ := os.ReadFile(path)

	turns := strings.Split(string(dat), "\n")
	numberOfZeros := 0
	zeros := 0

	for i := 0; i < len(turns); i++ {
		fmt.Printf("before making turn: rotation: %s position: %d\n", turns[i], position)
		position, zeros = makeTurn(position, turns[i])
		numberOfZeros += zeros
		if position == 0 {
			numberOfZeros++
		}
		fmt.Printf("After making turn: rotation: %s position: %d zeros: %d, total number of zeros: %d\n", turns[i], position, zeros, numberOfZeros)
	}

	fmt.Println("Number of times dial points at zero:", numberOfZeros)
	fmt.Println("Final position:", position)
}

func makeTurn(position int, turn string) (int, int) {
	numberOfZeros := 0
	rotation := turn[0:1]
	numberOfClicks, _ := strconv.Atoi(turn[1:])
	switch rotation {
	case "L":
		newPosition := position - numberOfClicks
		if newPosition < 0 {
			newPosition, numberOfZeros = makeLeftRotation(position, numberOfClicks)
		}
		position = newPosition
	case "R":
		newPosition := position + numberOfClicks
		if newPosition > 99 {
			newPosition, numberOfZeros = makeRightRotation(position, numberOfClicks)
		}
		position = newPosition
	}
	return position, numberOfZeros
}

func makeLeftRotation(position int, numberOfClicks int) (int, int) {
	numberOfZeros := 0
	for numberOfClicks > 0 {
		if numberOfClicks >= position {
			numberOfClicks = numberOfClicks - position
			old_position := position
			position = 0
			if numberOfClicks > 0 {
				numberOfClicks = numberOfClicks - 1
				position = 99
				if old_position != 0 {
					numberOfZeros++
				}
			}
		} else {
			position = position - numberOfClicks
			numberOfClicks = 0
		}
	}
	return position, numberOfZeros
}

func makeRightRotation(position int, numberOfClicks int) (int, int) {
	numberOfZeros := 0
	for numberOfClicks > 0 {
		if (numberOfClicks + position) > 99 {
			numberOfClicks = numberOfClicks - (99 - position)
			position = 99
			if numberOfClicks > 0 {
				numberOfClicks = numberOfClicks - 1
				position = 0
			}
		} else {
			position = position + numberOfClicks
			numberOfClicks = 0
		}
		if position == 0 && numberOfClicks > 0 {
			numberOfZeros++
		}
	}
	return position, numberOfZeros
}
