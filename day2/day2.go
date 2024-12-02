package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func readInput() [][]int {
	lines := utils.ReadInputLines()

	var reports [][]int

	for _, line := range lines {
		levels := strings.Fields(line)
		var report []int
		for i := 0; i < len(levels); i++ {
			num, _ := strconv.Atoi(levels[i])
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	return reports
}

func puzzle1() int {
	reports := readInput()

	count := 0

	for i := 0; i < len(reports); i++ {
		if isSafe(reports[i]) {
			count++
		}
	}

	return count
}

func isSafe(slice []int) bool {
	safe := true
	descending := slice[0] > slice[len(slice)-1]

	for j := 1; j < len(slice); j++ {
		curr := slice[j]
		prev := slice[j-1]

		if descending {
			if !isValid(prev, curr) {
				safe = false
				break
			}
		} else {
			if !isValid(curr, prev) {
				safe = false
				break
			}
		}
	}

	return safe
}

func isValid(a int, b int) bool {
	return (a-b) > 0 && (a-b) <= 3
}

func puzzle2() int {
	reports := readInput()

	count := 0

	for i := 0; i < len(reports); i++ {

		if isSafe(reports[i]) {
			count++
			continue
		}

		for k := 0; k < len(reports[i]); k++ {
			var slice []int
			slice = append(slice, reports[i][:k]...)
			slice = append(slice, reports[i][k+1:]...)

			if isSafe(slice) {
				count++
				break
			}
		}

	}

	return count
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
