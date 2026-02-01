package app

// inBounds returns true if idx is a valid index for a slice of the given length.
func inBounds(idx, length int) bool {
	return idx >= 0 && idx < length
}

// clampIndex normalizes an index after a slice modification (e.g., deletion).
// It ensures the index stays within valid bounds, returning 0 for empty slices
// and clamping to the last valid index otherwise.
func clampIndex(idx, length int) int {
	if length == 0 {
		return 0
	}
	if idx >= length {
		return length - 1
	}
	if idx < 0 {
		return 0
	}
	return idx
}
