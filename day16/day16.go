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
	defer utils.Duration(utils.Track("puzzle1"))

	matrix := readInput()

	reindeer := findReindeer(matrix)

	var distance = make([][]int, len(matrix))
	for y := range distance {
		distance[y] = make([]int, len(matrix[y]))
	}

	score := dfs(matrix, reindeer, 0, 0, distance)

	return int(score)
}

func getNext(reindeer Coordinate, direction Coordinate) Coordinate {
	return Coordinate{reindeer.y + direction.y, reindeer.x + direction.x}
}

func dfs(matrix [][]string, reindeer Coordinate, dir int, score float64, distance [][]int) float64 {

	if matrix[reindeer.y][reindeer.x] == "E" {
		return score
	}

	if matrix[reindeer.y][reindeer.x] != "." && matrix[reindeer.y][reindeer.x] != "S" {
		// Can't continue this direction
		return math.Inf(1)
	}

	if distance[reindeer.y][reindeer.x] > 0 && score > 0 && int(score) > distance[reindeer.y][reindeer.x] {
		// There is a better path
		return math.Inf(1)
	}

	distance[reindeer.y][reindeer.x] = int(score)

	// utils.DebugInt(distance)

	// Mark as visited
	matrix[reindeer.y][reindeer.x] = directionsChar[dir]

	// utils.Debug(matrix)

	current := dfs(matrix, getNext(reindeer, directions[dir]), dir, score+1, distance)
	left := dfs(matrix, getNext(reindeer, directions[left(dir)]), left(dir), score+1001, distance)
	right := dfs(matrix, getNext(reindeer, directions[right(dir)]), right(dir), score+1001, distance)

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
