package main

import (
	"adventofcode/2024/utils"
	"fmt"
)

type mem struct {
	position int
	size     int
}

func readInput() ([]mem, []mem, int) {
	lines := utils.ReadInputLines()
	diskMap := lines[0]

	var files []mem
	var spaces []mem

	var index int

	for i, rune := range diskMap {
		n := int(rune - '0')

		if i%2 == 0 {
			files = append(files, mem{index, n})
		} else {
			spaces = append(spaces, mem{index, n})
		}

		index += n
	}

	return files, spaces, index
}

func getRepresentation(files []mem, size int) []int {

	var representation []int = make([]int, size)

	for i := 0; i < size; i++ {
		representation[i] = -1
	}

	for id, info := range files {
		for j := 0; j < info.size; j++ {
			representation[info.position+j] = id
		}
	}

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
	files, _, size := readInput()
	representation := getRepresentation(files, size)
	compacted := compact(representation)
	return calculateChecksum(compacted)
}

func defrag(files []mem, spaces []mem) []mem {

	for f := len(files) - 1; f > 0; f-- {
		for s := 0; s < len(spaces) && spaces[s].position < files[f].position; s++ {
			if spaces[s].size >= files[f].size {
				files[f].position = spaces[s].position
				spaces[s].position += files[f].size
				spaces[s].size -= files[f].size
			}
		}
	}

	return files
}

func puzzle2() int {
	files, spaces, size := readInput()
	files = defrag(files, spaces)
	representation := getRepresentation(files, size)
	return calculateChecksum(representation)
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
