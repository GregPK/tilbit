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
			tilbits := []core.Tilbit{}
			if len(args) == 0 {
				allTilbits := core.AllTilbits()
				randTil := getRandomBit(allTilbits)
				tilbits = append(tilbits, randTil)
			} else {
				ids := core.ParseIdsFromString(args[0])

				var err error
				tilbits, err = core.ByIds(ids)
				if err != nil {
					panic(err)
				}
			}

			for _, tilbit := range tilbits {
				text, _ := core.GetBitString(tilbit, true)
				fmt.Println(text)
			}
		},
	}
)
