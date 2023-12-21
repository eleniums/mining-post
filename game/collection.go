package game

// Copies the values of a map into a slice, ignoring the key.
func MapValues[K comparable, V any](m map[K]V) []V {
	flattened := []V{}
	for _, v := range m {
		flattened = append(flattened, v)
	}
	return flattened
}

// Finds a value in a slice and if found, returns it.
func Find[T any](s []T, val T, compare func(a, b T) bool) (T, bool) {
	for _, v := range s {
		if compare(v, val) {
			return v, true
		}
	}
	var zeroValue T
	return zeroValue, false
}

// Filters the slice to only contain values that return true from the filter function.
func Filter[T any](s []T, filter func(val T) bool) []T {
	filtered := []T{}
	for _, v := range s {
		if filter(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
