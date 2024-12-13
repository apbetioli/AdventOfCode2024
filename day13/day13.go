package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type machine struct {
	a     [2]int
	b     [2]int
	prize [2]int
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
		claw.prize = readPrizeLine(lines[i+2])

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
			remX := (a*m.a[0] + b*m.b[0]) % m.prize[0]
			remY := (a*m.a[1] + b*m.b[1]) % m.prize[1]
			if remX == 0 && remY == 0 {
				cost := 3*a + b
				if float64(cost) < bestCost {
					bestCost = float64(cost)
				}
			}
		}
	}

	return bestCost
}

func puzzle1() int {
	machines := readInput()

	sum := 0

	for _, machine := range machines {
		cost := calculate(machine)
		if !math.IsInf(cost, 1) {
			sum += int(cost)
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
