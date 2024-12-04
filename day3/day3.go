package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func readInput() []string {
	return utils.ReadInputLines()
}

func puzzle1() int {
	lines := readInput()
	sum := 0
	for _, line := range lines {
		r, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
		instructions := r.FindAllString(line, -1)
		for i := 0; i < len(instructions); i++ {
			instruction := instructions[i]
			instruction = strings.Replace(instruction, "mul(", "", 1)
			instruction = strings.Replace(instruction, ")", "", 1)
			digits := strings.Split(instruction, ",")
			a, _ := strconv.Atoi(digits[0])
			b, _ := strconv.Atoi(digits[1])
			mul := a * b
			sum += mul
		}

	}
	return sum
}

func puzzle2() int {
	lines := readInput()
	sum := 0
	disabled := false
	for _, line := range lines {
		r, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)|do\(\)|don\'t\(\)`)
		instructions := r.FindAllString(line, -1)
		for i := 0; i < len(instructions); i++ {
			instruction := instructions[i]
			if instruction == "don't()" {
				disabled = true
			} else if instruction == "do()" {
				disabled = false
			} else if !disabled {
				instruction = strings.Replace(instruction, "mul(", "", 1)
				instruction = strings.Replace(instruction, ")", "", 1)
				digits := strings.Split(instruction, ",")
				a, _ := strconv.Atoi(digits[0])
				b, _ := strconv.Atoi(digits[1])
				mul := a * b
				sum += mul
			}
		}

	}
	return sum
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
