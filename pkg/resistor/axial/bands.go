package axial

//go:generate stringer -type=BandColor
type BandColor int

const (
	Invalid          = -1
	DefaultTolerance = 20
)

const (
	None BandColor = iota
	Black
	Brown
	Red
	Orange
	Yellow
	Green
	Blue
	Violet
	Grey
	White
	Gold
	Silver
	Pink
)

type Band struct {
	Code string `json:"code"`
	// SignificantFigures is a digit of the resistor value in ohms.
	SigFig float64 `json:"sig_fig,omitempty"`
	// Multiplier is the multiplier of the resistor value.
	Multiplier float64 `json:"multiplier,omitempty"`
	// Tolerance is the tolerance of the resistor value in percentage.
	Tolerance float64 `json:"tolerance,omitempty"`
	// TCR is the temperature coefficient of resistance in ppm/K.
	TCR int `json:"tcr,omitempty"`
}

// ColorToBand returns the Band struct for the given BandColor.
func ColorToBand(b BandColor) Band {
	//nolint:mnd // reference value
	return map[BandColor]Band{
		None: {
			Code:       "--",
			SigFig:     Invalid,
			Multiplier: Invalid,
			Tolerance:  Invalid,
			TCR:        Invalid,
		},
		Black: {
			Code:       "BK",
			SigFig:     0,
			Multiplier: 1,
			Tolerance:  Invalid,
			TCR:        250,
		},
		Brown: {
			Code:       "BN",
			SigFig:     1,
			Multiplier: 10,
			Tolerance:  1,
			TCR:        100,
		},
		Red: {
			Code:       "RD",
			SigFig:     2,
			Multiplier: 100,
			Tolerance:  2,
			TCR:        50,
		},
		Orange: {
			Code:       "OG",
			SigFig:     3,
			Multiplier: 1_000,
			Tolerance:  3,
			TCR:        15,
		},
		Yellow: {
			Code:       "YE",
			SigFig:     4,
			Multiplier: 10_000,
			Tolerance:  4,
			TCR:        25,
		},
		Green: {
			Code:       "GN",
			SigFig:     5,
			Multiplier: 100_000,
			Tolerance:  0.5,
			TCR:        20,
		},
		Blue: {
			Code:       "BL",
			SigFig:     6,
			Multiplier: 1_000_000,
			Tolerance:  0.25,
			TCR:        10,
		},
		Violet: {
			Code:       "VT",
			SigFig:     7,
			Multiplier: 10_000_000,
			Tolerance:  0.1,
			TCR:        5,
		},
		Grey: {
			Code:       "GY",
			SigFig:     8,
			Multiplier: 100_000_000,
			Tolerance:  0.05,
			TCR:        1,
		},
		White: {
			Code:       "WH",
			SigFig:     9,
			Multiplier: 1_000_000_000,
			Tolerance:  Invalid,
			TCR:        Invalid,
		},
		Gold: {
			Code:       "GD",
			SigFig:     Invalid,
			Multiplier: 0.1,
			Tolerance:  5,
			TCR:        Invalid,
		},
		Silver: {
			Code:       "SV",
			SigFig:     Invalid,
			Multiplier: 0.01,
			Tolerance:  10,
			TCR:        Invalid,
		},
		Pink: {
			Code:       "PK",
			SigFig:     Invalid,
			Multiplier: 0.001,
			Tolerance:  Invalid,
			TCR:        Invalid,
		},
	}[b]
}
