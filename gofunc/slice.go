package gofunc

// Returns the slice of the input without the last element.
func RemoveLastElement[T any](slice []T) []T {
	return append(slice[:len(slice)-1])
}

// Returns the slice of the input without the first element.
func RemoveFirstElement[T any](slice []T) []T {
	return append(slice[1:])
}

// Returns the slice of the input without the indicated element. Panics if index is less than 0.
func RemoveElement[T any](slice []T, index int) []T {
	if index < 0 {
		panic("Index can't be less than 0")
	}
	support := slice[:index]
	return append(support, slice[index+1:]...)
}

// Returns the slice reordered in reverse.
func ReverseSlice[T any](slice []T) []T {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
	return slice
}

// Returns the position of element in the slice. If there is no element, returns an empty slice.
func ResearchElement[T comparable](slice []T, element T) []int {
	support := make([]int, 0)
	for k, v := range slice {
		if v == element {
			support = append(support, k)
		}
	}
	return support
}

// Sorts the slice using InsertionSort.
func InsertionSort[T Ordered](slice []T) {
	for k := 1; k < len(slice); k++ {
		x := slice[k]
		var j int
		for j = k - 1; j >= 0 && slice[j] > x; j-- {
			slice[j+1] = slice[j]
		}
		slice[j+1] = x
	}
}

// Sorts the slice using SelectionSort.
func SelectionSort[T Ordered](slice []T) {
	for k := 0; k < len(slice)-1; k++ {
		m := k
		for j := k + 1; j < len(slice); j++ {
			if slice[j] < slice[m] {
				m = j
			}
		}
		slice[m], slice[k] = slice[k], slice[m]
	}
}

// Sorts the slice using MergeSort.
func MergeSort[T Ordered](slice []T) {
	leng := len(slice) / 2
	out := make([]T, 0)
	// Base Case
	if len(slice) < 2 {
		return
	}
	// Recursive
	MergeSort(slice[:leng])
	MergeSort(slice[leng:])
	// Merge
	i, j := 0, 0
	for i < len(slice[:leng]) && j < len(slice[leng:]) {
		if slice[:leng][i] <= slice[leng:][j] {
			out = append(out, slice[:leng][i])
			i++
		} else {
			out = append(out, slice[leng:][j])
			j++
		}
	}
	if i < len(slice[:leng]) {
		for _, v := range slice[:leng][i:] {
			out = append(out, v)
		}
	} else {
		for _, v := range slice[leng:][j:] {
			out = append(out, v)
		}
	}
	slice = out
}
