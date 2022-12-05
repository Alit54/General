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

// Returns the slice reordered in reverse
func ReverseSlice[T comparable](slice []T) []T {
	slice2 := make([]T, 0)
	for i := len(slice) - 1; i >= 0; i++ {
		slice2 = append(slice2, slice[i])
	}
	return slice2
}
