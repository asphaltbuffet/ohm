package axialcli

import (
	"github.com/spf13/cobra"

	"github.com/asphaltbuffet/ohm/pkg/resistor/axial"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "axial <colors...>",
		Short:   "calculate the value of axial resistors",
		Aliases: []string{"a"},
		Args:    cobra.MinimumNArgs(1),
		Run:     RunAxial,
	}

	cmd.Flags().BoolP("verbose", "v", false, "show work")

	return cmd
}

func RunAxial(cmd *cobra.Command, args []string) {
	var rev bool

	code, err := axial.Parse(args, rev)
	if err != nil {
		rev = !rev
		// try again with reversed order
		code, err = axial.Parse(args, rev)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		cmd.PrintErrln("reverse order detected")
	}

	verbose, _ := cmd.Flags().GetBool("verbose")

	if verbose {
		switch len(code.Bands) {
		case axial.Axial3Band, axial.Axial4Band:
			cmd.PrintErrf(
				"%0.1f[%s] * 10 + %0.1f[%s] * %0.1f[%s]\n",
				code.Bands[0].SigFig, code.Bands[0].Code,
				code.Bands[1].SigFig, code.Bands[1].Code,
				code.Bands[2].Multiplier, code.Bands[2].Code,
			)

			if len(code.Bands) == axial.Axial4Band {
				cmd.PrintErrf("  Tolerance: %0.2f%% [%s]\n", code.Bands[3].Multiplier, code.Bands[3].Code)
			}

		case axial.Axial5Band, axial.Axial6Band:
			cmd.PrintErrf(
				"((%0.1f[%s] * 100) + (%0.1f[%s] * 10) + %0.1f[%s]) * %0.1f[%s]\n",
				code.Bands[0].SigFig, code.Bands[0].Code,
				code.Bands[1].SigFig, code.Bands[1].Code,
				code.Bands[2].SigFig, code.Bands[2].Code,
				code.Bands[3].Multiplier, code.Bands[3].Code,
			)

			cmd.PrintErrf("  Tolerance: %0.2f%% [%s]\n", code.Bands[3].Multiplier, code.Bands[3].Code)

			if len(code.Bands) == axial.Axial6Band {
				cmd.PrintErrf("  Tolerance: %0.2f%% [%s]\n", code.Bands[3].Multiplier, code.Bands[3].Code)
			}

		default:
			cmd.PrintErrln("can't show work")
		}
	}

	val, err := code.Value()
	if err != nil {
		cmd.PrintErrln(err)
		return
	}

	cmd.Printf("%.0f Ω\n", val)
}
