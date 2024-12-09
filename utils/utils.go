package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInputLines() []string {
	filename := "input.txt"
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 {
		filename = argsWithoutProg[0]
	}

	dat, err := os.ReadFile(filename)
	check(err)
	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	return lines
}

func StringArrayToIntArray(strings []string) []int {
	var ints []int

	for _, item := range strings {
		value, _ := strconv.Atoi(item)
		ints = append(ints, value)
	}

	return ints
}

func Contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func Debug(matrix [][]string) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}
