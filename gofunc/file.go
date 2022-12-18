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
	file := File(name)
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

// Returns a slice of Bytes read from a file.
func ReadMatrixByte(name string) [][]byte {
	var input [][]byte
	var slice []byte
	file := File(name)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		for _, v := range row {
			slice = append(slice, v[0])
		}
		input = append(input, slice)
		slice = make([]byte, 0)
	}
	return input
}
