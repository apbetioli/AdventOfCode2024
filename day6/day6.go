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

func calculateStartPosition(matrix [][]string) (int, int) {
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] == "^" {
				return y, x
			}
		}
	}
	panic("Guard not found")
}

func puzzle1() int {

	matrix := readInput()
	y, x := calculateStartPosition(matrix)

	directions := [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	distinct := 1
	matrix[y][x] = "X"

	directionIndex := 0
	direction := directions[0]

	for {
		ny := y + direction[0]
		nx := x + direction[1]

		if ny < 0 || ny >= len(matrix) || nx < 0 || nx >= len(matrix[0]) {
			break
		}

		if matrix[ny][nx] == "#" {
			directionIndex++
			direction = directions[directionIndex%4]
		} else {
			y = ny
			x = nx
			if matrix[y][x] != "X" {
				matrix[y][x] = "X"
				distinct++
			}
		}
	}

	return distinct
}

func puzzle2() int {
	return 0
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
