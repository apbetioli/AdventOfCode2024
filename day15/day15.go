package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"strings"
)

func readInput() ([][]string, []string) {
	lines := utils.ReadInputLines()

	var matrix [][]string
	i := 0
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}
		matrix = append(matrix, strings.Split(lines[i], ""))
	}

	var movements []string

	for ; i < len(lines); i++ {
		movements = append(movements, strings.Split(lines[i], "")...)
	}

	return matrix, movements
}

func getRobot(matrix [][]string) [2]int {
	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] == "@" {
				return [2]int{y, x}
			}
		}
	}
	panic("Robot not found")
}

func getBoxes(matrix [][]string) [][2]int {
	var boxes [][2]int
	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] == "O" {
				boxes = append(boxes, [2]int{y, x})
			}
		}
	}
	return boxes
}

func puzzle1() int {
	matrix, movements := readInput()

	// fmt.Println("Initial state:")
	// utils.Debug(matrix)
	// fmt.Println()

	robot := getRobot(matrix)

	for _, move := range movements {
		y := robot[0]
		x := robot[1]

		if move == ">" {
			if matrix[y][x+1] == "O" {
				xi := x + 1
				for ; matrix[y][xi] == "O"; xi++ {
				}
				if matrix[y][xi] == "." {
					matrix[y][xi] = "O"
					matrix[y][x+1] = "."
				}
			}
			if matrix[y][x+1] == "." {
				robot[1] = x + 1
				matrix[y][x+1] = "@"
				matrix[y][x] = "."
			}
		} else if move == "<" {
			if matrix[y][x-1] == "O" {
				xi := x - 1
				for ; matrix[y][xi] == "O"; xi-- {
				}
				if matrix[y][xi] == "." {
					matrix[y][xi] = "O"
					matrix[y][x-1] = "."
				}
			}
			if matrix[y][x-1] == "." {
				robot[1] = x - 1
				matrix[y][x-1] = "@"
				matrix[y][x] = "."
			}
		} else if move == "^" {
			if matrix[y-1][x] == "O" {
				yi := y - 1
				for ; matrix[yi][x] == "O"; yi-- {
				}
				if matrix[yi][x] == "." {
					matrix[yi][x] = "O"
					matrix[y-1][x] = "."
				}
			}
			if matrix[y-1][x] == "." {
				robot[0] = y - 1
				matrix[y-1][x] = "@"
				matrix[y][x] = "."
			}
		} else if move == "v" {
			if matrix[y+1][x] == "O" {
				yi := y + 1
				for ; matrix[yi][x] == "O"; yi++ {
				}
				if matrix[yi][x] == "." {
					matrix[yi][x] = "O"
					matrix[y+1][x] = "."
				}
			}
			if matrix[y+1][x] == "." {
				robot[0] = y + 1
				matrix[y+1][x] = "@"
				matrix[y][x] = "."
			}
		}

		// fmt.Println("Move", move, ":")
		// utils.Debug(matrix)
		// fmt.Println()
	}

	boxes := getBoxes(matrix)

	sum := 0
	for _, box := range boxes {
		gps := box[0]*100 + box[1]
		sum += gps
	}

	return sum
}

func puzzle2() int {
	return 0
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
