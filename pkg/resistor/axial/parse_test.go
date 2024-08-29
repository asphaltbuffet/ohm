package axial_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/ohm/pkg/resistor/axial"
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
			assert.Equal(t, tt.want, axial.Tokenize(tt.args))
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		s []string
	}

	tests := []struct {
		name    string
		args    args
		want    *axial.Resistor
		wantErr error
	}{
		{"empty", args{[]string{}}, nil, axial.ErrBandCodeLength},
		{"one char", args{[]string{"B"}}, nil, axial.ErrBandCodeLength}, // can't tokenize odd number of characters
		{
			"zero ohm",
			args{[]string{"BK"}},
			&axial.Resistor{
				IsReversed: false,
				Bands:      []axial.Band{axial.Bands[axial.Black]},
			},
			nil,
		},
		{
			"multiple tokens",
			args{[]string{"BKBUBKSVSV"}},
			&axial.Resistor{
				IsReversed: false,
				Bands: []axial.Band{
					axial.Bands[axial.Black],
					axial.Bands[axial.Blue],
					axial.Bands[axial.Black],
					axial.Bands[axial.Silver],
					axial.Bands[axial.Silver],
				},
			},
			nil,
		},
		{
			"valid in reverse",
			args{[]string{"SVSVBKBUBK"}},
			&axial.Resistor{
				IsReversed: true,
				Bands: []axial.Band{
					axial.Bands[axial.Black],
					axial.Bands[axial.Blue],
					axial.Bands[axial.Black],
					axial.Bands[axial.Silver],
					axial.Bands[axial.Silver],
				},
			},
			nil,
		},
		{"invalid color", args{[]string{"ZZ"}}, nil, axial.ErrBandColor},
		{
			"pre-tokenized",
			args{[]string{"BK", "BU", "BK", "SV", "SV"}},
			&axial.Resistor{
				IsReversed: false,
				Bands: []axial.Band{
					axial.Bands[axial.Black],
					axial.Bands[axial.Blue],
					axial.Bands[axial.Black],
					axial.Bands[axial.Silver],
					axial.Bands[axial.Silver],
				},
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := axial.New(tt.args.s...)

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
		want axial.BandColor
	}{
		{"lowercase", "pk", axial.Pink},
		{"uppercase", "SV", axial.Silver},
		{"full name", "purple", axial.Violet},
		{"invalid", "ZZ", axial.None},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, axial.GetColor(tt.args))
		})
	}
}
