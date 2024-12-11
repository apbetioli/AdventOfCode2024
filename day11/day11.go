package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func readInput() []string {
	lines := utils.ReadInputLines()
	return strings.Split(lines[0], " ")
}

func blink(line []string) []string {
	var newLine []string
	for _, stone := range line {
		if stone == "0" {
			newLine = append(newLine, "1")
		} else if len(stone)%2 == 0 {
			mid := len(stone) / 2
			s1 := stone[:mid]
			s2 := stone[mid:]
			n1, _ := strconv.Atoi(s1)
			n2, _ := strconv.Atoi(s2)
			newLine = append(newLine, strconv.Itoa(n1), strconv.Itoa(n2))
		} else {
			n, _ := strconv.Atoi(stone)
			n *= 2024
			newLine = append(newLine, strconv.Itoa(n))
		}
	}
	return newLine
}

func puzzle1() int {
	line := readInput()

	for i := 0; i < 25; i++ {
		line = blink(line)
	}

	return len(line)
}

func puzzle2() int {
	return 0
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
