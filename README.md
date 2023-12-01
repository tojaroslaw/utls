# utls
The utls package is a package filled with functions that ought to be in golang, but aren't for reasons unknown.
Sometimes these can be found scattered in the exp package but since I often end up re-implementing these anyway, I
just wanted to centralize everything in one place.

This package contains the following functions:
- ToPtr: converts a value to a pointer
- ToVal: converts a pointer to its value if possible and leaves an ok bool if the pointer is not nil
- SliceContains: checks if a slice contains a value
- MapContains: checks if a map contains a key
- SliceToMap: converts a slice to a map where each value is a key and the value is true
- Min: returns the minimum of two ordered values
- Max: returns the maximum of two ordered values
