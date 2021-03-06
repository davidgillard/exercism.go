package cards

// test
// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if index < 0 || index >= len(slice) {
		return 0, false
	}
	return slice[index], true
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	switch {
	case index < 0 || index >= len(slice):
		slice = append(slice, value)

	default:
		slice[index] = value
	}
	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	switch {
	case length <= 0:
		return []int{}
	default:
		slice := make([]int, length)
		for i := range slice {
			slice[i] = value
		}
		return slice
	}
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	// Delete element at index index
	switch {
	case index > len(slice)-1:
		return slice
	case ((index >= len(slice)) || (index < 0)):
		return slice

		// The ... operator auto expands slice[i+1:] so that
		// its elements can be appended to slice[:i] one by one
	default:
		return append(slice[:index], slice[index+1:]...)
	}
}
