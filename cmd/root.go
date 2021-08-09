package cmd

import (
	"fmt"

	"github.com/GregPK/tilbit/core"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "tilbit",
		Short: "TILBit",
		Long: `A Fast and Flexible Static Site Generator built with
						love by spf13 and friends in Go.
						Complete documentation is available at http://hugo.spf13.com`,
		Run: func(cmd *cobra.Command, args []string) {
			tilbits := parseFile(privateDbFilename())

			randTil := getRandomBit(tilbits)

			fmt.Println(core.GetBitString(randTil))
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}