package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"strings"
)

func readInput() [][]string {
	var matrix [][]string
	lines := utils.ReadInputLines()
	for _, line := range lines {
		chars := strings.Split(line, "")
		matrix = append(matrix, chars)
	}
	return matrix
}

func countXmas(matrix [][]string, r int, c int) int {
	words := 0

	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	letters := []string{"X", "M", "A", "S"}

	for _, dir := range directions {

		s := "X"
		y := r
		x := c
		for i := 1; i < 4; i++ {
			y += dir[0]
			x += dir[1]
			if y < 0 || y >= len(matrix) || x < 0 || x >= len(matrix[0]) {
				break
			}
			if letters[i] != matrix[y][x] {
				break
			}
			s += matrix[y][x]
		}

		if s == "XMAS" {
			words++
		}
	}

	return words
}

func puzzle1() int {
	matrix := readInput()
	total := 0

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == "X" {
				total += countXmas(matrix, r, c)
			}
		}
	}

	return total
}

func countX_mas(matrix [][]string, r int, c int) int {

	directions := [4][2]int{
		{-1, -1}, {-1, 1},
		{1, -1}, {1, 1},
	}

	m := 2
	s := 2

	for _, dir := range directions {

		y := r + dir[0]
		x := c + dir[1]

		if y < 0 || y >= len(matrix) || x < 0 || x >= len(matrix[0]) {
			break
		}

		if matrix[y][x] == "M" {
			m--
		} else if matrix[y][x] == "S" {
			s--
		} else {
			break
		}
	}

	if m != 0 || s != 0 {
		return 0
	}

	if matrix[r-1][c-1] == matrix[r+1][c+1] { // MAM and SAS
		return 0
	}

	return 1
}

func puzzle2() int {
	matrix := readInput()
	total := 0

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == "A" {
				total += countX_mas(matrix, r, c)
			}
		}
	}

	return total
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
