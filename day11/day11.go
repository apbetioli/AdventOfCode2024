package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func readInput() []int {
	lines := utils.ReadInputLines()
	return utils.StringArrayToIntArray(strings.Split(lines[0], " "))
}

func blink(line []int) []int {
	var newLine []int
	for _, stone := range line {
		if stone == 0 {
			newLine = append(newLine, 1)
		} else if stoneStr := strconv.Itoa(stone); len(stoneStr)%2 == 0 {
			mid := len(stoneStr) / 2
			n1, _ := strconv.Atoi(stoneStr[:mid])
			n2, _ := strconv.Atoi(stoneStr[mid:])
			newLine = append(newLine, n1, n2)
		} else {
			newLine = append(newLine, stone*2024)
		}
	}
	return newLine
}

func puzzle1() int {
	defer duration(track("puzzle1"))

	line := readInput()

	for i := 0; i < 25; i++ {
		line = blink(line)
	}

	return len(line)
}

func puzzle2() int {
	defer duration(track("puzzle1"))

	line := readInput()

	for i := 0; i < 75; i++ {
		line = blink(line)
	}

	return len(line)
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
