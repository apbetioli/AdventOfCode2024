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

func readInput() (map[string][]coordinate, int, int, [][]string) {

	antennas := make(map[string][]coordinate)
	var matrix [][]string

	lines := utils.ReadInputLines()
	for lineIndex, line := range lines {
		chars := strings.Split(line, "")
		matrix = append(matrix, chars)
		for i := 0; i < len(chars); i++ {
			if chars[i] != "." {
				antennas[chars[i]] = append(antennas[chars[i]], coordinate{lineIndex, i})
			}
		}
	}

	height := len(lines)
	width := len(lines[0])

	fmt.Println(antennas)

	return antennas, height, width, matrix
}

func getAntinodes(p1 coordinate, p2 coordinate) (coordinate, coordinate) {

	distance := coordinate{p2.y - p1.y, p2.x - p1.x}

	a1 := coordinate{p1.y - distance.y, p1.x - distance.x}
	a2 := coordinate{p2.y + distance.y, p2.x + distance.x}

	fmt.Println("antinode", p1, p2, a1, a2)

	return a1, a2
}

func isValidAntinode(a coordinate, height int, width int) bool {
	return a.y >= 0 && a.y < height && a.x >= 0 && a.x < width
}

func puzzle1() int {
	antennas, height, width, matrix := readInput()

	count := 0

	for _, positions := range antennas {

		if len(positions) == 1 {
			continue
		}

		for p1 := 0; p1 < len(positions)-1; p1++ {
			for p2 := p1 + 1; p2 < len(positions); p2++ {
				a1, a2 := getAntinodes(positions[p1], positions[p2])
				if isValidAntinode(a1, height, width) {
					if matrix[a1.y][a1.x] != "#" {
						count++
					}
					matrix[a1.y][a1.x] = "#"
				}
				if isValidAntinode(a2, height, width) {
					if matrix[a2.y][a2.x] != "#" {
						count++
					}
					matrix[a2.y][a2.x] = "#"
				}
			}
		}
	}

	utils.Debug(matrix)

	return count
}

func puzzle2() int {
	return 0
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
