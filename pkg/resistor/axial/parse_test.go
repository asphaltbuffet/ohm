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
		s   []string
		rev bool
	}

	tests := []struct {
		name    string
		args    args
		want    *axial.BandCode
		wantErr error
	}{
		{"empty", args{[]string{}, false}, nil, axial.ErrBandCodeLength},
		{"one char", args{[]string{"B"}, false}, nil, axial.ErrBandCodeLength}, // can't tokenize odd number of characters
		{
			"zero ohm",
			args{[]string{"BK"}, false},
			&axial.BandCode{
				Reversed: false,
				Bands:    []axial.Band{axial.Bands[axial.Black]},
			},
			nil,
		},
		{
			"multiple tokens",
			args{[]string{"BKBUBKSVSV"}, false},
			&axial.BandCode{
				Reversed: false,
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
			args{[]string{"SVSVBKBUBK"}, false},
			&axial.BandCode{
				Reversed: true,
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
		{"invalid color", args{[]string{"ZZ"}, false}, nil, axial.ErrBandColor},
		{
			"pre-tokenized",
			args{[]string{"BK", "BU", "BK", "SV", "SV"}, false},
			&axial.BandCode{
				Reversed: false,
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
			got, err := axial.Parse(tt.args.s, tt.args.rev)

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
