package cmd

import (
	"fmt"

	"github.com/GregPK/tilbit/core"
	"github.com/spf13/cobra"
)

var (
	showCmd = &cobra.Command{
		Use:   "show",
		Short: "Show specific TILBit",
		Long:  `Shows a specific TILBit when given ID, shows randon otherwise`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tilbits, err := core.ByQuery(args[0], nil)
			if err != nil {
				panic(err)
			}

			for _, tilbit := range tilbits {
				text, _ := core.GetBitString(tilbit, Config.outputFormat == "box")
				fmt.Println(text)
			}
		},
	}
)
