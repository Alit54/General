package gofunc

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Returns the file already opened.
func File(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		panic("File not found")
	}
	return file
}

// Returns a slice with the entire file read.
func ReadRows(name string) []string {
	var input []string
	file := File(name)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

// Returns a slice of Integers read from a file.
func ReadMatrixInt(name string) [][]int {
	var input [][]int
	var sliceInt []int
	file := Alit.File(name)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		for _, v := range row {
			value, _ := strconv.Atoi(v)
			sliceInt = append(sliceInt, value)
		}
		input = append(input, sliceInt)
		sliceInt = make([]int, 0)
	}
	return input
}
