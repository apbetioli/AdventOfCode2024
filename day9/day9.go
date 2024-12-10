package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func readInput() []string {
	lines := utils.ReadInputLines()
	return strings.Split(lines[0], "")
}

func getRepresentation(diskMap []string) []string {
	var representation []string

	var f int
	for i := 0; i < len(diskMap); i++ {
		n, _ := strconv.Atoi(diskMap[i])
		if i%2 == 0 { //file
			for r := 0; r < n; r++ {
				representation = append(representation, strconv.Itoa(f))
			}
			f++
		} else { //space
			for r := 0; r < n; r++ {
				representation = append(representation, ".")
			}
		}
	}

	return representation
}

func compact(representation []string) ([]string, int) {

	compacted := make([]string, len(representation))
	copy(compacted, representation)

	var a int
	var b int = len(compacted) - 1

	for a < b {
		if compacted[a] != "." {
			a++
		} else if compacted[b] == "." {
			b--
		} else {
			compacted[a] = compacted[b]
			compacted[b] = "."
			a++
			b--
		}

	}

	return compacted, a
}

func calculateChecksum(compacted []string) int {
	var checksum int
	for i := 0; i < len(compacted) && compacted[i] != "."; i++ {
		n, _ := strconv.Atoi(compacted[i])
		checksum += i * n
	}
	return checksum
}

func puzzle1() int {
	diskMap := readInput()
	representation := getRepresentation(diskMap)
	compacted, _ := compact(representation)
	checksum := calculateChecksum(compacted)

	return checksum
}

func puzzle2() int {
	return 0
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
