package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func readInput() ([][]string, []string) {
	lines := utils.ReadInputLines()

	var matrix [][]string
	i := 0
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}
		matrix = append(matrix, strings.Split(lines[i], ""))
	}

	var movements []string

	for ; i < len(lines); i++ {
		movements = append(movements, strings.Split(lines[i], "")...)
	}

	return matrix, movements
}

func getRobot(matrix [][]string) [2]int {
	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] == "@" {
				return [2]int{y, x}
			}
		}
	}
	panic("Robot not found")
}

func getBoxes(matrix [][]string) [][2]int {
	var boxes [][2]int
	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] == "O" || matrix[y][x] == "[" {
				boxes = append(boxes, [2]int{y, x})
			}
		}
	}
	return boxes
}

func puzzle1() int {
	matrix, movements := readInput()

	// utils.Debug(matrix)

	robot := getRobot(matrix)

	for _, move := range movements {
		y := robot[0]
		x := robot[1]

		if move == ">" {
			if moveRight(matrix, y, x+1) {
				swap(matrix, y, x+1, y, x)
				robot[1] = x + 1
			}
		} else if move == "<" {
			if moveLeft(matrix, y, x-1) {
				swap(matrix, y, x-1, y, x)
				robot[1] = x - 1
			}
		} else if move == "^" {
			if moveUp(matrix, y-1, x) {
				swap(matrix, y-1, x, y, x)
				robot[0] = y - 1
			}
		} else if move == "v" {
			if moveDown(matrix, y+1, x) {
				swap(matrix, y+1, x, y, x)
				robot[0] = y + 1
			}
		}

		// utils.Debug(matrix)
	}

	return calculateGPS(matrix)
}

func calculateGPS(matrix [][]string) int {
	boxes := getBoxes(matrix)

	sum := 0
	for _, box := range boxes {
		gps := box[0]*100 + box[1]
		sum += gps
	}

	return sum
}

func resample(matrix [][]string) [][]string {
	for y := range matrix {
		var newLine []string
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] == "#" {
				newLine = append(newLine, "#", "#")
			} else if matrix[y][x] == "O" {
				newLine = append(newLine, "[", "]")
			} else if matrix[y][x] == "@" {
				newLine = append(newLine, "@", ".")
			} else if matrix[y][x] == "." {
				newLine = append(newLine, ".", ".")
			}
		}
		matrix[y] = newLine
	}
	return matrix
}

func swap(matrix [][]string, y1 int, x1 int, y2 int, x2 int) {
	temp := matrix[y1][x1]
	matrix[y1][x1] = matrix[y2][x2]
	matrix[y2][x2] = temp
}

func moveRight(matrix [][]string, y int, x int) bool {
	if matrix[y][x] == "." {
		return true
	}

	if matrix[y][x] == "[" || matrix[y][x] == "]" || matrix[y][x] == "O" || matrix[y][x] == "@" {
		ny := y
		nx := x + 1
		if moveRight(matrix, ny, nx) {
			swap(matrix, ny, nx, y, x)
			return true
		}
	}

	return false
}

func moveLeft(matrix [][]string, y int, x int) bool {
	if matrix[y][x] == "." {
		return true
	}

	if matrix[y][x] == "[" || matrix[y][x] == "]" || matrix[y][x] == "O" || matrix[y][x] == "@" {
		ny := y
		nx := x - 1
		if moveLeft(matrix, ny, nx) {
			swap(matrix, ny, nx, y, x)
			return true
		}
	}

	return false
}

func canMoveUp(matrix [][]string, y int, x int) bool {
	if matrix[y][x] == "." {
		return true
	}

	ny := y - 1
	nx := x

	if matrix[y][x] == "O" || matrix[y][x] == "@" {
		if canMoveUp(matrix, ny, nx) {
			return true
		}
	} else if matrix[y][x] == "[" {
		if canMoveUp(matrix, ny, nx) && canMoveUp(matrix, ny, nx+1) {
			return true
		}
	} else if matrix[y][x] == "]" {
		if canMoveUp(matrix, ny, nx) && canMoveUp(matrix, ny, nx-1) {
			return true
		}
	}

	return false
}

