package main

import (
	"fmt"
)

func main() {
	solutions := [][][]int{
		{
			{8, 2, 9, 3, 6, 5, 1, 4, 7},
			{6, 4, 3, 7, 1, 8, 5, 1, 9},
			{7, 5, 1, 4, 9, 2, 6, 3, 8},
			{3, 1, 8, 5, 2, 7, 4, 9, 6},
			{5, 9, 6, 1, 4, 1, 7, 8, 2},
			{4, 7, 2, 9, 8, 6, 3, 1, 5},
			{9, 8, 4, 6, 5, 1, 2, 7, 3},
			{2, 6, 7, 8, 3, 4, 9, 5, 1},
			{9, 3, 5, 2, 7, 9, 8, 6, 4},
		},
		{
			{5, 8, 6, 4, 3, 7, 1, 9, 2},
			{1, 9, 4, 5, 8, 2, 3, 6, 7},
			{7, 2, 3, 9, 6, 1, 4, 5, 8},
			{2, 4, 7, 1, 9, 8, 6, 3, 5},
			{8, 3, 9, 6, 2, 4, 7, 2, 1},
			{6, 5, 1, 7, 2, 3, 8, 4, 9},
			{9, 7, 5, 3, 1, 6, 7, 8, 4},
			{3, 1, 8, 2, 4, 5, 9, 7, 6},
			{4, 6, 2, 8, 7, 9, 5, 1, 3},
		},
		{
			{1, 8, 3, 2, 7, 4, 6, 5, 9},
			{9, 7, 4, 5, 8, 6, 3, 2, 1},
			{2, 6, 5, 1, 9, 3, 7, 4, 8},
			{5, 9, 2, 8, 3, 1, 4, 6, 7},
			{8, 4, 6, 7, 2, 5, 9, 1, 3},
			{7, 3, 1, 4, 6, 9, 2, 8, 5},
			{3, 5, 9, 6, 4, 8, 1, 7, 2},
			{6, 1, 7, 3, 5, 2, 8, 9, 4},
			{4, 2, 8, 9, 1, 7, 5, 3, 6},
		},
	}

	for _, sol := range solutions {
		validateSolution(sol)
	}
}

func validateSolution(solution [][]int) {
	for i, row := range solution {
		isValidRow := validSequence(row)
		if !isValidRow {
			fmt.Printf("Invalid - row %d\n", i+1)
			return
		}

		isValidCol := validCol(i, solution)
		if !isValidCol {
			return
		}
	}

	hasValidGrids := validGrids(solution)
	if !hasValidGrids {
		return
	}

	fmt.Println("Valid :)")
}

func validGrids(solution [][]int) (valid bool) {
	for i := 0; i < 9; i += 3 {
		rowSubset := solution[i : i+3]
		for j := 0; j < 9; j += 3 {
			var grid []int
			for _, row := range rowSubset {
				grid = append(grid, row[j], row[j+1], row[j+2])
			}
			isValidGrid := validSequence(grid)
			if !isValidGrid {
				fmt.Printf("Invalid - sub-grid (%d, %d)\n", (j/3)+1, (i/3)+1)
				return false
			}
		}
	}
	return true
}

func validCol(i int, solution [][]int) (valid bool) {
	var col []int
	for j := range solution {
		col = append(col, solution[j][i])
	}
	isValidCol := validSequence(col)
	if !isValidCol {
		fmt.Printf("Invalid - column %d\n", i+1)
	}
	return isValidCol
}

func validSequence(array []int) (result bool) {
	return (min(array) == 1) && (max(array) == 9) && !hasDuplicate(array)
}

func min(array []int) (min int) {
	for index, element := range array {
		if index == 0 || element < min {
			min = element
		}
	}
	return min
}

func max(array []int) (max int) {
	for index, element := range array {
		if index == 0 || element > max {
			max = element
		}
	}
	return max
}

func hasDuplicate(array []int) (result bool) {
	entries := make(map[int]bool, 0)
	for _, element := range array {
		if entries[element] {
			return true
		}
		entries[element] = true
	}
	return false
}
