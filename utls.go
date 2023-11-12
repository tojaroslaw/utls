package utls

// ToPtr takes any item and returns a pointer to that item
func ToPtr[T any](item T) *T {
	return &item
}
