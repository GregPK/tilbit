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
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tilbits := []core.Tilbit{}
			if len(args) == 0 {
				allTilbits := core.AllTilbits()
				randTil := getRandomBit(allTilbits)
				tilbits = append(tilbits, randTil)
			} else {
				tilbit, err := core.ById(args[0])
				if err != nil {
					panic(err)
				}
				tilbits = []core.Tilbit{tilbit}
			}

			for _, tilbit := range tilbits {
				text, _ := core.GetBitString(tilbit)
				fmt.Println(text)
			}
		},
	}
)
