package axial

import (
	"strings"
)

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
