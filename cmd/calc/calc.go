package calc

import (
	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"

	"github.com/asphaltbuffet/ohm/pkg/resistor"
	"github.com/asphaltbuffet/ohm/pkg/resistor/axial"
	"github.com/asphaltbuffet/ohm/pkg/resistor/smd"
)

func NewCommand() *cobra.Command {
	var isAxial, isSMD bool

	cmd := &cobra.Command{
		Use:     "calc",
		Short:   "calculate resistance value",
		Example: "ohm calc --smd 4k7",
		Aliases: []string{"c", "calculate"},
		Args:    cobra.MinimumNArgs(1),
		RunE:    runCalc,
	}

	cmd.Flags().BoolVarP(&isAxial, "axial", "a", false, "axial resistors")
	cmd.Flags().BoolVarP(&isSMD, "smd", "s", false, "SMD resistors")

	cmd.Flags().BoolP("tolerance", "t", false, "show tolerance")
	cmd.Flags().BoolP("TCR", "T", false, "show Temperature Coefficient of Resistance")
	cmd.Flags().BoolP("verbose", "v", false, "show calculations")

	cmd.Flags().BoolP("no-humanize", "H", false, "do not use SI units (e.g. 4700 instead of 4.7k)")

	cmd.MarkFlagsMutuallyExclusive("axial", "smd")
	cmd.MarkFlagsOneRequired("axial", "smd")
	cmd.Flags().SortFlags = false

	return cmd
}

func runCalc(cmd *cobra.Command, args []string) error {
	var r resistor.Resistor
	var err error

	switch {
	case cmd.Flags().Changed("smd"):
		r, err = smd.New(args[0])
		if err != nil {
			return err
		}

	case cmd.Flags().Changed("axial"):
		r, err = axial.New(args...)
		if err != nil {
			return err
		}

		// verbose, _ := cmd.Flags().GetBool("verbose")
		//
		// if verbose {
		// 	switch len(r.Bands) {
		// 	case axial.Axial3Band, axial.Axial4Band:
		// 		cmd.PrintErrf(
		// 			"%0.1f[%s] * 10 + %0.1f[%s] * %0.1f[%s]\n",
		// 			r.Bands[0].SigFig, r.Bands[0].Code,
		// 			r.Bands[1].SigFig, r.Bands[1].Code,
		// 			r.Bands[2].Multiplier, r.Bands[2].Code,
		// 		)

		// 		if len(r.Bands) == axial.Axial4Band {
		// 			cmd.PrintErrf("  Tolerance: %0.2f%% [%s]\n", r.Bands[3].Multiplier, r.Bands[3].Code)
		// 		}

		// 	case axial.Axial5Band, axial.Axial6Band:
		// 		cmd.PrintErrf(
		// 			"((%0.1f[%s] * 100) + (%0.1f[%s] * 10) + %0.1f[%s]) * %0.1f[%s]\n",
		// 			r.Bands[0].SigFig, r.Bands[0].Code,
		// 			r.Bands[1].SigFig, r.Bands[1].Code,
		// 			r.Bands[2].SigFig, r.Bands[2].Code,
		// 			r.Bands[3].Multiplier, r.Bands[3].Code,
		// 		)

		// 		cmd.PrintErrf("  Tolerance: %0.2f%% [%s]\n", r.Bands[3].Multiplier, r.Bands[3].Code)

		// 		if len(r.Bands) == axial.Axial6Band {
		// 			cmd.PrintErrf("  Tolerance: %0.2f%% [%s]\n", r.Bands[3].Multiplier, r.Bands[3].Code)
		// 		}

		// 	default:
		// 		cmd.PrintErrln("can't show work")
		// 	}
	}

	v, err := r.Value()
	if err != nil {
		return err
	}

	if ok, _ := cmd.Flags().GetBool("no-humanize"); ok {
		cmd.Println(humanize.Ftoa(v), "Ω")
	} else {
		cmd.Println(humanize.SI(v, "Ω"))
	}

	return nil
}
