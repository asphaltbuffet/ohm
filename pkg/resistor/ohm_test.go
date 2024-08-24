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
		name      string
		bands     []resistor.Band
		assertion require.ErrorAssertionFunc
	}{
		{"valid 3-band", []resistor.Band{black, red, red}, require.NoError},
		{"invalid first band", []resistor.Band{gold, red, red}, require.Error},
		{"invalid second band", []resistor.Band{red, gold, red}, require.Error},
		{"gold multiplier", []resistor.Band{red, red, gold}, require.NoError},
		{"valid 4-band", []resistor.Band{red, red, gold, gold}, require.NoError},
		{"invalid tolerance", []resistor.Band{red, red, gold, pink}, require.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := resistor.BandCode{
				Bands: tt.bands,
			}

			tt.assertion(t, bc.Validate())
		})
	}
}

func TestBandCode_Resistance(t *testing.T) {
	red := resistor.Bands[resistor.Red]
	gold := resistor.Bands[resistor.Gold]
	orange := resistor.Bands[resistor.Orange]

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
			name:      "valid 4-band",
			Bands:     []resistor.Band{red, red, orange, gold},
			want:      22_000,
			assertion: require.NoError,
		},
		{
			name:      "valid 5-band",
			Bands:     []resistor.Band{red, red, orange, orange, gold},
			want:      223_000,
			assertion: require.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := resistor.BandCode{Bands: tt.Bands}

			got, err := bc.Resistance()

			tt.assertion(t, err)
			assert.InEpsilon(t, tt.want, 0.0001, got)
		})
	}
}
