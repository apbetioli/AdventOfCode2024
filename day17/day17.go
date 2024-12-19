package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Computer struct {
	a                  int
	b                  int
	c                  int
	program            []string
	output             []string
	instructionPointer int
}

func (comp *Computer) combo(operand int) int {
	if operand <= 3 {
		return operand
	} else if operand == 4 {
		return comp.a
	} else if operand == 5 {
		return comp.b
	} else if operand == 6 {
		return comp.c
	} else {
		panic("Not implemented")
	}
}

func (comp *Computer) adv(operand int) {
	// fmt.Println("adv", operand)
	comp.a = comp.a / int(math.Pow(2.0, float64(comp.combo(operand))))
	comp.instructionPointer += 2
}
func (comp *Computer) bxl(operand int) {
	// fmt.Println("bxl", operand)
	comp.b = comp.b ^ operand
	comp.instructionPointer += 2
}
func (comp *Computer) bst(operand int) {
	// fmt.Println("bst", operand)
	comp.b = comp.combo(operand) % 8
	comp.instructionPointer += 2
}
func (comp *Computer) jnz(operand int) {
	// fmt.Println("jnz", operand)
	if comp.a == 0 {
		comp.instructionPointer += 2
	} else {
		comp.instructionPointer = operand
	}
}
func (comp *Computer) bxc(_ int) {
	// fmt.Println("bxc")
	comp.b = comp.b ^ comp.c
	comp.instructionPointer += 2
}
func (comp *Computer) out(operand int) {
	// fmt.Println("out", operand)
	out := comp.combo(operand) % 8
	comp.output = append(comp.output, strconv.Itoa(out))
	comp.instructionPointer += 2
}
func (comp *Computer) bdv(operand int) {
	// fmt.Println("bdv", operand)
	comp.b = comp.a / int(math.Pow(2.0, float64(comp.combo(operand))))
	comp.instructionPointer += 2
}
func (comp *Computer) cdv(operand int) {
	// fmt.Println("cdv", operand)
	comp.c = comp.a / int(math.Pow(2.0, float64(comp.combo(operand))))
	comp.instructionPointer += 2
}
func (comp *Computer) hasNext() bool {
	return comp.instructionPointer < len(comp.program)-1
}

func (comp *Computer) next() {
	opcode, _ := strconv.Atoi(comp.program[comp.instructionPointer])
	operand, _ := strconv.Atoi(comp.program[comp.instructionPointer+1])

	switch opcode {
	case 0:
		comp.adv(operand)
	case 1:
		comp.bxl(operand)
	case 2:
		comp.bst(operand)
	case 3:
		comp.jnz(operand)
	case 4:
		comp.bxc(operand)
	case 5:
		comp.out(operand)
	case 6:
		comp.bdv(operand)
	case 7:
		comp.cdv(operand)
	default:
		panic("Not implemented")
	}

	// fmt.Println("Computer", comp)
	// fmt.Println("instructionPointer", instructionPointer)
	// fmt.Println("opcode", opcode, "operand", operand, "state", comp)
	// fmt.Println("operand", operand)
	// fmt.Scanln()
}

func readInput() Computer {
	lines := utils.ReadInputLines()
	a, _ := strconv.Atoi(strings.TrimSpace(strings.Split(lines[0], ":")[1]))
	b, _ := strconv.Atoi(strings.TrimSpace(strings.Split(lines[1], ":")[1]))
	c, _ := strconv.Atoi(strings.TrimSpace(strings.Split(lines[2], ":")[1]))
	program := strings.Split(strings.TrimSpace(strings.Split(lines[4], ":")[1]), ",")
	return Computer{a, b, c, program, []string{}, 0}
}

func puzzle1() string {
	comp := readInput()
	// fmt.Println(comp)

	for comp.hasNext() {
		comp.next()
	}

	return strings.Join(comp.output, ",")
}

func puzzle2() int {

	// fmt.Println(-10 % 8)
	return 0
}

func main() {
	fmt.Println(puzzle1())
	fmt.Println(puzzle2())
}
