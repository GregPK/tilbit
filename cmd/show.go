package cmd

import (
	"fmt"

	"github.com/GregPK/tilbit/core"
	"github.com/spf13/cobra"
)

var showCmd = ShowCmd()

func ShowCmd(inputTilbits ...core.Tilbit) *cobra.Command {
	cmd := cobra.Command{
		Use:   "show",
		Short: "Show specific TILBit",
		Long:  `Shows a specific TILBit when given ID, shows random otherwise`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tilbits, err := core.GetTilbits(Config.sources, args[0], inputTilbits)
			if err != nil {
				panic(err)
			}
			for _, tilbit := range tilbits {
				text, _ := core.GetBitString(tilbit, Config.outputFormat)
				fmt.Fprintf(cmd.OutOrStdout(), text)
			}
		},
	}
	cmd.Flags().StringVarP(&Config.outputFormat, "output-format", "f", "box", "Output format for show command")
	cmd.Flags().StringSliceVarP(&Config.sources, "source", "i", make([]string, 0), "Source list. Currently only handles directories. Ultimately will handle arbitrary URIs")
	return &cmd
}
