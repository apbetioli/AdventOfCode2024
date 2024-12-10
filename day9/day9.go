package main

import (
	"adventofcode/2024/utils"
	"fmt"
)

func readInput() string {
	lines := utils.ReadInputLines()
	return lines[0]
}

func getRepresentation(diskMap string) []int {
	size := 0
	for _, rune := range diskMap {
		n := int(rune - '0')
		size += n
	}

	var representation []int = make([]int, size)

	var fileId int
	var index int

	for i, rune := range diskMap {
		n := int(rune - '0')

		if i%2 == 0 { //file
			for j := 0; j < n; j++ {
				representation[index] = fileId
				index++
			}
			fileId++
		} else { //space
			for r := 0; r < n; r++ {
				representation[index] = -1
				index++
			}
		}
	}

	fmt.Println(representation)

	return representation
}

func compact(compacted []int) []int {

	var a int
	var b int = len(compacted) - 1

	for a < b {
		if compacted[a] != -1 {
			a++
		} else if compacted[b] == -1 {
			b--
		} else {
			compacted[a] = compacted[b]
			compacted[b] = -1
			a++
			b--
		}

	}

	return compacted
}

func calculateChecksum(compacted []int) int {
	var checksum int
	for i := 0; i < len(compacted); i++ {
		if compacted[i] != -1 {
			checksum += i * compacted[i]
		}
	}
	return checksum
}

func puzzle1() int {
	diskMap := readInput()
	representation := getRepresentation(diskMap)
	compacted := compact(representation)
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
