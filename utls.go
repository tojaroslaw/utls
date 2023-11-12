package utls

// ToPtr takes any item and returns a pointer to that item
func ToPtr[T any](item T) *T {
	return &item
}

// MapContains takes in a map and an item to check for as a key in that map. If the item is a key in the map, it returns
// true; otherwise it returns false.
func MapContains[T comparable, U any](m map[T]U, item T) bool {
	_, ok := m[item]
	return ok
}
