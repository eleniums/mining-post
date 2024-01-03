package game

// Copies the values of a map into a slice, ignoring the key.
func MapValues[K comparable, V any](m map[K]V) []V {
	flattened := []V{}
	for _, v := range m {
		flattened = append(flattened, v)
	}
	return flattened
}

// Merges all given maps into a single new map. In the case of duplicate keys, the last value merged will overwrite any previous value.
func MapMerge[K comparable, V any](maps ...map[K]V) map[K]V {
	merged := map[K]V{}
	for _, m := range maps {
		for k, v := range m {
			merged[k] = v
		}
	}
	return merged
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

// Creates a shallow copy of the given object.
func Copy[T any](src *T) *T {
	new := *src
	return &new
}

// Creates a shallow copy of the objects in a slice.
func CopySlice[T any](src []*T) []*T {
	new := make([]*T, len(src))
	for i, v := range src {
		new[i] = Copy(v)
	}
	return new
}
