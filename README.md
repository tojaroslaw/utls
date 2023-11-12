# utls
The utls package is a package filled with functions that ought to be in golang, but aren't for reasons unknown.
Sometimes these can be found scattered in the exp package but since I often end up re-implementing these anyway, I
just wanted to centralize everything in one place.

This package contains the following functions:
- ToPtr: converts a value to a pointer
- SliceContains: checks if a slice contains a value
- MapContains: checks if a map contains a key
- Min: returns the minimum of two ordered values
- Max: returns the maximum of two ordered values
