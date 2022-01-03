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
		Long:  `Shows a specific TILBit when given ID, shows randon otherwise`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tilbits, err := core.ByQuery(args[0], inputTilbits)
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
	return &cmd
}
