package axial

import (
	"github.com/spf13/cobra"

	"github.com/asphaltbuffet/ohm/pkg/resistor"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "axial <colors...>",
		Short:   "calculate the value of axial resistors",
		Aliases: []string{"a"},
		Args:    cobra.MinimumNArgs(1),
		Run:     RunAxial,
	}

	return cmd
}

func RunAxial(cmd *cobra.Command, args []string) {
	code, err := resistor.Parse(args)
	if err != nil {
		cmd.PrintErrln(err)
		return
	}

	val, err := code.Resistance()
	if err != nil {
		cmd.PrintErrln(err)
		return
	}

	cmd.Printf("%.0f Ω\n", val)
}
