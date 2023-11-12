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

// SliceContains takes in a slice and an ptr to check for in the slice. If the ptr is in the slice, it returns true,
// otherwise it returns false.
func SliceContains[S ~[]T, T comparable](slice S, item T) bool {
	return slices.Contains(slice, item)
}

// MapContains takes in a map and an ptr to check for as a key in that map. If the ptr is a key in the map, it returns
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
