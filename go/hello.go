package main

import (
	"fmt"
)

func main() {
	solution1 := [][]int{
		{8, 2, 9, 3, 6, 5, 1, 4, 7},
		{6, 4, 3, 7, 1, 8, 5, 1, 9},
		{7, 5, 1, 4, 9, 2, 6, 3, 8},
		{3, 1, 8, 5, 2, 7, 4, 9, 6},
		{5, 9, 6, 1, 4, 1, 7, 8, 2},
		{4, 7, 2, 9, 8, 6, 3, 1, 5},
		{9, 8, 4, 6, 5, 1, 2, 7, 3},
		{2, 6, 7, 8, 3, 4, 9, 5, 1},
		{9, 3, 5, 2, 7, 9, 8, 6, 4},
	}
	solution2 := [][]int{
		{5, 8, 6, 4, 3, 7, 1, 9, 2},
		{1, 9, 4, 5, 8, 2, 3, 6, 7},
		{7, 2, 3, 9, 6, 1, 4, 5, 8},
		{2, 4, 7, 1, 9, 8, 6, 3, 5},
		{8, 3, 9, 6, 2, 4, 7, 2, 1},
		{6, 5, 1, 7, 2, 3, 8, 4, 9},
		{9, 7, 5, 3, 1, 6, 7, 8, 4},
		{3, 1, 8, 2, 4, 5, 9, 7, 6},
		{4, 6, 2, 8, 7, 9, 5, 1, 3},
	}
	solution3 := [][]int{
		{1, 8, 3, 2, 7, 4, 6, 5, 9},
		{9, 7, 4, 5, 8, 6, 3, 2, 1},
		{2, 6, 5, 1, 9, 3, 7, 4, 8},
		{5, 9, 2, 8, 3, 1, 4, 6, 7},
		{8, 4, 6, 7, 2, 5, 9, 1, 3},
		{7, 3, 1, 4, 6, 9, 2, 8, 5},
		{3, 5, 9, 6, 4, 8, 1, 7, 2},
		{6, 1, 7, 3, 5, 2, 8, 9, 4},
		{4, 2, 8, 9, 1, 7, 5, 3, 6},
	}

	fmt.Println(ValidateSolution(solution1))
	fmt.Println(ValidateSolution(solution2))
	fmt.Println(ValidateSolution(solution3))
}

func ValidateSolution(solution [][]int) (result bool) {
	for i, row := range solution {
		// check row
		isValidRow := ValidateArray(row)
		if !isValidRow {
			return false
		}

		// check col
		var col []int
		for j := range solution {
			col = append(col, solution[j][i])
		}
		isValidCol := ValidateArray(col)
		if !isValidCol {
			return false
		}
	}

	// check grids
	for i := 0; i < 9; i += 3 {
		rowSubset := solution[i : i+3]
		for j := 0; j < 9; j += 3 {
			var grid []int
			for _, row := range rowSubset {
				grid = append(grid, row[j], row[j+1], row[j+2])
			}
			isValidGrid := ValidateArray(grid)
			if !isValidGrid {
				return false
			}
		}
	}

	return true
}

func ValidateArray(array []int) (result bool) {
	return (Min(array) == 1) && (Max(array) == 9) && !ContainsDuplicate(array)
}

func Min(array []int) (min int) {
	for index, element := range array {
		if index == 0 || element < min {
			min = element
		}
	}
	return min
}

func Max(array []int) (max int) {
	for index, element := range array {
		if index == 0 || element > max {
			max = element
		}
	}
	return max
}

func ContainsDuplicate(array []int) (result bool) {
	entries := make(map[int]bool, 0)
	for _, element := range array {
		if entries[element] {
			return true
		}
		entries[element] = true
	}
	return false
}