func moveUp(matrix [][]string, y int, x int) bool {
	if !canMoveUp(matrix, y, x) {
		return false
	}

	if matrix[y][x] == "." {
		return true
	}

	ny := y - 1
	nx := x

	if matrix[y][x] == "O" || matrix[y][x] == "@" {
		if moveUp(matrix, ny, nx) {
			swap(matrix, ny, nx, y, x)
			return true
		}
	} else if matrix[y][x] == "[" {
		if moveUp(matrix, ny, nx) && moveUp(matrix, ny, nx+1) {
			swap(matrix, ny, nx, y, x)
			swap(matrix, ny, nx+1, y, x+1)
			return true
		}
	} else if matrix[y][x] == "]" {
		if moveUp(matrix, ny, nx) && moveUp(matrix, ny, nx-1) {
			swap(matrix, ny, nx, y, x)
			swap(matrix, ny, nx-1, y, x-1)
			return true
		}
	}

	return false
}

func canMoveDown(matrix [][]string, y int, x int) bool {
	if matrix[y][x] == "." {
		return true
	}

	ny := y + 1
	nx := x

	if matrix[y][x] == "O" || matrix[y][x] == "@" {
		if canMoveDown(matrix, ny, nx) {
			return true
		}
	} else if matrix[y][x] == "[" {
		if canMoveDown(matrix, ny, nx) && canMoveDown(matrix, ny, nx+1) {
			return true
		}
	} else if matrix[y][x] == "]" {
		if canMoveDown(matrix, ny, nx) && canMoveDown(matrix, ny, nx-1) {
			return true
		}
	}

	return false
}

func moveDown(matrix [][]string, y int, x int) bool {
	if !canMoveDown(matrix, y, x) {
		return false
	}

	if matrix[y][x] == "." {
		return true
	}

	ny := y + 1
	nx := x

	if matrix[y][x] == "O" || matrix[y][x] == "@" {
		if moveDown(matrix, ny, nx) {
			swap(matrix, ny, nx, y, x)
			return true
		}
	} else if matrix[y][x] == "[" {
		if moveDown(matrix, ny, nx) && moveDown(matrix, ny, nx+1) {
			swap(matrix, ny, nx, y, x)
			swap(matrix, ny, nx+1, y, x+1)
			return true
		}
	} else if matrix[y][x] == "]" {
		if moveDown(matrix, ny, nx) && moveDown(matrix, ny, nx-1) {
			swap(matrix, ny, nx, y, x)
			swap(matrix, ny, nx-1, y, x-1)
			return true
		}
	}

	return false
}

func puzzle2() int {
	matrix, movements := readInput()

	matrix = resample(matrix)

	// utils.Debug(matrix)

	robot := getRobot(matrix)

	for _, move := range movements {
		y := robot[0]
		x := robot[1]

		if move == ">" {
			if moveRight(matrix, y, x+1) {
				swap(matrix, y, x+1, y, x)
				robot[1] = x + 1
			}
		} else if move == "<" {
			if moveLeft(matrix, y, x-1) {
				swap(matrix, y, x-1, y, x)
				robot[1] = x - 1
			}
		} else if move == "^" {
			if moveUp(matrix, y-1, x) {
				swap(matrix, y-1, x, y, x)
				robot[0] = y - 1
			}
		} else if move == "v" {
			if moveDown(matrix, y+1, x) {
				swap(matrix, y+1, x, y, x)
				robot[0] = y + 1
			}
		}

		// utils.Debug(matrix)
	}

	return calculateGPS(matrix)
}

func interactive() {

	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	matrix, _ := readInput()

	matrix = resample(matrix)

	utils.Debug(matrix)

	robot := getRobot(matrix)

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		move := string(b)

		y := robot[0]
		x := robot[1]

		if move == "d" {
			if moveRight(matrix, y, x+1) {
				swap(matrix, y, x+1, y, x)
				robot[1] = x + 1
			}
		} else if move == "a" {
			if moveLeft(matrix, y, x-1) {
				swap(matrix, y, x-1, y, x)
				robot[1] = x - 1
			}
		} else if move == "w" {
			if moveUp(matrix, y-1, x) {
				swap(matrix, y-1, x, y, x)
				robot[0] = y - 1
			}
		} else if move == "s" {
			if moveDown(matrix, y+1, x) {
				swap(matrix, y+1, x, y, x)
				robot[0] = y + 1
			}
		}

		utils.Debug(matrix)
	}
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())

	// interactive()
}
