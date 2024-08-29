package smd_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/ohm/pkg/resistor"
	"github.com/asphaltbuffet/ohm/pkg/resistor/smd"
)

func TestNew(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		s       string
		want    *smd.Resistor
		wantErr error
	}{
		{"empty", "", nil, smd.ErrInvalidInput},
		{"non-numerical", "abc", nil, strconv.ErrSyntax},

		{"zero", "0", &smd.Resistor{Digits: 0, MultPower: 0, RType: resistor.SMDStandard}, nil},
		{"zero3", "000", &smd.Resistor{Digits: 0, MultPower: 0, RType: resistor.SMDStandard}, nil},
		{"zero4", "0000", &smd.Resistor{Digits: 0, MultPower: 0, RType: resistor.SMDPrecision}, nil},
		{"leading zero", "042", &smd.Resistor{Digits: 4, MultPower: 2, RType: resistor.SMDStandard}, nil},

		// standard smd resistors (3-digit)
		{"334", "334", &smd.Resistor{Digits: 33, MultPower: 4, RType: resistor.SMDStandard}, nil},
		{"222", "222", &smd.Resistor{Digits: 22, MultPower: 2, RType: resistor.SMDStandard}, nil},
		{"473", "473", &smd.Resistor{Digits: 47, MultPower: 3, RType: resistor.SMDStandard}, nil},
		{"105", "105", &smd.Resistor{Digits: 10, MultPower: 5, RType: resistor.SMDStandard}, nil},
		// less than 100, but greater than 10
		{"100", "100", &smd.Resistor{Digits: 10, MultPower: 0, RType: resistor.SMDStandard}, nil},
		{"220", "220", &smd.Resistor{Digits: 22, MultPower: 0, RType: resistor.SMDStandard}, nil},
		{"10", "10", &smd.Resistor{Digits: 10, MultPower: 0, RType: resistor.SMDStandard}, nil},
		{"22", "22", &smd.Resistor{Digits: 22, MultPower: 0, RType: resistor.SMDStandard}, nil},
		// less than 10 (with radix)
		{"4R7", "4R7", &smd.Resistor{Digits: 4.7, MultPower: 0, RType: resistor.SMDStandard}, nil},
		{"R300", "R300", &smd.Resistor{Digits: 0.3, MultPower: 0, RType: resistor.SMDStandard}, nil},
		{"0R22", "0R22", &smd.Resistor{Digits: 0.22, MultPower: 0, RType: resistor.SMDStandard}, nil},
		{"0R01", "0R01", &smd.Resistor{Digits: 0.01, MultPower: 0, RType: resistor.SMDStandard}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := smd.New(tt.s)

			require.ErrorIs(t, err, tt.wantErr)
			assert.EqualExportedValues(t, tt.want, got)
		})
	}
}

func TestValue(t *testing.T) {
	t.Parallel()

	type args struct {
		digits float64
		mult   int
		t      resistor.Marking
	}

	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr error
	}{
		// standard smd resistors (3-digit)
		{"334", args{digits: 33, mult: 4, t: resistor.SMDStandard}, 330_000, nil},
		{"222", args{digits: 22, mult: 2, t: resistor.SMDStandard}, 2_200, nil},
		{"473", args{digits: 47, mult: 3, t: resistor.SMDStandard}, 47_000, nil},
		{"105", args{digits: 10, mult: 5, t: resistor.SMDStandard}, 1_000_000, nil},
		// less than 100, but greater than 10
		{"100", args{digits: 10, mult: 0, t: resistor.SMDStandard}, 10, nil},
		{"220", args{digits: 22, mult: 0, t: resistor.SMDStandard}, 22, nil},
		{"10", args{digits: 10, mult: 0, t: resistor.SMDStandard}, 10, nil},
		{"22", args{digits: 22, mult: 0, t: resistor.SMDStandard}, 22, nil},
		// less than 10 (with radix)
		{"4R7", args{digits: 4.7, mult: 0, t: resistor.SMDStandard}, 4.7, nil},
		{"R300", args{digits: 0.3, mult: 0, t: resistor.SMDStandard}, 0.3, nil},
		{"0R22", args{digits: 0.22, mult: 0, t: resistor.SMDStandard}, 0.22, nil},
		{"0R01", args{digits: 0.01, mult: 0, t: resistor.SMDStandard}, 0.01, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := smd.Resistor{Digits: tt.args.digits, MultPower: tt.args.mult, RType: tt.args.t}
			got, err := r.Value()

			require.ErrorIs(t, err, tt.wantErr)
			assert.InEpsilon(t, tt.want, got, 0.00001)
		})
	}
}
