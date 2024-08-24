package resistor

// BandCode is a struct that represents a resistor 4, 5, or 6 band code.
type BandCode struct {
	Bands []Band
}

// Validate checks if the band code is valid.
func (bc BandCode) Validate() error {
	return validateBandOrder(bc.Bands)
}

func validateBandOrder(b []Band) error {
	l := len(b)

	//nolint:mnd // 3, 4, 5, or 6 bands
	switch {
	case l == 1 && b[0].Code != "BK": // 0 Ω resistors are single black band
		fallthrough
	case l < 3, l > 6:
		return &BandCodeLengthError{Length: l}

	case b[0].SigFig == Invalid:
		return &BandCodeColorError{ColorCode: b[0].Code, BandType: "significant figure"}

	case b[1].SigFig == Invalid:
		return &BandCodeColorError{ColorCode: b[1].Code, BandType: "significant figure"}

	// any color can be used for band 3 if it's a multiplier, but
	// 5/6 band resistors have a 3rd significant figure band
	case l >= 5 && b[2].SigFig == Invalid:
		return &BandCodeColorError{ColorCode: b[2].Code, BandType: "significant figure"}

	case l == 4 && b[3].Tolerance == Invalid:
		return &BandCodeColorError{ColorCode: b[3].Code, BandType: "tolerance"}

	case l >= 5 && b[4].Tolerance == Invalid:
		return &BandCodeColorError{ColorCode: b[4].Code, BandType: "tolerance"}
	case l >= 6 && b[5].TCR == Invalid:
		return &BandCodeColorError{ColorCode: b[5].Code, BandType: "TCR"}
	default:
		return nil
	}
}

func (bc BandCode) Resistance() (float64, error) {
	if err := bc.Validate(); err != nil {
		return 0, err
	}

	switch len(bc.Bands) {
	case 3, 4: //nolint:mnd // 3 or 4 bands
		return (bc.Bands[0].SigFig*10 + bc.Bands[1].SigFig) * bc.Bands[2].Multiplier, nil

	case 5, 6: //nolint:mnd // 5 or 6 bands
		return (bc.Bands[0].SigFig*100 + bc.Bands[1].SigFig*10 + bc.Bands[2].SigFig) * bc.Bands[3].Multiplier,
			nil

	default:
		return 0, &BandInvalidError{Bands: bc.Bands}
	}
}
