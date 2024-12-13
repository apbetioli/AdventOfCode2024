package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type machine struct {
	a [2]int
	b [2]int
	p [2]int
}

func readButtonLine(line string) [2]int {
	button := strings.Split(line, ":")
	eqs := strings.Split(button[1], ",")
	x := strings.Split(eqs[0], "+")
	y := strings.Split(eqs[1], "+")
	xi, _ := strconv.Atoi(x[1])
	yi, _ := strconv.Atoi(y[1])
	return [2]int{xi, yi}
}

func readPrizeLine(line string) [2]int {
	prize := strings.Split(line, ":")
	eqs := strings.Split(prize[1], ",")
	x := strings.Split(eqs[0], "=")
	y := strings.Split(eqs[1], "=")
	xi, _ := strconv.Atoi(x[1])
	yi, _ := strconv.Atoi(y[1])
	return [2]int{xi, yi}
}

func readInput() []machine {
	lines := utils.ReadInputLines()
	var machines []machine
	for i := 0; i < len(lines); i += 4 {
		claw := machine{}

		claw.a = readButtonLine(lines[i])
		claw.b = readButtonLine(lines[i+1])
		claw.p = readPrizeLine(lines[i+2])

		machines = append(machines, claw)
	}
	return machines
}

func calculate(m machine) float64 {

	bestCost := math.Inf(1)

	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			if a+b == 0 {
				continue
			}
			remX := (a*m.a[0] + b*m.b[0]) % m.p[0]
			remY := (a*m.a[1] + b*m.b[1]) % m.p[1]
			if remX == 0 && remY == 0 {
				cost := 3*a + b
				if float64(cost) < bestCost {
					bestCost = float64(cost)
				}
			}
		}
	}

	if math.IsInf(bestCost, 1) {
		return 0
	}

	return bestCost
}

func cramers(m machine) int {

	det := m.b[1]*m.a[0] - m.b[0]*m.a[1]

	if det == 0 {
		// No solution or multiple solutions
		panic("No solution or multiple solutions")
	}

	detA := m.p[0]*m.b[1] - m.p[1]*m.b[0]
	detB := m.a[0]*m.p[1] - m.a[1]*m.p[0]

	a := detA / det
	b := detB / det

	cost := 3*a + b

	x := a*m.a[0] + b*m.b[0]
	y := a*m.a[1] + b*m.b[1]

	if x == m.p[0] && y == m.p[1] {
		return cost
	} else {
		return 0
	}
}

func puzzle1() int {
	machines := readInput()

	sum := 0

	for _, machine := range machines {
		cost := calculate(machine)
		sum += int(cost)
	}

	return sum
}

func puzzle2() int {
	machines := readInput()

	sum := 0

	for _, machine := range machines {
		machine.p[0] += 10000000000000
		machine.p[1] += 10000000000000

		cost := cramers(machine)
		sum += int(cost)
	}

	return sum
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
