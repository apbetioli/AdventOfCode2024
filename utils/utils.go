package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
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
	time.Sleep(100 * time.Millisecond)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	buf := new(bytes.Buffer)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			buf.WriteString(matrix[i][j])
		}
		buf.WriteString("\n")
	}
	fmt.Println(buf)
}
func DebugInt(matrix [][]int) {
	time.Sleep(100 * time.Millisecond)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	// buf := new(bytes.Buffer)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j], " ")
			// buf.WriteString(strconv.Itoa(matrix[i][j]))
		}
		fmt.Println()
		// buf.WriteString("\n")
	}
	fmt.Println()
	// fmt.Println(buf)
}

func DebugNoClear(matrix [][]string) {
	buf := new(bytes.Buffer)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			buf.WriteString(matrix[i][j])
		}
		buf.WriteString("\n")
	}
	fmt.Println(buf)
}

func IsValidCoordinate[T any](matrix [][]T, y int, x int) bool {
	return y >= 0 && y < len(matrix) && x >= 0 && x < len(matrix[0])
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
