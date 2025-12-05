package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// they can't figure out which of their ingredients are fresh and which are spoiled.

type IngredientIdRange struct {
	lowerRange int
	upperRange int
}

// NOTES: Storing my first attempt here. This was an actual itration of values that would of taken hours to run.

// func main() {
// 	currentDir, _ := os.Getwd()
// 	path := filepath.Join(currentDir, "input.txt")
// 	dat, _ := os.ReadFile(path)
// 	fileRows := strings.Split(string(dat), "\n")
// 	ingredientIdRanges := []IngredientIdRange{}
// 	isIterationOfIds := false
// 	for i := 0; i < len(fileRows); i++ {
// 		if fileRows[i] == "" {
// 			isIterationOfIds = true
// 			break
// 		}
// 		if !isIterationOfIds {
// 			fileRowsParts := strings.Split(fileRows[i], "-")
// 			lowerRange, _ := strconv.Atoi(fileRowsParts[0])
// 			upperRange, _ := strconv.Atoi(fileRowsParts[1])
// 			ingredientIdRanges = append(ingredientIdRanges, IngredientIdRange{lowerRange: lowerRange, upperRange: upperRange})
// 		}
// 	}
// 	totalFreshIngredientIdsMap := make(map[int]struct{})
// 	for i := 0; i < len(ingredientIdRanges); i++ {
// 		fmt.Printf("Iterating through range: %d of %d\n", i+1, len(ingredientIdRanges))
// 		for j := ingredientIdRanges[i].lowerRange; j <= ingredientIdRanges[i].upperRange; j++ {
// 			if _, exists := totalFreshIngredientIdsMap[j]; !exists {
// 				//fmt.Println("Adding ID:", j)
// 				totalFreshIngredientIdsMap[j] = struct{}{}
// 			}
// 		}
// 	}
// 	fmt.Println("totalFreshIngredientIds", len(totalFreshIngredientIdsMap))
// }

func main() {
	currentDir, _ := os.Getwd()
	path := filepath.Join(currentDir, "input.txt")
	dat, _ := os.ReadFile(path)
	fileRows := strings.Split(string(dat), "\n")
	ingredientIdRanges := []IngredientIdRange{}

	isIterationOfIds := false
	for i := 0; i < len(fileRows); i++ {
		if fileRows[i] == "" {
			isIterationOfIds = true
			break
		}
		if !isIterationOfIds {
			fileRowsParts := strings.Split(fileRows[i], "-")
			lowerRange, _ := strconv.Atoi(fileRowsParts[0])
			upperRange, _ := strconv.Atoi(fileRowsParts[1])
			ingredientIdRanges = append(ingredientIdRanges, IngredientIdRange{lowerRange: lowerRange, upperRange: upperRange})
		}
	}

	sort.Slice(ingredientIdRanges, func(i, j int) bool {
		return ingredientIdRanges[i].lowerRange < ingredientIdRanges[j].lowerRange
	})
	//fmt.Println("ingredientIdRanges", ingredientIdRanges)

	// Then merge them together to have a much shorter amount of ranges.
	merged := []IngredientIdRange{}
	for _, r := range ingredientIdRanges {
		// If merged is empty or current range does not overlap, add it as new range
		if len(merged) == 0 || merged[len(merged)-1].upperRange < r.lowerRange-1 {
			merged = append(merged, r)
			continue
		}
		// Otherwise, merge with the last range
		last := &merged[len(merged)-1]
		if r.upperRange > last.upperRange {
			last.upperRange = r.upperRange
		}
	}
	//fmt.Println("merged", merged)

	total := 0
	for _, r := range merged {
		// 5 - 3 = 2. So need to add 1 for it to be 3.
		total += r.upperRange - r.lowerRange + 1
	}
	fmt.Println("totalFreshIngredientIds", total)
}
