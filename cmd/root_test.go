package cmd_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cli "github.com/asphaltbuffet/ohm/cmd"
)

func TestNewCommand(t *testing.T) {
	var a *cobra.Command
	var b *cobra.Command

	t.Run("instantiate", func(t *testing.T) {
		a = cli.RootCommand()
		require.NotNil(t, a)
	})

	t.Run("re-instantiate", func(t *testing.T) {
		b = cli.RootCommand()
		require.NotNil(t, b)

		assert.NotSame(t, a, b)
	})
}
