package es

// E3 returns the E3 value at the given position. It is 1-indexed.
//
// Returns -1 if the position is out of range.
func E3(p int) float64 {
	values := []float64{1.0, 2.2, 4.7}

	idx := p - 1

	if idx < 0 || idx >= len(values) {
		return -1
	}

	return values[idx]
}
