package ohm_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/ohm/pkg/ohm"
)

func TestBandCode_Validate(t *testing.T) {
	black := ohm.Bands[ohm.Black]
	red := ohm.Bands[ohm.Red]
	gold := ohm.Bands[ohm.Gold]
	pink := ohm.Bands[ohm.Pink]

	tests := []struct {
		name      string
		bands     []ohm.Band
		assertion require.ErrorAssertionFunc
	}{
		{"valid 3-band", []ohm.Band{black, red, red}, require.NoError},
		{"invalid first band", []ohm.Band{gold, red, red}, require.Error},
		{"invalid second band", []ohm.Band{red, gold, red}, require.Error},
		{"gold multiplier", []ohm.Band{red, red, gold}, require.NoError},
		{"valid 4-band", []ohm.Band{red, red, gold, gold}, require.NoError},
		{"invalid tolerance", []ohm.Band{red, red, gold, pink}, require.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := ohm.BandCode{
				Bands: tt.bands,
			}

			tt.assertion(t, bc.Validate())
		})
	}
}
