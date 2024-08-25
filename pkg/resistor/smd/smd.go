package smd

import (
	"errors"
	"fmt"
)

var ErrInvalidInput = errors.New("invalid input")

func Parse(s string) (int, error) {
	v := new(int)

	_, err := fmt.Sscanf(s, "%d", v)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrInvalidInput, err)
	}

	return *v, nil
}
