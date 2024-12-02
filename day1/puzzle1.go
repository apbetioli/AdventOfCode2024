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
	lines := utils.ReadLines("input.txt")

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

func main() {
	list1, list2 := readInput()
	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0

	for i := 0; i < len(list1); i++ {
		diff := int(math.Abs(float64(list2[i] - list1[i])))
		sum += diff
	}

	fmt.Print(sum)
}
