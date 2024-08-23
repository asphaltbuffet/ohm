package es

// E6 returns the E6 value at the given position. It is 1-indexed.
//
// Returns -1 if the position is out of range.
func E6(p int) float64 {
	values := []float64{1.0, 1.5, 2.2, 3.3, 4.7, 6.8}

	idx := p - 1

	if idx < 0 || idx >= len(values) {
		return -1
	}

	return values[idx]
}
