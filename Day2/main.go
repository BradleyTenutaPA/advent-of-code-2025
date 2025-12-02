package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

// input has a list of ranges, each item in the list contains upper and lower range.
// invalid product ids are the ids where each number is duplicated twice. 6464 (64 twice) 44 (4 twice).
// No ids have leading zeros
// Make a list of all invalid ids, then add them together to get a sum and thats the result.
// Test result is 1227775554.

func main() {
	currentDir, _ := os.Getwd()
	path := filepath.Join(currentDir, "input.txt")
	dat, _ := os.ReadFile(path)
	ranges := strings.Split(string(dat), ",")
	allInvalidIds := []int{}

	for i := 0; i < len(ranges); i++ {
		lowerRange, UpperRange := parseRange(ranges[i])
		invalidIds := findInvalidIdsInRange(lowerRange, UpperRange)
		allInvalidIds = append(allInvalidIds, invalidIds...)
	}

	fmt.Println("All invalid ids:", allInvalidIds)
	fmt.Println("Sum of all invalid ids:", sumList(allInvalidIds))
}

func parseRange(r string) (int, int) {
	parts := strings.Split(r, "-")
	lowerRange, _ := strconv.Atoi(parts[0])
	UpperRange, _ := strconv.Atoi(parts[1])
	return lowerRange, UpperRange
}

func findInvalidIdsInRange(lowerRange int, upperRange int) []int {
	invalidIds := []int{}
	for id := lowerRange; id <= upperRange; id++ {
		idNumbers := splitInt(id)
		if isInvalidId(idNumbers) {
			invalidIds = append(invalidIds, id)
		}
	}
	return invalidIds
}

func isInvalidId(idNumbers []int) bool {
	// split into two lists down the middle, if added together and equal then its invaid.
	if len(idNumbers) > 1 {
		length := len(idNumbers)
		numberOfSplits := 1
		for i := 0; i < length; i++ {
			splits := splitListIntoSections(numberOfSplits, idNumbers)
			numberOfSplits++
			if len(splits) > 1 && areAllSplitsEqual(splits) {
				return true
			}
		}
	}
	return false
}

func splitListIntoSections(numberOfSplits int, list []int) [][]int {
	var j int
	var result [][]int
	for i := 0; i < len(list); i += numberOfSplits {
		j += numberOfSplits
		if j > len(list) {
			j = len(list)
		}
		// do what do you want to with the sub-slice, here just printing the sub-slices
		result = append(result, list[i:j])
	}
	return result
}

func areAllSplitsEqual(splits [][]int) bool {
	areAllSplitsEqual := true
	for i := 0; i < len(splits); i++ {
		for j := i + 1; j < len(splits); j++ {
			if !slices.Equal(splits[i], splits[j]) {
				areAllSplitsEqual = false
			}
		}
	}
	return areAllSplitsEqual
}

func sumList(ids []int) int {
	sum := 0
	for _, value := range ids {
		sum += value
	}
	return sum
}

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}
	return slc
}
