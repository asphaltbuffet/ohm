package es

// E48 returns the E48 value at the given position. It is 1-indexed.
//
// Returns -1 if the position is out of range.
func E48(p int) float64 {
	values := []float64{
		1.00, 1.05, 1.10, 1.15, 1.21, 1.27, 1.33, 1.40, 1.47, 1.54, 1.62, 1.69,
		1.78, 1.87, 1.96, 2.05, 2.15, 2.26, 2.37, 2.49, 2.61, 2.74, 2.87, 3.01,
		3.16, 3.32, 3.48, 3.65, 3.83, 4.02, 4.22, 4.42, 4.64, 4.87, 5.11, 5.36,
		5.62, 5.90, 6.19, 6.49, 6.81, 7.15, 7.50, 7.87, 8.25, 8.66, 9.09, 9.53,
	}

	idx := p - 1

	if idx < 0 || idx >= len(values) {
		return -1
	}

	return values[idx]
}
