package axial

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

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

		bands = append(bands, Bands[c])
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

func GetColor(s string) BandColor {
	return map[string]BandColor{
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
	}[strings.ToLower(s)]
}
