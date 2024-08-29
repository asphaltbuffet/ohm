package smdcli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/asphaltbuffet/ohm/pkg/resistor/smd"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "smd <value>",
		Aliases:      []string{"s"},
		SuggestFor:   []string{"smt"},
		SilenceUsage: true,
		Short:        "calculate the value of smd resistors",
		Args:         cobra.ExactArgs(1),
		PreRunE: func(_ *cobra.Command, args []string) error {
			_, err := fmt.Sscanf(args[0], "%d", new(int))
			return err
		},
		Run: func(cmd *cobra.Command, args []string) {
			v, err := smd.New(args[0])
			if err != nil {
				cmd.PrintErrln(err)
				return
			}

			cmd.Println(v)
		},
	}

	return cmd
}
