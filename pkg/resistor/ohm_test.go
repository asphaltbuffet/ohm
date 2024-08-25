package resistor_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/ohm/pkg/resistor"
)

func TestBandCode_Validate(t *testing.T) {
	black := resistor.Bands[resistor.Black]
	red := resistor.Bands[resistor.Red]
	gold := resistor.Bands[resistor.Gold]
	pink := resistor.Bands[resistor.Pink]

	tests := []struct {
		name  string
		bands []resistor.Band
		want  error
	}{
		{"valid 3-band", []resistor.Band{black, red, red}, nil},
		{"invalid first band", []resistor.Band{gold, red, red}, resistor.ErrBandPosition},
		{"invalid second band", []resistor.Band{red, gold, red}, resistor.ErrBandPosition},
		{"gold multiplier", []resistor.Band{red, red, gold}, nil},
		{"valid 4-band", []resistor.Band{red, red, gold, gold}, nil},
		{"invalid tolerance", []resistor.Band{red, red, gold, pink}, resistor.ErrBandPosition},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := resistor.BandCode{
				Bands:    tt.bands,
				Reversed: false,
			}

			assert.ErrorIs(t, bc.Validate(), tt.want)
		})
	}
}

func TestBandCode_Value(t *testing.T) {
	red := resistor.Bands[resistor.Red]
	gold := resistor.Bands[resistor.Gold]
	orange := resistor.Bands[resistor.Orange]
	pink := resistor.Bands[resistor.Pink]

	tests := []struct {
		name      string
		Bands     []resistor.Band
		want      float64
		assertion require.ErrorAssertionFunc
	}{
		{
			name:      "valid 3-band",
			Bands:     []resistor.Band{red, red, orange},
			want:      22_000,
			assertion: require.NoError,
		},
		{
			name:      "invalid sig fig",
			Bands:     []resistor.Band{gold, red, orange, gold},
			want:      0,
			assertion: require.Error,
		},
		{
			name:      "valid 4-band",
			Bands:     []resistor.Band{red, red, orange, gold},
			want:      22_000,
			assertion: require.NoError,
		},
		{
			name:      "invalid 4-band tolerance",
			Bands:     []resistor.Band{red, red, orange, pink},
			want:      0,
			assertion: require.Error,
		},
		{
			name:      "valid 5-band",
			Bands:     []resistor.Band{red, red, orange, orange, gold},
			want:      223_000,
			assertion: require.NoError,
		},
		{
			name:      "invalid precision sig fig",
			Bands:     []resistor.Band{red, red, pink, orange, gold},
			want:      0,
			assertion: require.Error,
		},
		{
			name:      "invalid precision tolerance",
			Bands:     []resistor.Band{red, red, orange, orange, pink},
			want:      0,
			assertion: require.Error,
		},
		{
			name:      "invalid precision TCR",
			Bands:     []resistor.Band{red, red, orange, orange, gold, pink},
			want:      0,
			assertion: require.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := resistor.BandCode{Reversed: false, Bands: tt.Bands}

			got, err := bc.Value()

			tt.assertion(t, err)
			if err == nil {
				assert.InEpsilon(t, tt.want, 0.0001, got)
			}
		})
	}
}
