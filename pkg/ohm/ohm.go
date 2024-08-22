package ohm

// BandCode is a struct that represents a resistor 4, 5, or 6 band code.
type BandCode struct {
	Bands []Band
}

// Validate checks if the band code is valid.
func (bc BandCode) Validate() error {
	if len(bc.Bands) < 3 || len(bc.Bands) > 6 { // TODO: support 0 Ω resistors (single black)
		return &BandCodeLengthError{Length: len(bc.Bands)}
	}

	// gold/silver/pink bands are not allowed in the first two bands
	for i := 0; i <= 1; i++ {
		if bc.Bands[i].SignificantFigures == Invalid {
			return &BandCodeColorError{ColorCode: bc.Bands[i].Code, BandType: "SigFig"}
		}
	}

	// 3rd band is the multiplier
	if bc.Bands[2].Multiplier == Invalid {
		return &BandCodeColorError{ColorCode: bc.Bands[2].Code, BandType: "Mult"}
	}

	// 4th band is the tolerance
	if len(bc.Bands) == 4 && bc.Bands[3].Tolerance == Invalid {
		return &BandCodeColorError{ColorCode: bc.Bands[3].Code, BandType: "Tol"}
	}

	return nil
}

func (bc BandCode) Resistance() (float64, error) {
	if err := bc.Validate(); err != nil {
		return 0, err
	}

	switch len(bc.Bands) {
	case 3, 4: //nolint:mnd // 3 or 4 bands
		return (bc.Bands[0].SignificantFigures*10 + bc.Bands[1].SignificantFigures) * bc.Bands[2].Multiplier, nil
	default:
		return 0, &BandInvalidError{Bands: bc.Bands}
	}
}
