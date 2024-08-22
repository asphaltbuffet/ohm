package ohm

import "fmt"

type BandCodeLengthError struct {
	Length int
}

func (e BandCodeLengthError) Error() string {
	return fmt.Sprintf("invalid band code length: %d", e.Length)
}

type BandCodeColorError struct {
	ColorCode string
	BandType  string
}

func (e BandCodeColorError) Error() string {
	return fmt.Sprintf("%q band at %s: invalid band color", e.ColorCode, e.BandType)
}

type BandInvalidError struct {
	Bands []Band
}

func (e BandInvalidError) Error() string {
	return fmt.Sprintf("invalid band code: %v", e.Bands)
}
