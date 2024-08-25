package resistor

import "errors"

var (
	ErrBandCodeLength = errors.New("invalid length")
	ErrBandColor      = errors.New("unknown color")
	ErrBandPosition   = errors.New("invalid color at position")
)
