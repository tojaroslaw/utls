package utls

import "golang.org/x/exp/constraints"

// ToPtr takes any x and returns a pointer to that item
func ToPtr[T any](x T) *T {
	return &x
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
