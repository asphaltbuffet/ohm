package axial_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/ohm/pkg/resistor/axial"
)

func TestBandCode_Validate(t *testing.T) {
	black := axial.Bands[axial.Black]
	red := axial.Bands[axial.Red]
	gold := axial.Bands[axial.Gold]
	pink := axial.Bands[axial.Pink]

	tests := []struct {
		name  string
		bands []axial.Band
		want  error
	}{
		{"valid 3-band", []axial.Band{black, red, red}, nil},
		{"invalid first band", []axial.Band{gold, red, red}, axial.ErrBandPosition},
		{"invalid second band", []axial.Band{red, gold, red}, axial.ErrBandPosition},
		{"gold multiplier", []axial.Band{red, red, gold}, nil},
		{"valid 4-band", []axial.Band{red, red, gold, gold}, nil},
		{"invalid tolerance", []axial.Band{red, red, gold, pink}, axial.ErrBandPosition},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := axial.Resistor{
				Bands:      tt.bands,
				IsReversed: false,
			}

			assert.ErrorIs(t, bc.Validate(), tt.want)
		})
	}
}

func TestBandCode_Value(t *testing.T) {
	red := axial.Bands[axial.Red]
	gold := axial.Bands[axial.Gold]
	orange := axial.Bands[axial.Orange]
	pink := axial.Bands[axial.Pink]

	tests := []struct {
		name      string
		Bands     []axial.Band
		want      float64
		assertion require.ErrorAssertionFunc
	}{
		{
			name:      "valid 3-band",
			Bands:     []axial.Band{red, red, orange},
			want:      22_000,
			assertion: require.NoError,
		},
		{
			name:      "invalid sig fig",
			Bands:     []axial.Band{gold, red, orange, gold},
			want:      0,
			assertion: require.Error,
		},
		{
			name:      "valid 4-band",
			Bands:     []axial.Band{red, red, orange, gold},
			want:      22_000,
			assertion: require.NoError,
		},
		{
			name:      "invalid 4-band tolerance",
			Bands:     []axial.Band{red, red, orange, pink},
			want:      0,
			assertion: require.Error,
		},
		{
			name:      "valid 5-band",
			Bands:     []axial.Band{red, red, orange, orange, gold},
			want:      223_000,
			assertion: require.NoError,
		},
		{
			name:      "invalid precision sig fig",
			Bands:     []axial.Band{red, red, pink, orange, gold},
			want:      0,
			assertion: require.Error,
		},
		{
			name:      "invalid precision tolerance",
			Bands:     []axial.Band{red, red, orange, orange, pink},
			want:      0,
			assertion: require.Error,
		},
		{
			name:      "invalid precision TCR",
			Bands:     []axial.Band{red, red, orange, orange, gold, pink},
			want:      0,
			assertion: require.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := axial.Resistor{IsReversed: false, Bands: tt.Bands}

			got, err := bc.Value()

			tt.assertion(t, err)
			if err == nil {
				assert.InEpsilon(t, tt.want, 0.0001, got)
			}
		})
	}
}
