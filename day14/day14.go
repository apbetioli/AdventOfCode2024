package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"strconv"
	"strings"
)

type Robot struct {
	px int
	py int
	vx int
	vy int
}

func (r *Robot) update(seconds int, width int, height int) {

	for range seconds {
		r.px += r.vx
		r.px = r.px % width
		if r.px < 0 {
			r.px = width + r.px
		}

		r.py += r.vy
		r.py = r.py % height
		if r.py < 0 {
			r.py = height + r.py
		}
	}
}

func readInput() []Robot {
	var robots []Robot
	lines := utils.ReadInputLines()
	for _, line := range lines {
		eqs := strings.Split(line, " ")
		p := strings.Split(eqs[0], "=")
		p = strings.Split(p[1], ",")
		px, _ := strconv.Atoi(p[0])
		py, _ := strconv.Atoi(p[1])
		v := strings.Split(eqs[1], "=")
		v = strings.Split(v[1], ",")
		vx, _ := strconv.Atoi(v[0])
		vy, _ := strconv.Atoi(v[1])
		robots = append(robots, Robot{px, py, vx, vy})
	}
	return robots
}

func getMatrix(robots []Robot, width int, height int) [][]int {
	var matrix = make([][]int, height)
	for i := range matrix {
		matrix[i] = make([]int, width)
	}

	for _, robot := range robots {
		matrix[robot.py][robot.px] += 1
	}

	// utils.Debug(matrix)
	return matrix
}

func calculateSafetyFactor(robots []Robot, width int, height int) int {
	my := height / 2
	mx := width / 2

	var q1, q2, q3, q4 int

	for _, robot := range robots {
		if robot.px < mx && robot.py < my {
			q1++
		} else if robot.px > mx && robot.py < my {
			q2++
		} else if robot.px < mx && robot.py > my {
			q3++
		} else if robot.px > mx && robot.py > my {
			q4++
		}
	}

	return q1 * q2 * q3 * q4
}

func isChristmasTree(matrix [][]int) bool {
	height := len(matrix)
	width := len(matrix[0])
	mx := width / 2

	for y := 0; y < height; y++ {
		for x := 0; x < mx; x++ {
			if matrix[y][x] != matrix[y][width-x-1] {
				return false
			}
		}
	}

	return true
}

func puzzle1() int {
	robots := readInput()
	width := 101
	height := 103

	for r := range robots {
		robots[r].update(100, width, height)
	}

	return calculateSafetyFactor(robots, width, height)
}

func puzzle2() int {
	robots := readInput()
	width := 101
	height := 103

	var matrix = getMatrix(robots, width, height)

	i := 0
	for !isChristmasTree(matrix) {
		for r := range robots {
			matrix[robots[r].py][robots[r].px] -= 1
			robots[r].update(1, width, height)
			matrix[robots[r].py][robots[r].px] += 1
		}

		i++
	}

	utils.Debug(matrix)

	return i
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
