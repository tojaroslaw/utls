package utls

import (
	"golang.org/x/exp/constraints"
	"slices"
)

// ToPtr takes any x and returns a pointer to that ptr
func ToPtr[T any](x T) *T {
	return &x
}

// ToVal takes any ptr and returns the value of that pointer if it exists and sets ok to true. If the pointer is nil,
// it returns the zero value of the type and sets ok to false.
func ToVal[T any](ptr *T) (val T, ok bool) {
	if ptr == nil {
		return val, false
	}
	return *ptr, true
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

// SliceToMap takes in a slice and returns a map with the slice values as keys and each corresponding value set to true.
// If the slice contains duplicate values, the map will only contain one each duplicated value.
func SliceToMap[S ~[]T, T comparable](slice S) map[T]bool {
	m := make(map[T]bool)
	for _, item := range slice {
		m[item] = true
	}
	return m
}

// Min returns the minimum of two values.
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two values.
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
