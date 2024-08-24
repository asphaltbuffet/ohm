package resistor

import (
	"fmt"
	"slices"
	"strings"
)

func Parse(s []string, rev bool) (*BandCode, error) {
	var tokens []string

	if len(s) == 1 {
		// if only one argument is passed, assume it is a color code string with abbreviations
		tokens = Tokenize(s[0])
	} else {
		tokens = s
	}

	if tokens == nil {
		return nil, &BandCodeLengthError{Length: len(s)}
	}

	if rev {
		slices.Reverse(tokens)
	}

	bands, err := convertToBands(tokens)
	if err != nil {
		return nil, err
	}

	return &BandCode{Bands: bands}, nil
}

func convertToBands(tokens []string) ([]Band, error) {
	var bands []Band
	for i, token := range tokens {
		c := GetColor(token)
		band, ok := Bands[c]
		if !ok {
			return nil, &BandCodeColorError{ColorCode: token, BandType: fmt.Sprintf("token[%d]", i)}
		}
		bands = append(bands, band)
	}

	return bands, nil
}

func Tokenize(s string) []string {
	var tokens []string

	if len(s)%2 != 0 {
		// string must have an even number of characters
		return nil
	}

	// We assume that the colors are 2-character codes or full color names.
	for i := 0; i < len(s); i += 2 {
		tokens = append(tokens, s[i:i+2])
	}

	return tokens
}

//nolint:gochecknoglobals // internal lookup
var colorMap = map[string]BandColor{
	"bk":     Black,
	"black":  Black,
	"bn":     Brown,
	"brown":  Brown,
	"rd":     Red,
	"red":    Red,
	"og":     Orange,
	"orange": Orange,
	"ye":     Yellow,
	"yellow": Yellow,
	"gn":     Green,
	"green":  Green,
	"bu":     Blue,
	"blue":   Blue,
	"vt":     Violet,
	"violet": Violet,
	"pu":     Violet,
	"purple": Violet,
	"gy":     Grey,
	"grey":   Grey,
	"slate":  Grey,
	"wh":     White,
	"white":  White,
	"gd":     Gold,
	"gold":   Gold,
	"sv":     Silver,
	"silver": Silver,
	"pk":     Pink,
	"pink":   Pink,
}

func GetColor(s string) BandColor {
	if color, exists := colorMap[strings.ToLower(s)]; exists {
		return color
	}

	return Invalid
}
