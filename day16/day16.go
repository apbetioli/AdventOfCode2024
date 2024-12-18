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

	return score
}

func getNeighbor(reindeer Coordinate, direction Coordinate) Coordinate {
	return Coordinate{reindeer.y + direction.y, reindeer.x + direction.x}
}

func dfs(matrix [][]string, reindeer Coordinate, dir int, score int, distance [][]int) int {

	if matrix[reindeer.y][reindeer.x] == "E" {
		distance[reindeer.y][reindeer.x] = score
		return score
	}

	if matrix[reindeer.y][reindeer.x] != "." && matrix[reindeer.y][reindeer.x] != "S" {
		// Can't continue this direction
		return math.MaxInt
	}

	if distance[reindeer.y][reindeer.x] > 0 && score > 0 && score > distance[reindeer.y][reindeer.x] {
		// There is a better path
		return math.MaxInt
	}

	distance[reindeer.y][reindeer.x] = score

	// Mark as visited
	matrix[reindeer.y][reindeer.x] = directionsChar[dir]

	// utils.Debug(matrix)

	current := dfs(matrix, getNeighbor(reindeer, directions[dir]), dir, score+1, distance)
	left := dfs(matrix, getNeighbor(reindeer, directions[left(dir)]), left(dir), score+1001, distance)
	right := dfs(matrix, getNeighbor(reindeer, directions[right(dir)]), right(dir), score+1001, distance)

	//Undo visited
	matrix[reindeer.y][reindeer.x] = "."

	// utils.Debug(matrix)

	res := current
	if left < res && left < right {
		res = left
		distance[reindeer.y][reindeer.x] += 1000
	}
	if right < res && right < left {
		res = right
		distance[reindeer.y][reindeer.x] += 1000
	}

	return res
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

func findTarget(matrix [][]string) Coordinate {
	for y := range matrix {
		for x, val := range matrix[y] {
			if val == "E" {
				return Coordinate{y, x}
			}
		}
	}
	panic("Reindeer not found")
}

func puzzle2() int {
	defer utils.Duration(utils.Track("puzzle1"))

	matrix := readInput()

	reindeer := findReindeer(matrix)
	target := findTarget(matrix)

	var distance = make([][]int, len(matrix))
	for y := range distance {
		distance[y] = make([]int, len(matrix[y]))
	}

	score := dfs(matrix, reindeer, 0, 0, distance)
	distance[target.y][target.x] = score

	var queue utils.Queue[Coordinate]
	queue.Enqueue(target)
	count := 0

	for queue.Length > 0 {
		current := queue.Dequeue()

		if matrix[current.y][current.x] == "O" {
			continue
		}

		matrix[current.y][current.x] = "O"
		currentDist := distance[current.y][current.x] % 1000
		count++

		for _, dir := range directions {
			nPos := getNeighbor(current, dir)

			if distance[nPos.y][nPos.x] == 0 {
				continue
			}

			if distance[nPos.y][nPos.x]%1000 == currentDist-1 && distance[nPos.y][nPos.x] < distance[current.y][current.x] {
				queue.Enqueue(nPos)
			}
		}
	}

	// utils.Debug(matrix)
	// utils.DebugInt(distance)

	return count
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
