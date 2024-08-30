package axial

import (
	"errors"
	"fmt"
	"slices"

	"github.com/asphaltbuffet/ohm/pkg/resistor"
)

const (
	Axial1Band int = iota + 1 // 0 Ω resistors are single black band
	Axial3Band int = iota + 2
	Axial4Band
	Axial5Band
	Axial6Band
	// Axial7Band // 6-band with failure rate not implemented yet
)

// Resistor is an axial resistor with color bands to represent its value.
type Resistor struct {
	// IsReversed from input order.
	IsReversed bool `json:"reversed"`
	// Bands is the color bands on the resistor.
	Bands []Band `json:"bands"`
}

var (
	ErrBandCodeLength = errors.New("invalid length")
	ErrBandColor      = errors.New("unknown color")
	ErrBandPosition   = errors.New("invalid color at position")
)

var _ resistor.Resistor = new(Resistor)

func New(ss ...string) (*Resistor, error) {
	var (
		tokens []string
		err    error
	)

	if len(ss) == 1 {
		// if only one argument is passed, assume it is a color code string with abbreviations
		tokens = Tokenize(ss[0])
	} else {
		tokens = ss
	}

	if tokens == nil {
		return nil, fmt.Errorf("tokenized %v: %w", tokens, ErrBandCodeLength)
	}

	bands, err := convertToBands(tokens) // this only validates inputs are colors
	if err != nil {
		return nil, fmt.Errorf("convert to bands: %w", err)
	}

	rev := false

	// if order is invalid, reverse the order and try again
	if v := validateBandOrder(bands); v != nil {
		rev = !rev
		slices.Reverse(bands)

		revErr := validateBandOrder(bands)
		if revErr != nil {
			return nil, errors.Join(err, revErr)
		}
	}

	return &Resistor{IsReversed: rev, Bands: bands}, nil
}

func convertToBands(tokens []string) ([]Band, error) {
	var bands []Band
	for _, token := range tokens {
		c := GetColor(token)
		if c == None {
			return nil, fmt.Errorf("convert %q to band: %w", token, ErrBandColor)
		}

		bands = append(bands, ColorToBand(c))
	}

	return bands, nil
}

func Tokenize(s string) []string {
	var tokens []string

	// string must have an even number of characters
	if len(s)%2 != 0 {
		return nil
	}

	// We assume that the colors are 2-character codes or full color names.
	for i := 0; i < len(s); i += 2 {
		tokens = append(tokens, s[i:i+2])
	}

	return tokens
}

// Validate checks if the band code is valid.
func (r Resistor) Validate() error {
	return validateBandOrder(r.Bands)
}

// Value calculates the resistance of the band code in ohms.
func (r Resistor) Value() (float64, error) {
	if err := r.Validate(); err != nil {
		return 0, err
	}

	switch len(r.Bands) {
	case Axial3Band, Axial4Band:
		return (r.Bands[0].SigFig*10 + r.Bands[1].SigFig) * r.Bands[2].Multiplier, nil

	case Axial5Band, Axial6Band:
		return (r.Bands[0].SigFig*100 + r.Bands[1].SigFig*10 + r.Bands[2].SigFig) * r.Bands[3].Multiplier,
			nil

	default:
		// we've already validated the band count; this should never happen
		panic("invalid band code")
	}
}

// TCR returns the temperature coefficient of resistance in ppm/K.
func (r Resistor) TCR() int {
	if len(r.Bands) < Axial5Band {
		return 0
	}

	// TODO: this will be inaccurate for 7-band resistors
	return r.Bands[len(r.Bands)-1].TCR
}

func (r Resistor) Type() resistor.Marking {
	return resistor.Marking(len(r.Bands))
}

// Tolerance of the resistor value as ±%.
func (r Resistor) Tolerance() float64 {
	var t float64

	switch len(r.Bands) {
	case Axial4Band:
		t = r.Bands[3].Tolerance
	case Axial5Band, Axial6Band:
		t = r.Bands[4].Tolerance
	default:
		return DefaultTolerance
	}

	return t
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
