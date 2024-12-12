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

func change(stone int) []int {
	if stone == 0 {
		return []int{1}
	} else if stoneStr := strconv.Itoa(stone); len(stoneStr)%2 == 0 {
		mid := len(stoneStr) / 2
		n1, _ := strconv.Atoi(stoneStr[:mid])
		n2, _ := strconv.Atoi(stoneStr[mid:])
		return []int{n1, n2}
	} else {
		return []int{stone * 2024}
	}
}

func cachedChange() func(int) []int {

	var cache = make(map[int][]int)

	return func(stone int) []int {
		if len(cache[stone]) == 0 {
			cache[stone] = change(stone)
		}
		return cache[stone]
	}
}

func puzzle1() int {
	defer duration(track("puzzle1"))

	line := readInput()

	blink := cachedChange()

	for i := 0; i < 25; i++ {
		var newLine []int
		for _, stone := range line {
			newLine = append(newLine, blink(stone)...)
		}
		line = newLine
	}

	return len(line)
}

func puzzle2() int {
	defer duration(track("puzzle1"))

	line := readInput()

	blink := cachedChange()

	for i := 0; i < 75; i++ {
		var newLine []int
		for _, stone := range line {
			newLine = append(newLine, blink(stone)...)
		}
		line = newLine
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
