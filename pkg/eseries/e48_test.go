package es_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	es "github.com/asphaltbuffet/ohm/pkg/eseries"
)

func TestE48(t *testing.T) {
	tests := []struct {
		name string
		args int
		want float64
	}{
		{"under", 0, -1.00},
		{"middle", 2, 1.05},
		{"over", 49, -1.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.InEpsilon(t, tt.want, es.E48(tt.args), 0.0001)
		})
	}
}

func ExampleE48() {
	fmt.Printf("%.2f\n", es.E48(1))
	fmt.Printf("%.2f\n", es.E48(2))
	fmt.Printf("%.2f\n", es.E48(100)) // Out of range
	// Output:
	// 1.00
	// 1.05
	// -1.00
}
