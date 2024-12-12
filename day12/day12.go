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

func readInput() [][]string {
	var matrix [][]string

	lines := utils.ReadInputLines()
	for _, line := range lines {
		chars := strings.Split(line, "")
		matrix = append(matrix, chars)
	}

	return matrix
}

func calculateRegion(matrix [][]string, visited [][]bool, y int, x int) (int, int) {

	visited[y][x] = true

	var neighbors []coordinate

	directions := [4]coordinate{
		{-1, 0}, //up
		{0, 1},  //right
		{1, 0},  //left
		{0, -1}, //down
	}

	for _, direction := range directions {
		neighbor := coordinate{y + direction.y, x + direction.x}
		if !utils.IsValidCoordinate(matrix, neighbor.y, neighbor.x) {
			continue
		}
		if matrix[neighbor.y][neighbor.x] != matrix[y][x] {
			continue
		}
		neighbors = append(neighbors, neighbor)
	}

	area := 1
	perimeter := 4 - len(neighbors)

	for _, neighbor := range neighbors {
		if visited[neighbor.y][neighbor.x] {
			continue
		}

		na, np := calculateRegion(matrix, visited, neighbor.y, neighbor.x)
		area += na
		perimeter += np
	}

	return area, perimeter
}

func puzzle1() int {

	matrix := readInput()
	visited := make([][]bool, len(matrix))
	for i := range matrix {
		visited[i] = make([]bool, len(matrix[i]))
	}

	cost := 0

	for y := range matrix {
		for x := range matrix[y] {
			if !visited[y][x] {
				area, perimeter := calculateRegion(matrix, visited, y, x)
				cost += area * perimeter
			}
		}
	}

	return cost
}

func puzzle2() int {
	return 0
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
