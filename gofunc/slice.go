package gofunc

// Returns the slice of the input without the last element.
func RemoveLastElement[T comparable](slice []T) []T {
	return append(slice[:len(slice)-1])
}

// Returns the slice of the input without the first element.
func RemoveFirstElement[T comparable](slice []T) []T {
	return append(slice[1:])
}

// Returns the slice of the input without the indicated element. Panics if index is less than 0.
func RemoveElement[T comparable](slice []T, index int) []T {
	if index < 0 {
		panic("Index can't be less than 0")
	}
	support := slice[:index]
	return append(support, slice[index+1:]...)
}

// Returns the slice reordered in reverse.
func ReverseSlice[T comparable](slice []T) []T {
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
