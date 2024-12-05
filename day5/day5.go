package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"strings"
)

func readInput() (map[int][]int, [][]int) {
	lines := utils.ReadInputLines()

	precedence := make(map[int][]int)
	var updates [][]int

	readingPrecedence := true

	for _, line := range lines {

		if line == "" {
			readingPrecedence = false
			continue
		}

		if readingPrecedence {
			split := strings.Split(line, "|")
			values := utils.StringArrayToIntArray(split)
			precedence[values[0]] = append(precedence[values[0]], values[1])

		} else {
			array := strings.Split(line, ",")
			updates = append(updates, utils.StringArrayToIntArray(array))
		}

	}
	return precedence, updates
}

func isValid(precedence map[int][]int, update []int) bool {

	seen := make(map[int]bool)

	for i := 0; i < len(update); i++ {
		current := update[i]
		pages := precedence[current]
		for _, page := range pages {
			if seen[page] {
				return false
			}
		}

		seen[current] = true
	}

	return true
}

func puzzle1() int {
	precedence, updates := readInput()

	var valid [][]int

	for _, update := range updates {
		if isValid(precedence, update) {
			valid = append(valid, update)
		}
	}

	var sum int
	for _, update := range valid {
		mid := len(update) / 2
		sum += update[mid]
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
