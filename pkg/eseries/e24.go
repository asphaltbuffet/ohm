package es

// E24 returns the E24 value at the given position. It is 1-indexed.
//
// Returns -1 if the position is out of range.
func E24(p int) float64 {
	values := []float64{
		1.0, 1.1, 1.2, 1.3, 1.5, 1.6, 1.8, 2.0, 2.2, 2.4, 2.7, 3.0,
		3.3, 3.6, 3.9, 4.3, 4.7, 5.1, 5.6, 6.2, 6.8, 7.5, 8.2, 9.1,
	}

	idx := p - 1

	if idx < 0 || idx >= len(values) {
		return -1
	}

	return values[idx]
}
