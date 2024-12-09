package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func readInput() [][]int {
	var equations [][]int
	lines := utils.ReadInputLines()
	for _, line := range lines {
		equation := strings.Split(line, ":")
		result, _ := strconv.Atoi(equation[0])
		var operands []int
		operands = append(operands, result)
		operandsStr := strings.Split(strings.TrimSpace(equation[1]), " ")
		for _, opStr := range operandsStr {
			op, _ := strconv.Atoi(opStr)
			operands = append(operands, op)
		}
		equations = append(equations, operands)
	}
	return equations
}

func isValid(result int, operands []int) bool {

	if len(operands) == 1 {
		return result == operands[0]
	}

	sum := operands[0] + operands[1]
	sumOps := append([]int{sum}, operands[2:]...)

	mul := operands[0] * operands[1]
	mulOps := append([]int{mul}, operands[2:]...)

	return isValid(result, sumOps) || isValid(result, mulOps)
}

func puzzle1() int {
	equations := readInput()

	sum := 0

	for _, equation := range equations {
		if isValid(equation[0], equation[1:]) {
			sum += equation[0]
		}
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
