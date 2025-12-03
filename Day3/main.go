package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	totalJoltage := 0
	currentDir, _ := os.Getwd()
	path := filepath.Join(currentDir, "test-input.txt")
	dat, _ := os.ReadFile(path)
	banks := strings.Split(string(dat), "\n")
	for i := 0; i < len(banks); i++ {
		selected := getLargestCombinationOfBatteries(banks[i])
		value, _ := strconv.Atoi(selected)
		totalJoltage += value
	}
	fmt.Println("Total joltage:", totalJoltage)
}

// e.g. bank = "234234234234278" (length = 15)
func getLargestCombinationOfBatteries(bank string) string {
	// TODO:
	//   - Create a stack
	//   - Track the number of removals we can do
	//   - Replace element at the top of the stack when current element is greater and we have removals left.
	//   - When all removals are gone, add all to stack.
	totalToRemove := len(bank) - 12
	stack := make([]byte, 0, len(bank))
	for i := 0; i < len(bank); i++ {
		battery := bank[i] // indexing a string does not return a charcter but a byte object.
		// We can compare characters as they are represented as numbers.
		for len(stack) > 0 && totalToRemove > 0 && stack[len(stack)-1] < battery {
			//fmt.Println("charatcer to remove from stack:", string(stack[len(stack)-1]))
			stack = stack[:len(stack)-1]
			totalToRemove--
		}
		stack = append(stack, battery)
	}
	if totalToRemove > 0 {
		stack = stack[:len(stack)-totalToRemove]
	}
	if len(stack) > 12 {
		stack = stack[:12]
	}
	return string(stack)
}
