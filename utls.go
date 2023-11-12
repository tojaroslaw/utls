package utls

import (
	"golang.org/x/exp/constraints"
	"slices"
)

// ToPtr takes any x and returns a pointer to that item
func ToPtr[T any](x T) *T {
	return &x
}

// SliceContains takes in a slice and an item to check for in the slice. If the item is in the slice, it returns true,
// otherwise it returns false.
func SliceContains[S ~[]T, T comparable](slice S, item T) bool {
	return slices.Contains(slice, item)
}

// MapContains takes in a map and an item to check for as a key in that map. If the item is a key in the map, it returns
// true; otherwise it returns false.
func MapContains[T comparable, U any](m map[T]U, item T) bool {
	_, ok := m[item]
	return ok
}

// Min returns the minimum of two values
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two values
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
