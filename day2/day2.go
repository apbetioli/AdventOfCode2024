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

		safe := true
		descending := reports[i][0] > reports[i][1]

		for j := 1; j < len(reports[i]); j++ {
			curr := reports[i][j]
			prev := reports[i][j-1]

			if descending {
				if (prev-curr) <= 0 || (prev-curr) > 3 {
					safe = false
					break
				}
			} else {
				if (curr-prev) <= 0 || (curr-prev) > 3 {
					safe = false
					break
				}
			}
		}

		if safe {
			count++
		}

	}

	return count
}

func main() {
	fmt.Println(puzzle1())
}
