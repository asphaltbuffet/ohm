package resistor_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/ohm/pkg/resistor"
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
			assert.Equal(t, tt.want, resistor.Tokenize(tt.args))
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		s   []string
		rev bool
	}

	tests := []struct {
		name      string
		args      args
		want      *resistor.BandCode
		assertion require.ErrorAssertionFunc
	}{
		{"empty", args{[]string{""}, false}, nil, require.Error},
		{"one char", args{[]string{"B"}, false}, nil, require.Error},
		{
			"one token",
			args{[]string{"BK"}, false},
			&resistor.BandCode{
				Bands: []resistor.Band{resistor.Bands[resistor.Black]},
			},
			require.NoError,
		},
		{
			"multiple tokens",
			args{[]string{"BKBUPKSVSV"}, false},
			&resistor.BandCode{
				Bands: []resistor.Band{
					resistor.Bands[resistor.Black],
					resistor.Bands[resistor.Blue],
					resistor.Bands[resistor.Pink],
					resistor.Bands[resistor.Silver],
					resistor.Bands[resistor.Silver],
				},
			},
			require.NoError,
		},
		{
			"valid in reverse",
			args{[]string{"SVSVPKBUBK"}, false},
			&resistor.BandCode{
				Bands: []resistor.Band{
					resistor.Bands[resistor.Black],
					resistor.Bands[resistor.Blue],
					resistor.Bands[resistor.Pink],
					resistor.Bands[resistor.Silver],
					resistor.Bands[resistor.Silver],
				},
			},
			require.NoError,
		},
		{"invalid color", args{[]string{"ZZ"}, false}, nil, require.Error},
		{
			"pre-tokenized",
			args{[]string{"BK", "BU", "PK", "SV", "SV"}, false},
			&resistor.BandCode{
				Bands: []resistor.Band{
					resistor.Bands[resistor.Black],
					resistor.Bands[resistor.Blue],
					resistor.Bands[resistor.Pink],
					resistor.Bands[resistor.Silver],
					resistor.Bands[resistor.Silver],
				},
			},
			require.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resistor.Parse(tt.args.s, tt.args.rev)
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
		want resistor.BandColor
	}{
		{"lowercase", "pk", resistor.Pink},
		{"uppercase", "SV", resistor.Silver},
		{"full name", "purple", resistor.Violet},
		{"invalid", "ZZ", resistor.Invalid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, resistor.GetColor(tt.args))
		})
	}
}
