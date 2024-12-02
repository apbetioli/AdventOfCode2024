package utils

import (
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadLines(filename string) []string {
	dat, err := os.ReadFile(filename)
	check(err)
	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	return lines
}
