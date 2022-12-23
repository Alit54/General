package gofunc

import "math"

// Returns maximun value of a list of integers. If no element is insert, it returns minimun value available.
func Max(values ...int) int {
	if len(values) == 0 {
		return math.MinInt
	}
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

// Returns minimun value of a list of integers. If no element is insert, it returns maximun value available.
func Min(values ...int) int {
	if len(values) == 0 {
		return math.MaxInt
	}
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

// Returns sum of values.
func Sum[T Number](values ...T) T {
	var sum T = 0
	for _, v := range values {
		sum += v
	}
	return sum
}

// Returns product of values.
func Product[T Number](values ...T) T {
	var product T = 1
	for _, v := range values {
		product *= v
	}
	return product
}

// Returns the Absolute Value of a number.
func Absolute[T Number](value T) T {
	if value < 0 {
		return -value
	}
	return value
}

// Returns the Square Root of a positive number.
func Sqrt[T Number](value T) float64 {
	if value < 0 {
		panic("Number must be positive.")
	}
	return math.Sqrt(float64(value))
}

// Returns TRUE if all values are TRUE. FALSE otherwise.
func AND(values ...bool) bool {
	for _, v := range values {
		if !v {
			return false
		}
	}
	return true
}

// Returns FALSE if all values are FALSE. TRUE otherwise.
func OR(values ...bool) bool {
	for _, v := range values {
		if v {
			return true
		}
	}
	return false
}
