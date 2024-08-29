package calc_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/asphaltbuffet/ohm/cmd/calc"
)

func TestNewCommand(t *testing.T) {
	var a *cobra.Command
	var b *cobra.Command

	t.Run("instantiate", func(t *testing.T) {
		a = calc.NewCommand()
		require.NotNil(t, a)
	})

	t.Run("re-instantiate", func(t *testing.T) {
		b = calc.NewCommand()
		require.NotNil(t, b)

		assert.NotSame(t, a, b)
	})
}
