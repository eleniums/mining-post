package game

// Performs a deep copy on an array and returns the copy.
func DeepCopy[T any](src []T) []T {
	dest := make([]T, len(src))
	copy(dest, src)
	return dest
}
