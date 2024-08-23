package smd_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/ohm/pkg/resistor/smd"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		s       string
		want    int
		wantErr error
	}{
		{"empty", "", 0, smd.ErrInvalidInput},
		{"non-numerical", "abc", 0, smd.ErrInvalidInput},
		{"too large", fmt.Sprint(math.MaxInt64, "0"), 0, smd.ErrInvalidInput},
		{"zero", "0", 0, nil},
		{"leading zero", "042", 42, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := smd.Parse(tt.s)

			require.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
