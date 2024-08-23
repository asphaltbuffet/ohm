package es

// E12 returns the E12 value at the given position. It is 1-indexed.
//
// Returns -1 if the position is out of range.
func E12(p int) float64 {
	values := []float64{1.0, 1.2, 1.5, 1.8, 2.2, 2.7, 3.3, 3.9, 4.7, 5.6, 6.8, 8.2}

	idx := p - 1

	if idx < 0 || idx >= len(values) {
		return -1
	}

	return values[idx]
}
