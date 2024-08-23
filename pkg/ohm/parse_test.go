package ohm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/ohm/pkg/ohm"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{"empty", "", nil},
		{"one char", "B", nil},
		{"one token", "BK", []string{"BK"}},
		{"multiple tokens", "BKBUPKSVSV", []string{"BK", "BU", "PK", "SV", "SV"}},
		{"invalid color", "ZZ", []string{"ZZ"}},
		{"invalid length", "BKB", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ohm.Tokenize(tt.args))
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		args      string
		want      *ohm.BandCode
		assertion require.ErrorAssertionFunc
	}{
		{"empty", "", nil, require.Error},
		{"one char", "B", nil, require.Error},
		{"one token", "BK", &ohm.BandCode{Bands: []ohm.Band{ohm.Bands[ohm.Black]}}, require.NoError},
		{
			"multiple tokens",
			"BKBUPKSVSV",
			&ohm.BandCode{
				Bands: []ohm.Band{
					ohm.Bands[ohm.Black],
					ohm.Bands[ohm.Blue],
					ohm.Bands[ohm.Pink],
					ohm.Bands[ohm.Silver],
					ohm.Bands[ohm.Silver],
				},
			},
			require.NoError,
		},
		{"invalid color", "ZZ", nil, require.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ohm.Parse(tt.args)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetColor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args string
		want ohm.BandColor
	}{
		{"lowercase", "pk", ohm.Pink},
		{"uppercase", "SV", ohm.Silver},
		{"full name", "purple", ohm.Violet},
		{"invalid", "ZZ", ohm.Invalid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, ohm.GetColor(tt.args))
		})
	}
}
