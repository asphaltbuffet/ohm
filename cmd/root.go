// Package cmd contains all CLI commands used by the application.
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/asphaltbuffet/ohm/cmd/axial"
	"github.com/asphaltbuffet/ohm/cmd/man"
	versionCmd "github.com/asphaltbuffet/ohm/cmd/version"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := NewCommand().Execute()
	if err != nil {
		os.Exit(1)
	}
}

// GetRootCommand returns the root command for the CLI.
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "ohm",
		Short:             "ohm calculates values for axial resistors",
		Args:              cobra.NoArgs,
		CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
	}

	cmd.AddCommand(axial.NewCommand())

	cmd.AddCommand(man.NewCommand())
	cmd.AddCommand(versionCmd.NewCommand())

	return cmd
}
