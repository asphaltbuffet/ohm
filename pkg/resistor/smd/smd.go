package smd

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/asphaltbuffet/ohm/pkg/resistor"
)

var ErrInvalidInput = errors.New("invalid input")

var _ resistor.Resistor = new(Resistor)

type Resistor struct {
	Digits    float64          `json:"digits"`
	MultPower int              `json:"multiplier"`
	RType     resistor.Marking `json:"type"`
}

func New(s string) (*Resistor, error) {
	const standardMaxLength = 3

	var d float64
	var m int
	var err error

	var t resistor.Marking

	switch {
	case len(s) == 0:
		return nil, ErrInvalidInput

	case s == "0", s == "00", s == "000":
		t = resistor.SMDStandard
		d = 0
		m = 0

	case s == "0000":
		t = resistor.SMDPrecision
		d = 0
		m = 0

	case strings.Contains(s, "R"):
		t = resistor.SMDStandard
		m = 0

		d, err = strconv.ParseFloat(strings.Replace(s, "R", ".", 1), 64)
		if err != nil {
			return nil, fmt.Errorf("digits %q: %w", s, err)
		}

	case len(s) < standardMaxLength:
		t = resistor.SMDStandard
		m = 0

		d, err = strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, fmt.Errorf("digits %q: %w", s[:2], err)
		}

	case len(s) == standardMaxLength:
		t = resistor.SMDStandard

		d, err = strconv.ParseFloat(s[:2], 64)
		if err != nil {
			return nil, fmt.Errorf("digits %q: %w", s[:2], err)
		}

		m, err = strconv.Atoi(s[2:])
		if err != nil {
			return nil, fmt.Errorf("multiplier %q: %w", s[2:], err)
		}

	default:
		return nil, ErrInvalidInput
	}

	return &Resistor{Digits: d, MultPower: m, RType: t}, nil
}

func (r Resistor) Value() (float64, error) {
	return r.Digits * math.Pow10(r.MultPower), nil
}

func (r Resistor) Tolerance() float64 {
	// TODO: not implemented yet
	return 0
}

func (r Resistor) TCR() int {
	// TODO: not implemented yet
	return 0
}

func (r Resistor) Type() resistor.Marking {
	return r.RType
}
