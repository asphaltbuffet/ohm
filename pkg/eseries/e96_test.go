package es_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	es "github.com/asphaltbuffet/ohm/pkg/eseries"
)

func TestE96(t *testing.T) {
	tests := []struct {
		name string
		args int
		want float64
	}{
		{"under", 0, -1.00},
		{"middle", 48, 3.09},
		{"over", 97, -1.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.InEpsilon(t, tt.want, es.E96(tt.args), 0.0001)
		})
	}
}

func ExampleE96() {
	fmt.Printf("%.2f\n", es.E96(1))
	fmt.Printf("%.2f\n", es.E96(25))
	fmt.Printf("%.2f\n", es.E96(100)) // Out of range
	// Output:
	// 1.00
	// 1.78
	// -1.00
}
