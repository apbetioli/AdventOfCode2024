package main

import (
	"adventofcode/2024/utils"
	"fmt"
)

func readInput() [][]int {
	lines := utils.ReadInputLines()

	var matrix [][]int = make([][]int, len(lines))

	for r, line := range lines {
		var row []int = make([]int, len(line))

		for c, rune := range line {
			n := int(rune - '0')
			row[c] = n
		}

		matrix[r] = row
	}

	return matrix
}

func countTrails(matrix [][]int, r int, c int, level int) int {

	if !utils.IsValidCoordinate(matrix, r, c) {
		return 0
	}

	if matrix[r][c] != level {
		return 0
	}

	if matrix[r][c] == 9 {
		matrix[r][c] = -1 // remove from results
		return 1
	}

	return countTrails(matrix, r-1, c, level+1) +
		countTrails(matrix, r+1, c, level+1) +
		countTrails(matrix, r, c-1, level+1) +
		countTrails(matrix, r, c+1, level+1)
}

func copyMatrix(matrix [][]int) [][]int {
	matrixCopy := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		rowCopy := make([]int, len(matrix[i]))
		copy(rowCopy, matrix[i])
		matrixCopy[i] = rowCopy
	}
	return matrixCopy
}

func puzzle1() int {
	matrix := readInput()

	total := 0

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == 0 {
				matrixCopy := copyMatrix(matrix)
				score := countTrails(matrixCopy, r, c, 0)
				total += score
			}
		}
	}

	return total
}

func puzzle2() int {
	return 0
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
