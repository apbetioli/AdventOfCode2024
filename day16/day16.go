package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"math"
	"strings"
)

type Coordinate struct {
	y int
	x int
}

var directions = [4]Coordinate{
	{0, 1},  //east
	{1, 0},  //south
	{0, -1}, //west
	{-1, 0}, //north
}
var directionsChar = [4]string{
	">", //east
	"v", //south
	"<", //west
	"^", //north
}

func right(index int) int {
	return (index + 1) % 4
}

func left(index int) int {
	return (index + 3) % 4
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

func puzzle1() int {

	matrix := readInput()

	reindeer := findReindeer(matrix)

	score := dfs(matrix, reindeer, 0, 0)

	return int(score)
}

func getNext(reindeer Coordinate, direction Coordinate) Coordinate {
	return Coordinate{reindeer.y + direction.y, reindeer.x + direction.x}
}

func dfs(matrix [][]string, reindeer Coordinate, dir int, score float64) float64 {

	if matrix[reindeer.y][reindeer.x] == "E" {
		return score
	}

	if matrix[reindeer.y][reindeer.x] != "." && matrix[reindeer.y][reindeer.x] != "S" {
		// Can't continue this direction
		return math.Inf(1)
	}

	// Mark as visited
	matrix[reindeer.y][reindeer.x] = directionsChar[dir]

	// utils.Debug(matrix)

	current := dfs(matrix, getNext(reindeer, directions[dir]), dir, score+1)
	left := dfs(matrix, getNext(reindeer, directions[left(dir)]), left(dir), score+1001)
	right := dfs(matrix, getNext(reindeer, directions[right(dir)]), right(dir), score+1001)

	//Undo visited
	matrix[reindeer.y][reindeer.x] = "."

	// utils.Debug(matrix)

	return math.Min(math.Min(current, left), right)
}

func findReindeer(matrix [][]string) Coordinate {
	for y := range matrix {
		for x, val := range matrix[y] {
			if val == "S" {
				return Coordinate{y, x}
			}
		}
	}
	panic("Reindeer not found")
}

func puzzle2() int {
	return 0
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
