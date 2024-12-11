package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"strings"
)

type coordinate struct {
	y int
	x int
}

func (position *coordinate) sum(direction coordinate) coordinate {
	return coordinate{position.y + direction.y, position.x + direction.x}
}

func readInput() [][]string {
	var matrix [][]string
	lines := utils.ReadInputLines()
	for _, line := range lines {
		chars := strings.Split(line, "")
		matrix = append(matrix, chars)
	}
	return matrix
}

func calculateStartPosition(matrix [][]string) coordinate {
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] == "^" {
				return coordinate{y, x}
			}
		}
	}
	panic("Guard not found")
}

func directioner() (coordinate, func() coordinate) {
	directions := [4]coordinate{
		{-1, 0}, //up
		{0, 1},  //right
		{1, 0},  //left
		{0, -1}, //down
	}

	directionIndex := 0
	direction := directions[0]

	return direction, func() coordinate {
		directionIndex++
		direction = directions[directionIndex%4]
		return direction
	}

}

func puzzle1() int {

	matrix := readInput()
	position := calculateStartPosition(matrix)

	distinct := 1
	matrix[position.y][position.x] = "X"

	direction, turnRight := directioner()

	for {
		next := position.sum(direction)

		if !utils.IsValidCoordinate(matrix, next.y, next.x) {
			break
		}

		if matrix[next.y][next.x] == "#" {
			// There is a wall in front of the guard, change direction
			direction = turnRight()
		} else {
			position = next
			if matrix[position.y][position.x] != "X" {
				matrix[position.y][position.x] = "X"
				distinct++
			}
		}
	}

	return distinct
}

/*
*
The idea is to check if at any step if changing direction would match
a previously run path in the same direction. If so, the next position can be
an obstruction (if not already)

Create a generic search function that detects loops.
For each position simulate an obstacle and check for loop.
*/
func puzzle2() int {
	matrix := readInput()
	position := calculateStartPosition(matrix)

	mark := [4]string{"|", "-", "|", "-"}

	obstructions := 0

	direction, turnRight := directioner()
	directionIndex := 0

	for {
		next := position.sum(direction)

		if !utils.IsValidCoordinate(matrix, next.y, next.x) {
			break
		}

		if matrix[next.y][next.x] == "#" {
			// There is a wall in front of the guard, change direction
			matrix[position.y][position.x] = "+"
			direction = turnRight()
			directionIndex++
			continue
		}

		//Suppose if we change direction now, would it be a loop?
		// if loopFound(matrix, directions, directionIndex, y, x) {
		// 	obstructions++
		// 	matrix[ny][nx] = "O"
		// }

		position = next
		if matrix[position.y][position.x] == "." {
			matrix[position.y][position.x] = mark[directionIndex%4]
		} else if matrix[position.y][position.x] != "^" {
			matrix[position.y][position.x] = "+"
		}
	}

	utils.Debug(matrix)

	return obstructions
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
