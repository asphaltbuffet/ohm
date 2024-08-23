package es_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	es "github.com/asphaltbuffet/ohm/pkg/eseries"
)

func TestE3(t *testing.T) {
	tests := []struct {
		name string
		args int
		want float64
	}{
		{"under", 0, -1.0},
		{"middle", 2, 2.2},
		{"over", 4, -1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.InEpsilon(t, tt.want, es.E3(tt.args), 0.0001)
		})
	}
}

func ExampleE3() {
	fmt.Printf("%.1f\n", es.E3(1))
	fmt.Printf("%.1f\n", es.E3(2))
	fmt.Printf("%.1f\n", es.E3(100)) // Out of range
	// Output:
	// 1.0
	// 2.2
	// -1.0
}
