package game

import (
	"log/slog"
	"sync"
)

// Performs a deep copy on an array and returns the copy.
func DeepCopy[T any](src []T) []T {
	dest := make([]T, len(src))
	copy(dest, src)
	return dest
}

// Helper function to get a value from a sync.Map of the correct type.
func MapLoad[K any, V any](m *sync.Map, key K) (V, bool) {
	val, ok := m.Load(key)
	if !ok {
		slog.Error("value not found in sync.Map")
	}
	result, ok := val.(V)
	return result, ok
}
