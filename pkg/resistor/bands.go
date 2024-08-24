package resistor

type BandColor int

const (
	Invalid          = -1
	DefaultTolerance = 20
)

const (
	Black BandColor = iota
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
	Code       string
	SigFig     float64 // SignificantFigures is a digit of the resistor value in ohms.
	Multiplier float64 // Multiplier is the multiplier of the resistor value.
	Tolerance  float64 // Tolerance is the tolerance of the resistor value in percentage.
	TCR        int     // TCR is the temperature coefficient of resistance in ppm/K.
}

//nolint:mnd,gochecknoglobals // reference value
var Bands = map[BandColor]Band{
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
}
