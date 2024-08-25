package axial

import "fmt"

const (
	Axial1Band int = iota + 1 // 0 Ω resistors are single black band
	Axial3Band int = iota + 2
	Axial4Band
	Axial5Band
	Axial6Band
	// Axial7Band // 6-band with failure rate not implemented yet
)

// BandCode is a struct that represents a resistor 4, 5, or 6 band code.
type BandCode struct {
	Reversed bool   // Reversed is true if the band is reversed from input order.
	Bands    []Band // Bands is a slice of bands that make up the code.
}

// Validate checks if the band code is valid.
func (bc BandCode) Validate() error {
	return validateBandOrder(bc.Bands)
}

func validateBandOrder(b []Band) error {
	l := len(b)

	switch {
	case l == 1 && b[0].Code == "BK": // 0 Ω resistors are single black band
		// this is a valid 0 Ω resistor
		return nil
	case l < Axial3Band, l > Axial6Band:
		return fmt.Errorf("validate %v: %w", b, ErrBandCodeLength)

	case b[0].SigFig == Invalid:
		return fmt.Errorf("%q as digit: %w", b[0].Code, ErrBandPosition)

	case b[1].SigFig == Invalid:
		return fmt.Errorf("%q as digit: %w", b[1].Code, ErrBandPosition)

	// any color can be used for band 3 if it's a multiplier, but
	// 5/6 band resistors have a 3rd significant figure band
	case l >= Axial5Band && b[2].SigFig == Invalid:
		return fmt.Errorf("%q as digit: %w", b[2].Code, ErrBandPosition)

	case l == Axial4Band && b[3].Tolerance == Invalid:
		return fmt.Errorf("%q as tolerance: %w", b[3].Code, ErrBandPosition)

	case l >= Axial5Band && b[4].Tolerance == Invalid:
		return fmt.Errorf("%q as tolerance: %w", b[4].Code, ErrBandPosition)

	case l >= Axial6Band && b[5].TCR == Invalid:
		return fmt.Errorf("%q as TCR: %w", b[5].Code, ErrBandPosition)

	default:
		return nil
	}
}

// Value calculates the resistance of the band code in ohms.
func (bc BandCode) Value() (float64, error) {
	if err := bc.Validate(); err != nil {
		return 0, err
	}

	switch len(bc.Bands) {
	case Axial3Band, Axial4Band:
		return (bc.Bands[0].SigFig*10 + bc.Bands[1].SigFig) * bc.Bands[2].Multiplier, nil

	case Axial5Band, Axial6Band:
		return (bc.Bands[0].SigFig*100 + bc.Bands[1].SigFig*10 + bc.Bands[2].SigFig) * bc.Bands[3].Multiplier,
			nil

	default:
		// we've already validated the band count; this should never happen
		panic("invalid band code")
	}
}
