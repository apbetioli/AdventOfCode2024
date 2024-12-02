package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func readInput() ([]int, []int) {
	lines := utils.ReadInputLines()

	var list1 []int
	var list2 []int

	for _, line := range lines {
		parts := strings.Fields(line)
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	return list1, list2
}

func puzzle1() int {
	list1, list2 := readInput()
	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0

	for i := 0; i < len(list1); i++ {
		diff := int(math.Abs(float64(list2[i] - list1[i])))
		sum += diff
	}

	return sum
}

func puzzle2() int {
	list1, list2 := readInput()

	count := make(map[int]int)

	for i := 0; i < len(list2); i++ {
		key := list2[i]
		count[key] = count[key] + 1
	}

	score := 0

	for i := 0; i < len(list1); i++ {
		key := list1[i]
		score += key * count[key]

	}

	return score
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
