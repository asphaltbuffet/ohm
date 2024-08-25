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
		name    string
		args    args
		want    *resistor.BandCode
		wantErr error
	}{
		{"empty", args{[]string{}, false}, nil, resistor.ErrBandCodeLength},
		{"one char", args{[]string{"B"}, false}, nil, resistor.ErrBandCodeLength}, // can't tokenize odd number of characters
		{
			"zero ohm",
			args{[]string{"BK"}, false},
			&resistor.BandCode{
				Reversed: false,
				Bands:    []resistor.Band{resistor.Bands[resistor.Black]},
			},
			nil,
		},
		{
			"multiple tokens",
			args{[]string{"BKBUBKSVSV"}, false},
			&resistor.BandCode{
				Reversed: false,
				Bands: []resistor.Band{
					resistor.Bands[resistor.Black],
					resistor.Bands[resistor.Blue],
					resistor.Bands[resistor.Black],
					resistor.Bands[resistor.Silver],
					resistor.Bands[resistor.Silver],
				},
			},
			nil,
		},
		{
			"valid in reverse",
			args{[]string{"SVSVBKBUBK"}, false},
			&resistor.BandCode{
				Reversed: true,
				Bands: []resistor.Band{
					resistor.Bands[resistor.Black],
					resistor.Bands[resistor.Blue],
					resistor.Bands[resistor.Black],
					resistor.Bands[resistor.Silver],
					resistor.Bands[resistor.Silver],
				},
			},
			nil,
		},
		{"invalid color", args{[]string{"ZZ"}, false}, nil, resistor.ErrBandColor},
		{
			"pre-tokenized",
			args{[]string{"BK", "BU", "BK", "SV", "SV"}, false},
			&resistor.BandCode{
				Reversed: false,
				Bands: []resistor.Band{
					resistor.Bands[resistor.Black],
					resistor.Bands[resistor.Blue],
					resistor.Bands[resistor.Black],
					resistor.Bands[resistor.Silver],
					resistor.Bands[resistor.Silver],
				},
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resistor.Parse(tt.args.s, tt.args.rev)

			require.ErrorIs(t, err, tt.wantErr)
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
		{"invalid", "ZZ", resistor.None},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, resistor.GetColor(tt.args))
		})
	}
}
